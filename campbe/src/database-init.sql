-- Create the database if it does not exist
CREATE DATABASE IF NOT EXISTS mydb;

-- Use the newly created database
USE mydb;

-- Drop existing tables if they exist
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS delivers;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS bases;

-- Create the 'users' table with a JSON column
CREATE TABLE users (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    enabled TINYINT NOT NULL DEFAULT 1,
    history JSON -- Renamed the JSON column to history
);

-- Create the 'bases' table
CREATE TABLE bases (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    base_address VARCHAR(255) NOT NULL,
    base_city VARCHAR(100) NOT NULL,
    base_zip_code VARCHAR(100) NOT NULL,
    enabled TINYINT NOT NULL DEFAULT 1
);

-- Create the 'delivers' table with a foreign key reference to 'bases'
CREATE TABLE delivers (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    base_id INT NOT NULL,
    deliver_type VARCHAR(100) NOT NULL,
    deliver_duration VARCHAR(100),
    deliver_status VARCHAR(100),
    enabled TINYINT NOT NULL DEFAULT 1,
    FOREIGN KEY (base_id) REFERENCES bases(id)
);

-- Create the 'orders' table
CREATE TABLE orders (
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
    total_weight DECIMAL(10, 2), NOT NULL,
    status VARCHAR(100),
    order_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    price DECIMAL(10, 2),
    price_id INT,
    deliver VARCHAR(100),
    enabled TINYINT NOT NULL DEFAULT 1
);

-- Insert records into the 'bases' table
INSERT INTO bases (base_address, base_city, base_zip_code)
VALUES ('153-44 S Conduit Ave', 'Jamaica', '11434');

INSERT INTO bases (base_address, base_city, base_zip_code)
VALUES ('900 Turnbull Canyon Rd', 'City Of Industry', '91745');

INSERT INTO bases (base_address, base_city, base_zip_code)
VALUES ('1657 N Kostner Ave', 'Chicago', '60639');
