CREATE TABLE IF NOT EXISTS users (
     id INT AUTO_INCREMENT PRIMARY KEY,
     email VARCHAR(50) NOT NULL UNIQUE,
     password VARCHAR(60) NOT NULL,
     home_folder_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS folders (
       folder_id INT AUTO_INCREMENT PRIMARY KEY,
       folder_name VARCHAR(50) NOT NULL,
       parent_folder_id INT,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
       UNIQUE (parent_folder_id, folder_name)
);

CREATE TABLE IF NOT EXISTS files (
     file_id INT AUTO_INCREMENT PRIMARY KEY,
     folder_id INT NOT NULL,
     file_name VARCHAR(50) NOT NULL,
     owner_id INT NOT NULL,
     mime_type VARCHAR(50) NOT NULL,
     size BIGINT UNSIGNED NOT NULL,
     compress_algo VARCHAR(20),
     s3_link VARCHAR(2083),
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     UNIQUE (folder_id, file_name)
);

CREATE TABLE IF NOT EXISTS folder_file (
           folder_id INT NOT NULL,
           file_id INT NOT NULL,
           PRIMARY KEY (folder_id, file_id)
);