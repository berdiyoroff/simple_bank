package db

import (
	"context"
	"testing"
	"time"

	"github.com/berdiyoroff/simple_bank/pkg/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func createEntryTest(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomInt(-100, 100),
	}
	entry, err := testStore.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createAccountTest(t)
	createEntryTest(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createAccountTest(t)
	entry1 := createEntryTest(t, account)
	entry2, err := testStore.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)
}

func TestAccountEntries(t *testing.T) {
	account1 := createAccountTest(t)
	for i := 0; i < 10; i++ {
		createEntryTest(t, account1)
	}
	arg := AccountEntriesParams{
		AccountID: account1.ID,
		Limit:     5,
		Offset:    0,
	}

	entries, err := testStore.AccountEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotZero(t, entry.ID)
		require.NotZero(t, entry.CreatedAt)
		require.NotZero(t, entry.Amount)
		require.Equal(t, entry.AccountID, account1.ID)
	}
}

func TestUpdateEntry(t *testing.T) {
	account := createAccountTest(t)
	entry1 := createEntryTest(t, account)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomInt(-100, 100),
	}
	entry2, err := testStore.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry2.Amount, arg.Amount)
	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	account := createAccountTest(t)
	entry1 := createEntryTest(t, account)
	err := testStore.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	entry2, err := testStore.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.Equal(t, err, pgx.ErrNoRows)
	require.Empty(t, entry2)
}
