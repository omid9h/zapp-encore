package queryhelper

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ApplyRangeFilter applies a range filter on query if from/to values are set
func ApplyRangeFilter(query *gorm.DB, column string, from, to any) *gorm.DB {

	if !isZero(from) && !isZero(to) {
		// Both from and to filters given
		return query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", column), from, to)

	} else if !isZero(from) {
		// Only from filter given
		return query.Where(fmt.Sprintf("%s > ?", column), from)

	} else if !isZero(to) {
		// Only to filter given
		return query.Where(fmt.Sprintf("%s < ?", column), to)
	}

	return query
}

func isZero(val any) bool {
	switch v := val.(type) {
	case time.Time:
		return v.IsZero()
	case string:
		return v == ""
	case int, int8, int16, int32, int64:
		return v == 0
	case uint, uint8, uint16, uint32, uint64:
		return v == 0
	case float32, float64:
		return v == 0
	default:
		return false
	}
}
