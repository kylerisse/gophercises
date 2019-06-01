#!/usr/bin/env bash

docker stop postgres
sleep 3
docker rm postgres
