#!/usr/bin/env bash

docker run --name postgres -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres 

sleep 3

docker exec -ti postgres \
    psql postgresql://postgres:password@localhost/postgres \
    -c 'CREATE TABLE public.items (Path varchar(255) constraint firstkey primary key, URL varchar(255));'

docker exec -ti postgres \
    psql postgresql://postgres:password@localhost/postgres \
    -c "INSERT INTO items (path, url) VALUES ('/yaks', 'https://techyaks.com/'), ('/godoc', 'https://godoc.org/');"

docker exec -ti postgres \
    psql postgresql://postgres:password@localhost/postgres \
    -c 'SELECT * from items'
