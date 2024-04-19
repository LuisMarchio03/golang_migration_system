package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const migrationsDir = "./migrations"

// Configuração do banco de dados
type CfgMySql struct {
	User   string
	Passwd string
	Net    string
	Addr   string
	DBName string
}

// Schema representa um esquema de tabela
type Schema struct {
	TableName string
	Fields    map[string]string // Mapa de nome de campo para tipo de dados
}

func ConfigDB(dbDriver string, cfg interface{}) error {
	if dbDriver == "MySql" {
		DbMysql(cfg)
	}

	return nil
}

func DbMysql(cfg CfgMySql) {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.User, cfg.Passwd, cfg.Net, cfg.Addr, cfg.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer db.Close()
}

// GenerateMigration cria uma nova migração com base nas estruturas de dados fornecidas.
// Ele cria um arquivo .sql com um nome que inclui um timestamp para garantir unicidade.
// Retorna o nome do arquivo de migração criado e um possível erro, se houver.
func GenerateMigration(schemas ...Schema) (string, error) {
	// 1. Criar o arquivo .sql da migration
	// - Gera um timestamp do momento atual.
	// - Cria um nome para a migração usando o timestamp.
	// - Cria um arquivo com o nome da migração no diretório de migrações.
	timestamp := time.Now().Format("20060102150405")
	migrationFileName := fmt.Sprintf("migration_%s.sql", timestamp)
	file, err := os.Create(filepath.Join(migrationsDir, migrationFileName))
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 2. Definir o conteúdo da migração SQL
	migrationContent := ""
	for _, schema := range schemas {
		migrationContent += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", schema.TableName)
		for fieldName, fieldType := range schema.Fields {
			migrationContent += fmt.Sprintf("    %s %s,\n", fieldName, fieldType)
		}
		migrationContent += ");\n\n"
	}

	// 3. Escrever o conteúdo da migração no arquivo
	_, err = file.WriteString(migrationContent)
	if err != nil {
		return "", err
	}

	return migrationFileName, nil
}

// RunMigrations executa todas as migrações encontradas no diretório migrationsDir no banco de dados especificado.
// Retorna um possível erro, se houver.
func RunMigrations(db *sql.DB) error {
	// 1. Verificar se o diretório de migrações existe
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return fmt.Errorf("O diretório de migrações não existe")
	}

	// 2. Listar arquivos de migração
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("Erro ao listar arquivos de migração: %v", err)
	}

	// 3. Executar migrações
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".sql" {
			migrationPath := filepath.Join(migrationsDir, f.Name())
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

func main() {

	// Gerar e executar migrações
	userSchema := Schema{
		TableName: "users",
		Fields: map[string]string{
			"id":       "INT AUTO_INCREMENT PRIMARY KEY",
			"username": "VARCHAR(50)",
			"email":    "VARCHAR(100)",
		},
	}

	// 4. Chamada da função para gerar a migração
	migrationFile, err := GenerateMigration(userSchema)
	if err != nil {
		fmt.Println("Erro ao gerar migração:", err)
		return
	}

	fmt.Println("Migração gerada com sucesso:", migrationFile)

	// 5. Executar migrações
	err = RunMigrations(db)
	if err != nil {
		fmt.Println("Erro ao executar migrações:", err)
		return
	}

	fmt.Println("Todas as migrações foram executadas com sucesso.")
}