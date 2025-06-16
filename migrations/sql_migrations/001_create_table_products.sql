-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    purchase_price INT NOT NULL,
    selling_price INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)

-- +migrate StatementEnd