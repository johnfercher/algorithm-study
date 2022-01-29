# Linked List

| Type | PushBack | PushFront | PushAt | PopBack | PopFront | PopAt | Search | Annotation |
|---|---|---|---|---|---|---|---|---|
| Single With Head & Tail | `O(1)` | `O(1)` | `O(n)` | `O(1)` | `O(n)` | `O(n)` |`O(n)` | PushAt, PopAt e Search precisam de busca transversal. PopFront precisa de busca transversal pois não é possivel saber qual será o novo Tail. |
| Double With Head & Tail | `O(1)` | `O(1)` | `O(n/2)` | `O(1)` | `O(1)` | `O(n/2)` |`O(n/2)` | PushAt, PopAt e Search precisam de busca transvsal, porém, é possivel buscar e diferentes direções para indices menores ou maiores que a metade do tamanho. |