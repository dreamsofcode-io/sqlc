-- name: AddItem :exec
INSERT INTO inventory (player_id, item_id)
VALUES ($1, $2);

-- name: RemoveItem :exec
DELETE FROM inventory
WHERE player_id = $1
AND item_id = $2;

-- name: ItemsForPlayer :many
SELECT item.*
FROM inventory
JOIN item ON item.id = item_id
WHERE player_id = $1;
