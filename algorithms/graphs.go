package main

// T  comparable - имя ноды может быть только типом, поддерживающим операции сравнения
type Node[T comparable] struct {
	// Имя вершины
	Name T
	// Флаг - 1 - помечена, 0 - не помечена
	Mark byte
	// Степень вершины.
	Power int
}

// AbstractGraph Абстрактное представление графа. Граф задан списком смежности в виде отображения (map) вершин. Вершины заданы структурами Node.
type AbstractGraph[T comparable] struct {
	Graph    map[*Node[T]]map[*Node[T]]int
	Vertexes []*Node[T]
}

// New - создаёт новую структуру AbstractGraph из переданной map.
func New[T comparable](graph map[T]map[T]int) *AbstractGraph[T] {
	output := make(map[*Node[T]]map[*Node[T]]int, len(graph))
	vertexes := make([]*Node[T], len(graph))
	g := &AbstractGraph[T]{Graph: output}
	i := 0
	for vert := range graph {
		n := &Node[T]{Name: vert, Mark: 0, Power: len(graph[vert])}
		vertexes[i] = n
		i++
	}
	for vert, list := range graph {
		var parentNode *Node[T]
		childList := make(map[*Node[T]]int, len(list))
		for _, v := range vertexes {
			if v.Name == vert {
				parentNode = v
				g.Graph[v] = childList
				break
			}
		}
		for vertex, weight := range list {
			for _, n := range vertexes {
				if n.Name == vertex {
					childList[n] = weight
				}
			}
		}
		g.Graph[parentNode] = childList
	}
	return g
}

// DFS Поиск в глубину.
func (g *AbstractGraph[T]) DFS(start T, compare func(want T) bool) ([]*Node[T], bool) {
	var searchStack []*Node[T]
	for vert := range g.Graph {
		if vert.Name == start {
			searchStack = append(searchStack, vert)
			break
		}
	}
	for len(searchStack) != 0 {
		var vertex = searchStack[len(searchStack)-1]
		searchStack = searchStack[:len(searchStack)-1]
		if vertex.Mark != 1 {
			if compare(vertex.Name) {
				g.Unmark()
				g.Vertexes = append(g.Vertexes, vertex)
				return g.Vertexes, true
			}
			vertex.Mark = 1
			g.Vertexes = append(g.Vertexes, vertex)
			searchStack = append(searchStack, g.GetAdjacentVertices(vertex)...)
		}
	}
	g.Unmark()
	return nil, false
}

// BFS Поиск в ширину
func (g *AbstractGraph[T]) BFS(start T, compare func(want T) bool) ([]*Node[T], bool) {
	var searchQueue []*Node[T]
	for vert := range g.Graph {
		if vert.Name == start {
			searchQueue = append(searchQueue, vert)
			break
		}
	}
	for len(searchQueue) != 0 {
		var vertex = searchQueue[0]
		searchQueue = searchQueue[1:]
		if vertex.Mark != 1 {
			if compare(vertex.Name) {
				g.Unmark()
				g.Vertexes = append(g.Vertexes, vertex)
				return g.Vertexes, true
			}
			vertex.Mark = 1
			g.Vertexes = append(g.Vertexes, vertex)
			searchQueue = append(searchQueue, g.GetAdjacentVertices(vertex)...)
		}

	}
	g.Unmark()
	return nil, false
}

func (g *AbstractGraph[T]) GetAdjacentVertices(n *Node[T]) []*Node[T] {
	keys := make([]*Node[T], len(g.Graph[n]))
	i := 0
	for k := range g.Graph[n] {
		keys[i] = k
		i++
	}

	return keys
}

// Unmark - устанавливает все пройденные вершины в изначальное состояние.
func (g *AbstractGraph[T]) Unmark() {
	for _, v := range g.Vertexes {
		v.Mark = 0
	}
}

func main() {

}
