package form

// AccountSearch represents search form fields for "/api/v1/accounts".
type AccountSearch struct {
	Query  string `form:"q"`
	Count  int    `form:"count" binding:"required"`
	Offset int    `form:"offset"`
	Order  string `form:"order"`
}

func (f *AccountSearch) GetQuery() string {
	return f.Query
}

func (f *AccountSearch) SetQuery(q string) {
	f.Query = q
}

func (f *AccountSearch) ParseQueryString() error {
	return ParseQueryString(f)
}

func NewAccountSearch(query string) AccountSearch {
	return AccountSearch{Query: query}
}
