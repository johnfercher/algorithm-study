# Big O

| Notação | Escalabilidade | Comum em |
|---|---|---|
| `O(1)`, `Constant` | Excelente | Acessar um elemento de um array, obter um elemento de um mapa |
| `O(log N)`| Excelente | ??? |
| `O(N)`, `Linear`| OK | Procurar um valor em uma lista (incremental) |
| `O(N log N)`| Ruim | Ordenar um array (merge sort, dividir e conquistar) |
| `O(N^2)`| Horrivel | Ordenar um array (insert sort, incremental) |
| `O(2^N)`| Horrivel | ??? |
| `O(N!)`| Horrivel | ??? |

* Não se diferencia `O(1)` e `O(42)`. Sempre é mostrado c `O(1)` não importando o valor da unidade. 