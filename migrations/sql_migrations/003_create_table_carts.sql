-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE carts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT NOT NULL,
    selling_price INT NOT NULL,
    quantity INT NOT NULL,
    subtotal INT NOT NULL,
    cashier_id INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (cashier_id) REFERENCES cashiers(id)
)

-- +migrate StatementEnd