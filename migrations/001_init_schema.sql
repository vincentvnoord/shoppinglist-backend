
CREATE TABLE lists (
    id SERIAL PRIMARY KEY,
    public_code VARCHAR(16) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    amount VARCHAR(255),
    completed BOOLEAN DEFAULT FALSE,
    notes TEXT,
    list_id INT NOT NULL,
    FOREIGN KEY (list_id) REFERENCES lists (id) ON DELETE CASCADE
);