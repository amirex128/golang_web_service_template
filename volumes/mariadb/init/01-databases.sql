CREATE DATABASE IF NOT EXISTS `selloora`;
CREATE DATABASE IF NOT EXISTS `search_db`;
CREATE DATABASE IF NOT EXISTS `vendor_scoring_system`;

CREATE USER 'search_db_user'@'%' IDENTIFIED BY '9xz3jrd8wf';
CREATE USER 'user'@'%' IDENTIFIED BY '123456';
CREATE USER 'selloora'@'%' IDENTIFIED BY 'q6766581Amirex';

GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';
GRANT ALL PRIVILEGES ON *.* TO 'search_db_user'@'%';
GRANT ALL PRIVILEGES ON *.* TO 'selloora'@'%';
GRANT ALL PRIVILEGES ON *.* TO 'user'@'%';

