CREATE TABLE IF NOT EXISTS user_account (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) UNIQUE NOT NULL,
    email_verified_at DATE,
    password VARCHAR(255),
    authentication_id INT NOT NULL,

    created_at DATETIME NOT NULL,
    created_by INT,
    updated_at DATETIME,
    updated_by INT,
    deleted_at DATETIME,
    deleted_by INT
);