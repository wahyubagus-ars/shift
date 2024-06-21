CREATE TABLE IF NOT EXISTS user_profiles (
    id INT PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    profile_picture VARCHAR(255),
    dob INT,
    pob VARCHAR(255),
    status INT,

    created_at DATE NOT NULL,
    created_by INT,
    updated_at DATE,
    updated_by INT,
    deleted_at DATE,
    deleted_by INT
);