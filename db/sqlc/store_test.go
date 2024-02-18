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

	require.Equal(t, recipeArg.RecipeName, result.Recipe.RecipeName)
	require.Equal(t, recipeArg.Link, result.Recipe.Link)
	require.Equal(t, recipeArg.AuthorID, result.Recipe.AuthorID)
	require.Equal(t, recipeArg.PrepTime, result.Recipe.PrepTime)
	require.Equal(t, recipeArg.PrepTimeUnit, result.Recipe.PrepTimeUnit)
	require.Equal(t, recipeArg.UserCreated, result.Recipe.UserCreated)

	require.Equal(t, len(result.RecipeIngredients), len(recipeIngredientsArg))
	require.Equal(t, len(result.RecipeSteps), len(recipeStepsArg))

	for _, resultI := range result.RecipeIngredients {
		require.Equal(t, resultI.RecipeID, result.Recipe.ID)
		for _, createI := range recipeIngredientsArg {
			if resultI.IngredientName == createI.IngredientName {
				require.Equal(t, resultI.Amount, createI.Amount)
				require.Equal(t, resultI.Unit, createI.Unit)
				require.Equal(t, resultI.Rank, createI.Rank)
			}
		}
	}

	for _, resultS := range result.RecipeSteps {
		require.Equal(t, resultS.RecipeID, result.Recipe.ID)
		for _, createI := range recipeStepsArg {
			if resultS.StepDescription == createI.StepDescription {
				require.Equal(t, resultS.Rank, createI.Rank)
			}
		}
	}

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

	require.Equal(t, createRecipeResult.Recipe.AuthorID, result.Recipe.AuthorID)

	require.Equal(t, createRecipeResult.Recipe.RecipeName, result.Recipe.RecipeName)
	require.Equal(t, createRecipeResult.Recipe.Link, result.Recipe.Link)
	require.Equal(t, createRecipeResult.Recipe.AuthorID, result.Recipe.AuthorID)
	require.Equal(t, createRecipeResult.Recipe.PrepTime, result.Recipe.PrepTime)
	require.Equal(t, createRecipeResult.Recipe.PrepTimeUnit, result.Recipe.PrepTimeUnit)
	require.Equal(t, createRecipeResult.Recipe.UserCreated, result.Recipe.UserCreated)
	// TODO: Check values of recipeIngredients and recipeSteps
}

func TestGetRecipeTx(t *testing.T) {
	createRecipeResult := createRandomRecipeTx(t)

	store := NewStore(testDB)

	result, err := store.GetRecipeTx(context.Background(), createRecipeResult.Recipe.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, createRecipeResult.Recipe.RecipeName, result.Recipe.RecipeName)
	require.Equal(t, createRecipeResult.Recipe.Link, result.Recipe.Link)
	require.Equal(t, createRecipeResult.Recipe.AuthorID, result.Recipe.AuthorID)
	require.Equal(t, createRecipeResult.Recipe.PrepTime, result.Recipe.PrepTime)
	require.Equal(t, createRecipeResult.Recipe.PrepTimeUnit, result.Recipe.PrepTimeUnit)
	require.Equal(t, createRecipeResult.Recipe.UserCreated, result.Recipe.UserCreated)
	// TODO: Check values of recipeIngredients and recipeSteps
}

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

	updateResult, err := store.UpdateRecipeTx(context.Background(), createRecipeResult.Recipe.ID, newRecipe, newRecipeIngredients, newRecipeSteps)
	require.NoError(t, err)
	require.NotEmpty(t, updateResult)

	require.Equal(t, newRecipe.RecipeName, updateResult.Recipe.RecipeName)
	require.Equal(t, newRecipe.Link, updateResult.Recipe.Link)
	require.Equal(t, newRecipe.AuthorID, updateResult.Recipe.AuthorID)
	require.Equal(t, newRecipe.PrepTime, updateResult.Recipe.PrepTime)
	require.Equal(t, newRecipe.PrepTimeUnit, updateResult.Recipe.PrepTimeUnit)
	// TODO: Check values of recipeIngredients and recipeSteps
}
