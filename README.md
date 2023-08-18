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

Used [dbdiagrams](https://dbdiagram.io) to design a database schema, downloaded 
and installed docker and also table plus, which is a gui for managing database

Day 2: 

Had issues installing docker, turns out I had to install [wsl2](https://docs.docker.com/desktop/wsl/) and a linux distribution, Ubuntu in this case before i could get docker to run properly. also had issues connecting to table plus with postgres, to solve this, I had to uninstall the previous postgres installed on my system.

Day 3:

Unit testing, CRUD Operations: wrote series and ran series of unit test to test database functionality, installed [golang testify](https://github.com/stretchr/testify), had issues with postgres driver, hence failed the first unit test, but was able to fix it by installing [golang lib pq](https://github.com/lib/pq). That wraps up day 3.

Day 4:
Database transaction and how to implement database transaction. the need for a db transaction is as follows;
1. To provide isolation between programs that access the database concurrently
2. To provide a reliable and consistent unit of work, even in case of system failure

Also learnt about ACID property; 
1. Atomicity: Either an operation complete successfully or the whole transaction fails and the db remains unchanged
2. Consistency: The db state must be valid after the transaction. All constraints must be satisfied
3. Isolation: Concurrent transactions must not affect each other
4. Durability: Data written by a successful transaction must be recorded in a persistent storage

created a store to store which provides all functions to execute db queries and transactions, also wrote and ran several unit test to for the store package.

Day 5:
DB TX LOCK: how to debug a transaction deadlock
implemented the updateAccount feature for the store and also wrote and ran a test for it.

Day 6:
Transaction Isolation level and how they are achieved in mysql and postgres db;
Read Phenomena:
1. Dirty Read: Occurs when a transaction reads data written by other concurrent uncommitted transaction
2. Non-Repeatable Read: Occurs when a transaction reads the same row twice and see different values because it has been modified by other committed transaction
3. Phantom Read: A transaction re-execute a query to find rows that satisfy a condition and sets a different set of rows due to changes by other committed transaction
4. Serialization Anomaly: result of a group concurrent committed transactions is impossible to achieve if we try them sequentially in any order without overlapping

## The Four (4) standard isolation levels
1. READ UNCOMMITTED: can see data written by uncommitted transaction
2. READ COMMITTED: Only see data written by committed transaction
3. REPEATABLE READ: Same read query always return same result
4. SERIALIZABLE : Can achieve same result if execute transactions serially in some order instead of concurrently

## MySQL vs Postgres
|               |  MySQL             | Postgres 
| ------------- | -------------      | -------------             |
| `isolation`   | 4 levels           | 3 Levels                  |
| `operation`   | locking  mechanism | Dependencies mechanism    |
| `lowest lv`   | Repeatable Read    | Read Committed            |

In Postgres, READ UNCOMMITTED behaves exactly like READ COMMITTED, making postgres to have 3 levels of isolation by default.

Day 7:
Github Actions + Postgres
setting up CI/CD using github actions, created the ci.yml locally before pushing to create the workflows for the project.
Added postgres service and also map the port for the github to connect to.

Day 8: 
Implemented HTTP API using Gin
installed [Gin](https://github.com/gin-gonic/gin) using go get command 
Gin has a martini-like API and claims to be up to 40 times faster. Gin allows you to build web applications and microservices in Go. It contains a set of commonly used functionalities (e.g., routing, middleware support, rendering, etc.)

* Created An API folder at the root of the project
* Added a server.go file to house all the endpoints
* Added an account.go file to house all handlers used in the account endpoints
* Implemented three (3) endpoints:
    - create account "/account" using POST method
    - get account  "/account/:id" using GET method
    - list accounts "/accounts" using GET method and the query params on post man to generate queries, added page id and also page size as the params

Day 9: 
Loading Config from file and ENV VARS with viper
## Why File?
    - Development: 
        - Easily Specify default configuration for local development and testing

## Why ENV VARS?
    - Deployment: 
        - Easily Override the default configuration for local when deploy with docker containers

## Why VIper?
    - Can Find, Load, unmarshal config file 
        * JSON, TOML, YAML, ENV, INI
    - Read config from environment variables or flag
        * Override existing values, set default values
    - Read config from remote system 
        * Etcd, Consul
    - Live watching and writing config file
        * Reread changed file, save any modification.

Day 10:
Mock DB for testing HTTP API in GO
## Why mock DB
    - Independent Test
        - Isolated Test data
        - avoid conflicts
    - Faster Test
        - Reduce a lot of time talking to the database
    - 100% Coverage
        - Easily setup edge cases: unexpected errors
## Is it good enough to API with a mock ?
    - Yes: our real db store is already tested.

    Mock DB and Real DB should implement the same interface.

## How to mock 
    - Use Fake DB; Memory implement a fake version of DB; Store data in the memory 
    - Use DB Studs: GOMOCK
        - Generate and build studs that return hard coded values

Day 11:
Using mock db to test api endpoints, set gin mode to test mode, so as to enable easy reading of test result
first test coverage: 48.8%
test left: CreateAccount, ListAccount.


<!-- If you want to more clear explanation, see my blog [House Price Prediction using Flask for Beginners](https://techyscientists.blogspot.com/2021/07/house-price-prediction-using-flask.html) -->

Day 12:
Custom params validator in Go
* implemented the createAccount test
* implemented a new endpoint (transfers) and a create transfer handler.
* created a new util function (currency.go)
* the util function is used to validate the currency when make  a transfer instead of hard coding the values.
* TODO: Test transfer endpoint.

Day 13:
* Added users table with unique and foreign key constraints 
* users can have multiple account with different currencies
 - to add the new database, we can either replace the entire  initial schema with the newly generated one but the approach is not standard, since requirements changes all the time, the other alternative is to create a new migration; hence the add users schema files in the migration folder.
* performed migrate up and down few times, and had to manually change the schema migration status to force, to take care of the dirty version.
* Added new migrateup1 and migratedown1 command to the make file

Day 14:
* Created a new db connection using table plus
* tested the newly established connection
* ran unit test for all functionalities till date
