CREATE TABLE users
(
    id                 serial PRIMARY KEY,
    email              text UNIQUE NOT NULL,
    first_name         text        NOT NULL,
    last_name          text        NOT NULL,
    password_hash      text        NOT NULL,
    stripe_customer_id text,
    email_verified     bool      DEFAULT FALSE,
    created_at         timestamp DEFAULT now(),
    updated_at         timestamp DEFAULT now(),
    last_login         timestamp DEFAULT now()
);

CREATE TABLE sessions
(
    id         serial PRIMARY KEY,
    token_hash text UNIQUE NOT NULL,
    user_id    serial      NOT NULL REFERENCES users (id),
    expires_at timestamp   NOT NULL,
    revoked    bool      DEFAULT FALSE,
    created_at timestamp DEFAULT now()
);

CREATE TABLE plans
(
    id                   serial PRIMARY KEY,
    name                 text    NOT NULL,
    tokens               integer NOT NULL, -- How many tokens they get per month (-1 for unlimited)
    price                integer NOT NULL, -- Price in cents
    priority             text    NOT NULL, -- (standard or express)
    stripe_product_id    text UNIQUE,
    concurrent_jobs      integer NOT NULL, -- How many jobs they can run at the same time
    max_resolution       text    NOT NULL, -- width * height
    file_retention_hours integer NOT NULL, -- How long we keep the files for them
    watermark            bool    NOT NULL  -- Whether the plan has a watermark or not
);

CREATE TABLE subscriptions
(
    id                     serial PRIMARY KEY,
    user_id                serial      NOT NULL REFERENCES users (id),
    stripe_subscription_id text UNIQUE NOT NULL,
    stripe_price_id        text        NOT NULL,
    plan_id                serial      NOT NULL REFERENCES plans (id),
    status                 text        NOT NULL,
    current_period_start   timestamp   NOT NULL,
    current_period_end     timestamp   NOT NULL,
    created_at             timestamp DEFAULT now(),
    updated_at             timestamp DEFAULT now()
);

CREATE TABLE jobs
(
    id                           serial PRIMARY KEY,
    user_id                      serial NOT NULL REFERENCES users (id),
    created_at                   timestamp DEFAULT now(),
    updated_at                   timestamp DEFAULT now(),
    file_uploaded                bool      DEFAULT FALSE,
    file_name                    text,
    status                       text,
    input_codec                  text,
    input_container              text,
    input_resolution_horizontal  integer,
    input_resolution_vertical    integer,
    input_size                   bigint,
    output_codec                 text,
    output_container             text,
    output_resolution_horizontal integer,
    output_resolution_vertical   integer,
    output_size                  bigint
);

CREATE TABLE token_balances
(
    user_id       serial    NOT NULL REFERENCES users (id),
    token_balance INTEGER   NOT NULL DEFAULT 0,
    free_tokens   INTEGER   NOT NULL DEFAULT 100,
    period_start  timestamp NOT NULL,
    period_end    timestamp NOT NULL
)