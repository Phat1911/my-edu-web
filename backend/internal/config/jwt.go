package config

import "os"

func JWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "eduweb-secret-key-2026"
	}
	return []byte(secret)
}
