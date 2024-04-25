package exec

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// RunMigrations executa todas as migrações encontradas no diretório MigrationsDir no banco de dados especificado.
// Retorna um possível erro, se houver.
func RunMigrations(db *sql.DB) error {
	// 1. Verificar se o diretório de migrações existe
	if _, err := os.Stat(MigrationsDir); os.IsNotExist(err) {
		return fmt.Errorf("O diretório de migrações não existe")
	}

	// 2. Listar arquivos de migração
	files, err := ioutil.ReadDir(MigrationsDir)
	if err != nil {
		return fmt.Errorf("Erro ao listar arquivos de migração: %v", err)
	}

	// 3. Executar migrações
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".sql" {
			migrationPath := filepath.Join(MigrationsDir, f.Name())
			fmt.Println("Executando migração:", migrationPath)

			// Lê o conteúdo do arquivo de migração
			query, err := ioutil.ReadFile(migrationPath)
			if err != nil {
				return fmt.Errorf("Erro ao ler arquivo de migração %s: %v", migrationPath, err)
			}

			// Executa a migração
			_, err = db.Exec(string(query))
			if err != nil {
				return fmt.Errorf("Erro ao executar migração %s: %v", migrationPath, err)
			}

			fmt.Println("Migração concluída com sucesso.")
		}
	}

	return nil
}
