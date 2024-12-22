-- +goose Up
-- +goose StatementBegin
    ALTER TABLE admins 
    ADD admin_status VARCHAR(30);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    ALTER TABLE admins 
    DROP COLUMN admin_status;
-- +goose StatementEnd
