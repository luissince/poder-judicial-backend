package pdf

type Data struct {
	NombreSistema  string `json:"nombreSistema,omitempty"`
	VersionSistema string `json:"versionSistema,omitempty"`
	UsuarioNombre  string `json:"usuarioNombre,omitempty"`
	CelularAxeso   string `json:"celularAxeso,omitempty"`
	Sede           string `json:"sede,omitempty"`
	Cargo          string `json:"cargo,omitempty"`
	PersonaReporte string `json:"personaReporte,omitempty"`
	CelularPersona string `json:"celularPersona,omitempty"`
	Fecha          string `json:"fecha,omitempty"`
	Descripcion    string `json:"descripcion,omitempty"`
	Descartes      string `json:"descartes,omitempty"`
}
