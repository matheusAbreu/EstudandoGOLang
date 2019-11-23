package Estruturas

type Grafo struct {
	Vertice
	valorAcumulado int
}

func NovoGrafo(nome string) *Grafo {
	vertice := Vertice{nome: nome, ligacoes: make(map[string]Aresta)}
	grafo := Grafo{Vertice: vertice, valorAcumulado: 0}
	return &grafo
}

func (grafo Grafo) AdicionarLigacao(nomeFilho string,
	filho Grafo, distancia int, valorAcumulado int) {
	caminho := Aresta{filho, distancia}
	grafo.ligacoes[nomeFilho] = caminho
	grafo.valorAcumulado = valorAcumulado
}

func (grafo Grafo) Imprimir() {
	//Não implementado. Te juro.
}
