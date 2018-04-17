package main

import (
"firebase.google.com/go"
"google.golang.org/api/option"
	"context"
	"log"
	"fmt"
	"google.golang.org/api/iterator"
)
func main()  {

	sa:= option.WithCredentialsFile("filestore.json")
	app,err:=firebase.NewApp(context.Background(),nil,sa)
	client,err:=app.Firestore(context.Background())
	if err!=nil {
		log.Fatalln(err)
	}

	// read cloud firestore data.
	iter := client.Collection("Football").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

	// add to data cloud firestore.
	_, _, err = client.Collection("Football").Add(context.Background(), map[string]interface{}{
		"id": 3,
		"club": "Göztepespor",
	})
	if err != nil {
		log.Fatalf("Failed adding club: %v", err)
	}


	// condition data.

	mostClub:=client.Collection("Football").Where("Club","==","Galatasaray").Documents(context.Background())
	for {
		doc, err := mostClub.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

	defer client.Close()
}



