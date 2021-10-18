create table if not exists items
(
    id          integer primary key,
    name        varchar not null,
    price       float   not null,
    description varchar not null,
    created_at   datetime not null

)