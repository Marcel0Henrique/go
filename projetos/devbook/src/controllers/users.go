package controllers

import "net/http"

//*Cria um novo Usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}

//* Busca todos os usuários cadastrados
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

//* Busca apenas um Usuario
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

//* Atualiza o usuário cadastrado
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário"))
}

//* Deleta usuário do banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuario"))
}
