package utils

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) ModificarEmail(emailAtual string) {
	c.Email = emailAtual
}

func AdicionarContato(contato Contato, contatos *[5]*Contato) bool {

	for i := 0; i < len(contatos); i++ {
		if contatos[i] == nil {
			contatos[i] = &contato
			return true
		}
	}
	return false
}

func RemoverContato(contatos *[5]*Contato) {
	for i := len(contatos) - 1; i >= 0; i-- { //começa do último elemento do array e avança para o começo
		if contatos[i] != nil {
			contatos[i] = nil
			return
		}
	}
}

func VisualizarListaContatos(contatos [5]*Contato) []string {
	var resultado []string
	for _, c := range contatos {
		if c != nil {
			informacoesContato := "Nome: " + c.Nome + " | Email: " + c.Email
			resultado = append(resultado, informacoesContato)
		}
	}
	return resultado
}
