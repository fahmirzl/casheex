-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transaction_details (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT NOT NULL,
    purchase_price INT NOT NULL,
    selling_price INT NOT NULL,
    quantity INT NOT NULL,
    subtotal INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id)
)

-- +migrate StatementEnd