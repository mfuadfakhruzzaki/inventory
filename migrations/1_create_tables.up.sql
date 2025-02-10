CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    quantity INTEGER NOT NULL DEFAULT 0,
    location VARCHAR(255) NOT NULL,
    category_id INTEGER REFERENCES categories(id)
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    item_id INTEGER REFERENCES items(id),
    transaction_type VARCHAR(10) NOT NULL CHECK (transaction_type IN ('in', 'out')),
    quantity INTEGER NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT now(),
    notes TEXT
);
