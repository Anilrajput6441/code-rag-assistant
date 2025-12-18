package user

import (
	"context"
	"errors"

	"rag-go/internal/auth"
	"rag-go/internal/security"
)

func GetUserAPIKey(uid, secret string) (string, error) {
	ctx := context.Background()
	doc, err := auth.Firestore.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		return "", err
	}

	enc, ok := doc.Data()["gemini_api_key"].(string)
	if !ok {
		return "", errors.New("api key not set")
	}

	return security.Decrypt(enc, secret)
}
