package Estruturas

type ArvoreDeBusca struct {
	vertice        *Vertice
	valorAcumulado int
}

func NovaArvore(nome string) *ArvoreDeBusca {
	verticeNovinho := NovoVertice(nome)
	arvore := ArvoreDeBusca{vertice: verticeNovinho, valorAcumulado: 0}
	return &arvore
}

func (arvore ArvoreDeBusca) AdicionarLigacao(nomeFilho string,
	filho ArvoreDeBusca, distancia int, valorAcumulado int) {
	//Usar funções do grafo.go
	//caminho := Aresta{filho, distancia}
	//arvore.ligacoes[nomeFilho] = caminho
	arvore.valorAcumulado = valorAcumulado
}

func (arvore ArvoreDeBusca) Imprimir() {
	//Não implementado. Te juro.
}
