CREATE TABLE IF NOT EXISTS user_workspace (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    workspace_id INT NOT NULL,
    workspace_role_id INT NOT NULL
);