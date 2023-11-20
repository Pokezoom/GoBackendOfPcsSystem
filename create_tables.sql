CREATE TABLE video (
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       name VARCHAR(255) NOT NULL UNIQUE,
                       user_id INT,
                       duration INT NOT NULL,
                       url VARCHAR(512) NOT NULL,
                       class VARCHAR(50),
                       academic_year VARCHAR(50),
                       subject VARCHAR(50),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       deleted TINYINT(1) DEFAULT 0
);

CREATE TABLE user (
                      userId INT PRIMARY KEY AUTO_INCREMENT,
                      name VARCHAR(255) NOT NULL UNIQUE,
                      password VARCHAR(255) NOT NULL,
                      email VARCHAR(255),
                      phone_number VARCHAR(20),
                      user_type ENUM('1', '2', '3') NOT NULL,
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      deleted TINYINT(1) DEFAULT 0
);

CREATE TABLE video_analysis (
                                id INT PRIMARY KEY AUTO_INCREMENT,
                                video_id INT NOT NULL,
                                student_attendance INT NOT NULL DEFAULT 0,
                                facial_data JSON,
                                fatigue_data JSON,
                                limb_data JSON,
                                study_status_data JSON,
                                image_url VARCHAR(512) NOT NULL DEFAULT '',
                                video_url VARCHAR(512) NOT NULL DEFAULT '',
                                uploader_id INT NOT NULL DEFAULT 0,
                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                deleted TINYINT(1) DEFAULT 0
);