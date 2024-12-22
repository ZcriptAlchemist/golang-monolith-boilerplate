-- name: CreateStatusCheck :exec
INSERT INTO status_checks (
    account_number, transaction_status, created_at
) VALUES (
    $1, $2, NOW() 
);

-- name: ListStatusChecks :many
SELECT transaction_status, account_number FROM status_checks ORDER BY created_at ASC;
