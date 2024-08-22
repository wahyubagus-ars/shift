CREATE TABLE IF NOT EXISTS project (
     id INT AUTO_INCREMENT PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
     description TEXT,
     client_id INT,
     profile_picture VARCHAR(255),
     metadata JSON
);