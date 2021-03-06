// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: account.sql

package db

import (
	"context"
)

const addAccountBalance = `-- name: AddAccountBalance :one
UPDATE accounts SET balance = balance + $1
WHERE account_id = $2
RETURNING account_id, owner, balance, currency, created_at
`

type AddAccountBalanceParams struct {
	Amount    int64 `json:"amount"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error) {
	row := q.queryRow(ctx, q.addAccountBalanceStmt, addAccountBalance, arg.Amount, arg.AccountID)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING account_id, owner, balance, currency, created_at
`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.queryRow(ctx, q.createAccountStmt, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts WHERE account_id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, accountID int64) error {
	_, err := q.exec(ctx, q.deleteAccountStmt, deleteAccount, accountID)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT account_id, owner, balance, currency, created_at FROM accounts
WHERE account_id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, accountID int64) (Account, error) {
	row := q.queryRow(ctx, q.getAccountStmt, getAccount, accountID)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT account_id, owner, balance, currency, created_at FROM accounts
WHERE account_id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, accountID int64) (Account, error) {
	row := q.queryRow(ctx, q.getAccountForUpdateStmt, getAccountForUpdate, accountID)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT account_id, owner, balance, currency, created_at FROM accounts
WHERE owner = $1
ORDER BY account_id
LIMIT $2
OFFSET $3
`

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.query(ctx, q.listAccountsStmt, listAccounts, arg.Owner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.AccountID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
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

const updateAccounts = `-- name: UpdateAccounts :one
UPDATE accounts SET balance = $2
WHERE account_id = $1
RETURNING account_id, owner, balance, currency, created_at
`

type UpdateAccountsParams struct {
	AccountID int64 `json:"account_id"`
	Balance   int64 `json:"balance"`
}

func (q *Queries) UpdateAccounts(ctx context.Context, arg UpdateAccountsParams) (Account, error) {
	row := q.queryRow(ctx, q.updateAccountsStmt, updateAccounts, arg.AccountID, arg.Balance)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
