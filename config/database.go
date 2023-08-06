package config

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/voltgizerz/rest-api-go-firestore/logger"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Database - client
type Database struct {
	FirestoreClient *firestore.Client
}

// InitDB - .
func InitDB() Database {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./credential/sa-key.json") // * Set your service account json here
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	logger.Log.Info("Database Firestore connected succesfully...")

	return Database{
		FirestoreClient: client,
	}
}
