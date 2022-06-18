-- +migrate Up
CREATE TABLE IF NOT EXISTS owner (
    id bigint PRIMARY KEY,
    first_name text,
    last_name text,
    email text UNIQUE,
    gender text,
    address text,
    city text
);

CREATE TABLE IF NOT EXISTS  car (
    id bigserial PRIMARY KEY,
    car_manufactur text,
    car_model text,
    car_model_year integer,
    owner_id bigint REFERENCES owner(id)
);

-- +migrate Down
DROP TABLE IF EXISTS car;
DROP TABLE IF EXISTS owner;
