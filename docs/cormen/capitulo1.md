# 1. O Papel dos Algoritmos na Computação

## 1.1 Algoritmos

Um algoritmo é qualquer procedimento computacional que transforma valores de `entrada` em valores de `saída` com base em
uma sequência de procedimentos computacionais.

A `entrada` do problema computacional também é chamada de `instancia`.

O melhor algoritmo para resolução de um problema depende de muitos fatores: número de itens a ordenar, grau de ordenação
prévia, arquitetura de computador e do tipo de armazenamento (memória) utilizado.

Um algoritmo correto é aquele que resolve todas as instâncias de um problema da forma correta. Algoritmos incorretos podem
possuir saidas erradas ou podem não retornar.

Algoritmos incorretos podem ser uteis, caso a taxa de erro seja aceitável

É importante conhecer os problemas NP-Completos para evitar gastar muito tempo procurando uma solução correta. Nesse caso,
é melhor uma solução boa do que uma solução ótima.

## 1.2 Algoritmos como Tecnologia

Computadores não possuem velocidade e memoria infinita.

A parte váriavel do BigO é mais importante que a parte constante, assim para problemas suficientemente grandes sempre vale
mais a pena os algoritmos com o `N` menor.

* Algortimo A: `8*n^2`
* Algoritmo B: `64n`
* Quando A fica melhor que B?
* `8*n^2 <= 64n` (divide by `8n`)
* R: `n <= 8`

