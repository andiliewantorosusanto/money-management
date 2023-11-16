package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Transaction struct {
	ID           int
	WalletIDFrom sql.NullInt64
	WalletIDTo   sql.NullInt64
	Amount       float64
	CreatedAt    time.Time
	CategoryID   int
}

type TransactionRepository struct {
	dbPool *pgxpool.Pool
}

func NewTransactionRepository(dbPool *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{
		dbPool: dbPool,
	}
}
func (r *TransactionRepository) GetAllTransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := r.dbPool.Query(ctx, "SELECT * FROM transaction")
	if err != nil {
		log.Fatal("Error selecting from transaction {}", err)
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		err := rows.Scan(&t.ID, &t.WalletIDFrom, &t.WalletIDTo, &t.Amount, &t.CreatedAt, &t.CategoryID)
		if err != nil {
			log.Fatal("Error scanning result {}", err)
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
