package cloudinary

import (
	"log"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/cloudinary/cloudinary-go/v2"
)

func Init(cfg *configs.Config) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(
		cfg.Cloudinary.CloudName,
		cfg.Cloudinary.APIKey,
		cfg.Cloudinary.APISecret,
	)

	if err != nil {
		log.Fatalf("failed to init cloudinary : %v", err)
		return nil, err
	}

	return cld, nil
}
