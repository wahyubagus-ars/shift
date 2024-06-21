CREATE TABLE IF NOT EXISTS authentications (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(25) NOT NULL, -- Adjust maximum length as needed
    enabled BOOLEAN NOT NULL DEFAULT TRUE, -- Set default enabled state
    created_at DATE,
    created_by INT,
    updated_at DATE,
    updated_by INT,
    deleted_at DATE,
    deleted_by INT
);