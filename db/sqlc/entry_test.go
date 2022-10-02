package db

import (
	"context"
	"testing"

	"github.com/smelton01/bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	want := createRandomEntry(t)

	got, err := testQueries.GetEntry(context.Background(), want.ID)
	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, want, got)
}

func TestListEntries(t *testing.T) {
	acc := createRandomEntry(t)
	arg := ListEntriesParams{
		AccountID: acc.AccountID,
		Limit:     2,
		Offset:    0,
	}
	got, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, got, 1)

	for _, entry := range got {
		require.NotEmpty(t, entry)
	}
}

func createRandomEntry(t *testing.T) Entry {
	acc := createRandomAcc(t)
	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomInt(0, 1000),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}
