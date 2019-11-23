package Estruturas

type Vertice struct {
	nome     string
	ligacoes map[string]Aresta
}

type Aresta struct {
	vertice   interface{}
	distancia int
}

func NovoVertice(nome string) *Vertice {
	vertice := Vertice{nome: nome, ligacoes: make(map[string]Aresta)}
	return &vertice
}

func (vertice Vertice) AdicionarLigacao(nomefilho string,
	filho Vertice, distancia int) {

	caminho := Aresta{filho, distancia}
	vertice.ligacoes[nomefilho] = caminho
}

func (vertice Vertice) Imprimir() {
	//Não implementado. Te juro.
}
