CREATE TABLE IF NOT EXISTS user_account (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_active TINYINT(1) DEFAULT 1,

    created_at DATE NOT NULL,
    created_by INT,
    updated_at DATE,
    updated_by INT,
    deleted_at DATE,
    deleted_by INT
);