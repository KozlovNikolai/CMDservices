-- +goose Up
-- +goose StatementBegin
CREATE TABLE services
(
    id   serial,
    name varchar(100) not null,
    price DECIMAL          not null,
    PRIMARY KEY (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table services;
-- +goose StatementEnd
