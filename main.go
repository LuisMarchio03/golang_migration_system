package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var MigrationsDir = "./migrations"

// Configuração do banco de dados
type Cfg struct {
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

// ConfigDB configura o banco de dados com base no driver especificado e nas configurações fornecidas.
// Ele recebe o nome do driver do banco de dados e as configurações do banco de dados como parâmetros.
// Se o driver for "MySql", chama a função DbMysql para configurar o banco de dados MySQL.
// Retorna um possível erro, se houver.
func ConfigDB(dbDriver string, cfg Cfg) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch dbDriver {
	case "MySql":
		db, err = DbMysql(cfg)
	// Adicione mais cases aqui para outros drivers de banco de dados
	default:
		return nil, fmt.Errorf("Driver de banco de dados não suportado: %s", dbDriver)
	}

	return db, err
}

// DbMysql estabelece uma conexão com um banco de dados MySQL utilizando as configurações fornecidas.
// Recebe um struct Cfg contendo os detalhes de configuração do banco de dados, como usuário, senha, endereço e nome do banco de dados.
// Retorna um ponteiro para sql.DB, que representa a conexão com o banco de dados, e um possível erro, se houver.
func DbMysql(cfg Cfg) (*sql.DB, error) {
	cfgMysql := mysql.Config{
		User:   cfg.User,
		Passwd: cfg.Passwd,
		Net:    cfg.Net,
		Addr:   cfg.Addr,
		DBName: cfg.DBName,
	}

	db, err := sql.Open("mysql", cfgMysql.FormatDSN())
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}

	return db, nil
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
	file, err := os.Create(filepath.Join(MigrationsDir, migrationFileName))
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 2. Definir o conteúdo da migração SQL
	migrationContent := ""
	for _, schema := range schemas {
		migrationContent += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", schema.TableName)
		fields := schema.Fields
		numFields := len(fields)
		i := 0
		for fieldName, fieldType := range fields {
			// Adiciona o campo com o tipo correspondente
			migrationContent += fmt.Sprintf("    %s %s", fieldName, fieldType)
			// Se não for o último campo, adiciona vírgula
			// Poís a vírgula no ultimo campo, ocasiona um  erro de sintaxe no SQL
			if i < numFields-1 {
				migrationContent += ","
			}
			migrationContent += "\n"
			i++
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

// TODO: No momento essa func é apenas para testar a aplicação!
func main() {
	cfg := Cfg{
		User:   "meu_app_user",
		Passwd: "meu_app_password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "meu_app_db",
	}

	// Declarar db
	db, err := ConfigDB("MySql", cfg)
	if err != nil {
		fmt.Println("Erro ao conectar com o db:", err)
		return
	}

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
