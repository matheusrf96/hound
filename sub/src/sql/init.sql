CREATE DATABASE hound;

CREATE TABLE person (
    id SERIAL
    , full_name VARCHAR(64) NOT NULL
    , email VARCHAR(64) NOT NULL
    , birth_date DATE
);