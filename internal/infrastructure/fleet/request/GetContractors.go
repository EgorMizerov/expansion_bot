package request

type GetContractorsRequest struct {
	Limit int
}

func (self *GetContractorsRequest) ToBody() GetContractorsRequestBody {
	return GetContractorsRequestBody{Limit: self.Limit}
}

type GetContractorsRequestBody struct {
	Query struct{} `json:"query"`
	Limit int      `json:"limit"`
}
