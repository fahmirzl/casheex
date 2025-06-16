-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE cashiers (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    gender ENUM('male', 'female') NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
)

-- +migrate StatementEnd