CREATE TABLE IF NOT EXISTS user_accounts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) UNIQUE NOT NULL,
    email_verified_at DATE,
    password VARCHAR(255),
    authentication_id INT NOT NULL,

    created_at DATE NOT NULL,
    created_by INT,
    updated_at DATE,
    updated_by INT,
    deleted_at DATE,
    deleted_by INT
);