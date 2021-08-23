package usecases

import "mse-subscriber/domain"

type Recipe struct {
	Id           int
	Name         string
	Ingredients  []string
	IsHalal      bool
	IsVegetarian bool
	Description  string
	Rating       float64
}

type RecipeInteractor struct {
	RecipeRepository domain.RecipeRepository
}

func (interactor *RecipeInteractor) InsertRecipe(recipe Recipe) error {
	err := interactor.RecipeRepository.Create(domain.Recipe{
		Id:           recipe.Id,
		Name:         recipe.Name,
		Ingredients:  recipe.Ingredients,
		IsHalal:      recipe.IsHalal,
		IsVegetarian: recipe.IsVegetarian,
		Description:  recipe.Description,
		Rating:       recipe.Rating,
	})
	if err != nil {
		return err
	}
	return nil
}

func (interactor *RecipeInteractor) UpdateRecipe(recipe Recipe) error {
	err := interactor.RecipeRepository.Update(domain.Recipe{
		Id:           recipe.Id,
		Name:         recipe.Name,
		Ingredients:  recipe.Ingredients,
		IsHalal:      recipe.IsHalal,
		IsVegetarian: recipe.IsVegetarian,
		Description:  recipe.Description,
		Rating:       recipe.Rating,
	})
	if err != nil {
		return err
	}
	return nil
}

func (interactor *RecipeInteractor) DeleteRecipe(id int) error {
	err := interactor.RecipeRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
