package EncontrandoCaminho

import (
	"container/list"

	"./Estruturas"
)

func ExpandirOrdenada(grafo *Estruturas.ArvoreDeBusca,
	listDeAbertos *list.List, listDeFechados *list.List) arvResultado *Estruturas.ArvoreDeBusca {
	arvResultado := Estruturas.NovaArvore("result")


	if &grafo.vertice.ligacoes == nil{

	}
	return 
}

/*
def expandir_ordenada(self, grafo, lista_de_abertos, lista_de_fechados):
	no_grafo = grafo[self.nome]

	if self.filhos == []:
		for no, custo in no_grafo.lista_de_ligacoes:
			if self.verifica_se_esta_na_lista(no.nome, lista_de_abertos):
				self.expandir_no_ordenado(no, custo, lista_de_abertos, lista_de_fechados)
			else:
				self.expandir_no(no, custo, [], lista_de_fechados)

def expandir_no_ordenado(self, no_grafo, custo, lista_de_abertos, lista_de_fechados):
	for no_aberto in lista_de_abertos:
		if no_grafo.nome == no_aberto.nome:
			custo_parcial_novo_no = self.custo_parcial + custo
			novo_no = ArvoreBusca(no_grafo.nome, [], custo_parcial_novo_no)
			item_a_ser_fechado = (
					self.mantem_item_de_menor_custo_e_retorna_o_de_maior_custo(
						novo_no, no_aberto, lista_de_abertos)
				)
			lista_de_fechados.append(item_a_ser_fechado)

def mantem_item_de_menor_custo_e_retorna_o_de_maior_custo(self, novo_no, no_antigo, lista):
	if novo_no.custo_parcial < no_antigo.custo_parcial:
		no_com_maior_custo = lista.remove(no_antigo)
		lista.append(novo_no)
		self.filhos.append(novo_no)
	else:
		no_com_maior_custo = novo_no
		ja_existe = self.verifica_se_esta_na_lista(novo_no.nome, self.filhos)
		if ja_existe == False:
			self.filhos.append(novo_no)
	return no_com_maior_custo*/
