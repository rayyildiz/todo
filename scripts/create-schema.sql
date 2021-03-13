CREATE TABLE IF NOT EXISTS todos
(
    id         uuid primary key,
    content    text not null,
    completed  boolean   default false,
    created_at timestamp default now()
);
