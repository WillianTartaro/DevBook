package banco

import (
	"DevBook/api/src/config"
	"DevBook/api/src/respostas"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver MYSQL
	"net/http"
)

// Conectar abre a conexão com o banco de dados e a retorna
func conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}

func ConexaoBanco(w http.ResponseWriter) *sql.DB {
	db, erro := conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return nil
	}
	defer db.Close()

	return db
}
