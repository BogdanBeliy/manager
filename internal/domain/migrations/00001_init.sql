-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id integer primary key ,
    email    varchar(256) unique not null,
    password varchar,
    first_name varchar(256),
    last_name varchar(256),
    age integer,
    owner bool default false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
