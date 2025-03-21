CREATE DATABASE IF NOT EXISTS laravel_api CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS laravel_frontend CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER IF NOT EXISTS 'dev'@'%' IDENTIFIED BY '12345678';

GRANT ALL PRIVILEGES ON laravel_api.* TO 'dev'@'%';
GRANT ALL PRIVILEGES ON laravel_frontend.* TO 'dev'@'%';

FLUSH PRIVILEGES;