package backend_masterclass

import (
	"context"
	"database/sql"
	"fmt"
)

// Store 
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// store provides all functions to execute SQL queries and transactions
type SQLStore struct{
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db: db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error ) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}


// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

// TransferTxResult is the result of the transaction
type TransferTxResult struct {
	Transfer Transfers `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entries `json:"from_entry"`
	ToEntry Entries `json:"to_entry"`
}


//TransferTx performs a money transfer form one account to the other
// It creates the transfer, add account entires, and update accounts' balance within a database transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error){
	var result TransferTxResult



	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil{
			return err
		}

		
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil{
			return err
		}

		// get account => update it balance
		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		//
		
		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		 




		return nil
	})

	return result, err
}
