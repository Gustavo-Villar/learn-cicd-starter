-- name: CreateNote :exec
INSERT INTO notes (id, created_at, updated_at, note, user_id)
VALUES ($1, $2, $3, $4, $5);
--

-- name: GetNote :one
SELECT * FROM notes WHERE id = $1;
--

-- name: GetNotesForUser :many
SELECT * FROM notes WHERE user_id = $1;
--
