package domain

type Dentista struct {
	ID        int    `json:"id"`
	Apellido  string `json:"apellido" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
}
