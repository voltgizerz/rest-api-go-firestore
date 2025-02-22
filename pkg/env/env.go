package env

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"

	"github.com/voltgizerz/rest-api-go-firestore/pkg/logger"
)

// LoadENV - load env file.
func LoadENV(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "pkg.Env.LoadENV")
	defer span.Finish()

	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("[LoadENV] No .env file found")
	}
}

func GetENV() string {
	envName, exists := os.LookupEnv("GO_ENV")
	if !exists {
		logger.Log.Fatalln("[NewConfiGetEnvg] GO_ENV is not set")
	}

	return envName
}
