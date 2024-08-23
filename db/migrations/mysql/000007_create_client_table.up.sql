CREATE TABLE IF NOT EXISTS client (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    profile_picture VARCHAR(255),
    created_at DATETIME NOT NULL,
    created_by INT,
    updated_at DATETIME,
    updated_by INT,
    deleted_at DATETIME,
    deleted_by INT
);