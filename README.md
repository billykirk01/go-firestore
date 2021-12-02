# ✨ go-firestore

A thin wrapper around the Google's [Firestore client library for Go](https://github.com/googleapis/google-cloud-go/tree/master/firestore).

## 📦 Installing

```go
go get github.com/wkirk01/go-firestore
```

## 📖 Example Usage

```go
package main

import (
   "fmt"
   log "github.com/sirupsen/logrus"
   firestore "github.com/wkirk01/go-firestore"
)

type person struct {
   First string
   Last  string
}

func main() {
   client, err := firestore.NewClient("~/path/to/firebase/credentials.json")
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

   documents, err := client.GetDocuments(client.Db.Collection("users"))
   if err != nil {
      log.Fatalf("Failed adding document: %v", err)
   }

   fmt.Printf("Documents in users collection: %v\n", documents)

   queryResults, err := client.Query(client.Db.Collection("users").Where("First", "==", "foo"))
   if err != nil {
      log.Fatalf("Failed adding document: %v", err)
   }

   fmt.Printf("Documents in users collection with foo for first name: %v\n", queryResults)

   err = client.DeleteDocument(client.Db.Collection("users"), docID)
   if err != nil {
      log.Fatalf("Failed deleting document: %v", err)
   }

   fmt.Printf("Deleted document with ID: %s\n", docID)
}
```
