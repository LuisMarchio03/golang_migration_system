package config

// Configuração do banco de dados
type Cfg struct {
	User     string
	Passwd   string
	Net      string
	Addr     string
	DBName   string
	Port     string
	Keyspace string
	Service  string
}

// Schema representa um esquema de tabela
type Schema struct {
	DbType    string
	TableName string
	Fields    map[string]string // Mapa de nome de campo para tipo de dados
}

// Documents representa um documento a ser inserido em uma migração MongoDB
//   - Exemplo de uso de ExecGenerateMigrationMongoDB
//     schemas := config.Documents{
//     CollectionName: "users",
//     Documents: []map[string]interface{}{
//     {"name": "John", "age": 30, "city": "New York"},
//     {"name": "Alice", "age": 25, "city": "San Francisco"},
//     },
//     }
type Documents struct {
	CollectionName string                   // Nome da coleção onde o documento será inserido
	Documents      []map[string]interface{} // Lista de documentos a serem inseridos (cada documento é um mapa de chaves e valores)
}
