-- name: CreateAccount :one
INSERT INTO accounts (
owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE owner = $1
ORDER BY account_id
LIMIT $2
OFFSET $3;

-- name: UpdateAccounts :one
UPDATE accounts SET balance = $2
WHERE account_id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts SET balance = balance + sqlc.arg(amount)
WHERE account_id = sqlc.arg(account_id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE account_id = $1;