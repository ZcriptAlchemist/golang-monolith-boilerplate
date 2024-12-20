-- name: ListAllMerchants :many
SELECT * FROM merchants ORDER BY created_at ASC;

-- name: CreateMerchant :exec
INSERT INTO admins (
    name, email, created_at 
) VALUES (
    $1, $2, NOW() 
);