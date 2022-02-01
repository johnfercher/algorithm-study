## Binary Search Tree

* Um nó pode apontar somente para esquerda e direita;
  * A direita são os valores maior que o nó;
  * A esquerda são os valores menores que o nó;
* Um nó tem somente um pai;
* Ordenado;
* Uma das mais equilibradas estruturas de dados em tempo de execução para operações de CRUD; `O(log n)`
* Necessário manter a árvore balanceada, pois se não vira uma Linked List; `O(n)`
* Operação de remove é complexa e é necessário re-computar as sub árvores
* AVL Tree e Red Black Tree se balanceiam automaticamente.
  * AVL: https://medium.com/basecs/the-little-avl-tree-that-could-86a3cae410c7
  * Red-Black: https://medium.com/basecs/painting-nodes-black-with-red-black-trees-60eacb2be9a5

## Binary Heap
* Um nó pode apontar somente para esquerda a direita;
  * Ambos direita e esquerda são menores que o valor do nó;
* Um nó tem somente um pai;
* Desordenado;
* Não há necessidade de balancear a árbore
* Utilizado em Priority Queues
* Busca é `O(n)`
* Inserção é rapida


## Trie
* Um nó aponta para um array
* Armazena letras
* Utilizado para complemento de texto em buscas