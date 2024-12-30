-- name: CreateAdmin :exec
INSERT INTO admins (
    name, email, created_at, admin_status
) VALUES (
    $1, $2, NOW(), $3
);

-- name: ListAdmins :many
SELECT * FROM admins ORDER BY id ASC;

SELECT 