package auth

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// ================== FIREBASE CLIENT ==================
var client *auth.Client

func InitFirebase() error {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil,
		option.WithCredentialsFile("firebase-service-account.json"),
	)
	if err != nil {
		return err
	}

	client, err = app.Auth(ctx)
	return err
}

func VerifyToken(idToken string) (string, error) {
	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return "", err
	}

	return token.UID, nil
}
