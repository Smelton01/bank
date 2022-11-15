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
	"github.com/stretchr/testify/suite"
)

type AccountSuite struct {
	MainSuite
}

func (s *AccountSuite) SetupSuite() {
	s.MainSuite.SetupSuite()
}

func (a *AccountSuite) TestCreateAccount() {
	a.createRandomAcc()
}

func (a *AccountSuite) TestGetAccount() {
	want := a.createRandomAcc()
	got, err := a.queries.GetAccount(context.Background(), want.ID)
	a.Require().NoError(err)
	a.Require().NotEmpty(got)

	a.Equal(want, got)
}

func (a *AccountSuite) TestDeleteAcc() {
	acc := a.createRandomAcc()
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
	_ = []Account{a.createRandomAcc(), a.createRandomAcc(), a.createRandomAcc()}
	got, err := a.queries.ListAccounts(context.Background(), arg)
	require.NoError(a.T(), err)

	require.Len(a.T(), got, 2)
	for _, account := range got {
		require.NotEmpty(a.T(), account)
	}

}

func (a *AccountSuite) TestUpdateAcc() {
	acc1 := a.createRandomAcc()
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

func (a *MainSuite) createRandomAcc() Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomCash(),
		Currency: util.RandomCurrency(),
	}
	acc, err := a.queries.CreateAccount(context.Background(), arg)
	a.Require().NoError(err)
	a.Require().NotEmpty(acc)

	a.Require().Equal(arg.Owner, acc.Owner)
	a.Require().Equal(arg.Balance, acc.Balance)
	a.Require().Equal(arg.Currency, acc.Currency)

	a.Require().NotZero(acc.ID)
	a.Require().NotZero(acc.CreatedAt)

	return acc
}

func TestAccounts(t *testing.T) {
	suite.Run(t, &AccountSuite{})
}
