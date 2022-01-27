# Big O

| Notação | Escalabilidade | Comum em |
|---|---|---|
| `O(1)`, `Constant` | Excelente | Acessar um elemento de um array, obter um elemento de um mapa. Sem loops |
| `O(log N)`| Excelente | Geralmente algoritmos de busca se os dados estiverem ordenados (binary search) (não é o caso de um hashmap) |
| `O(N)`, `Linear`| OK | Procurar um valor em uma lista (incremental) |
| `O(N log N)`| Ruim | Ordenar um array (merge sort, dividir e conquistar) |
| `O(N^2)`| Horrivel | Ordenar um array (insert sort, incremental). Todos os elementos precisam ser comparados |
| `O(2^N)`| Horrivel | Algoritmos recursivos que resolvem problemas N |
| `O(N!)`| Horrivel | Um loop para cada elemento |

* Não se diferencia `O(1)` e `O(42)`. Sempre é mostrado c `O(1)` não importando o valor da unidade. 

## Calculando Big O (Time Complexity)
  
3. Diferentes termos de entrada
4. Ignorar não dominantes

### Operações simples fora de iterações possuem complexidade `O(1)`
```golang
func fun(arr []int) {
	a := len(arr) // O(1)
}
```

### Iterações com base no input possuem complexidade `O(n)`
```golang
func fun(arr []int) {
	for _, value := range arr { // O(n) 
		// doing something 
	}
}
```

### Operações simples dentro de iterações com base no input possuem complexidade `O(n)`
```golang
func fun(arr []int) {
	for _, value := range arr { // O(n) 
		a := value % 2 // O(n)
	}
}
```

### Soma-se operações simples fora de iterações
```golang
func fun(arr []int) {
	a := len(arr) // O(1)
	b := b % 2 // O(1)
}
```

1. `Big O = O(1) + O(1)`
2. `Big O(2)`
3. `Big O(1)` (Simplifica para 1)

### Soma-se operações simples dentro de uma iteração com a própria iteração
```golang
func fun(arr []int) {
	for _, value := range arr { // O(n) 
		a := value % 2 // O(n)
	}
}
```

1. `Big O = O(n) + O(n)`
2. `Big O(2n)`
3. `Big O(n)` (Simplifica para n)

### Soma-se duas iterações sequenciais
```golang
func fun(arr []int) {
	for _, value := range arr { // O(n) 
		a := value % 2 // O(n)
	}
	
	for _, value := range arr { // O(n) 
		a := value % 2 // O(n)
	}
}
```

1. `Big O = [O(n) + O(n)] + [O(n) + O(n)]`
2. `Big O(4n)`
3. `Big O(n)` (Simplifica para n)

### Sempre se utiliza o pior caso
```golang
func fun(arr []int) {
	for _, value := range arr { // O(n)
		if value == 42 { // O(n) 
			break
		}   
	}
}
```

1. `Big O = O(n) + O(n)` (Embora o loop possa demorar menos que n, no pior cenário ele vai iterar em n)
2. `Big O(2n)`
3. `Big O(n)` (Simplifica para n)

### Sempre se ignora constantes
```golang
func fun(arr []int) {
	for _, value := range arr { // O(n) 
		a := value % 2 // O(n)
	}
	
	for i := 0 ; i < 100 ; i++ { // O(100) 
		a := value % 2 // O(100)
	}
}
``` 

1. `Big O = [O(n) + O(n)] + [O(100) + O(100)]`
2. `Big O(2n) + O(200)`
3. `Big O(2n)` (Ignora a constante)
4. `Big O(n)` (Simplifica para n)

### Iterações em diferentes entradas são representadas por variaveis diferentes
```golang
func fun(arr []int, arr2 []int) {
	for _, value := range arr { // O(n) 
		a := value % 2 // O(n)
	}
	
	for _, value := range arr2 { // O(m) 
		a := value % 2 // O(m) 
	}
}
```

1. `Big O = [O(n) + O(n)] + [O(m) + O(m)]`
2. `Big O(2n) + O(2m)`
3. `Big O(n + n)` (Simplifica para n)

### Iterações aninhadas devem ser multiplicadas
```golang
func fun(arr []int) {
	for _, value := range arr { // O(n) 
		for _, value := range arr { // O(n)
		}
	}
}
```

1. `Big O = O(n) * O(n)`
2. `Big O(n^2)` (Representando multiplicação como exponenciação)

### Iterações aninhadas devem ser multiplicadas. Caso venham de entradas diferentes, deve-se representar cada um por letras diferentes
```golang
func fun(arr []int, arr2 []int) {
	for _, value := range arr { // O(n) 
		for _, value := range arr2 { // O(m)
		}
	}
}
```

1. `Big O = O(n) * O(m)`
2. `Big O(n*m)` (Representando multiplicação como exponenciação)

### Deve-se pegar somente o BigO dominante
```golang
func fun(arr []int) {
	for _, value := range arr { // O(n) 
		for _, value := range arr { // O(n)
		}
	}
	
	for _, value := range arr { // O(n)
	}
}
```

1. `Big O = [O(n) * O(n)] + O(n)`
2. `Big O = O(n^2)+ O(n)`
3. `Big O(n^2)` (Utilizando somente n^2, pois o crescimento é muito maior do que n)

## Complexidade de Espaço

`BigO` pode ser utilizado para definir o custo de memória, em geral diferentes estruturas de dados podem possuir diferentes
custos de memória de de tempo. Em geral, há uma competição entre ambos: estruturas que ajudam na rapidez gastam mais memória e
estruturas que focam em ocupar menos espaço oneram a velocidade.

## Calculando Big O (Space Complexity)

### Variáveis consomem espaço
```golang
func fun(arr []int) {
	a := len(arr) // O(1)
}
```

### Estruturas de dados consomem espaço
```golang
func fun(qtd int) {
	arr := []int{}
	for i := 0 ; i < qtd ; i++ { 
		arr = append(arr, i) // O(n)
	}
}
```