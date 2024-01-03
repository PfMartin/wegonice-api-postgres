// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: recipe_step.sql

package db

import (
	"context"
)

const createRecipeStep = `-- name: CreateRecipeStep :one
INSERT INTO recipe_steps (
  rank,
  step_description,
  recipe_id
) VALUES (
  $1, $2, $3
) RETURNING id, rank, step_description, recipe_id, created_at
`

type CreateRecipeStepParams struct {
	Rank            int32  `json:"rank"`
	StepDescription string `json:"step_description"`
	RecipeID        int64  `json:"recipe_id"`
}

func (q *Queries) CreateRecipeStep(ctx context.Context, arg CreateRecipeStepParams) (RecipeStep, error) {
	row := q.db.QueryRowContext(ctx, createRecipeStep, arg.Rank, arg.StepDescription, arg.RecipeID)
	var i RecipeStep
	err := row.Scan(
		&i.ID,
		&i.Rank,
		&i.StepDescription,
		&i.RecipeID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRecipeStepById = `-- name: DeleteRecipeStepById :one
DELETE FROM recipe_steps
WHERE id = $1
RETURNING id, rank, step_description, recipe_id, created_at
`

func (q *Queries) DeleteRecipeStepById(ctx context.Context, id int64) (RecipeStep, error) {
	row := q.db.QueryRowContext(ctx, deleteRecipeStepById, id)
	var i RecipeStep
	err := row.Scan(
		&i.ID,
		&i.Rank,
		&i.StepDescription,
		&i.RecipeID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRecipeStepsByRecipeId = `-- name: DeleteRecipeStepsByRecipeId :many
DELETE FROM recipe_steps
WHERE recipe_id = $1
RETURNING id, rank, step_description, recipe_id, created_at
`

func (q *Queries) DeleteRecipeStepsByRecipeId(ctx context.Context, recipeID int64) ([]RecipeStep, error) {
	rows, err := q.db.QueryContext(ctx, deleteRecipeStepsByRecipeId, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RecipeStep{}
	for rows.Next() {
		var i RecipeStep
		if err := rows.Scan(
			&i.ID,
			&i.Rank,
			&i.StepDescription,
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

const listRecipeStepsByRecipeId = `-- name: ListRecipeStepsByRecipeId :many
SELECT id, rank, step_description, recipe_id, created_at FROM recipe_steps
WHERE recipe_id = $1
ORDER BY rank
`

func (q *Queries) ListRecipeStepsByRecipeId(ctx context.Context, recipeID int64) ([]RecipeStep, error) {
	rows, err := q.db.QueryContext(ctx, listRecipeStepsByRecipeId, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RecipeStep{}
	for rows.Next() {
		var i RecipeStep
		if err := rows.Scan(
			&i.ID,
			&i.Rank,
			&i.StepDescription,
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

const updateRecipeStepById = `-- name: UpdateRecipeStepById :one
UPDATE recipe_steps
SET 
  rank = $2,
  step_description = $3
WHERE 
  id = $1
RETURNING id, rank, step_description, recipe_id, created_at
`

type UpdateRecipeStepByIdParams struct {
	ID              int64  `json:"id"`
	Rank            int32  `json:"rank"`
	StepDescription string `json:"step_description"`
}

func (q *Queries) UpdateRecipeStepById(ctx context.Context, arg UpdateRecipeStepByIdParams) (RecipeStep, error) {
	row := q.db.QueryRowContext(ctx, updateRecipeStepById, arg.ID, arg.Rank, arg.StepDescription)
	var i RecipeStep
	err := row.Scan(
		&i.ID,
		&i.Rank,
		&i.StepDescription,
		&i.RecipeID,
		&i.CreatedAt,
	)
	return i, err
}

const updateRecipeStepByRecipeId = `-- name: UpdateRecipeStepByRecipeId :one
UPDATE recipe_steps
SET 
  rank = $2,
  step_description = $3
WHERE 
  recipe_id = $1
RETURNING id, rank, step_description, recipe_id, created_at
`

type UpdateRecipeStepByRecipeIdParams struct {
	RecipeID        int64  `json:"recipe_id"`
	Rank            int32  `json:"rank"`
	StepDescription string `json:"step_description"`
}

func (q *Queries) UpdateRecipeStepByRecipeId(ctx context.Context, arg UpdateRecipeStepByRecipeIdParams) (RecipeStep, error) {
	row := q.db.QueryRowContext(ctx, updateRecipeStepByRecipeId, arg.RecipeID, arg.Rank, arg.StepDescription)
	var i RecipeStep
	err := row.Scan(
		&i.ID,
		&i.Rank,
		&i.StepDescription,
		&i.RecipeID,
		&i.CreatedAt,
	)
	return i, err
}
