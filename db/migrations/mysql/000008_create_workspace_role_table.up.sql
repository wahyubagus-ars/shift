CREATE TABLE IF NOT EXISTS workspace_role (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    policies JSON,
    created_at DATETIME NOT NULL,
    created_by INT,
    updated_at DATETIME,
    updated_by INT,
    deleted_at DATETIME,
    deleted_by INT
);