package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/PfMartin/wegonice-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomAuthor(t *testing.T) Author {
	createdUser := createRandomUser(t)

	arg := CreateAuthorParams{
		AuthorName:  util.RandomString(6),
		UserCreated: createdUser.Email,
	}

	author, err := testQueries.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, arg.AuthorName, author.AuthorName)
	require.Equal(t, arg.UserCreated, author.UserCreated)
	require.Equal(t, author.Instagram, "")
	require.Equal(t, author.Youtube, "")
	require.Equal(t, author.Website, "")

	require.NotZero(t, author.CreatedAt)

	return author
}

func TestCreateAuthor(t *testing.T) {
	createRandomAuthor(t)
}

func TestDeleteAuthorById(t *testing.T) {
	author := createRandomAuthor(t)

	deletedAuthor, err := testQueries.DeleteAuthorById(context.Background(), author.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedAuthor)

	require.Equal(t, author.AuthorName, deletedAuthor.AuthorName)
	require.Equal(t, author.Instagram, deletedAuthor.Instagram)
	require.Equal(t, author.Youtube, deletedAuthor.Youtube)
	require.Equal(t, author.Website, deletedAuthor.Website)
	require.WithinDuration(t, author.CreatedAt, deletedAuthor.CreatedAt, time.Second)

	gotAccount, err := testQueries.GetAuthor(context.Background(), author.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, gotAccount)
}

func TestGetAuthor(t *testing.T) {
	author := createRandomAuthor(t)

	gotAuthor, err := testQueries.GetAuthor(context.Background(), author.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotAuthor)

	require.Equal(t, author.AuthorName, gotAuthor.AuthorName)
	require.Equal(t, author.Instagram, gotAuthor.Instagram)
	require.Equal(t, author.Youtube, gotAuthor.Youtube)
	require.Equal(t, author.Website, gotAuthor.Website)
	require.WithinDuration(t, author.CreatedAt, gotAuthor.CreatedAt, time.Second)
}

func TestListAuthors(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = createRandomAuthor(t)
	}

	arg := ListAuthorsParams{
		Limit:  5,
		Offset: 0,
	}

	authors, err := testQueries.ListAuthors(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, authors)
	require.Equal(t, len(authors), 5)

}

func TestUpdateAuthorById(t *testing.T) {
	author := createRandomAuthor(t)

	arg := UpdateAuthorByIdParams{
		ID:         author.ID,
		AuthorName: util.RandomString(6),
		Website:    util.RandomString(10),
		Youtube:    util.RandomString(10),
		Instagram:  util.RandomString(10),
	}

	updatedAuthor, err := testQueries.UpdateAuthorById(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAuthor)

	require.Equal(t, updatedAuthor.AuthorName, arg.AuthorName)
	require.Equal(t, updatedAuthor.Instagram, arg.Instagram)
	require.Equal(t, updatedAuthor.Youtube, arg.Youtube)
	require.Equal(t, updatedAuthor.Website, arg.Website)
}
