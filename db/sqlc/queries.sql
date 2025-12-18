-- name: CreateUser :execresult 
INSERT INTO users (name, dob)
VALUES (?, ?);


-- name: GetUserById :one
SELECT id, name, dob FROM users WHERE id = ?;

-- name: ListUsers :many
SELECT id, name, dob FROM users;

-- name: UpdateUser :execresult
UPDATE users
SET name = ?, dob = ?
WHERE id = ?;

-- name: DeleteUser :execresult
DELETE FROM users WHERE id = ?;
