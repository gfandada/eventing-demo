/*
Copyright 2020 gfandada@gmail.com.
*/
package main

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/google/uuid"
	"log"
)

type ReceiveData struct {
	// Msg holds the message from the event
	Msg string `json:"msg,omitempty"`
}

type EventingDemo struct {
	// Msg holds the message from the event
	Msg string `json:"msg,omitempty"`
}

func receive(ctx context.Context, event cloudevents.Event, response *cloudevents.EventResponse) error {
	log.Printf("Event received. Context: %v\n", event.Context)

	data := &ReceiveData{}
	if err := event.DataAs(data); err != nil {
		log.Printf("Error while extracting cloudevent Data: %s\n", err.Error())
		return err
	}
	log.Printf("Event received. Msg: %s\n", data.Msg)

	newEvent := cloudevents.NewEvent()
	newEvent.SetID(uuid.New().String())
	newEvent.SetSource("github.com/gfandada/eventing-demo")
	newEvent.SetType("gfandada.eventing-demo.eventingdemo")
	newEvent.SetData(EventingDemo{Msg: "event from github.com/gfandada/eventing-demo"})
	response.RespondWith(200, &newEvent)

	log.Printf("Responded with event %v", newEvent)

	return nil
}

func main() {
	log.Print("eventing-demo started.")
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), receive))
}
