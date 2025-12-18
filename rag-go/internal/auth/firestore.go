package auth

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// ================== FIRESTORE CLIENT ==================
var Firestore *firestore.Client

func InitFirestore(projectID string) error {
	ctx := context.Background()
	
	client, err := firestore.NewClient(ctx, projectID,
		option.WithCredentialsFile("firebase-service-account.json"),
	)
	if err != nil {
		return err
	}
	
	Firestore = client
	return nil
}
