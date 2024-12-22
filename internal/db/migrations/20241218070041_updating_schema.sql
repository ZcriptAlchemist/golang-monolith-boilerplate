-- +goose Up
-- +goose StatementBegin
    ALTER TABLE partners 
    ADD status VARCHAR(30);
    
    ALTER TABLE partners 
    ADD username VARCHAR(30);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE partners
DROP COLUMN status;

ALTER TABLE partners 
DROP COLUMN username;
-- +goose StatementEnd
