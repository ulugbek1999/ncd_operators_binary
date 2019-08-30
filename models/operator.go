package models

type Operator struct {
	Id            int    `json:"id"`
	Phone         string `json:"phone"`
	Operator      int    `json:"operator"`
	Channel       string `json:"channel"`
	OperatorGroup `db:"op_group"`
}
