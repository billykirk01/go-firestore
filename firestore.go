package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/fatih/structs"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

//FiresoreClient provides access to the Firestore service.
type FiresoreClient struct {
	ctx context.Context
	Db  *firestore.Client
}

// NewClient creates a new Cloud Firestore Database Client
func NewClient(crednetialsPath string) (*FiresoreClient, error) {
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
		Db:  client,
	}, nil
}

//Close closes any resources held by the client
func (client *FiresoreClient) Close() {
	client.Db.Close()
}

//CreateDocument creates a new document in a collection
func (client *FiresoreClient) CreateDocument(collectionRef *firestore.CollectionRef, data interface{}) (string, error) {
	docRef, _, err := collectionRef.Add(client.ctx, structs.Map(data))
	if err != nil {
		log.Warnf("Failed adding document: %v", err)
		return "", err
	}
	return docRef.ID, nil
}

//DeleteDocument deletes a document in a collection
func (client *FiresoreClient) DeleteDocument(collectionRef *firestore.CollectionRef, docID string) error {
	_, err := collectionRef.Doc(docID).Delete(client.ctx)
	if err != nil {
		log.Warnf("Failed adding document: %v", err)
		return err
	}
	return nil
}

//GetDocument gets a document in a collection
func (client *FiresoreClient) GetDocument(collectionRef *firestore.CollectionRef, docID string) (interface{}, error) {
	dsnap, err := collectionRef.Doc(docID).Get(client.ctx)
	if err != nil {
		log.Warnf("Failed getting document: %v", err)
		return nil, err
	}
	return dsnap.Data(), nil
}

//GetDocuments gets documents in a collection reference
func (client *FiresoreClient) GetDocuments(collectionRef *firestore.CollectionRef) ([]interface{}, error) {
	documents := []interface{}{}
	iter := collectionRef.Documents(client.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc.Data())
	}
	return documents, nil
}

//Query gets documents in a collection reference
func (client *FiresoreClient) Query(query firestore.Query) ([]interface{}, error) {
	documents := []interface{}{}
	iter := query.Documents(client.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc.Data())
	}
	return documents, nil
}
