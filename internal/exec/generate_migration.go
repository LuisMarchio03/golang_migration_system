package exec

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
)

// GenerateMigration cria uma nova migração com base nas estruturas de dados fornecidas.
// Ele cria um arquivo .sql com um nome que inclui um timestamp para garantir unicidade.
// Retorna o nome do arquivo de migração criado e um possível erro, se houver.
func GenerateMigration(migrationsDir string, schemas ...config.Schema) (string, error) {
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
		if schema.DbType == "FirebirdSql" {
			migrationContent = fmt.Sprintf("CREATE TABLE %s (\n", schema.TableName)
		}
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

// GenerateMigrationMongoDB cria uma nova "migração" para o MongoDB com base nos documentos fornecidos.
// Ele cria um arquivo com um nome que inclui um timestamp para garantir unicidade.
// Retorna o nome do arquivo de migração criado e um possível erro, se houver.
func GenerateMigrationMongoDB(migrationsDir string, documents ...map[string]interface{}) (string, error) {
	timestamp := time.Now().Format("20060102150405")
	migrationFileName := fmt.Sprintf("migration_%s.js", timestamp) // Usando extensão .js para scripts do MongoDB
	file, err := os.Create(filepath.Join(migrationsDir, migrationFileName))
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Escrever operações no arquivo de migração (por exemplo, inserções de documentos)
	for _, doc := range documents {
		// Construir operação de inserção para o MongoDB
		insertOperation := fmt.Sprintf("db.collection('%s').insertOne(%#v);\n", doc["collection"], doc["document"])
		_, err := file.WriteString(insertOperation)
		if err != nil {
			return "", err
		}
	}

	return migrationFileName, nil
}
