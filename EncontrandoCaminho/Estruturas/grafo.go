package Estruturas

type Vertice struct {
	nome     string
	ligacoes map[string]*Aresta
}

type Aresta struct {
	vertice   interface{}
	distancia int
}

func NovaAresta(vertice interface{}, distancia int) *Aresta {
	aresta := Aresta{vertice: vertice, distancia: distancia}
	return &aresta
}

func NovoVertice(nome string) *Vertice {
	vertice := Vertice{nome: nome, ligacoes: make(map[string]*Aresta)}
	return &vertice
}

func (vertice Vertice) AdicionarLigacao(filho *Vertice, distancia int) {
	caminho := NovaAresta(filho, distancia)
	vertice.ligacoes[filho.nome] = caminho
}

func (vertice Vertice) ImprimirNo() {
	//Não implementado. Te juro.
}
