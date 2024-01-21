package db

import (
	"context"
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

func getRandomRecipeIngredients(amount int) []CreateRecipeIngredientParams {
	var recipeIngredients []CreateRecipeIngredientParams
	for i := 1; i < amount+1; i++ {
		recipeIngredients = append(recipeIngredients, CreateRecipeIngredientParams{
			IngredientName: util.RandomString(10),
			Rank:           int32(i),
			Unit:           util.RandomString(5),
			Amount:         float64(util.RandomInt(0, 10)),
		})
	}

	return recipeIngredients
}

func getRandomRecipeSteps(amount int) []CreateRecipeStepParams {
	var recipeSteps []CreateRecipeStepParams
	for i := 1; i < amount+1; i++ {
		recipeSteps = append(recipeSteps, CreateRecipeStepParams{
			Rank:            int32(i),
			StepDescription: util.RandomString(20),
		})
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
