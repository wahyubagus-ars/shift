CREATE TABLE IF NOT EXISTS user_profile (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_account_id INT,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    dob INT,
    pob VARCHAR(255),
    status INT,
    metadata JSON,
    created_at DATETIME NOT NULL,
    created_by INT,
    updated_at DATETIME,
    updated_by INT,
    deleted_at DATETIME,
    deleted_by INT
);