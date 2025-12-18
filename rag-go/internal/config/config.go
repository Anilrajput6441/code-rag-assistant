package config

import (
	"os"

	"github.com/joho/godotenv"
)

// ================== APPLICATION CONFIG ==================
type Config struct {
	Port              string
	FirebaseProjectID string
	EncryptionSecret  string
}

func Load() Config {
	_ = godotenv.Load()

	// ========== LOAD ENVIRONMENT VARIABLES ==========
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	
	firebaseProjectID := os.Getenv("FIREBASE_PROJECT_ID")
	if firebaseProjectID == "" {
		firebaseProjectID = "code-rag-assistant"
	}
	
	encryptionSecret := os.Getenv("ENCRYPTION_SECRET")
	if encryptionSecret == "" {
		encryptionSecret = "default-secret-key-change-in-production"
	}

	return Config{
		Port:              port,
		FirebaseProjectID: firebaseProjectID,
		EncryptionSecret:  encryptionSecret,
	}
}
