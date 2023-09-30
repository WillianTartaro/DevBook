package controllers

import (
	"DevBook/api/src/autenticacao"
	"DevBook/api/src/banco"
	"DevBook/api/src/modelos"
	"DevBook/api/src/repositorios"
	"DevBook/api/src/respostas"
	"DevBook/api/src/seguranca"
	"encoding/json"
	"io"
	"net/http"
)

// Login é responsavel por autenticar um usuario na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db := banco.ConexaoBanco(w)

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarioSalvo, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvo.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvo.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))

}
