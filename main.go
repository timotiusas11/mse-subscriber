package main

import (
	"log"
	"mse-subscriber/infrastructure"
	"mse-subscriber/interfaces"
	"mse-subscriber/usecases"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

func main() {
	client, err := infrastructure.GetESClient()
	if err != nil {
		return
	}
	log.Println("ES initialized.")

	repo := interfaces.NewRecipeRepository(client)
	recipeInteractor := new(usecases.RecipeInteractor)
	recipeInteractor.RecipeRepository = repo

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("recipes", "synch-elasticsearch", config)
	if err != nil {
		return
	}

	log.Println("Listening to NSQ.")

	consumer.AddHandler(&interfaces.RecipeHandler{RecipeInteractor: recipeInteractor})

	err = consumer.ConnectToNSQLookupd("http://127.0.0.1:4161")

	if err != nil {
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	consumer.Stop()
}
