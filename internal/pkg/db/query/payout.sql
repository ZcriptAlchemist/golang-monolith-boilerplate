-- name: ListAllPayouts :many
SELECT * FROM payouts ORDER BY crn ASC;

-- name: ListAllPendingPayouts :many
SELECT created_at, crn, transaction_status from payouts where transaction_status = 'pending' ORDER BY created_at ASC;
