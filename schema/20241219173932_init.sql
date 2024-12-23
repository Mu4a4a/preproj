-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id serial not null primary key,
    name varchar(255) not null,
    email varchar(255) not null unique,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE products (
    id serial not null primary key,
    name varchar(255) not null,
    description varchar(255) not null,
    price float not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id integer,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;

DROP TABLE users;
-- +goose StatementEnd
