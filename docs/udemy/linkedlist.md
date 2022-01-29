# Linked List

| Type | PushBack | PushFront | PushAt | PopBack | PopFront | PopAt | Search By Index | Annotation |
|---|---|---|---|---|---|---|---|---|
| Array | `O(1)` | `O(n)` | `O(n)` | `O(1)` | `O(n)` | `O(n)` |`O(1)` | Iteração mais rapida e leve que Linked list, por estar agrupado na memoria (cache) e por não precisar de um objeto node.  |
| Single With Head & Tail | `O(1)` | `O(1)` | `O(n)` | `O(1)` | `O(n)` | `O(n)` |`O(n)` | PushAt, PopAt e Search precisam de busca transversal. PopFront precisa de busca transversal pois não é possivel saber qual será o novo Tail. |
| Double With Head & Tail | `O(1)` | `O(1)` | `O(n/2)` | `O(1)` | `O(1)` | `O(n/2)` |`O(n/2)` | PushAt, PopAt e Search precisam de busca transvsal, porém, é possivel buscar e diferentes direções para indices menores ou maiores que a metade do tamanho. |

* Single Linked List With Tail: Faz com que a operação ambas a operação de PushBack e PushFront sejam `O(n)`
* Single Linked List Without Tail: Para adicionar um elemento no final a operação vira `O(n)`
* Double Linked List With Tail: Faz com que todas as operações (push, pop) nas extremidades sejam `O(n)`. 