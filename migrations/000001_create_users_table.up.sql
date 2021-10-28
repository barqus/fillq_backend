CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    token VARCHAR(100) NOT NULL,
    token_generated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);