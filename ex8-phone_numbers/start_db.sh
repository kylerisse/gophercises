#!/usr/bin/env bash

docker run --name mariadb -e MYSQL_ROOT_PASSWORD=password -d -p 3306:3306 mariadb

sleep 10 

docker exec -ti mariadb \
    mysql -ppassword -e "CREATE DATABASE numbers; \
        USE numbers; \
        CREATE TABLE numbers (numbers INT NOT NULL PRIMARY KEY); \
        DESCRIBE numbers;"
