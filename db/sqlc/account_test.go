package db

import (
	"context"
	"testing"
	"time"

	"github.com/berdiyoroff/simple_bank/pkg/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func createAccountTest(t *testing.T, owner string) Account {
	arg := CreateAccountParams{
		Owner:    owner,
		Balance:  util.RandomInt(0, 1000),
		Currency: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	user := createUserTest(t)
	createAccountTest(t, user.Username)
}

func TestGetAccount(t *testing.T) {
	user := createUserTest(t)
	account1 := createAccountTest(t, user.Username)
	account2, err := testStore.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	user := createUserTest(t)
	account1 := createAccountTest(t, user.Username)

	params := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomInt(0, 1000),
	}

	account2, err := testStore.UpdateAccount(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.Balance, params.Balance)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	user := createUserTest(t)
	account1 := createAccountTest(t, user.Username)

	err := testStore.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err := testStore.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.Equal(t, err, pgx.ErrNoRows)
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	user := createUserTest(t)

	createAccountTest(t, user.Username)

	arg := ListAccountsParams{
		Limit:  1,
		Offset: 0,
	}

	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 1)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestAddAccountBalance(t *testing.T) {
	user := createUserTest(t)
	account1 := createAccountTest(t, user.Username)

	params := AddAccountBalanceParams{
		ID:     account1.ID,
		Amount: util.RandomInt(0, 100),
	}

	account2, err := testStore.AddAccountBalance(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.Balance, account1.Balance+params.Amount)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}
