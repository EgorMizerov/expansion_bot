package request

type GetDriversRequestOrderBy string
type GetDriversRequestOrderDirection string

const (
	GetDriversRequestOrderByCreatedAt GetDriversRequestOrderBy = "driver_profile.created_date"

	OrderByAsc  GetDriversRequestOrderDirection = "asc"
	OrderByDesc GetDriversRequestOrderDirection = "desc"
)

type GetDriversRequest struct {
	OrderBy        GetDriversRequestOrderBy
	OrderDirection GetDriversRequestOrderDirection
	Limit          int
	Offset         int
}

func (self *GetDriversRequest) ToBody(parkID string) GetDriversRequestBody {
	return GetDriversRequestBody{
		Query: GetDriversRequestBodyQuery{GetDriversRequestBodyQueryPark{parkID}},
		SortOrder: []GetDriversRequestBodyQuerySortOrder{{
			Field:     string(self.OrderBy),
			Direction: string(self.OrderDirection),
		}},
		Limit:  self.Limit,
		Offset: self.Offset,
	}
}

type GetDriversRequestBody struct {
	Query     GetDriversRequestBodyQuery            `json:"query"`
	SortOrder []GetDriversRequestBodyQuerySortOrder `json:"sort_order"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetDriversRequestBodyQuery struct {
	Park GetDriversRequestBodyQueryPark `json:"park"`
}

type GetDriversRequestBodyQueryPark struct {
	Id string `json:"id"`
}

type GetDriversRequestBodyQuerySortOrder struct {
	Direction string `json:"direction"`
	Field     string `json:"field"`
}
