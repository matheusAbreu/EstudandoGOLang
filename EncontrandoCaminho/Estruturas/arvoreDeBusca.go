package Estruturas

import (
	"fmt"
)

type ArvoreDeBusca struct {
	Vertice
	valorAcumulado int
}

func NovaArvore(nome string) *ArvoreDeBusca {
	vertice := Vertice{nome: nome, ligacoes: make(map[string]*Aresta)}
	arvore := ArvoreDeBusca{Vertice: vertice, valorAcumulado: 0}
	return &arvore
}

func (arvore ArvoreDeBusca) AdicionarLigacao(nomeFilho string,
	distancia int, valorAcumulado int) {

	filho := NovaArvore(nomeFilho)
	caminho := NovaAresta(filho, distancia)
	arvore.ligacoes[nomeFilho] = caminho
	arvore.valorAcumulado = valorAcumulado
}

func (arvore ArvoreDeBusca) Imprimir(resposta ArvoreDeBusca) {
	alturaAtual := 0
	solucao := "G"
	arvore.ImprimirNo(alturaAtual, solucao)
}

func (arvore ArvoreDeBusca) ImprimirNo(alturaAtual int, solucao string) {
	for index := 0; index < alturaAtual; index++ {
		fmt.Printf("| ")
	}
	if arvore.nome == solucao {
		fmt.Printf("|-+%s_[%b] <-- Resposta\n", arvore.nome, arvore.valorAcumulado)
	} else {
		fmt.Printf("|-+%s_[%b]\n", arvore.nome, arvore.valorAcumulado)
	}
	alturaAtual++
	for _, filho := range arvore.ligacoes {
		filho.vertice.ImprimirNo(alturaAtual, solucao)
	}
	alturaAtual--
}
