package controllers

import "net/http"


func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}