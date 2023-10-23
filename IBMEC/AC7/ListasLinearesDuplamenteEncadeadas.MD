
# Lista Duplamente Encadeada


*******

Sumário Executivo
1. [Partes Envolvidas](#partesenvolvidas)
2. [Contexto do Problema](#contextoproblema)
3. [Resolução do Problema](#final)
4. [Referências Bibliográficas e Documentação Suporte de Apoio](#ref)

*******


<div id='partesenvolvidas'/> 

## Partes Envolvidas

#### Disciplina de Estrutura de Dados - IBMEC

Professor:

- Victor - Professor da Disciplina de Estrutura de Dados (IBMEC)

Aluno:

- Leonardo Costa;

<div id='contextoproblema'/> 

## Contexto do Problema

Para este estudo dirigido, estude os casos particulares de listas lineares duplamente encadeadas. O material deverá apresentar:


> ###### Um arquivo .md descrevendo uma lista linear duplamente encadeada com os seus nós ordenados, incluindo os seguintes algoritmos em pseudocódigo
- > Exibição dos nós em uma lista duplamente encadeadas;
- > Busca de um nó;
- > Inserção de um nó;
- > Remoção de um nó;

<div id='resolucao'/>

## Resolução do Problema

- > Exibição dos nós em uma Lista Duplamente Encadeada;


Obs.: nó_cabeça -> primeiro nó

``` 

func exibir(lista, valor) {
    var valor
    ponteiro := nó_cabeça
    enquanto ponteiro != valor:
        retorne ponteiro.lista[valor]
        ponteiro := ponteiro.proximo 
}

   
```

- > Busca de um nó;

```

Pseudocódigo: 

func busca(valor) {
    
    var valor
    var z
     
    anterior := nó_cabeça
    ponteiro := nó_cabeça.proximo
    enquanto ponteiro[index] != z:
        se ponteiro[index] < valor:
            anterior = ponteiro
            ponteiro = ponteiro.proximo
        caso contrário:
            ponteiro := z
}
```

- > Inserção de um nó em uma Lista Duplamente Encadeada;

``` 
se busca(valor, x) = z
    ponteiro.valor := valor
    ponteiro[index] := x
    ponteiro.próximo = nó_anterior.proximo
    nó_anterior.próximo := ponteiro

```

- > Remoção de um nó;

``` 
resultado := busca(valor)
se ponteiro != resultado
    nó_anterior := nó_seguinte
    valor := ponteiro.valor
    remove ponteiro

```