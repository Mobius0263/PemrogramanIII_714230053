package config

var allowedOrigins = []string{
	"http://localhost:3000",
	"http://Mobius0263.github.io",
	"http://localhost:5173/",
	"http://127.0.0.1:8099/",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
