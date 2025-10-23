package utils

import (
	"context"
	"mime/multipart"
	"path"
	"strings"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func ExtractPublicID(url string) string {
	parts := strings.Split(url, "/upload/")
	if len(parts) < 2 {
		return ""
	}

	publicPath := parts[1]

	if strings.HasPrefix(publicPath, "v") {
		if idx := strings.Index(publicPath, "/"); idx != -1 {
			publicPath = publicPath[idx+1:]
		}
	}

	publicPath = strings.TrimSuffix(publicPath, path.Ext(publicPath))

	publicPath = strings.Trim(publicPath, "/")

	return publicPath
}

func UploadToCloudinary(ctx context.Context, file *multipart.FileHeader, cld *cloudinary.Cloudinary, cfg *configs.Config) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	res, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{
		Folder: cfg.Cloudinary.Folder,
	})
	if err != nil {
		return "", err
	}

	return res.SecureURL, nil
}

func DeleteFromCloudinary(ctx context.Context, cld *cloudinary.Cloudinary, imageURL string) error {
	publicID := ExtractPublicID(imageURL)
	if publicID == "" {
		return nil
	}

	_, err := cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
		PublicIDs: []string{publicID},
	})

	return err
}
