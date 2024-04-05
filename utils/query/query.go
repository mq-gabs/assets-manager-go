package query

type IQuery struct {
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
	Name     string `json:"name"`
}
