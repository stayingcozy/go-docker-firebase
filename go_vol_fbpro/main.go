package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func write_hello() {
	// Path within the container corresponding to the mounted volume
	filePath := "/test/hello.txt"

	// Open the file in write mode, create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write a line to the file
	line := "Hello from inside the container; We're doing some Firebase Writing\n"
	_, err = file.WriteString(line)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message written to %s\n", filePath)
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func addDocAsMap(ctx context.Context, client *firestore.Client) error {
	_, err := client.Collection("cities").NewDoc().Set(ctx, map[string]interface{}{
		"name":    "Los Angeles",
		"state":   "CA",
		"country": "USA",
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

func main() {

	// Init values
	path_to_serviceAccountKey := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	const project_id = os.Getenv("PROJECT_ID")
	var app *firebase.App
	var err error

	if len(path_to_serviceAccountKey) > 0 {
		fmt.Println("Using Service Account key path")
		opt := option.WithCredentialsFile(path_to_serviceAccountKey)
		config := &firebase.Config{ProjectID: project_id}
		app, err = firebase.NewApp(context.Background(), config, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
	} else {
		fmt.Println("By pass service acc. path. In production or check GOOGLE_APPLICATION_CREDENTIALS")
		config := &firebase.Config{ProjectID: project_id} // project id is required to access Firestore
		app, err = firebase.NewApp(context.Background(), config)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
	}

	// Initialize the Firebase Admin SDK
	ctx := context.Background()

	// Setup Context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Write to firebase //
	addDocAsMap(ctx, client)
}