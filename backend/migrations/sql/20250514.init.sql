CREATE TABLE users
(
    id            serial PRIMARY KEY,
    email         text UNIQUE NOT NULL,
    first_name    text,
    last_name     text,
    password_hash text NOT NULL
);

CREATE TABLE sessions
(
    id serial PRIMARY KEY,
    token_hash text UNIQUE NOT NULL,
    user_id serial NOT NULL REFERENCES users(id),
    expires_at timestamp NOT NULL,
    revoked bool DEFAULT FALSE,
    created_at timestamp DEFAULT now()
);