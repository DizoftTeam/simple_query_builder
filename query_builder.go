package sqb

import (
	"log"
	"strings"
)

// Query struct for QB
type Query struct {
	tableName string // Имя таблицы

	selectFields string // Выбираемые поля
	limit        int    // Лимит
	offset       int    // Отступ

	where []string // Выборка
}

// --------------- PUBLIC ---------------

// NewQuery creates a new query object
func NewQuery(tableName string) *Query {
	return &Query{
		tableName: tableName,
		where:     []string{},
	}
}

// ----------- STRUCT METHODS -----------

// Select adds a select expression
func (q *Query) Select(expr string) *Query {
	q.selectFields = expr

	return q
}

// Where adds a where
func (q *Query) Where(whereType string, condition string) *Query {
	whereType = strings.ToUpper(whereType)

	if whereType != "AND" && whereType != "OR" {
		log.Printf("Unknown WHERE type: %v\n", whereType)

		return q
	}

	q.where = append(q.where, whereType+" "+condition)

	return q
}

// Limit sets a limit
func (q *Query) Limit(limit int) *Query {
	q.limit = limit

	return q
}

// Offset sets a offset
func (q *Query) Offset(offset int) *Query {
	q.offset = offset

	return q
}

// Generate sql statement
func (q *Query) Generate() string {
	result := ""

	if q.selectFields != "" {
		result += "SELECT " + q.selectFields
	} else {
		result += "SELECT *"
	}

	result += " FROM " + q.tableName

	result += " WHERE 1 = 1"

	for _, cond := range q.where {
		result += " " + cond
	}

	return result
}

// --------------- PRIVATE ---------------
