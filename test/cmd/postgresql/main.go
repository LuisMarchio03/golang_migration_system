package main

import (
	"fmt"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"github.com/LuisMarchio03/golang_migration_system/internal/exec"
)

func main() {
	// Configurações para o PostgreSQL
	cfg := config.Cfg{
		User:   "meu_app_user",
		Passwd: "meu_app_password",
		Net:    "tcp",
		Addr:   "localhost", // Porta modificada para 5438
		Port:   "5438",
		DBName: "meu_app_db",
	}

	// Conectar ao PostgreSQL
	db, err := exec.ConfigDB("PostgreSQL", cfg)
	if err != nil {
		fmt.Println("Erro ao conectar com o PostgreSQL:", err)
		return
	}

	// Gerar e executar migrações
	userSchema := config.Schema{
		TableName: "users",
		Fields: map[string]string{
			"id":       "SERIAL PRIMARY KEY", // Modificado para SERIAL para PostgreSQL
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
