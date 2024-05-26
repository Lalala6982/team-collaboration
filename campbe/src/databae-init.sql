CREATE DATABASE IF NOT EXISTS mydb;
USE mydb;

DROP TABLE IF EXISTS delivers;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS bases;


CREATE TABLE users(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    password VARCHAR(100) NOT NULL,
    enabled TINYINT NOT NULL DEFAULT 1
);

CREATE TABLE bases(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    base_address VARCHAR(255) NOT NULL,
    -- num_of_robots INT NOT NULL,
    -- num_of_drones INT NOT NULL,
    enabled TINYINT NOT NULL DEFAULT 1
);

CREATE TABLE delivers(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    base_id INT NOT NULL,
    deliver_type VARCHAR(100) NOT NULL,
    deliver_speed INT NOT NULL,
    deliver_status VARCHAR(100),
    enabled TINYINT NOT NULL DEFAULT 1,
    FOREIGN KEY (base_id) REFERENCES bases(id)
);

CREATE TABLE orders(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    shipper VARCHAR(50) NOT NULL,
    from_address VARCHAR(255) NOT NULL,
    from_zip_code VARCHAR(100) NOT NULL,
    from_city VARCHAR(100) NOT NULL,
    from_county VARCHAR(100) NOT NULL,
    from_phone VARCHAR(100),
    from_email VARCHAR(100),
    consignee VARCHAR(100) NOT NULL,
    to_address VARCHAR(255) NOT NULL,
    to_zip_code VARCHAR(100) NOT NULL,
    to_city VARCHAR(100) NOT NULL,
    to_county VARCHAR(100) NOT NULL,
    to_phone VARCHAR(100),
    to_email VARCHAR(100),
    total_weight VARCHAR(100) NOT NULL,
    status VARCHAR(100),
    order_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    price DECIMAL(10, 2),
    price_id INT,
    deliver_id INT,
    enabled TINYINT NOT NULL DEFAULT 1,
    FOREIGN KEY (deliver_id) REFERENCES delivers(id)
);


