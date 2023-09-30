package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

func NewSuccessRespone(data, paging, filter interface{}) *successRes {
	return &successRes{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}
func SimpleSuccessRespone(data interface{}) *successRes {
	return NewSuccessRespone(data, nil, nil)
}
