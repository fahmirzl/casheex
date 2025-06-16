-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE detail_transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT NOT NULL,
    purchase_price INT NOT NULL,
    selling_price INT NOT NULL,
    quantity INT NOT NULL,
    subtotal INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
)

-- +migrate StatementEnd