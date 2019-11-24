package EncontrandoPalindromo

import (
	"container/list"
	"strconv"
	str "strings"
)

type MeuTexto struct {
	conteudo string
	posEsp   *list.List //qnt de espaço: len(posEsp)

}

func CriaTexto() *MeuTexto {
	textoNovo := MeuTexto{conteudo: "", posEsp: list.New()}
	return &textoNovo
}
func (text *MeuTexto) LimpandoMeuTexto() {
	text.posEsp = nil
	text.conteudo = ""
}
func (text *MeuTexto) AddTexto(novoConteudo string) {
	text.conteudo = novoConteudo

	for i := 0; i < len(text.conteudo); i++ {
		if text.conteudo[i] == ' ' {
			text.posEsp.PushBack(i)
			//fmt.Println("tem um espaço aqui")
		}
	}
}
func (text MeuTexto) ImprimeMeuTexto() string {
	LocalEsp := "Temos " + strconv.Itoa(text.posEsp.Len()) + " espaços na Lista:{"
	for i := text.posEsp.Front(); i != nil; i = i.Next() {
		LocalEsp += strconv.Itoa(i.Value.(int)) + "; "
	}
	LocalEsp += "}"
	LocalEsp = str.Replace(LocalEsp, "; }", "}", -1)
	return text.conteudo + ". " + LocalEsp
}
func VerificandoCaracterEspecial(text string) (result bool) {
	result = false
	for i := 0; i < len(text); i++ {
		if text[i] == '.' || text[i] == ',' || text[i] == '?' ||
			text[i] == '!' || text[i] == ';' || text[i] == ':' ||
			text[i] == '[' || text[i] == ' ' || text[i] == '-' ||
			text[i] == '@' || text[i] == '#' || text[i] == ']' ||
			text[i] == '$' || text[i] == '%' || text[i] == '&' ||
			text[i] == '*' || text[i] == '(' || text[i] == ')' ||
			text[i] == '{' || text[i] == '}' || text[i] == '[' {
			result = true
		}
	}
	return
}

/*
void identificandoEspacos(MeuTexto *x)
{
    int temp[x->tam], i, qntEsp =0;

    for(i =0; i<x->tam;i++)
        if(x->vetor[i] == ' ')
            temp[qntEsp++] = i;

    free(x->posEsp);
    x->posEsp = (int*)malloc(qntEsp*sizeof(int));

    if(x->posEsp != NULL)
    {
        for(i=0;i < qntEsp;i++)
            x->posEsp[i] = temp[i];

        x->qntEsp = qntEsp ;
    }
    else
        printf("\nHouve um erro na alocacao - Funcao:criandoEspaco\n");


}
void escrevendoMeuTexto(MeuTexto *x)
{
    int tam,i;
    char temp[10000];

    gets(temp);
    LIMP;

    tam = strlen(temp);
    if(tam > 2)
    {
        limpandoMeuTexto(x);
        x->vetor = (char*)malloc((tam+1)*sizeof(char));

        if(x->vetor != NULL)
        {
            x->tam = tam;
            strcpy(x->vetor, temp);
            identificandoEspacos(x);
        }
        else
            printf("\nHouve um erro na alocacao - Funcao:escrevendoMeuTexto\n");
    }
    else
            printf("\nHouve um erro na alocacao - Funcao:escrevendoMeuTexto\n");
}
void copiandoMeuTexto(MeuTexto *dest, MeuTexto *ori, int ini, int fim)
{
    int i;

        limpandoMeuTexto(dest);
        dest->vetor = (char*)malloc(((fim-ini)+1)*sizeof(char));
        if(dest->vetor != NULL)
        {
            dest->tam = (fim-ini);
            for(i =ini; i<fim;i++)
                dest->vetor[i-ini] = ori->vetor[i];

            identificandoEspacos(dest);
        }
        else
            printf("\nHouve um erro na alocacao - Funcao:escrevendoMeuTexto\n");

}
void imprimindoMeuTexto(MeuTexto *x)
{
    int i;
    printf("\n\nTamanho da String: %d\nA String armazenada e: \"%s\"\nContendo %d espacos",x->tam, x->vetor,x->qntEsp);
    if(x->qntEsp > 0)
    {
        printf(", nas casas:");
        for(i = 0;i<x->qntEsp;i++)
            printf((i==(x->qntEsp-1))?(" %d.\n\n"):(" %d,"),x->posEsp[i]);
    }
}
int verificandoCaracterEspecial(char c)
{
    /*Retorna 1 se for um caracter especial
      Retorna 0 caso não seja
    if((c == '.' || c == ',' || c == '?' ||
        c == '!' || c == ';' || c == ':' ||
        c == '[' || c == ' ' || c == '-' ||
        c == '@' || c == '#' || c == ']' ||
        c == '$' || c == '%' || c == '&' ||
        c == '*' || c == '(' || c == ')' ||
        c == '{' || c == '}' || c == '[' ))
    {
        return 1;
    }
    else
        return 0;
}
int verificandoPalindromo(MeuTexto *x, int ini, int fim)
{
    /*Retorna 1 se for palindro
      Retorna 0 caso não seja
    int i,j=0,k;
    char temp[(fim-ini)+1];

    for(i=0; i<(fim-ini)+1;i++)
    {
        if(!verificandoCaracterEspecial(x->vetor[ini+i]))
            temp[i-j] = x->vetor[ini+i];
        else
            j++;
    }

    for(i=0,k=(fim-ini)-j; i<=k;i++,k--)
        if(temp[i] != temp[k])
        {return 0;}

    return 1;
}
MeuTexto *procurandoMaiorPalindro(MeuTexto *x)
{
    MeuTexto *pali = criandoMeuTexto();
    int maiorIntr[] = {0,0}, j, i,espI =0, espJ;

    for(i=0; i< x->tam;i++)
    {
        espJ =0;
        if( !verificandoCaracterEspecial(x->vetor[i]))
        {
            for(j = x->tam - 1; j > i; j--)
            {
               if( !verificandoCaracterEspecial(x->vetor[j]))
               {
                    if(x->vetor[i] == x->vetor[j] && (j-i) > (maiorIntr[1]-maiorIntr[0]))
                    {
                        if(verificandoPalindromo(x, i, j)==1)
                        {
                            copiandoMeuTexto(pali, x, i,j+1);
                            maiorIntr[0] = i;
                            maiorIntr[1] = j;
                            j = x->posEsp[(x->qntEsp-(++espJ))];
                        }
                    }
                    else if((x->qntEsp - espJ) >= 0)
                        j = x->posEsp[(x->qntEsp-(++espJ))];
               }
            }
            if(espI <= (x->qntEsp-1))
                i = x->posEsp[espI++];
        }
    }
    return pali;
}*/
