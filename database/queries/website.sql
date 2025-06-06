-- name: ListWebsites :many
SELECT * FROM website;

-- name: CreateWebsite :one
INSERT INTO website (url, "createdAt")
VALUES ($1, $2)
RETURNING *;