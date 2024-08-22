CREATE TABLE IF NOT EXISTS workspace (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    profile_picture VARCHAR(255),
    metadata JSON
);