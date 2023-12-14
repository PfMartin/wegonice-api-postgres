package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/PfMartin/wegonice-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomRecipe(t *testing.T) Recipe {
	createdUser := createRandomUser(t)
	createdAuthor := createRandomAuthor(t)

	arg := CreateRecipeParams{
		RecipeName:   util.RandomString(6),
		Link:         util.RandomString(10),
		AuthorID:     createdAuthor.ID,
		PrepTime:     float64(10),
		PrepTimeUnit: "Minutes",
		UserCreated:  createdUser.Email,
	}

	recipe, err := testQueries.CreateRecipe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipe)

	require.Equal(t, arg.RecipeName, recipe.RecipeName)
	require.Equal(t, arg.Link, recipe.Link)
	require.Equal(t, arg.AuthorID, recipe.AuthorID)
	require.Equal(t, recipe.PrepTime, float64(10))
	require.Equal(t, recipe.PrepTimeUnit, "Minutes")
	require.Equal(t, arg.UserCreated, recipe.UserCreated)

	require.NotZero(t, recipe.CreatedAt)

	return recipe
}

func TestCreateRecipe(t *testing.T) {
	createRandomRecipe(t)
}

func TestDeleteRecipeById(t *testing.T) {
	recipe := createRandomRecipe(t)

	deletedRecipe, err := testQueries.DeleteRecipeById(context.Background(), recipe.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedRecipe)

	require.Equal(t, recipe.RecipeName, deletedRecipe.RecipeName)
	require.Equal(t, recipe.Link, deletedRecipe.Link)
	require.Equal(t, recipe.AuthorID, deletedRecipe.AuthorID)
	require.Equal(t, recipe.UserCreated, deletedRecipe.UserCreated)
	require.Equal(t, recipe.PrepTime, deletedRecipe.PrepTime)
	require.Equal(t, recipe.PrepTimeUnit, deletedRecipe.PrepTimeUnit)
	require.WithinDuration(t, recipe.CreatedAt, deletedRecipe.CreatedAt, time.Second)

	gotRecipe, err := testQueries.GetRecipe(context.Background(), recipe.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, gotRecipe)
}

func TestGetRecipe(t *testing.T) {
	recipe := createRandomRecipe(t)

	gotRecipe, err := testQueries.GetRecipe(context.Background(), recipe.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotRecipe)

	require.Equal(t, recipe.RecipeName, gotRecipe.RecipeName)
	require.Equal(t, recipe.Link, gotRecipe.Link)
	require.Equal(t, recipe.AuthorID, gotRecipe.AuthorID)
	require.Equal(t, recipe.UserCreated, gotRecipe.UserCreated)
	require.Equal(t, recipe.PrepTime, gotRecipe.PrepTime)
	require.Equal(t, recipe.PrepTimeUnit, gotRecipe.PrepTimeUnit)
	require.WithinDuration(t, recipe.CreatedAt, gotRecipe.CreatedAt, time.Second)
}

func TestListRecipes(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = createRandomRecipe(t)
	}

	arg := ListRecipesParams{
		Limit:  5,
		Offset: 0,
	}

	recipes, err := testQueries.ListRecipes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipes)
	require.Equal(t, len(recipes), 5)

}

func TestUpdateRecipeById(t *testing.T) {
	recipe := createRandomRecipe(t)

	newAuthor := createRandomAuthor(t)

	arg := UpdateRecipeByIdParams{
		ID:           recipe.ID,
		RecipeName:   util.RandomString(6),
		Link:         util.RandomString(10),
		AuthorID:     newAuthor.ID,
		PrepTime:     float64(1),
		PrepTimeUnit: "Hours",
	}

	updatedRecipe, err := testQueries.UpdateRecipeById(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRecipe)

	require.Equal(t, updatedRecipe.RecipeName, arg.RecipeName)
	require.Equal(t, updatedRecipe.Link, arg.Link)
	require.Equal(t, updatedRecipe.AuthorID, arg.AuthorID)
	require.Equal(t, updatedRecipe.PrepTime, arg.PrepTime)
	require.Equal(t, updatedRecipe.PrepTimeUnit, arg.PrepTimeUnit)
}
