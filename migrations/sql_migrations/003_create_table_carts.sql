-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE carts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT NOT NULL,
    selling_price INT NOT NULL,
    quantity INT NOT NULL,
    subtotal INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)

-- +migrate StatementEnd