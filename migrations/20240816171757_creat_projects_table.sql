-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists projects(
    id integer primary key,
    title varchar(255) not null,
    link varchar(255) not null,
    description varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists projects;
-- +goose StatementEnd
