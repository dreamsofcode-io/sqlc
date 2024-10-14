package main

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close(ctx)

	playerID := int64(1)
	itemID := uuid.MustParse("74feda80-8c1a-428a-af1e-52a2a721284f")

	SellItem(ctx, conn, playerID, itemID)
}
