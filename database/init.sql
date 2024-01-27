CREATE DATABASE carapp;
\c carapp;

CREATE TABLE car
(
    id varchar(10) not null unique ,
    name varchar(255) not null ,
    brand varchar(255) not null ,
    year varchar(4) not null ,
    price decimal not null ,
    primary key (id)
);

INSERT INTO car(id, name, brand, "year", price) VALUES
    ('A001', 'Innova Zennix', 'Toyota', '2023', 614000000),
    ('A002', 'Civic Turbo', 'Honda', '2022', 680000000);