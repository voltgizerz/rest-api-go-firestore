package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/env"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/logger"
)

// Database - client
type Database struct {
	FirestoreClient *firestore.Client
}

func InitDB(ctx context.Context, cfg *config.Config) *Database {
	span, ctx := opentracing.StartSpanFromContext(ctx, "database.InitDB")
	defer span.Finish()

	// * Make sure your service account credential json correct.
	serviceAccountFirestoreFilePath := fmt.Sprintf("./config/credential/sa-%s.json", env.GetENV())
	sa := option.WithCredentialsFile(serviceAccountFirestoreFilePath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		logger.Log.Fatalf("[InitDB.NewApp] err: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		logger.Log.Fatalf("[InitDB.Firestore] err: %v", err)
	}

	logger.Log.Info("[InitDB] Database firestore connected succesfully...")

	return &Database{
		FirestoreClient: client,
	}
}

// CloseFirestoreClient - clone firestore client
func (d *Database) CloseFirestoreClient() {
	d.FirestoreClient.Close()
}
