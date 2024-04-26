package main

import (
	"fmt"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"github.com/LuisMarchio03/golang_migration_system/internal/exec"
	_ "github.com/nakagami/firebirdsql"
)

func main() {
	// Configurações para o Firebird
	cfg := config.Cfg{
		User:   "sysdba",
		Passwd: "masterkey",
		Net:    "tcp",
		Addr:   "localhost:3050", // Porta padrão do Firebird
		DBName: "meu_app_db",     // Nome do banco de dados Firebird
	}

	// Conectar ao Firebird
	db, err := exec.ConfigDB("FirebirdSql", cfg)
	if err != nil {
		fmt.Println("Erro ao conectar com o Firebird:", err)
		return
	}

	// Gerar e executar migrações
	userSchema := config.Schema{
		DbType:    "FirebirdSql",
		TableName: "users",
		Fields: map[string]string{
			"id":       "INT PRIMARY KEY", // Modifique conforme a sintaxe do Firebird
			"username": "VARCHAR(50)",
			"email":    "VARCHAR(100)",
		},
	}

	// Gerar migração
	migrationFile, err := exec.GenerateMigration(userSchema)
	if err != nil {
		fmt.Println("Erro ao gerar migração:", err)
		return
	}

	fmt.Println("Migração gerada com sucesso:", migrationFile)

	// Executar migrações
	err = exec.RunMigrations(db)
	if err != nil {
		fmt.Println("Erro ao executar migrações:", err)
		return
	}

	fmt.Println("Todas as migrações foram executadas com sucesso.")
}
