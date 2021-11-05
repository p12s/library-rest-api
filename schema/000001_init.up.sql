CREATE TABLE IF NOT EXISTS users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS book
(
    id          serial       not null unique,
    title       varchar(255) not null
);

CREATE TABLE IF NOT EXISTS author
(
    id          serial       not null unique,
    first_name  varchar(255) not null,
    second_name   varchar(255) not null
);

CREATE TABLE IF NOT EXISTS book_author
(
    id          serial                                       not null unique,
    book_id     int references book (id) on delete cascade   not null,
    author_id   int references author (id) on delete cascade not null
);
