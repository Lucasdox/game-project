CREATE table "user" (
    id uuid not null primary key,
    name text UNIQUE not null,
    games_played int null,
    score int null
);

