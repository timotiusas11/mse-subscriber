package interfaces

import (
	"context"
	"encoding/json"
	"log"
	"mse-subscriber/domain"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type recipeRepo struct {
	client *elasticsearch.Client
}

func NewRecipeRepository(client *elasticsearch.Client) domain.RecipeRepository {
	return &recipeRepo{client: client}
}

func (repo *recipeRepo) Create(recipe domain.Recipe) error {
	dataJson, err := json.Marshal(recipe)
	if err != nil {
		return err
	}
	js := string(dataJson)

	request := esapi.IndexRequest{Index: "recipes", DocumentID: strconv.Itoa(recipe.Id), Body: strings.NewReader(js)}

	res, err := request.Do(context.Background(), repo.client)

	defer res.Body.Close()

	if err != nil {
		return err
	}

	log.Println(res)

	return nil
}

func (repo *recipeRepo) Update(recipe domain.Recipe) error {
	dataJson, err := json.Marshal(recipe)
	if err != nil {
		return err
	}
	js := string(dataJson)

	request := esapi.IndexRequest{Index: "recipes", DocumentID: strconv.Itoa(recipe.Id), Body: strings.NewReader(js)}

	res, err := request.Do(context.Background(), repo.client)

	defer res.Body.Close()

	if err != nil {
		return err
	}

	log.Println(res)

	return nil
}

func (repo *recipeRepo) Delete(id int) error {
	request := esapi.DeleteRequest{Index: "recipes", DocumentID: strconv.Itoa(id)}

	res, err := request.Do(context.Background(), repo.client)

	defer res.Body.Close()

	if err != nil {
		return err
	}

	return nil
}
