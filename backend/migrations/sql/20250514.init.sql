CREATE TABLE users
(
    id            serial PRIMARY KEY,
    email         text UNIQUE NOT NULL,
    first_name    text,
    last_name     text,
    password_hash text NOT NULL,
    stripe_customer_id text
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

CREATE TABLE subscriptions
(
    id serial PRIMARY KEY,
    user_id serial NOT NULL REFERENCES users(id),
    stripe_subscription_id text UNIQUE NOT NULL,
    stripe_price_id text NOT NULL,
    status text NOT NULL,
    current_period_start timestamp NOT NULL,
    current_period_end timestamp NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
);