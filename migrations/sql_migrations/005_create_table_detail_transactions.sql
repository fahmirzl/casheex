-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transaction_details (
    id INT PRIMARY KEY AUTO_INCREMENT,
    transaction_id INT NOT NULL,
    product_id INT NOT NULL,
    purchase_price INT NOT NULL,
    selling_price INT NOT NULL,
    quantity INT NOT NULL,
    subtotal INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (transaction_id) REFERENCES transactions(id)
)

-- +migrate StatementEnd