-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payouts (
    id SERIAL PRIMARY KEY,
    crn SERIAL,
    account_number VARCHAR(50) UNIQUE NOT NULL,
    transaction_status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create an index on the crn column
CREATE INDEX IF NOT EXISTS idx_payouts_crn ON payouts (crn);

-- Start the crn sequence from 1000
ALTER SEQUENCE payouts_crn_seq RESTART WITH 1000;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
