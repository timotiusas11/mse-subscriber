package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mse-subscriber/usecases"

	"github.com/nsqio/go-nsq"
)

type RecipeInteractor interface {
	InsertRecipe(recipe usecases.Recipe) error
	UpdateRecipe(recipe usecases.Recipe) error
	DeleteRecipe(id int) error
}

type RecipeHandler struct {
	RecipeInteractor RecipeInteractor
}

func (handler *RecipeHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return errors.New("Body is empty.")
	}

	data := struct {
		Action string
		Data   usecases.Recipe
	}{}

	err := json.Unmarshal(m.Body, &data)

	if err != nil {
		log.Println("Error when Unmarshaling the message body, Err : ", err)
		// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
		return err
	}

	if data.Action == "create" {
		err = handler.RecipeInteractor.InsertRecipe(data.Data)
	} else if data.Action == "update" {
		err = handler.RecipeInteractor.UpdateRecipe(data.Data)
	} else if data.Action == "delete" {
		err = handler.RecipeInteractor.DeleteRecipe(data.Data.Id)
	}

	if err != nil {
		return err
	}

	log.Println("===============")
	log.Println("Data :")
	log.Println(fmt.Sprintf("Action : %s", data.Action))
	log.Println(fmt.Sprintf("Name : %s", data.Data.Name))
	log.Println(fmt.Sprintf("Description : %v", data.Data.Description))
	log.Println("===============")

	return nil
}
