package models

import (
	s "ncd_operators/pkg/settings"
)

// city
type City struct {
	Id        int        `json:"id"`
	NameRu    string     `json:"name_ru" db:"name_ru"`
	NameEn    string     `json:"name_en" db:"name_en"`
	Code      string     `json:"code"`
	Districts []District `json:"districts,omitempty"`
}

// district
type District struct {
	Id     int    `json:"id"`
	NameRu string `json:"name_ru" db:"name_ru"`
	NameEn string `json:"name_en" db:"name_en"`
	City   `db:"city"`
}

//ProfessionAvailable is a struct of available options
type ProfessionAvailable struct {
	Id     int    `json:"id"`
	NameRu string `json:"name_ru" db:"name_ru"`
	NameEn string `json:"name_en" db:"name_en"`
	Score  int    `json:score`
}

//ProfessionAvailableList is responsible for getting data from professions_available
func ProfessionAvailableList() []*ProfessionAvailable {
	var p []*ProfessionAvailable
	_ = s.UDB.Select(&p, `SELECT * FROM professions_available`)
	return p
}

// country
type Country struct {
	Id     int    `json:"id"`
	NameRu string `json:"name_ru" db:"name_ru"`
	NameEn string `json:"name_en" db:"name_en"`
	Score  int    `json:"score"`
}

func CountryList() []*Country {
	var c []*Country
	_ = s.UDB.Select(&c, `SELECT * FROM directory_country`)
	return c
}

// language
type DLanguage struct {
	Id           int    `json:"id"`
	NameRu       string `json:"name_ru" db:"name_ru"`
	NameEn       string `json:"name_en" db:"name_en"`
	Elementary   int    `json:"elementary"`
	Intermediate int    `json:"intermediate"`
	Advanced     int    `json:"advanced"`
}

//EducationType of employee
type EducationType struct {
	Id     int    `json:"id"`
	NameRu string `json:"name_ru" db:"name_ru"`
	NameEn string `json:"name_en" db:"name_en"`
	Score  int    `json:"score"`
}
