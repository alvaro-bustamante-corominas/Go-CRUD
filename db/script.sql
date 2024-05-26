CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    status ENUM('Created', 'InProgress', 'Completed') NOT NULL,
    date DATE NOT NULL
);