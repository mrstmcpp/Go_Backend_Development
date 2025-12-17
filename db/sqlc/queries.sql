-- name: CreateUser :execresult
INSERT INTO users (name, dob)
VALUES (?, ?);

-- name: GetUserById :one
SELECT id, name, dob FROM users WHERE id = ?;

-- name: ListUsers :many
SELECT id, name, dob FROM users;

-- name: UpdateUser :exec
UPDATE users
SET name = ?, dob = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
