package config

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/voltgizerz/rest-api-go-firestore/logger"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const (
	SERVICE_ACCOUNT_CREDENTIAL_FILE_PATH = "./credential/sa-key.json"
)

// Database - client
type Database struct {
	FirestoreClient *firestore.Client
}

// * InitDB - Make sure your service account credential json correct.
func InitDB() *Database {
	ctx := context.Background()

	sa := option.WithCredentialsFile(SERVICE_ACCOUNT_CREDENTIAL_FILE_PATH)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	logger.Log.Info("Database firestore connected succesfully...")

	return &Database{
		FirestoreClient: client,
	}
}

// CloseFirestoreClient - clone firestore client
func (d *Database) CloseFirestoreClient() {
	d.FirestoreClient.Close()
}
