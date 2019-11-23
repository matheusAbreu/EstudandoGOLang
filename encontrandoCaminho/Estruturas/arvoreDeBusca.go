package Estruturas

type ArvoreDeBusca struct {
	Vertice
	valorAcumulado int
}

func NovaArvore(nome string) *ArvoreDeBusca {
	vertice := Vertice{nome: nome, ligacoes: make(map[string]Aresta)}
	arvore := ArvoreDeBusca{Vertice: vertice, valorAcumulado: 0}
	return &arvore
}

func (arvore ArvoreDeBusca) AdicionarLigacao(nomeFilho string,
	filho ArvoreDeBusca, distancia int, valorAcumulado int) {

	caminho := Aresta{filho, distancia}
	arvore.ligacoes[nomeFilho] = caminho
	arvore.valorAcumulado = valorAcumulado
}

func (arvore ArvoreDeBusca) Imprimir() {
	//Não implementado. Te juro.
}
