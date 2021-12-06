CREATE TABLE customers (
    id bigint(20) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name varchar(100)  NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;
ALTER TABLE customers
    ADD COLUMN email varchar(100),
    ADD COLUMN balance INTEGER DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT FALSE;
DESC customers;
INSERT INTO customers(name,email,balance,rating,birth_date,married) VALUES 
    ('Wahidin', 'wahidin@gmail.com', 200000, 95.0, '1999-9-9', true),
    ('Aji', 'aji@gmail.com', 100000, 85.5, '1998-8-8', true),
    ('joko', 'joko@gmail.com', 100000, 87.5, '1877-07-15', true);
SELECT * FROM customers;
UPDATE customers SET email = NULL, birth_date = NULL WHERE name='joko';
DELETE FROM customers;
DROP TABLE customers;

/** CREATE TABLE users */
CREATE TABLE users (
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY(username)
)ENGINE =InnoDB;
INSERT INTO users (username, password) VALUES ('admin','admin');
SELECT * FROM users;
DELETE FROM users;

/** CREATE TABLE comments */
CREATE TABLE comments(
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100)  NOT NULL,
    comment TEXT,
    PRIMARY KEY(id)
)ENGINE =InnoDB;
DESC comments;
SELECT * FROM comments;






/** customer tanpa s */
CREATE TABLE customer (
    id varchar(100) NOT NULL,
    name varchar(100)  NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB;
ALTER TABLE customer
    ADD COLUMN email varchar(100),
    ADD COLUMN balance INTEGER DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT FALSE;
DESC customer;


INSERT INTO customer(id, name,email,balance,rating,birth_date,married) VALUES 
    (1, 'Wahidin', 'wahidin@gmail.com', 200000, 95.0, '1999-9-9', true),
    (2, 'Aji', 'aji@gmail.com', 100000, 85.5, '1998-8-8', true);

SELECT * FROM customer;
DELETE FROM customer;
DROP TABLE customer;