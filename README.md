# Learning GORM 📚

## Overview 🌟

This project demonstrates various concepts and functionalities of the GORM package in Go.

## GORM Basics 🛠️

GORM is an ORM (Object Relational Mapping) for Go, used to query and execute commands to a relational database. It's the equivalent of Entity Framework Core from .NET.

### Key Packages 📦

-   **GORM package**: `go get -u gorm.io/gorm`
-   **SQL Server**: `go get -u gorm.io/driver/sqlserver`
-   **SQLite**: `go get -u gorm.io/driver/sqlite`
-   **PostgreSQL**: `go get -u gorm.io/driver/postgres`
-   **MySQL**: `go get -u gorm.io/driver/mysql`

## CRUD Operations 🔄

-   **Create**: Add new records to the database.
-   **Read**: Retrieve records from the database.
-   **Update**: Modify existing records in the database.
-   **Delete**: Remove records from the database.

## Relationships 🔗

-   **One-to-One**: A single record in one table is related to a single record in another table.
-   **One-to-Many**: A single record in one table is related to multiple records in another table.
-   **Many-to-Many**: Multiple records in one table are related to multiple records in another table.

## Transactions 🔒

Executing database operations within a transaction to ensure data integrity.

## Raw SQL and ToSQL 📝

-   **Raw SQL**: Executing raw SQL queries directly.
-   **ToSQL**: Generating SQL from GORM queries.

## Views 👁️

Creating and querying database views to simplify complex queries.

## Stored Procedures 📜

Creating and executing stored procedures for reusable database logic.

## Paging, Filtering, Searching, and Ordering 🔍

Implementing pagination, filtering, searching, and ordering in queries to manage large datasets efficiently.
