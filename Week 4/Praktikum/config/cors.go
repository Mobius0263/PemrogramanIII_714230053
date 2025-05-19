package config

var allowedOrigins = []string{
	"http://localhost:3000",
	"http://Mobius0263.github.io",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
