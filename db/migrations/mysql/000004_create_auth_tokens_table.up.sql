CREATE TABLE IF NOT EXISTS auth_tokens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_account_id INT NOT NULL,
    access_token INT NOT NULL,
    refresh_token INT NOT NULL,
    expires_in DATE NOT NULL,
    is_active TINYINT(1) DEFAULT 0,
    created_at DATE,
    created_by INT,
    updated_at DATE,
    updated_by INT,
    deleted_at DATE,
    deleted_by INT
);
