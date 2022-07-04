package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/db/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandOwner(),
		Balance:  util.RandMoney(),
		Currency: util.RandCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.AccountID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.AccountID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.AccountID, account1.AccountID)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, account1.Balance)
	require.Equal(t, account2.Currency, account1.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {

	accountOld := createRandomAccount(t)

	arg := UpdateAccountsParams{
		AccountID: accountOld.AccountID,
		Balance:   util.RandMoney(),
	}

	accountNew, err := testQueries.UpdateAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accountNew)
	require.Equal(t, accountNew.AccountID, accountOld.AccountID)
	require.Equal(t, accountNew.Owner, accountOld.Owner)
	require.Equal(t, accountNew.Balance, arg.Balance)
	require.Equal(t, accountNew.Currency, accountOld.Currency)
	require.WithinDuration(t, accountNew.CreatedAt, accountOld.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.AccountID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.AccountID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, int(arg.Limit))

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
