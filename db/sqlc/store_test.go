package db

import (
	"context"
	"testing"

	"github.com/PfMartin/wegonice-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomRecipeTx(t *testing.T) CreateRecipeTxResult {
	store := NewStore(testDB)

	author := createRandomAuthor(t)

	recipeArg := CreateRecipeParams{
		RecipeName:   util.RandomString(6),
		Link:         util.RandomString(10),
		AuthorID:     author.ID,
		PrepTime:     float64(util.RandomInt(0, 100)),
		PrepTimeUnit: util.RandomString(5),
		UserCreated:  author.UserCreated,
	}

	var recipeIngredientsArg []CreateRecipeIngredientParams
	for i := 1; i < 11; i++ {
		recipeIngredientsArg = append(recipeIngredientsArg, CreateRecipeIngredientParams{
			IngredientName: util.RandomString(10),
			Rank:           int32(i),
			Unit:           util.RandomString(5),
			Amount:         float64(util.RandomInt(0, 10)),
		})
	}

	var recipeStepsArg []CreateRecipeStepParams
	for i := 1; i < 9; i++ {
		recipeStepsArg = append(recipeStepsArg, CreateRecipeStepParams{
			Rank:            int32(i),
			StepDescription: util.RandomString(20),
		})
	}

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
