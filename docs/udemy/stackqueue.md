# Stacks & Queues

* Versões limitadas de Linked List com intuito de serem simples e rapidas;

## Stacks
* LIFO (Last In First Out)
* Não há a necessidade de manter um `tail` computado, pois não existem operações feitas direto no primeiro elemento
* Implementação com Array.
    1. Mais rapido: os valores de um array são alocados próximos uns aos outros, isso possibilidade o uso de cache pelo computador.
    2. Mais leve: não há necessidade de implementar a struct node
    3. Menor limite de elementos: por estar agrupado na memoria
* Implementação com Linked List.
    1. Não há beneficio de se utilizar double linked list, pois a implementação só trabalha em uma direção
    2. Mais lento: os valores são distribuidos aleatoriamente na memoria
    3. Mais pesado: há a necessidade de implementar a struct node
    4. Maior limite de elementos: por estar espalhado na memoria

## Queues
* FIFO (First In First Out)
* Nunca se usa arrays, pois a operação de `dequeue()` é `O(n)`;
* A implementação utilizando Single Linked List (mapeando o primeiro) faz com que a operação de `enqueue` seja `O(n)`, pois para encontrar o ultimo da filha é necessário percorrer toda a fila
* A implementação utilizando Double Linked List (mapeando o primeiro e o segundo), faz com que a operação de `enqueue` seja `O(n)`, pois sabesse sempre o ultimo da fila