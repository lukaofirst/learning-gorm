<div align="center">
  <h1>Learning GORM</h1>
</div>

This project demonstrates various concepts and functionalities of the GORM package in Go

-   **Basics**

    -   It's the ORM (Object Relational Mapping) for Go
    -   Used to query and execute commands to a relational database
    -   It's the equivalent of Entity Framework Core from .NET
    -   Packages

        -   GORM package `go get -u gorm.io/gorm`
        -   SQL Server `gorm.io/driver/sqlserver`
        -   SQLite `gorm.io/driver/sqlite`
        -   PostgreSQL `gorm.io/driver/postgres`
        -   MySQL `gorm.io/driver/mysql`

-   **CRUD Operations**

    -   Create
    -   Read
    -   Update
    -   Delete

-   **Relationships**

    -   One-to-One
    -   One-to-Many
    -   Many-to-Many

-   **Transactions**

    -   Executing database operations within a transaction

-   **Raw SQL and ToSQL**

    -   Executing raw SQL queries and generating SQL from GORM queries

-   **Views**

    -   Creating and querying database views

-   **Stored Procedures**

    -   Creating and executing stored procedures

-   **Paging, Filtering, Searching, and Ordering**
    -   Implementing pagination, filtering, searching, and ordering in queries
