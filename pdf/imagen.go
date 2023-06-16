package pdf

type Imagen struct {
	Base64String string `json:"base64String"`
	Extension    string `json:"extension"`
	Descripcion  string `json:"descripcion"`
}
