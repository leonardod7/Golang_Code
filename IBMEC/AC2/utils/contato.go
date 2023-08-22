package utils

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) ModificarEmail(emailAtual string) {
	c.Email = emailAtual
}
