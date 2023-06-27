package pdf

type Data struct {
	NombreSIJ      string   `json:"nombreSIJ,omitempty"`
	NombreWEB      string   `json:"nombreWEB,omitempty"`
	VersionSistema string   `json:"versionSistema,omitempty"`
	UsuarioNombre  string   `json:"usuarioNombre,omitempty"`
	CelularAxeso   string   `json:"celularAxeso,omitempty"`
	Sede           string   `json:"sede,omitempty"`
	Cargo          string   `json:"cargo,omitempty"`
	PersonaReporte string   `json:"personaReporte,omitempty"`
	CelularPersona string   `json:"celularPersona,omitempty"`
	Fecha          string   `json:"fecha,omitempty"`
	Descripcion    string   `json:"descripcion,omitempty"`
	PreguntaUno    string   `json:"preguntaUno,omitempty"`
	PreguntaDos    string   `json:"preguntaDos,omitempty"`
	PreguntaTres   string   `json:"preguntaTres,omitempty"`
	PreguntaCuatro string   `json:"preguntaCuatro,omitempty"`
	PreguntaCinco  string   `json:"preguntaCinco,omitempty"`
	Imagenes       []Imagen `json:"imagenes"`
	NombreServicio string `json:"nombreServicio,omitempty"`
	PasswordBaseDatos string `json:"passwordBaseDatos,omitempty"`
}
