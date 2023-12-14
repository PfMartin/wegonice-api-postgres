-- name: CreateAuthor :one
INSERT INTO authors (
  author_name,
  website,
  instagram,
  youtube,
  user_created
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY author_name
LIMIT $1
OFFSET $2;

-- name: UpdateAuthorById :one
UPDATE authors
SET 
  author_name = $2,
  website = $3,
  instagram = $4,
  youtube = $5
WHERE 
  id = $1
RETURNING *;

-- name: DeleteAuthorById :one
DELETE FROM authors
WHERE id = $1
RETURNING *;