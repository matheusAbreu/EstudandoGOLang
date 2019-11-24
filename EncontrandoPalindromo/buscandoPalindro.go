package EncontrandoPalindromo

import (
	"container/list"
	"fmt"
	"strconv"
	str "strings"
)

type MeuTexto struct {
	conteudo string
	posEsp   *list.List //qnt de espaço: len(posEsp)
}

func (text *MeuTexto) GetConteudo() string {
	return text.conteudo
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
		}
	}
}
func (text *MeuTexto) VerificaExistePosEsp(valor int) bool {
	for i := text.posEsp.Front(); i != nil; i = i.Next() {
		if i.Value.(int) == valor {
			return true
		}
	}
	return false
}
func (text MeuTexto) ImprimeMeuTexto() string {
	LocalEsp := "Temos " + strconv.Itoa(text.posEsp.Len()) + " espaços no texto:{"
	for i := text.posEsp.Front(); i != nil; i = i.Next() {
		LocalEsp += strconv.Itoa(i.Value.(int)) + "; "
	}
	LocalEsp += "}"
	LocalEsp = str.Replace(LocalEsp, "; }", "}", -1)
	return text.conteudo + ". " + LocalEsp
}
func VerificandoCaracterEspecial(text string) bool {

	for i := 0; i < len(text); i++ {
		if text[i] == '.' || text[i] == ',' || text[i] == '?' ||
			text[i] == '!' || text[i] == ';' || text[i] == ':' ||
			text[i] == '[' || text[i] == ']' || text[i] == '-' ||
			text[i] == '@' || text[i] == '#' || text[i] == '}' ||
			text[i] == '$' || text[i] == '%' || text[i] == '&' ||
			text[i] == '*' || text[i] == '(' || text[i] == ')' ||
			text[i] == '{' {
			return true
		}
	}
	return false
}

func (text *MeuTexto) VerificandoPalindromo() bool {
	k := len(text.conteudo)

	if !VerificandoCaracterEspecial(text.conteudo) && k > 0 {
		for i := 0; i <= k; i++ {
			k--
			if text.VerificaExistePosEsp(i) {
				i++
			}
			if text.VerificaExistePosEsp(k) {
				k--
			}
			if text.conteudo[i] != text.conteudo[k] {
				return false
			}
		}
	} else {
		fmt.Println("Existem caracters especiais, não posso buscar um palindromo")
		return false
	}
	return true
} /*
func (text *MeuTexto) ProcurandoMaiorPalindro(x string){
   // MeuTexto *pali = criandoMeuTexto();
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
}

/*
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
