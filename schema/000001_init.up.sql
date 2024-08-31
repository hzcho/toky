create table if not exists file_metadata(
    id bigserial primary key,
    file_name varchar(255) not null,
    path text not null,
    size bigint not null,
    created_at timestamp not null
);

create table if not exists users(
    id bigserial primary key,
    email varchar(255) not null unique,
    pass_hash text not null
);