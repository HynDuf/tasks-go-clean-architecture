USE tasks_go_clean_architecture;

-- DOWN
DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS users;

-- UP
CREATE TABLE
  `tasks` (
    id INT NOT NULL AUTO_INCREMENT,
    title TEXT NOT NULL,
    user_id INT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
  );

CREATE TABLE
  `users` (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
  );
