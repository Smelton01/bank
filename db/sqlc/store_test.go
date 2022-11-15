package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	MainSuite
}

func (s *StoreSuite) SetupSuite() {
	s.MainSuite.SetupSuite()
}

func (s *StoreSuite) TestTransferTX() {
	store := NewStore(s.db)

	acc1 := s.createRandomAcc()
	acc2 := s.createRandomAcc()

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: acc1.ID,
				ToAccountID:   acc2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)
	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		s.Require().NoError(err)

		result := <-results
		s.Require().NotEmpty(result)

		// check transfer
		transfer := result.Transfer
		s.Require().NotEmpty(transfer)
		s.Require().Equal(acc1.ID, transfer.FromAccountID)
		s.Require().Equal(acc2.ID, transfer.ToAccountID)
		s.Require().Equal(amount, transfer.Amount)
		s.Require().NotZero(transfer.ID)
		s.Require().NotZero(transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		s.Require().NoError(err)

		// check entries
		fromEntry := result.FromEntry
		s.Require().NotEmpty(fromEntry)
		s.Require().Equal(acc1.ID, fromEntry.AccountID)
		s.Require().Equal(-amount, fromEntry.Amount)
		s.Require().NotZero(fromEntry.ID)
		s.Require().NotZero(fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		s.Require().NoError(err)

		toEntry := result.ToEntry
		s.Require().NotEmpty(toEntry)
		s.Require().Equal(acc2.ID, toEntry.AccountID)
		s.Require().Equal(amount, toEntry.Amount)
		s.Require().NotZero(toEntry.ID)
		s.Require().NotZero(toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		s.Require().NoError(err)

		// check accounts
		fromAccount := result.FromAccount
		s.Require().NotEmpty(fromAccount)
		s.Require().Equal(acc1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		s.Require().NotEmpty(toAccount)
		s.Require().Equal(acc2.ID, toAccount.ID)

		// check balances
		diff1 := acc1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - acc2.Balance
		s.Require().Equal(diff1, diff2)
		s.Require().True(diff1 > 0)
		s.Require().True(diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount

		k := int(diff1 / amount)
		s.Require().True(k >= 1 && k <= n)
		s.Require().NotContains(existed, k)
		existed[k] = true

	}

	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), acc1.ID)
	s.Require().NoError(err)

	updatedAccount2, err := store.GetAccount(context.Background(), acc2.ID)
	s.Require().NoError(err)

	s.Require().Equal(acc1.Balance-int64(n)*amount, updatedAccount1.Balance)
	s.Require().Equal(acc2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestStore(t *testing.T) {
	suite.Run(t, &StoreSuite{})
}
