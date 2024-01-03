-- name: CreateRecipeIngredient :one
INSERT INTO recipe_ingredients (
  ingredient_name,
  rank,
  unit,
  amount,
  recipe_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListRecipeIngredientsByRecipeId :many
SELECT * FROM recipe_ingredients
WHERE recipe_id = $1
ORDER BY rank;

-- name: UpdateRecipeIngredientById :one
UPDATE recipe_ingredients
SET 
  rank = $2,
  ingredient_name = $3,
  unit = $4,
  amount = $5
WHERE 
  id = $1
RETURNING *;

-- name: DeleteRecipeIngredientById :one
DELETE FROM recipe_ingredients
WHERE id = $1
RETURNING *;

-- name: DeleteRecipeIngredientsByRecipeId :many
DELETE FROM recipe_ingredients
WHERE recipe_id = $1
RETURNING *;