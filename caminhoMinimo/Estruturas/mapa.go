package Estruturas

type Vertice struct {
	nome     string
	ligacoes map[string]Aresta
}

type Aresta struct {
	cidade    interface{}
	distancia int
}

func NovaCidade(nome string) *Vertice
{
	cidade := Vertice{nome: nome, ligacoes: make(map[string]Aresta)}
	return &cidade
}

func (cidade Vertice) AdicionarLigacao(nomeCidadeVizinha string,
	 cidadeVizinha Vertice, distancia int) 
{
	caminho := Aresta{cidadeVizinha, distancia}
	cidade.ligacoes[nomeCidadeVizinha] = caminho
}

func (cidade Vertice) Imprimir() 
{
	//Não implementado. Te juro.
}
