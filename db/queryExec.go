package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type QueryExecutor struct {
	db *sql.DB
}

func NewQueryExecutor(db *sql.DB) *QueryExecutor {
	return &QueryExecutor{db: db}
}

func (qe *QueryExecutor) Query(ctx context.Context, query string, args ...any) ([]map[string]interface{}, error) {
	rows, err := qe.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuesPtrs := make([]interface{}, len(columns))

		for i := range values {
			valuesPtrs[i] = &values[i]
		}

		if err := rows.Scan(valuesPtrs...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			row[col] = values[i]
		}
		results = append(results, row)
	}

	return results, nil
}

func (qe *QueryExecutor) Exec(ctx context.Context, query string, args ...any) (int64, error) {
	result, err := qe.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
