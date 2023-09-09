package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	_ "github.com/yuqihu1103/SimpleBank/util"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.randomOwner(),
		Balance:  util.randomMoney(),
		Currency: util.randomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
