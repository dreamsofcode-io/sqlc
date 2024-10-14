-- name: FindByID :one
SELECT * FROM item WHERE id = $1;

-- name: FindByName :one
SELECT * FROM item WHERE name = $1;

-- name: UpdateName :exec
UPDATE item SET name = $1 WHERE id = $2;
