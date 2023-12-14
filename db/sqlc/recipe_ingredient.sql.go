// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: recipe_ingredient.sql

package db

import (
	"context"
)

const createRecipeIngredient = `-- name: CreateRecipeIngredient :one
INSERT INTO recipe_ingredients (
  ingredient_name,
  rank,
  unit,
  amount,
  recipe_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, rank, ingredient_name, unit, amount, recipe_id, created_at
`

type CreateRecipeIngredientParams struct {
	IngredientName string  `json:"ingredient_name"`
	Rank           int32   `json:"rank"`
	Unit           string  `json:"unit"`
	Amount         float64 `json:"amount"`
	RecipeID       int64   `json:"recipe_id"`
}

func (q *Queries) CreateRecipeIngredient(ctx context.Context, arg CreateRecipeIngredientParams) (RecipeIngredient, error) {
	row := q.db.QueryRowContext(ctx, createRecipeIngredient,
		arg.IngredientName,
		arg.Rank,
		arg.Unit,
		arg.Amount,
		arg.RecipeID,
	)
	var i RecipeIngredient
	err := row.Scan(
		&i.ID,
		&i.Rank,
		&i.IngredientName,
		&i.Unit,
		&i.Amount,
		&i.RecipeID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRecipeIngredientById = `-- name: DeleteRecipeIngredientById :one
DELETE FROM recipe_ingredients
WHERE id = $1
RETURNING id, rank, ingredient_name, unit, amount, recipe_id, created_at
`

func (q *Queries) DeleteRecipeIngredientById(ctx context.Context, id int64) (RecipeIngredient, error) {
	row := q.db.QueryRowContext(ctx, deleteRecipeIngredientById, id)
	var i RecipeIngredient
	err := row.Scan(
		&i.ID,
		&i.Rank,
		&i.IngredientName,
		&i.Unit,
		&i.Amount,
		&i.RecipeID,
		&i.CreatedAt,
	)
	return i, err
}

const listRecipeIngredientsByRecipeId = `-- name: ListRecipeIngredientsByRecipeId :many
SELECT id, rank, ingredient_name, unit, amount, recipe_id, created_at FROM recipe_ingredients
WHERE recipe_id = $1
ORDER BY rank
`

func (q *Queries) ListRecipeIngredientsByRecipeId(ctx context.Context, recipeID int64) ([]RecipeIngredient, error) {
	rows, err := q.db.QueryContext(ctx, listRecipeIngredientsByRecipeId, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RecipeIngredient{}
	for rows.Next() {
		var i RecipeIngredient
		if err := rows.Scan(
			&i.ID,
			&i.Rank,
			&i.IngredientName,
			&i.Unit,
			&i.Amount,
			&i.RecipeID,
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

const updateRecipeIngredientById = `-- name: UpdateRecipeIngredientById :one
UPDATE recipe_ingredients
SET 
  rank = $2,
  ingredient_name = $3,
  unit = $4,
  amount = $5
WHERE 
  id = $1
RETURNING id, rank, ingredient_name, unit, amount, recipe_id, created_at
`

type UpdateRecipeIngredientByIdParams struct {
	ID             int64   `json:"id"`
	Rank           int32   `json:"rank"`
	IngredientName string  `json:"ingredient_name"`
	Unit           string  `json:"unit"`
	Amount         float64 `json:"amount"`
}

func (q *Queries) UpdateRecipeIngredientById(ctx context.Context, arg UpdateRecipeIngredientByIdParams) (RecipeIngredient, error) {
	row := q.db.QueryRowContext(ctx, updateRecipeIngredientById,
		arg.ID,
		arg.Rank,
		arg.IngredientName,
		arg.Unit,
		arg.Amount,
	)
	var i RecipeIngredient
	err := row.Scan(
		&i.ID,
		&i.Rank,
		&i.IngredientName,
		&i.Unit,
		&i.Amount,
		&i.RecipeID,
		&i.CreatedAt,
	)
	return i, err
}
