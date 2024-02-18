// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: recipe.sql

package db

import (
	"context"
)

const createRecipe = `-- name: CreateRecipe :one
INSERT INTO recipes (
  recipe_name,
  link,
  author_id,
  prep_time,
  prep_time_unit,
  user_created
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, recipe_name, link, author_id, prep_time, prep_time_unit, user_created, created_at
`

type CreateRecipeParams struct {
	RecipeName   string  `json:"recipe_name"`
	Link         string  `json:"link"`
	AuthorID     int64   `json:"author_id"`
	PrepTime     float64 `json:"prep_time"`
	PrepTimeUnit string  `json:"prep_time_unit"`
	UserCreated  string  `json:"user_created"`
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, createRecipe,
		arg.RecipeName,
		arg.Link,
		arg.AuthorID,
		arg.PrepTime,
		arg.PrepTimeUnit,
		arg.UserCreated,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.RecipeName,
		&i.Link,
		&i.AuthorID,
		&i.PrepTime,
		&i.PrepTimeUnit,
		&i.UserCreated,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRecipeById = `-- name: DeleteRecipeById :one
DELETE FROM recipes
WHERE id = $1
RETURNING id, recipe_name, link, author_id, prep_time, prep_time_unit, user_created, created_at
`

func (q *Queries) DeleteRecipeById(ctx context.Context, id int64) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, deleteRecipeById, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.RecipeName,
		&i.Link,
		&i.AuthorID,
		&i.PrepTime,
		&i.PrepTimeUnit,
		&i.UserCreated,
		&i.CreatedAt,
	)
	return i, err
}

const getRecipe = `-- name: GetRecipe :one
SELECT id, recipe_name, link, author_id, prep_time, prep_time_unit, user_created, created_at FROM recipes
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRecipe(ctx context.Context, id int64) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, getRecipe, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.RecipeName,
		&i.Link,
		&i.AuthorID,
		&i.PrepTime,
		&i.PrepTimeUnit,
		&i.UserCreated,
		&i.CreatedAt,
	)
	return i, err
}

const listRecipes = `-- name: ListRecipes :many
SELECT id, recipe_name, link, author_id, prep_time, prep_time_unit, user_created, created_at FROM recipes
ORDER BY recipe_name
LIMIT $1
OFFSET $2
`

type ListRecipesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRecipes(ctx context.Context, arg ListRecipesParams) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, listRecipes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Recipe{}
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.RecipeName,
			&i.Link,
			&i.AuthorID,
			&i.PrepTime,
			&i.PrepTimeUnit,
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

const updateRecipeById = `-- name: UpdateRecipeById :one
UPDATE recipes
SET 
  recipe_name = $2,
  link = $3,
  author_id = $4,
  prep_time = $5,
  prep_time_unit = $6
WHERE 
  id = $1
RETURNING id, recipe_name, link, author_id, prep_time, prep_time_unit, user_created, created_at
`

type UpdateRecipeByIdParams struct {
	ID           int64   `json:"id"`
	RecipeName   string  `json:"recipe_name"`
	Link         string  `json:"link"`
	AuthorID     int64   `json:"author_id"`
	PrepTime     float64 `json:"prep_time"`
	PrepTimeUnit string  `json:"prep_time_unit"`
}

func (q *Queries) UpdateRecipeById(ctx context.Context, arg UpdateRecipeByIdParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, updateRecipeById,
		arg.ID,
		arg.RecipeName,
		arg.Link,
		arg.AuthorID,
		arg.PrepTime,
		arg.PrepTimeUnit,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.RecipeName,
		&i.Link,
		&i.AuthorID,
		&i.PrepTime,
		&i.PrepTimeUnit,
		&i.UserCreated,
		&i.CreatedAt,
	)
	return i, err
}
