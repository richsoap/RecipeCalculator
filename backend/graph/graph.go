package graph

type Graph struct {
	Nodes           map[uint64]*Node
	Degree          map[uint64]int64
	ZeroDegreeNodes map[uint64]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes:           make(map[uint64]*Node),
		Degree:          make(map[uint64]int64),
		ZeroDegreeNodes: make(map[uint64]*Node),
	}
}

func (g *Graph) IsExist(ID uint64) bool {
	_, ok := g.Nodes[ID]
	return ok
}

func (g *Graph) AddNode(node *Node) {
	g.Nodes[node.ID] = node
	if _, exist := g.Degree[node.ID]; !exist {
		g.Degree[node.ID] = 0
	}
	if g.Degree[node.ID] == 0 {
		g.ZeroDegreeNodes[node.ID] = node
	}
	for child := range node.Depends {
		if degree, exist := g.Degree[child]; exist {
			g.Degree[child] = degree + 1
		} else {
			g.Degree[child] = 1
		}
	}
}

func (g *Graph) RemoveNode(node *Node) {
	delete(g.ZeroDegreeNodes, node.ID)
	delete(g.Degree, node.ID)
	if _, exist := g.Nodes[node.ID]; exist {
		for child := range node.Depends {
			g.Degree[child]--
			if childNode, nodeExist := g.Nodes[child]; g.Degree[child] == 0 && nodeExist {
				g.ZeroDegreeNodes[child] = childNode
			}
		}
	}
	delete(g.Nodes, node.ID)
}

func (g *Graph) GetZeroDegreeNode() *Node {
	for _, node := range g.ZeroDegreeNodes {
		return node
	}
	return nil
}

func (g *Graph) AdjustAmount(id uint64, amount int) {
	if node, exist := g.Nodes[id]; exist {
		node.Amount += amount
	}
}

func (g *Graph) ToOrderedNodes() OrderedNodes {
	result := make(OrderedNodes, 0, len(g.Nodes))
	for _, node := range g.Nodes {
		result = append(result, node)
	}
	return result
}

func (g *Graph) IsEmpty() bool {
	return len(g.Nodes) == 0
}
