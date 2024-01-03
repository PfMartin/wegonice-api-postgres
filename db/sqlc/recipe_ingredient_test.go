package db

import (
	"context"
	"testing"
	"time"

	"github.com/PfMartin/wegonice-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomRecipeIngredient(t *testing.T, params ...int64) RecipeIngredient {
	createdRecipe := createRandomRecipe(t)

	recipeId := createdRecipe.ID

	if len(params) > 0 {
		recipeId = params[0]
	}

	arg := CreateRecipeIngredientParams{
		IngredientName: util.RandomString(6),
		Rank:           1,
		Unit:           util.RandomString(4),
		Amount:         float64(util.RandomInt(0, 1000)),
		RecipeID:       recipeId,
	}

	recipeIngredient, err := testQueries.CreateRecipeIngredient(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipeIngredient)

	require.Equal(t, arg.IngredientName, recipeIngredient.IngredientName)
	require.Equal(t, arg.Rank, recipeIngredient.Rank)
	require.Equal(t, arg.Unit, recipeIngredient.Unit)
	require.Equal(t, arg.Amount, recipeIngredient.Amount)
	require.Equal(t, arg.RecipeID, recipeIngredient.RecipeID)

	require.NotZero(t, recipeIngredient.CreatedAt)

	return recipeIngredient
}

func TestCreateRecipeIngredient(t *testing.T) {
	createRandomRecipeIngredient(t)
}

func TestDeleteRecipeIngredientById(t *testing.T) {
	recipeIngredient := createRandomRecipeIngredient(t)

	deletedRecipeIngredient, err := testQueries.DeleteRecipeIngredientById(context.Background(), recipeIngredient.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedRecipeIngredient)

	require.Equal(t, recipeIngredient.IngredientName, deletedRecipeIngredient.IngredientName)
	require.Equal(t, recipeIngredient.Rank, deletedRecipeIngredient.Rank)
	require.Equal(t, recipeIngredient.Unit, deletedRecipeIngredient.Unit)
	require.Equal(t, recipeIngredient.Amount, deletedRecipeIngredient.Amount)
	require.Equal(t, recipeIngredient.RecipeID, deletedRecipeIngredient.RecipeID)
	require.WithinDuration(t, recipeIngredient.CreatedAt, deletedRecipeIngredient.CreatedAt, time.Second)
}

func TestListRecipeIngredients(t *testing.T) {
	recipe := createRandomRecipe(t)

	for i := 0; i < 10; i++ {
		_ = createRandomRecipeIngredient(t, recipe.ID)
	}

	recipeIngredients, err := testQueries.ListRecipeIngredientsByRecipeId(context.Background(), recipe.ID)
	require.NoError(t, err)
	require.NotEmpty(t, recipeIngredients)
	require.Equal(t, len(recipeIngredients), 10)
}

func TestUpdateRecipeIngredientById(t *testing.T) {
	recipeIngredient := createRandomRecipeIngredient(t)

	arg := UpdateRecipeIngredientByIdParams{
		ID:             recipeIngredient.ID,
		Rank:           int32(util.RandomInt(1, 10)),
		IngredientName: util.RandomString(6),
		Unit:           util.RandomString(4),
		Amount:         float64(util.RandomInt(0, 1000)),
	}

	updatedRecipeIngredient, err := testQueries.UpdateRecipeIngredientById(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRecipeIngredient)

	require.Equal(t, updatedRecipeIngredient.Rank, arg.Rank)
	require.Equal(t, updatedRecipeIngredient.IngredientName, arg.IngredientName)
	require.Equal(t, updatedRecipeIngredient.Unit, arg.Unit)
	require.Equal(t, updatedRecipeIngredient.Amount, arg.Amount)
}

func TestUpdateRecipeIngredientByRecipeId(t *testing.T) {
	recipeIngredient := createRandomRecipeIngredient(t)

	arg := UpdateRecipeIngredientByRecipeIdParams{
		RecipeID:       recipeIngredient.RecipeID,
		Rank:           int32(util.RandomInt(1, 10)),
		IngredientName: util.RandomString(6),
		Unit:           util.RandomString(4),
		Amount:         float64(util.RandomInt(0, 1000)),
	}

	updatedRecipeIngredient, err := testQueries.UpdateRecipeIngredientByRecipeId(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRecipeIngredient)

	require.Equal(t, updatedRecipeIngredient.Rank, arg.Rank)
	require.Equal(t, updatedRecipeIngredient.IngredientName, arg.IngredientName)
	require.Equal(t, updatedRecipeIngredient.Unit, arg.Unit)
	require.Equal(t, updatedRecipeIngredient.Amount, arg.Amount)
}

func TestDeleteRecipeIngredientsByRecipeId(t *testing.T) {
	recipeId := util.RandomInt(1, 100)

	var recipeIngredients []RecipeIngredient

	for i := 0; i < 10; i++ {
		recipeIngredient := createRandomRecipeIngredient(t, recipeId)
		recipeIngredients = append(recipeIngredients, recipeIngredient)
	}

	deletedRecipeIngredients, err := testQueries.DeleteRecipeIngredientsByRecipeId(context.Background(), recipeId)
	require.NoError(t, err)
	require.NotEmpty(t, deletedRecipeIngredients)

	for i, deletedRecipeIngredient := range deletedRecipeIngredients {
		require.Equal(t, recipeIngredients[i].ID, deletedRecipeIngredient.ID)
		require.Equal(t, recipeIngredients[i].IngredientName, deletedRecipeIngredient.IngredientName)
		require.Equal(t, recipeIngredients[i].Rank, deletedRecipeIngredient.Rank)
		require.Equal(t, recipeIngredients[i].Unit, deletedRecipeIngredient.Unit)
		require.Equal(t, recipeIngredients[i].Amount, deletedRecipeIngredient.Amount)
		require.WithinDuration(t, recipeIngredients[i].CreatedAt, deletedRecipeIngredient.CreatedAt, time.Second)
	}
}
