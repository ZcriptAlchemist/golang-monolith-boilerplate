-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS status_checks (
    id SERIAL PRIMARY KEY,
    crn SERIAL,
    account_number VARCHAR(50) UNIQUE NOT NULL,
    transaction_status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS status_checks;
-- +goose StatementEnd
