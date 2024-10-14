-- name: FindAllPlayers :many
SELECT * FROM player
ORDER BY level DESC;

-- name: InsertItem :one
INSERT INTO item (id, name, value)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING *;

-- name: FindItemByID :one
SELECT * FROM item
WHERE id = $1;

-- name: GetInventoryAndPlayer :many
SELECT player.*, item.*
FROM inventory
JOIN player ON player.id = player_id
JOIN item ON item.id = item_id;

-- name: GetInventoryItem :one
SELECT item.*
FROM inventory
JOIN item ON item.id = item_id
WHERE player_id = $1
AND item_id = $2;

-- name: RemoveInventoryItem :execrows
DELETE FROM inventory
WHERE player_id = $1
AND item_id = $2;

-- name: IncrPlayerGold :one
UPDATE player
SET gold = gold + sqlc.arg(amount)::int
WHERE id = sqlc.arg(player_id)::bigint
RETURNING *;
