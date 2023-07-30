package query

type IBackpackQuery interface {
	GetItems(backpackId string) ([]GetItemResponse, error)
}
