package query

type IQuery struct {
	Page     uint16 `json:"page"`
	PageSize uint16 `json:"pageSize"`
	Name     string `json:"name"`
}
