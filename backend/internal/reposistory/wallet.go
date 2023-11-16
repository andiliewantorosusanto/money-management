package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Wallet struct {
	ID      int
	Name    string
	Balance float64
}

type WalletRepository struct {
	dbPool *pgxpool.Pool
}

func NewWalletRepository(dbPool *pgxpool.Pool) *WalletRepository {
	return &WalletRepository{
		dbPool: dbPool,
	}
}

func (r *WalletRepository) GetAllWallets(ctx context.Context) ([]Wallet, error) {
	rows, err := r.dbPool.Query(ctx, "SELECT * FROM wallet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID, &w.Name, &w.Balance)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}

	return wallets, nil
}
