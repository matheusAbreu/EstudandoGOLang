package EncontrandoPalindromo

import (
	"container/list"
	"strconv"
	str "strings"
)

type MeuTexto struct {
	conteudo string
	posEsp   *list.List //qnt de espaço: len(posEsp)
}

func (text *MeuTexto) GetConteudo() string {
	return text.conteudo
}
func CriaTexto() *MeuTexto {
	textoNovo := MeuTexto{conteudo: "", posEsp: list.New()}
	return &textoNovo
}
func (text *MeuTexto) LimpandoMeuTexto() {
	text.posEsp = nil
	text.conteudo = ""
}
func (text *MeuTexto) AddTexto(novoConteudo string) {
	text.conteudo = novoConteudo

	for i := 0; i < len(text.conteudo); i++ {
		if text.conteudo[i] == ' ' {
			text.posEsp.PushBack(i)
		}
	}
}
func (text *MeuTexto) VerificaExistePosEsp(valor int) bool {
	for i := text.posEsp.Front(); i != nil; i = i.Next() {
		if i.Value.(int) == valor {
			return true
		}
	}
	return false
}
func (text MeuTexto) ImprimeMeuTexto() string {
	LocalEsp := "Temos " + strconv.Itoa(text.posEsp.Len()) + " espaços no texto:{"
	for i := text.posEsp.Front(); i != nil; i = i.Next() {
		LocalEsp += strconv.Itoa(i.Value.(int)) + "; "
	}
	LocalEsp += "}"
	LocalEsp = str.Replace(LocalEsp, "; }", "}", -1)
	return text.conteudo + ". " + LocalEsp
}

func (text *MeuTexto) VerificandoPalindromo() bool {
	k := len(text.conteudo)

	for i := 0; i <= k; i++ {
		k--

		if text.conteudo[i] == '.' || text.conteudo[i] == ',' || text.conteudo[i] == '?' ||
			text.conteudo[i] == '!' || text.conteudo[i] == ';' || text.conteudo[i] == ':' ||
			text.conteudo[i] == '[' || text.conteudo[i] == ']' || text.conteudo[i] == '-' ||
			text.conteudo[i] == '@' || text.conteudo[i] == '#' || text.conteudo[i] == '}' ||
			text.conteudo[i] == '$' || text.conteudo[i] == '%' || text.conteudo[i] == '&' ||
			text.conteudo[i] == '*' || text.conteudo[i] == '(' || text.conteudo[i] == ')' ||
			text.conteudo[i] == '{' {
			i++
		}
		if text.VerificaExistePosEsp(i) {
			i++
		}
		if text.conteudo[k] == '.' || text.conteudo[k] == ',' || text.conteudo[k] == '?' ||
			text.conteudo[k] == '!' || text.conteudo[k] == ';' || text.conteudo[k] == ':' ||
			text.conteudo[k] == '[' || text.conteudo[k] == ']' || text.conteudo[k] == '-' ||
			text.conteudo[k] == '@' || text.conteudo[k] == '#' || text.conteudo[k] == '}' ||
			text.conteudo[k] == '$' || text.conteudo[k] == '%' || text.conteudo[k] == '&' ||
			text.conteudo[k] == '*' || text.conteudo[k] == '(' || text.conteudo[k] == ')' ||
			text.conteudo[k] == '{' {
			k--
		}
		if text.VerificaExistePosEsp(k) {
			k--
		}
		if text.conteudo[i] != text.conteudo[k] {
			return false
		}
	}

	return true
}
