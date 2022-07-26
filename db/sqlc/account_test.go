package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/rewhatsmore/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
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
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, account.Owner, lastAccount.Owner)
	}
}
