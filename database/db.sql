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
    ADD COLUMN birth_day DATE,
    ADD COLUMN married BOOLEAN DEFAULT FALSE;
DESC customers;

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
    ADD COLUMN birth_day DATE,
    ADD COLUMN married BOOLEAN DEFAULT FALSE;
DESC customer;


INSERT INTO customer(id, name,email,balance,rating,birth_day,married) VALUES 
    (1, 'Wahidin', 'wahidin@gmail.com', 200000, 5.0, '1999-9-9', true),
    (2, 'Aji', 'aji@gmail.com', 200000, 5.0, '1999-9-9', true);
DELETE FROM customers;

INSERT INTO customers(name,email,balance,rating,birth_day,married) VALUES 
    ('Wahidin', 'wahidin@gmail.com', 200000, 5.0, '1999-9-9', true),
    ('Aji', 'aji@gmail.com', 200000, 5.0, '1999-9-9', true);
DELETE FROM customer;

SELECT * FROM customers;
SELECT * FROM customer;