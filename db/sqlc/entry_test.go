package db

import (
	"context"
	"testing"

	"github.com/smelton01/bank/util"
	"github.com/stretchr/testify/suite"
)

type EntrySuite struct {
	MainSuite
}

func (s *EntrySuite) SetupSuite() {
	s.MainSuite.SetupSuite()
}

func (s *EntrySuite) TestCreateEntry() {
	s.createRandomEntry()
}

func (s *EntrySuite) TestGetEntry() {
	want := s.createRandomEntry()

	got, err := s.queries.GetEntry(context.Background(), want.ID)
	s.Require().NoError(err)
	s.Require().NotEmpty(got)

	s.Require().Equal(want, got)
}

func (s *EntrySuite) TestListEntries() {
	acc := s.createRandomEntry()
	arg := ListEntriesParams{
		AccountID: acc.AccountID,
		Limit:     2,
		Offset:    0,
	}
	got, err := s.queries.ListEntries(context.Background(), arg)
	s.Require().NoError(err)

	s.Require().Len(got, 1)

	for _, entry := range got {
		s.Require().NotEmpty(entry)
	}
}

func (s *EntrySuite) createRandomEntry() Entry {
	acc := s.createRandomAcc()
	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomInt(0, 1000),
	}
	entry, err := s.queries.CreateEntry(context.Background(), arg)
	s.Require().NoError(err)
	s.Require().NotEmpty(acc)

	s.Require().Equal(arg.AccountID, entry.AccountID)
	s.Require().Equal(arg.Amount, entry.Amount)

	s.Require().NotZero(entry.ID)
	s.Require().NotZero(entry.CreatedAt)

	return entry
}

func TestEntries(t *testing.T) {
	suite.Run(t, &EntrySuite{})
}
