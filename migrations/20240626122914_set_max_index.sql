-- +goose Up
-- +goose StatementBegin
SELECT setval(pg_get_serial_sequence('services', 'id'), max(id)) FROM services;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
