package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/smelton01/bank/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAcc(t)
}

func TestGetAccount(t *testing.T) {
	want := createRandomAcc(t)
	got, err := testQueries.GetAccount(context.Background(), want.ID)
	require.NoError(t, err)
	require.NotEmpty(t, got)

	assert.Equal(t, want, got)
}

func TestDeleteAcc(t *testing.T) {
	acc := createRandomAcc(t)
	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc2)

}

func TestListAccounts(t *testing.T) {
	arg := ListAccountsParams{
		Limit:  2,
		Offset: 1,
	}
	_ = []Account{createRandomAcc(t), createRandomAcc(t), createRandomAcc(t)}
	got, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, got, 2)
	for _, account := range got {
		require.NotEmpty(t, account)
	}

}

func TestUpdateAcc(t *testing.T) {
	acc1 := createRandomAcc(t)
	arg := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: util.RandomCash(),
	}
	acc2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.NotEqual(t, acc1.Balance, acc2.Balance)
	assert.Equal(t, arg.Balance, acc2.Balance)
	assert.True(t, cmp.Equal(acc1, acc2, cmpopts.IgnoreFields(Account{}, "Balance")))

}

func createRandomAcc(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomCash(),
		Currency: util.RandomCurrency(),
	}
	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}
