package entity

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestBackPackTestSuite(t *testing.T) {
	suite.Run(t, new(BackPackTestSuite))
}

type BackPackTestSuite struct {
	suite.Suite
}

func (s *BackPackTestSuite) SetupTest() {
	//suite.VariableThatShouldStartAtFive = 5
}

func (s *BackPackTestSuite) Test_WhenPutItem_ItemShouldBeRecorded() {
	b := NewBackpack("bid")
	id := "Id"
	actual, _ := b.PutItem(id, 1)

	s.Equal("Id", actual.ItemId)
	s.Equal(1, actual.Count)
}
func (s *BackPackTestSuite) Test_WhenPutItem_ShouldIncreaseCountByGivenQuantity() {
	b := NewBackpack("bid")

	itemId := "Id"
	_, _ = b.PutItem(itemId, 1)
	actual, _ := b.PutItem(itemId, 1)

	s.Equal("Id", actual.ItemId)
	s.Equal(2, actual.Count)
}
func (s *BackPackTestSuite) Test_WhenPutItem_ItemOver999_ShouldReturnErrorAnd999() {
	b := NewBackpack("bid")

	itemId := "Id"
	_, _ = b.PutItem(itemId, 999)
	actual, errOver := b.PutItem(itemId, 1)

	s.Equal("Id", actual.ItemId)
	s.Equal(999, actual.Count)
	s.Error(errOver, "Count is over 999")
}

func (s *BackPackTestSuite) Test_TakeOutItem_ItemShouldDecreaseByQtn() {
	// arrange
	b := NewBackpack("bid")

	itemId := "Id"
	_, _ = b.PutItem(itemId, 999)

	// act
	_, err := b.TakeOut(itemId, 1)
	actual := b.GetSlotByItemId(itemId)

	expected := Slot{
		ItemId: "Id",
		Count:  998,
	}

	// assert
	s.Equal(expected.ItemId, actual.ItemId)
	s.Equal(expected.Count, actual.Count)
	s.Nil(err)
}
