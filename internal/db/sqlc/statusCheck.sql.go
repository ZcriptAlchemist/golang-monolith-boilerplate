// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: statusCheck.sql

package sqlc

import (
	"context"
)

const createStatusCheck = `-- name: CreateStatusCheck :exec
INSERT INTO status_checks (
    account_number, transaction_status, created_at
) VALUES (
    $1, $2, NOW() 
)
`

type CreateStatusCheckParams struct {
	AccountNumber     string
	TransactionStatus string
}

func (q *Queries) CreateStatusCheck(ctx context.Context, arg CreateStatusCheckParams) error {
	_, err := q.db.Exec(ctx, createStatusCheck, arg.AccountNumber, arg.TransactionStatus)
	return err
}

const listStatusChecks = `-- name: ListStatusChecks :many
SELECT transaction_status, account_number FROM status_checks ORDER BY created_at ASC
`

type ListStatusChecksRow struct {
	TransactionStatus string
	AccountNumber     string
}

func (q *Queries) ListStatusChecks(ctx context.Context) ([]ListStatusChecksRow, error) {
	rows, err := q.db.Query(ctx, listStatusChecks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListStatusChecksRow
	for rows.Next() {
		var i ListStatusChecksRow
		if err := rows.Scan(&i.TransactionStatus, &i.AccountNumber); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
