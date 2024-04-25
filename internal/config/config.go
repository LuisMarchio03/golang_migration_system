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
	TableName string
	Fields    map[string]string // Mapa de nome de campo para tipo de dados
}
