-- name: CreateUser :one
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING id;



-- name: GetUserById :one
SELECT id, name, dob FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, dob FROM users;

-- name: UpdateUser :execresult
UPDATE users
SET name = $1, dob = $2
WHERE id = $3;

-- name: DeleteUser :execresult
DELETE FROM users WHERE id = $1;


-- name: ListAllUsersWithPagination :many
SELECT id, name, dob 
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT COUNT(*) AS count FROM users;