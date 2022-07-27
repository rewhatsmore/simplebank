// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
from_account_id, to_account_id, amount
) VALUES (
$1, $2, $3
)
RETURNING transfer_id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.queryRow(ctx, q.createTransferStmt, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.TransferID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT transfer_id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE transfer_id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, transferID int64) (Transfer, error) {
	row := q.queryRow(ctx, q.getTransferStmt, getTransfer, transferID)
	var i Transfer
	err := row.Scan(
		&i.TransferID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT transfer_id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE from_account_id = $1 OR to_account_id = $2
ORDER BY transfer_id
LIMIT $3
OFFSET $4
`

type ListTransfersParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error) {
	rows, err := q.query(ctx, q.listTransfersStmt, listTransfers,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.TransferID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
