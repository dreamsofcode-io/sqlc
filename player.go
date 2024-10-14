package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/dreamsofcode-io/sqlc/internal/repository"
)

func SellItem(
	ctx context.Context,
	conn *pgx.Conn,
	playerID int64,
	itemID uuid.UUID,
) error {
	tx, _ := conn.Begin(ctx)
	defer tx.Rollback(ctx)

	repo := repository.New(tx)

	item, err := repo.GetInventoryItem(ctx, repository.GetInventoryItemParams{
		PlayerID: playerID,
		ItemID:   itemID,
	})
	if err != nil {
		return err
	}

	count, _ := repo.RemoveInventoryItem(ctx, repository.RemoveInventoryItemParams{
		PlayerID: playerID,
		ItemID:   itemID,
	})
	if count == 0 {
		return fmt.Errorf("no item to remove")
	}

	player, err := repo.IncrPlayerGold(ctx, repository.IncrPlayerGoldParams{
		PlayerID: playerID,
		Amount:   item.Value,
	})
	if err != nil {
		fmt.Println("IncrPlayerGold query failed!")
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	fmt.Println("Item sold! Player Gold:", player.Gold.Int64)

	return nil
}
