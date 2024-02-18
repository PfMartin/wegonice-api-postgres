package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/PfMartin/wegonice-api/util"
	"github.com/stretchr/testify/require"
)

func getRandomRecipe(author Author) CreateRecipeParams {
	return CreateRecipeParams{
		RecipeName:   util.RandomString(6),
		Link:         util.RandomString(10),
		AuthorID:     author.ID,
		PrepTime:     float64(util.RandomInt(0, 100)),
		PrepTimeUnit: util.RandomString(5),
		UserCreated:  author.UserCreated,
	}
}

func getRandomRecipeIngredients(amount int, params ...int64) []CreateRecipeIngredientParams {
	var recipeIngredients []CreateRecipeIngredientParams
	for i := 1; i < amount+1; i++ {
		newIngredient := CreateRecipeIngredientParams{
			IngredientName: util.RandomString(10),
			Rank:           int32(i),
			Unit:           util.RandomString(5),
			Amount:         float64(util.RandomInt(0, 10)),
		}

		if len(params) > 0 {
			newIngredient.RecipeID = params[0]
		}

		recipeIngredients = append(recipeIngredients, newIngredient)
	}

	return recipeIngredients
}

func getRandomRecipeSteps(amount int, params ...int64) []CreateRecipeStepParams {
	var recipeSteps []CreateRecipeStepParams
	for i := 1; i < amount+1; i++ {
		newStep := CreateRecipeStepParams{
			Rank:            int32(i),
			StepDescription: util.RandomString(20),
		}

		if len(params) > 0 {
			newStep.RecipeID = params[0]
		}

		recipeSteps = append(recipeSteps, newStep)
	}

	return recipeSteps
}

func createRandomRecipeTx(t *testing.T) CreateRecipeTxResult {
	store := NewStore(testDB)

	author := createRandomAuthor(t)

	recipeArg := getRandomRecipe(author)

	recipeIngredientsArg := getRandomRecipeIngredients(10)
	recipeStepsArg := getRandomRecipeSteps(8)

	result, err := store.CreateRecipeTx(context.Background(), recipeArg, recipeIngredientsArg, recipeStepsArg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// TODO: Check values of inserted recipe, recipeIngredients and recipeSteps

	return result
}

func TestCreateRecipeTx(t *testing.T) {
	createRandomRecipeTx(t)
}

func TestDeleteRecipeTx(t *testing.T) {
	createRecipeResult := createRandomRecipeTx(t)

	store := NewStore(testDB)

	result, err := store.DeleteRecipeTx(context.Background(), createRecipeResult.Recipe.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// TODO: Check values of inserted recipe, recipeIngredients and recipeSteps
}

func TestGetRecipeTx(t *testing.T) {
	createRecipeResult := createRandomRecipeTx(t)

	store := NewStore(testDB)

	result, err := store.GetRecipeTx(context.Background(), createRecipeResult.Recipe.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// TODO: Check values of inserted recipe, recipeIngredients and recipeSteps
}

// TODO: Create Test for UpdateRecipeTx
func TestUpdateRecipeTx(t *testing.T) {
	createRecipeResult := createRandomRecipeTx(t)

	store := NewStore(testDB)

	newRecipe := UpdateRecipeByIdParams{
		ID:           createRecipeResult.Recipe.ID,
		RecipeName:   "Updated recipe name",
		Link:         "https://updated-website.com",
		AuthorID:     createRecipeResult.Recipe.AuthorID,
		PrepTime:     float64(2),
		PrepTimeUnit: "Hours",
	}

	// Only new recipe ingredients
	// TODO: Check updating recipe ingredients and updating recipe steps
	newRecipeIngredients := getRandomRecipeIngredients(5, createRecipeResult.Recipe.ID)
	for _, i := range newRecipeIngredients {
		i.RecipeID = createRecipeResult.Recipe.ID
		fmt.Println(i.RecipeID)
	}

	newRecipeSteps := getRandomRecipeSteps(5, createRecipeResult.Recipe.ID)
	for _, s := range newRecipeSteps {
		s.RecipeID = createRecipeResult.Recipe.ID
	}

	updatedRecipe, err := store.UpdateRecipeTx(context.Background(), createRecipeResult.Recipe.ID, newRecipe, newRecipeIngredients, newRecipeSteps)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRecipe)
}
