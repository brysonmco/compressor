CREATE TABLE users
(
    id            serial PRIMARY KEY,
    email         text UNIQUE NOT NULL,
    first_name    text,
    last_name     text,
    password_hash text
);

CREATE TABLE refresh_tokens
(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    token text UNIQUE NOT NULL,
    user_id serial NOT NULL REFERENCES users(id),
    expires_at timestamp NOT NULL,
    revoked bool DEFAULT FALSE,
    created_at timestamp DEFAULT now()
);