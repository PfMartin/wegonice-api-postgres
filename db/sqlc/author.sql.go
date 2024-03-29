// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: author.sql

package db

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
  author_name,
  website,
  instagram,
  youtube,
  user_created
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, author_name, website, instagram, youtube, user_created, created_at
`

type CreateAuthorParams struct {
	AuthorName  string `json:"author_name"`
	Website     string `json:"website"`
	Instagram   string `json:"instagram"`
	Youtube     string `json:"youtube"`
	UserCreated string `json:"user_created"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor,
		arg.AuthorName,
		arg.Website,
		arg.Instagram,
		arg.Youtube,
		arg.UserCreated,
	)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.AuthorName,
		&i.Website,
		&i.Instagram,
		&i.Youtube,
		&i.UserCreated,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAuthorById = `-- name: DeleteAuthorById :one
DELETE FROM authors
WHERE id = $1
RETURNING id, author_name, website, instagram, youtube, user_created, created_at
`

func (q *Queries) DeleteAuthorById(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, deleteAuthorById, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.AuthorName,
		&i.Website,
		&i.Instagram,
		&i.Youtube,
		&i.UserCreated,
		&i.CreatedAt,
	)
	return i, err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, author_name, website, instagram, youtube, user_created, created_at FROM authors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.AuthorName,
		&i.Website,
		&i.Instagram,
		&i.Youtube,
		&i.UserCreated,
		&i.CreatedAt,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, author_name, website, instagram, youtube, user_created, created_at FROM authors
ORDER BY author_name
LIMIT $1
OFFSET $2
`

type ListAuthorsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.AuthorName,
			&i.Website,
			&i.Instagram,
			&i.Youtube,
			&i.UserCreated,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthorById = `-- name: UpdateAuthorById :one
UPDATE authors
SET 
  author_name = $2,
  website = $3,
  instagram = $4,
  youtube = $5
WHERE 
  id = $1
RETURNING id, author_name, website, instagram, youtube, user_created, created_at
`

type UpdateAuthorByIdParams struct {
	ID         int64  `json:"id"`
	AuthorName string `json:"author_name"`
	Website    string `json:"website"`
	Instagram  string `json:"instagram"`
	Youtube    string `json:"youtube"`
}

func (q *Queries) UpdateAuthorById(ctx context.Context, arg UpdateAuthorByIdParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, updateAuthorById,
		arg.ID,
		arg.AuthorName,
		arg.Website,
		arg.Instagram,
		arg.Youtube,
	)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.AuthorName,
		&i.Website,
		&i.Instagram,
		&i.Youtube,
		&i.UserCreated,
		&i.CreatedAt,
	)
	return i, err
}
