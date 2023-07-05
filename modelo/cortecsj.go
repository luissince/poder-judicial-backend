package modelo

type CorteCsj struct {
	Id             string `json:"id"`
	NombreServicio string `json:"nombre_servicio"`
	PasswordDb     string `json:"password_db"`
	NombreDb       string `json:"nombre_db"`
	Ip             string `json:"ip"`
	UserId         string `json:"user_id"`
	Puerto         string `json:"puerto"`
}
