-- name: CreateRecipe :one
INSERT INTO recipes (
  recipe_name,
  link,
  author_id,
  prep_time,
  prep_time_unit,
  user_created
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetRecipe :one
SELECT * FROM recipes
WHERE id = $1 LIMIT 1;

-- name: ListRecipes :many
SELECT * FROM recipes
ORDER BY recipe_name
LIMIT $1
OFFSET $2;

-- name: UpdateRecipeById :one
UPDATE recipes
SET 
  recipe_name = $2,
  link = $3,
  author_id = $4,
  prep_time = $5,
  prep_time_unit = $6
WHERE 
  id = $1
RETURNING *;

-- name: DeleteRecipeById :one
DELETE FROM recipes
WHERE id = $1
RETURNING *;