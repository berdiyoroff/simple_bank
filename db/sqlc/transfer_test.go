package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/berdiyoroff/simple_bank/pkg/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func createTransferTest(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomInt(1, 100),
	}

	transfer, err := testStore.q.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createAccountTest(t)
	account2 := createAccountTest(t)
	createTransferTest(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createAccountTest(t)
	account2 := createAccountTest(t)
	transfer1 := createTransferTest(t, account1, account2)

	transfer2, err := testStore.q.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createAccountTest(t)
	account2 := createAccountTest(t)
	for range 5 {
		createTransferTest(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        0,
	}

	transfers, err := testStore.q.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
		require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	}
}

func TestUpdateTransfer(t *testing.T) {
	account1 := createAccountTest(t)
	account2 := createAccountTest(t)
	transfer1 := createTransferTest(t, account1, account2)

	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomInt(1, 100),
	}

	transfer2, err := testStore.q.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	account1 := createAccountTest(t)
	account2 := createAccountTest(t)
	transfer1 := createTransferTest(t, account1, account2)

	err := testStore.q.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	transfer2, err := testStore.q.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.Equal(t, err, pgx.ErrNoRows)
	require.Empty(t, transfer2)
}
