package query

type GetItemListQuery struct {
	repo IBackpackQuery
}

func NewGetItemListQuery(repo IBackpackQuery) *GetItemListQuery {
	return &GetItemListQuery{repo: repo}
}

func (g GetItemListQuery) Query(backpackId string) ([]GetItemResponse, error) {
	return g.repo.GetItems(backpackId)
}

type GetItemResponse struct {
	Id    string
	Name  string
	Count int
}
