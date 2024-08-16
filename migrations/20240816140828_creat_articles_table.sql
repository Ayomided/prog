-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists articles(
    id integer primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title varchar(255) not null,
    slug varchar(255) not null,
    filename varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists articles;
-- +goose StatementEnd
