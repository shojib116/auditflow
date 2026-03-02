-- +goose up
-- +goose statementbegin
create extension if not exists "pgcrypto";

create or replace function update_updated_at_column()
returns trigger as $$
begin
    new.updated_at = now();
    return new;
end;
$$ language 'plpgsql';
-- +goose statementend

-- +goose down
-- +goose statementbegin
drop function if exists update_updated_at_column();
-- +goose statementend
