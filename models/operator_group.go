package models

type OperatorGroup struct {
	Id        int    `json:"id"`
	GroupName string `json:"group_name" db:"group_name"`
	District  `db:"district"`
	//Operator1 *Operator `json:"operator_1"`
	//Operator2 *Operator `json:"operator_2"`
	//Operator3 *Operator `json:"operator_3"`
	//Operator4 *Operator `json:"operator_4"`
}
