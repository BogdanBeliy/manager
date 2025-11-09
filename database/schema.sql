create table if not exists users
(
    id integer primary key ,
    email    varchar(256) unique ,
    password varchar,
    first_name varchar(256),
    last_name varchar(256),
    age integer,
    owner bool default false
)