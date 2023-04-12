#!/bin/bash

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "CREATE TABLE IF NOT EXISTS TEST(
    id int(10) AUTO_INCREMENT NOT NULL primary key,
    text varchar(50) NOT NULL
);"
$CMD_MYSQL -e "INSERT INTO TEST VALUES (1,'first text');"