package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTX(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	fmt.Print(">> before:", account1.Balance, account2.Balance)

	//run some concurrent tests to ensure robustness
	n := 5
	amount := int64(100)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	//check the results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check entries on from and to accounts
		from_entry := result.FromEntry
		require.NotEmpty(t, from_entry)
		require.Equal(t, account1.ID, from_entry.AccountID)
		require.Equal(t, -amount, from_entry.Amount)
		require.NotZero(t, from_entry.ID)
		require.NotZero(t, from_entry.CreatedAt)

		_, err = store.GetEntry(context.Background(), from_entry.ID)
		require.NoError(t, err)

		to_entry := result.ToEntry
		require.NotEmpty(t, to_entry)
		require.Equal(t, account2.ID, to_entry.AccountID)
		require.Equal(t, amount, to_entry.Amount)
		require.NotZero(t, to_entry.ID)
		require.NotZero(t, to_entry.CreatedAt)

		_, err = store.GetEntry(context.Background(), to_entry.ID)
		require.NoError(t, err)

		//check accounts
		from_account := result.FromAccount
		require.NotEmpty(t, from_account)
		require.Equal(t, from_account.ID, account1.ID)

		to_account := result.ToAccount
		require.NotEmpty(t, to_account)
		require.Equal(t, to_account.ID, account2.ID)

		//check balance on the accounts
		fmt.Print(">> tx:", account1.Balance, account2.Balance)
		diff1 := account1.Balance - from_account.Balance
		diff2 := to_account.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	//check final updated balances
	updated_account1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.Equal(t, updated_account1.Balance, account1.Balance-int64(n)*amount)

	updated_account2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.Equal(t, updated_account2.Balance, account2.Balance+int64(n)*amount)

	fmt.Print(">> after:", account1.Balance, account2.Balance)
}

func TestTransferTXDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	fmt.Print(">> before:", account1.Balance, account2.Balance)

	//run some concurrent tests to ensure robustness
	n := 10
	amount := int64(100)

	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	//check the results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	//check final updated balances
	updated_account1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.Equal(t, updated_account1.Balance, account1.Balance)

	updated_account2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.Equal(t, updated_account2.Balance, account2.Balance)

	fmt.Print(">> after:", account1.Balance, account2.Balance)
}
