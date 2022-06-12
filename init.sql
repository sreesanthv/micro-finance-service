CREATE TYPE sex AS ENUM('Male', 'Female', 'Other');

CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    location TEXT,
    pan VARCHAR(10),
    address TEXT,
    phone VARCHAR(15),
    sex sex,
    nationality VARCHAR(50),
    email TEXT UNIQUE,
    password TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT,
    debit DOUBLE PRECISION DEFAULT 0,
    credit DOUBLE PRECISION DEFAULT 0,
    narration TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);