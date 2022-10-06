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

type AccountSuite struct {
	MainSuite
}

func (a *AccountSuite) TestCreateAccount() {
	a.createRandomAcc(a.T())
}

func (a *AccountSuite) TestGetAccount(t *testing.T) {
	want := a.createRandomAcc(t)
	got, err := a.queries.GetAccount(context.Background(), want.ID)
	require.NoError(t, err)
	require.NotEmpty(t, got)

	assert.Equal(t, want, got)
}

func (a *AccountSuite) TestDeleteAcc() {
	acc := a.createRandomAcc(a.T())
	err := a.queries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(a.T(), err)

	acc2, err := a.queries.GetAccount(context.Background(), acc.ID)
	require.Error(a.T(), err)
	require.EqualError(a.T(), err, sql.ErrNoRows.Error())
	require.Empty(a.T(), acc2)

}

func (a *AccountSuite) TestListAccounts() {
	arg := ListAccountsParams{
		Limit:  2,
		Offset: 1,
	}
	_ = []Account{a.createRandomAcc(a.T()), a.createRandomAcc(a.T()), a.createRandomAcc(a.T())}
	got, err := a.queries.ListAccounts(context.Background(), arg)
	require.NoError(a.T(), err)

	require.Len(a.T(), got, 2)
	for _, account := range got {
		require.NotEmpty(a.T(), account)
	}

}

func (a *AccountSuite) TestUpdateAcc() {
	acc1 := a.createRandomAcc(a.T())
	arg := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: util.RandomCash(),
	}
	acc2, err := a.queries.UpdateAccount(context.Background(), arg)
	require.NoError(a.T(), err)
	require.NotEmpty(a.T(), acc2)

	require.NotEqual(a.T(), acc1.Balance, acc2.Balance)
	assert.Equal(a.T(), arg.Balance, acc2.Balance)
	assert.True(a.T(), cmp.Equal(acc1, acc2, cmpopts.IgnoreFields(Account{}, "Balance")))

}

func (a *AccountSuite) createRandomAcc(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomCash(),
		Currency: util.RandomCurrency(),
	}
	acc, err := a.queries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}
