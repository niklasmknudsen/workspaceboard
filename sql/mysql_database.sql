use test;

-- User creation and granting permissions to access db.
CREATE USER 'nmk'@'localhost' IDENTIFIED BY 'passpass';
grant all privileges on *.* to 'nmk'@'localhost' with grant option;


-- Defining tables 
CREATE TABLE IF NOT EXISTS workspaces (
    workspace_id INT AUTO_INCREMENT PRIMARY KEY,
    workspace_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS boards (
    board_id INT AUTO_INCREMENT PRIMARY KEY,
    board_name VARCHAR(255) NOT NULL,
    workspace_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (workspace_id) 
    REFERENCES workspaces (workspace_id) 
    ON UPDATE RESTRICT 
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tasks (
    task_id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    board_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (board_id) 
    REFERENCES boards (board_id) 
    ON UPDATE RESTRICT 
    ON DELETE CASCADE
);


-- seed data
insert into workspaces(workspace_name) values ('TEST-Workspace-1');
insert into workspaces(workspace_name) values ('TEST-Workspace-2');
insert into boards(board_name, workspace_id) values ('Test-Board-1', 1);
insert into boards(board_name, workspace_id) values ('Test-Board-2', 2);
insert into tasks(description, board_id) values('Task-1', 1);
insert into tasks(description, board_id) values('Task-2', 2);
