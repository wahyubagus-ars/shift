CREATE TABLE IF NOT EXISTS auth_token (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_account_id INT NOT NULL,
    access_token VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    expires_in DATETIME NOT NULL,
    is_active TINYINT(1) DEFAULT 0,
    created_at DATETIME NOT NULL,
    created_by INT,
    updated_at DATETIME,
    updated_by INT,
    deleted_at DATETIME,
    deleted_by INT
);
