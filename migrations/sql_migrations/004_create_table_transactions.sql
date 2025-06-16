-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    date DATE NOT NULL,
    user_id INT NOT NULL,
    total INT NOT NULL,
    paid INT NOT NULL,
    `change` INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
)

-- +migrate StatementEnd