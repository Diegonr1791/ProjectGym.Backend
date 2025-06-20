package model

import "time"

type Login struct {
    ID        int       `json:"id"`
    UsuarioID int       `json:"usuario_id"`
    FechaHora time.Time `json:"fecha_hora"`
    IP        string    `json:"ip"`
}