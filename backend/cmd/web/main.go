package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	repository "github.com/andiliewantorosusanto/money-management/internal/reposistory"
	"github.com/andiliewantorosusanto/money-management/pkg/config"
	"github.com/andiliewantorosusanto/money-management/pkg/handlers"
	"github.com/jackc/pgx/v4/pgxpool"
)

const portNumber = ":8080"

func main() {

	connString := "user=postgres dbname=money_management sslmode=disable password=postgres"
	dbPool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Error Connecting to Database {}", err)
	}
	defer dbPool.Close()

	repo := repository.NewTransactionRepository(dbPool)
	transactions, err := repo.GetAllTransactions(context.Background())
	if err != nil {
		log.Fatal("Error Getting All Transaction {}", err)
	}

	// Print or do something with the transactions
	for _, t := range transactions {
		fmt.Printf("ID: %d, WalletIDFrom: %d, WalletIDTo: %d, Amount: %.2f, CreatedAt: %s, CategoryID: %d\n",
			t.ID, t.WalletIDFrom, t.WalletIDTo, t.Amount, t.CreatedAt.Format("2006-01-02 15:04:05"), t.CategoryID)
	}

	var app config.AppConfig
	repos := handlers.NewRepository(&app)
	handlers.NewHandlers(repos)

	handler := routes(&app)

	s := http.Server{
		Addr:    portNumber,
		Handler: handler,
	}

	log.Fatal(s.ListenAndServe())
}
