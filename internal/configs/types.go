package configs

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Service struct {
		SecretKey string `yaml:"secretKey"`
	} `yaml:"service"`
	Database struct {
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
	Cloudinary struct {
		APIKey    string `yaml:"apiKey"`
		APISecret string `yaml:"apiSecret"`
		CloudName string `yaml:"cloudName"`
		Folder    string `yaml:"folder"`
	} `yaml:"cloudinary"`
}
