# Go-Firestore

A thin wrapper around the Google's [Firestore client library for Go](https://github.com/googleapis/google-cloud-go/tree/master/firestore).

Examples below:

```go
package main

import (
   "fmt"
   "log"

   fire "github.com/wkirk01/Go-Firestore"
)

type person struct {
   First string
   Last  string
}

func main() {
   client, err := fire.NewFirestoreClient("/path/to/serviceaccount.json")
   if err != nil {
	   log.Fatalf("Could not initialize Cloud Firestore: %v", err)
   }
   defer client.Close()

   p1 := person{
	   First: "Billy",
	   Last:  "Kirk",
   }

   docID, err := client.CreateDocument(client.Db.Collection("users"), p1)
   if err != nil {
	   log.Fatalf("Failed adding document: %v", err)
   }

   fmt.Println("Saved document with ID:", docID)

   err = client.DeleteDocument(client.Db.Collection("users"), docID)
   if err != nil {
 	   log.Fatalf("Failed deleting document: %v", err)
   }

   fmt.Println("Deleted document with ID:", docID)
}
```