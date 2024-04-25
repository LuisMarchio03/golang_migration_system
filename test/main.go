package main

import (
	"fmt"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"github.com/LuisMarchio03/golang_migration_system/internal/exec"
)

// essa func é apenas para testar a aplicação!
func main() {
	cfg := config.Cfg{
		User:   "meu_app_user",
		Passwd: "meu_app_password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "meu_app_db",
	}

	// Declarar db
	db, err := exec.ConfigDB("MySql", cfg)
	if err != nil {
		fmt.Println("Erro ao conectar com o db:", err)
		return
	}

	// Gerar e executar migrações
	userSchema := config.Schema{
		TableName: "users",
		Fields: map[string]string{
			"id":       "INT AUTO_INCREMENT PRIMARY KEY",
			"username": "VARCHAR(50)",
			"email":    "VARCHAR(100)",
		},
	}

	// 4. Chamada da função para gerar a migração
	migrationFile, err := exec.GenerateMigration(userSchema)
	if err != nil {
		fmt.Println("Erro ao gerar migração:", err)
		return
	}

	fmt.Println("Migração gerada com sucesso:", migrationFile)

	// 5. Executar migrações
	err = exec.RunMigrations(db)
	if err != nil {
		fmt.Println("Erro ao executar migrações:", err)
		return
	}

	fmt.Println("Todas as migrações foram executadas com sucesso.")
}
