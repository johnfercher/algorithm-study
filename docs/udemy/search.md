# Search

## Linear Search
* `O(n)` em arrays ou listas

## Binary Search
* `O(log n)` em arrays

## Depth First Search (DFS)
* `O(n)` em árvores ou grafos
* Não utiliza memória
* Descobrir se um caminho existe
* Pode ficar lento

## Breadth First Search (BFS)
* `O(n)` em árvores ou grafos
* Utiliza memória
* Achar o menor caminho

## Perguntas e Respostas (DFS x BFS)
1. Se você sabe que uma solução não está longe da raiz
   * BFS
   * Pois o algoritmo procura os nós mais proximos antes de aprofundar na árvore
2. Se uma árvore é muito profunda e soluções são raras.
   * BFS
   * Se a solução existe, não tem o por que usar DFS nesse caso. DFS iria demorar muito nesse cenário.
3. Se uma árvore é muito larga
   * DFS
   * Não tem muitos ganhos entre DFS e BFS nesse cenário, por isso a vantagem de usar DFS, pois não gasta memória
4. Se soluções são frequentes mas localizadas no fundo da árvore
   * DFS
5. Determinar se existe um caminho entre dois nós
   * DFS
6. Encontrar o menor caminho
   * BFS

## Dijkstra x Bellman-Ford
* Possibilita encontrar o menor caminho em um grafo com pesos nas arestas.
* Dijkstra não suporta valores negativos, Bellman suporta
* Dijkstra é um pouco melhor do que Bellman em relação a tempo
* BFS e DFS não levam em consideração pesos nas arestas