package Estruturas

import (
	"strconv"
)

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

func InicializarGrafo(grafo map[string]*Vertice) {
	alfabeto, err := strconv.Atoi("A")
	if err == nil {
		for index := 0; index < 7; index++ {
			nome := strconv.Itoa(alfabeto)
			grafo[nome] = NovoVertice(nome)
			alfabeto++
		}
		DefinirCaminhoDoGrafo(grafo)
	}
}

func DefinirCaminhoDoGrafo(grafo map[string]*Vertice) {
	for nome, vertice := range grafo {
		switch {
		case nome == "A":
			vertice.AdicionarLigacao(grafo["B"], 9)
			vertice.AdicionarLigacao(grafo["C"], 5)
			vertice.AdicionarLigacao(grafo["D"], 13)
		case nome == "B":
			vertice.AdicionarLigacao(grafo["A"], 9)
			vertice.AdicionarLigacao(grafo["D"], 3)
			vertice.AdicionarLigacao(grafo["E"], 10)
		case nome == "C":
			vertice.AdicionarLigacao(grafo["A"], 5)
			vertice.AdicionarLigacao(grafo["F"], 12)
		case nome == "D":
			vertice.AdicionarLigacao(grafo["A"], 13)
			vertice.AdicionarLigacao(grafo["B"], 3)
			vertice.AdicionarLigacao(grafo["E"], 6)
			vertice.AdicionarLigacao(grafo["G"], 14)
		case nome == "E":
			vertice.AdicionarLigacao(grafo["B"], 10)
			vertice.AdicionarLigacao(grafo["D"], 6)
			vertice.AdicionarLigacao(grafo["G"], 7)
		case nome == "F":
			vertice.AdicionarLigacao(grafo["C"], 12)
			vertice.AdicionarLigacao(grafo["G"], 10)
		case nome == "G":
			vertice.AdicionarLigacao(grafo["D"], 14)
			vertice.AdicionarLigacao(grafo["E"], 7)
			vertice.AdicionarLigacao(grafo["F"], 10)
		}
	}
}

func (vertice Vertice) AdicionarLigacao(filho *Vertice, distancia int) {
	caminho := NovaAresta(filho, distancia)
	vertice.ligacoes[filho.nome] = caminho
}

func (vertice Vertice) Imprimir() {
	//Não implementado. Te juro.
}
