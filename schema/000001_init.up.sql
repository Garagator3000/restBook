CREATE TABLE books
(
    id serial not null unique,
    title varchar(255) not null,
    author varchar(255) not null,
    price numeric(7, 2) not null,
    year_of_publish integer not null,
    count integer not null
);