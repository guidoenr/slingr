package main

import (
	"context"
	"fmt"
	"log"
	"unicode/utf8"

	"cloud.google.com/go/storage"
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/iterator"
)

func main() {
	validString()
}

func validString() {
	bad_chars := []string{"â", "€", "™"}
	for i := range bad_chars {
		fmt.Printf("valid string: %v -> %v \n", bad_chars[i], utf8.ValidString(bad_chars[i]))
	}

}


// implicit uses Application Default Credentials to authenticate.
func implicit() {
	ctx := context.Background()

	// For API packages whose import path is starting with "cloud.google.com/go",
	// such as cloud.google.com/go/storage in this case, if there are no credentials
	// provided, the client library will look for credentials in the environment.
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer storageClient.Close()

	it := storageClient.Buckets(ctx, "project-id")
	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bucketAttrs.Name)
	}

	// For packages whose import path is starting with "google.golang.org/api",
	// such as google.golang.org/api/cloudkms/v1, use NewService to create the client.
	kmsService, err := cloudkms.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_ = kmsService
}
