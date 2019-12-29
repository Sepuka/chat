#!/usr/bin/env sh

mysql -e "CREATE USER 'user'@'localhost'"
mysql -e "CREATE DATABASE db"
mysql -e "GRANT ALL PRIVILEGES ON db.* TO 'user'@'localhost'"
mysql -e "FLUSH PRIVILEGES"

#mv /var/www/joomla/installation/configuration.php-dist /var/www/joomla/installation/configuration.php
#chmod +w /var/www/joomla/installation/configuration.php
#sed -i 's/public \x24user = \x27\x27;/public \x24user = \x27user\x27;/' /var/www/joomla/installation/configuration.php
#sed -i 's/public \x24db = \x27\x27;/public \x24db = \x27db\x27;/' /var/www/joomla/installation/configuration.php
