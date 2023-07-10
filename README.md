# Golang Master Class Project

This repository documents my learning experiences in progressively complex topics as I created a Golang backend application. K8, AWS, and Docker are all referenced. A lot of testing and other effort is required.

## Requirements
Basic Understanding of Golang

golang installed

Docker

Postgres (it wasn't easy installing docker and postgres on windows machine, most especially Postgres, I will add the work around method I used in achieving this.)

sqlc

## Progress Report

Day 1:

used [dbdiagrams](https://dbdiagram.io) to design a database schema, downloaded 
and installed docker and also table plus, which is a gui for managing database

Day 2: 

had issues installing docker, turns out I had to install [wsl2](https://docs.docker.com/desktop/wsl/) and a linux distribution, Ubuntu in this case before i could get docker to run properly. also had issues connecting to table plus with postgres, to solve this, I had to uninstall the previous postgres installed on my system.

Day 3:

Unit testing, CRUD Operations: wrote series and ran series of unit test to test database functionality, installed [golang testify](https://github.com/stretchr/testify), had issues with postgres driver, hence failed the first unit test, but was able to fix it by installing [golang lib pq](https://github.com/lib/pq). That wraps up day 3.