// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func CreateBeacon(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if r.Method == http.MethodPost {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var d struct {
			Source  string `json:"source"`
			Content string `json:"content"`
		}
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if d.Source == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Sets your Google Cloud Platform project ID.
		// projectID := os.Getenv("project-id")
		bucketName := os.Getenv("bucket-name")

		// Creates a client.
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Creates a Bucket instance.
		bucket := client.Bucket(bucketName)

		// Write file
		fileName := fmt.Sprintf("%s-%d", d.Source, time.Now().Unix())
		wc := bucket.Object(fileName).NewWriter(ctx)
		_, err = fmt.Fprintf(wc, d.Content)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Close, just like writing a file.
		err = wc.Close()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
