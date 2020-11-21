package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/fatih/structs"
	"google.golang.org/api/option"
)

//FiresoreClient provides access to the Firestore service.
type FiresoreClient struct {
	ctx context.Context
	db  *firestore.Client
}

// NewFirestoreClient creates a new Cloud Firestore Database Client
func NewFirestoreClient(crednetialsPath string) (*FiresoreClient, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile(crednetialsPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &FiresoreClient{
		ctx: ctx,
		db:  client,
	}, nil
}

//CreateDocument creates a new document in a collection
func (client *FiresoreClient) CreateDocument(collectionRef *firestore.CollectionRef, data interface{}) (string, error) {
	docRef, _, err := collectionRef.Add(client.ctx, structs.Map(data))
	if err != nil {
		log.Fatalf("Failed adding document: %v", err)
		return "", err
	}
	return docRef.ID, nil
}

//DeleteDocument deletes a document in a collection
func (client *FiresoreClient) DeleteDocument(collectionRef *firestore.CollectionRef, docID string) error {
	_, err := collectionRef.Doc(docID).Delete(client.ctx)
	if err != nil {
		log.Fatalf("Failed adding document: %v", err)
		return err
	}
	return nil
}

//Close closes any resources held by the client
func (client *FiresoreClient) Close() {
	client.db.Close()
}
