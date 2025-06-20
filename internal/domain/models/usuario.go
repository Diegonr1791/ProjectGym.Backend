package model

type Usuario struct {
    ID       int    `json:"id"`
    Nombre   string `json:"nombre"`
    Email    string `json:"email"`
    Password string `json:"password"`
}