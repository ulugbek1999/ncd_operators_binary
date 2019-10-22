package utils

import (
	"database/sql"
	"strconv"
	"time"
	"github.com/lib/pq"
)

func NewNullInt(s string) sql.NullInt64 {
	if i, err := strconv.Atoi(s); err != nil {
		return sql.NullInt64{}
	} else {
		return sql.NullInt64 {
			Int64: int64(i),
			Valid: true,
		}
	}
}

func NewNullFloat(s string) sql.NullFloat64 {
	if i, err := strconv.ParseFloat(s, 64); err != nil {
		return sql.NullFloat64{}
	} else {
		return sql.NullFloat64 {
			Float64: i,
			Valid: true,
		}
	}
}

func NewNullTime(s string) pq.NullTime {
	if t, err := time.Parse("2006-01-02", s); err != nil {
		return pq.NullTime{}
	} else {
		return pq.NullTime {
			Time: t,
			Valid: true,
		}
	}
}