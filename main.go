package main

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	generateV4GetObjectSignedURL("gcs-dev-poc-8ede431f7e54b0f5", "appeltaart.png")
}

// generateV4GetObjectSignedURL generates object signed URL.
func generateV4GetObjectSignedURL(bucket, object string) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("auth.json"))
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(1 * time.Minute),
	}

	u, err := client.Bucket(bucket).SignedURL(object, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucket, err)
	}

	fmt.Println("Generated GET signed URL:")
	fmt.Printf("%q\n", u)
	return u, nil
}
