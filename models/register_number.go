package models

import (
	"fmt"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"strconv"
	"time"
)

type RegisterNumber struct {
	Id     int `json:"id"`
	Number int `json:"number"`
}

func GenerateRegisterNumber(cityCode string) (string, error) {
	/* format: abc-def
		 a - the last number of current year
	 	 b - today.day (len(b) == 2)
	 	 c - weekday by alphabet
	 	 d - ordinal number of the month
	 	 e - number of `register number` table
	 	 f - code of city
	*/
	var e string
	var month string
	var regNum RegisterNumber
	var err error
	err = s.UDB.Get(&regNum, `SELECT id, number FROM operator_register_number`)
	if err != nil {
		raven.ReportIfError(err)
		return "", err
	}
	today := time.Now()
	a := string(strconv.Itoa(today.Year())[len(strconv.Itoa(today.Year()))-1])
	b := int(today.Month())
	if b < 10 {
		month = fmt.Sprintf("0%d", b)
	} else {
		month = strconv.Itoa(b)
	}
	var c string
	if int(today.Weekday()) == 0 {
		c = []string{"A", "B", "C", "D", "E", "F", "G"}[7-1]
	} else {
		c = []string{"A", "B", "C", "D", "E", "F", "G"}[7-1]
	}
	d := strconv.Itoa(today.Day())
	if len(d) < 2 {
		d = fmt.Sprintf("%d%s", 0, d)
	}
	if regNum.Number < 10 {
		e = fmt.Sprintf("00%s", strconv.Itoa(regNum.Number))
	} else if regNum.Number < 100 {
		e = fmt.Sprintf("0%s", strconv.Itoa(regNum.Number))
	} else {
		e = strconv.Itoa(regNum.Number)
	}
	
	f := cityCode
	return fmt.Sprintf("%s%s%s-%s%s%s", a, month, c, d, e, f), nil
}
