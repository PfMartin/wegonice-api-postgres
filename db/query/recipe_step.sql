-- name: CreateRecipeStep :one
INSERT INTO recipe_steps (
  rank,
  step_description,
  recipe_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ListRecipeStepsByRecipeId :many
SELECT * FROM recipe_steps
WHERE recipe_id = $1
ORDER BY rank;

-- name: UpdateRecipeStepById :one
UPDATE recipe_steps
SET 
  rank = $2,
  step_description = $3
WHERE 
  id = $1
RETURNING *;

-- name: DeleteRecipeStepById :one
DELETE FROM recipe_steps
WHERE id = $1
RETURNING *;

-- name: DeleteRecipeStepsByRecipeId :many
DELETE FROM recipe_steps
WHERE recipe_id = $1
RETURNING *;