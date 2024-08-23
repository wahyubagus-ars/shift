CREATE TABLE IF NOT EXISTS authentication (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(25) NOT NULL, -- Adjust maximum length as needed
    enabled BOOLEAN NOT NULL DEFAULT TRUE, -- Set default enabled state
    created_at DATETIME NOT NULL,
    created_by INT,
    updated_at DATETIME,
    updated_by INT,
    deleted_at DATETIME,
    deleted_by INT
);