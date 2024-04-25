package exec

import (
	"fmt"
)

// MigrationsDir é o diretório onde estão localizados os arquivos de migração.
var MigrationsDir = "../cmd/migrations"

// InitHandle inicializa o pacote exec.
func InitHandle() {
	fmt.Println("Handle inicializado")
	// Aqui você pode adicionar outras tarefas de inicialização, se necessário
}
