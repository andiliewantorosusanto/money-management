-- Create Wallet table
CREATE TABLE wallet (
    id SERIAL PRIMARY KEY,
    name TEXT,
    balance MONEY
);

-- Create Category table
CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name TEXT
);

-- Create Transaction table
CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    wallet_id_from INTEGER REFERENCES wallet(id),
    wallet_id_to INTEGER REFERENCES wallet(id),
    amount MONEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    category_id INTEGER REFERENCES category(id)
);
