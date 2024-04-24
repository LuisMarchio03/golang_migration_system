# Golang Migration System

## Description

The Golang Migration System is a Go library that facilitates the management of MySQL database migrations.

## Installation

To install the library, you can use the go get command:

```bash
go get github.com/LuisMarchio03/golang_migration_system
```

## Usage

The library offers simple functionalities to configure and execute database migrations. Here's a basic example of how you can use it:

```bash
package main

import (
    "fmt"
    "github.com/LuisMarchio03/golang_migration_system"
)

func main() {
    // Database configuration
    cfg := golang_migration_system.Cfg{
        User:   "root",
        Passwd: "password",
        Net:    "tcp",
        Addr:   "localhost:3306",
        DBName: "my_database",
    }

    // Configure the database
    db, err := golang_migration_system.ConfigDB("MySql", cfg)
    if err != nil {
        fmt.Println("Error configuring the database:", err)
        return
    }

    // Define the table schema
    schema := golang_migration_system.Schema{
        TableName: "users",
        Fields: map[string]string{
            "id":         "INT NOT NULL AUTO_INCREMENT PRIMARY KEY",
            "username":   "VARCHAR(50) NOT NULL",
            "email":      "VARCHAR(100) NOT NULL",
            "created_at": "TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
            "updated_at": "TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
        },
    }

    // Generate the migration
    migrationFileName, err := golang_migration_system.GenerateMigration(schema)
    if err != nil {
        fmt.Println("Error generating migration:", err)
        return
    }

    fmt.Println("Migration generated successfully:", migrationFileName)

    // Execute migrations
    err = golang_migration_system.RunMigrations(db)
    if err != nil {
        fmt.Println("Error executing migrations:", err)
        return
    }

    fmt.Println("Migrations completed successfully.")
}
```

## Contributions

Contributions are welcome! If you find an issue or have an idea to improve the library, feel free to open an issue or submit a pull request.

## License 

Just replace the `[MIT License](LICENSE)` link with the appropriate link to your license file or the license text itself.
