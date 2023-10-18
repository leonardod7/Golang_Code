
# Lista Encadeada


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

Para este estudo dirigido, estude os casos particulares de listas circulares simples encadeadas. O material deverá apresentar:

> ###### Um arquivo .md descrevendo uma lista circular encadeada, incluindo os seguintes algoritmos em pseudocódigo:
- > Exibição dos nós em uma lista circular;
- > Inserção de um nó no início da lista;
- > Exclusão do primeiro nó da lista;
  
<div id='resolucao'/>

## Resolução do Problema

Importante lembrar que, antes de entendermos a definição de uma Lista Circular Encadeada, precisamos entender as diferenças entre a alocação sequencial e alocação encadeada.
Na Alocação Sequencial, o acesso ao k-ésimo elemento de uma lista é imediato, onde cada nó é formado por campos, responsáveis por armazenar características de cada elemento da lista.
O uso de mais de duas listas de forma simultânea pode se tornar um problema de memória. Para resolver o problema, temos a aplicação da alocação encadeada, onde as posições da memória são alocadas e desalocadas conforme necessidade de uso. Na lista encadeada, os nós são intergligados por ponteiros que indicam a posição do próximo elemento. Todo o percurso da lista seá feito a partir desse ponteiro.
Na Lista Circular Encadeada, os elementos estão organizados de forma circular, similar a Lista Encadeada Simples. A maneira mais fácil e rápida de convertermos uma Lista Simples Encadeada em Lista Circular Simples Encadeada é fazer o próximo elemento do último elemento da lista apontar para o primeiro elemento da lista.
Existe um custo de performance do algoritmo em função da complexidade e, no caso de uma inserção ou remocão, teríamos que percorrer toda a lista para atualizarmos o próximo elemento do último elemento, que acaba sendo uma operação custosa.
Uma vantagem é que o último nó ao invés de ser null, fica apontado para o último elemento

Vejamos um exemplo em Pseudocódigo.

- > Exibição dos nós em uma lista circular;

```

func ExibicaoCircular (nó) {
    

    
}






```

- > Inserção de um nó no início da lista;


```






```

- > Exclusão do primeiro nó da lista;

```






```