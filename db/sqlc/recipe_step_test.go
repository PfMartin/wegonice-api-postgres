package db

import (
	"context"
	"testing"
	"time"

	"github.com/PfMartin/wegonice-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomRecipeStep(t *testing.T, params ...int64) RecipeStep {
	createdRecipe := createRandomRecipe(t)

	recipeId := createdRecipe.ID
	if len(params) > 0 {
		recipeId = params[0]
	}

	rank := int64(1)
	if len(params) > 1 {
		rank = params[1]
	}

	arg := CreateRecipeStepParams{
		StepDescription: util.RandomString(50),
		Rank:            int32(rank),
		RecipeID:        recipeId,
	}

	recipeStep, err := testQueries.CreateRecipeStep(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipeStep)

	require.Equal(t, arg.StepDescription, recipeStep.StepDescription)
	require.Equal(t, arg.Rank, recipeStep.Rank)
	require.Equal(t, arg.RecipeID, recipeStep.RecipeID)

	require.NotZero(t, recipeStep.CreatedAt)

	return recipeStep
}

func TestCreateRecipeStep(t *testing.T) {
	createRandomRecipeStep(t)
}

func TestDeleteRecipeStepById(t *testing.T) {
	recipeStep := createRandomRecipeStep(t)

	deletedRecipeStep, err := testQueries.DeleteRecipeStepById(context.Background(), recipeStep.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedRecipeStep)

	require.Equal(t, recipeStep.StepDescription, deletedRecipeStep.StepDescription)
	require.Equal(t, recipeStep.Rank, deletedRecipeStep.Rank)
	require.Equal(t, recipeStep.RecipeID, deletedRecipeStep.RecipeID)
	require.WithinDuration(t, recipeStep.CreatedAt, deletedRecipeStep.CreatedAt, time.Second)
}

func TestListRecipeSteps(t *testing.T) {
	recipe := createRandomRecipe(t)

	for i := 0; i < 10; i++ {
		_ = createRandomRecipeStep(t, recipe.ID, int64(i))
	}

	recipeSteps, err := testQueries.ListRecipeStepsByRecipeId(context.Background(), recipe.ID)
	require.NoError(t, err)
	require.NotEmpty(t, recipeSteps)
	require.Equal(t, len(recipeSteps), 10)
}

func TestUpdateRecipeStepById(t *testing.T) {
	recipeStep := createRandomRecipeStep(t)

	arg := UpdateRecipeStepByIdParams{
		ID:              recipeStep.ID,
		Rank:            int32(util.RandomInt(1, 10)),
		StepDescription: util.RandomString(50),
	}

	updatedRecipeStep, err := testQueries.UpdateRecipeStepById(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRecipeStep)

	require.Equal(t, updatedRecipeStep.Rank, arg.Rank)
	require.Equal(t, updatedRecipeStep.StepDescription, arg.StepDescription)
}

func TestDeleteRecipeStepsByRecipeId(t *testing.T) {
	recipeId := util.RandomInt(1, 100)

	var recipeSteps []RecipeStep

	for i := 0; i < 10; i++ {
		recipeStep := createRandomRecipeStep(t, recipeId)
		recipeSteps = append(recipeSteps, recipeStep)
	}

	deletedRecipeSteps, err := testQueries.DeleteRecipeStepsByRecipeId(context.Background(), recipeId)
	require.NoError(t, err)
	require.NotEmpty(t, deletedRecipeSteps)

	for i, deletedRecipeStep := range deletedRecipeSteps {
		require.Equal(t, recipeSteps[i].ID, deletedRecipeStep.ID)
		require.Equal(t, recipeSteps[i].StepDescription, deletedRecipeStep.StepDescription)
		require.Equal(t, recipeSteps[i].Rank, deletedRecipeStep.Rank)
		require.WithinDuration(t, recipeSteps[i].CreatedAt, deletedRecipeStep.CreatedAt, time.Second)
	}
}
