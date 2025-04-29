CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    document VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transaction_type (
    id INT PRIMARY KEY,
    description VARCHAR NOT NULL
);

INSERT INTO transaction_type (id, description) VALUES
    ( 1, 'CASH_IN'),
    ( 2, 'CASH_OUT') ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    transaction_type_id INT NOT NULL REFERENCES transaction_type(id),
    account_id INT NOT NULL REFERENCES accounts(id),
    amount FLOAT8 NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);