package Estruturas

type Vertice struct {
	nome     string
	ligacoes map[string]*Aresta
}

type no interface {
	ImprimirNo(alturaAtual int, solucao string)
}

type Aresta struct {
	vertice   no
	distancia int
}

func NovaAresta(vertice no, distancia int) *Aresta {
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

func (vertice Vertice) ImprimirNo(alturaAtual int, solucao string) {
	//Não implementado. Te juro.
}
