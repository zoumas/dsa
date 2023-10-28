class Graph:
    def __init__(self):
        self.graph = {}

    def add_edge(self, u, v):
        if u in self.graph:
            self.graph[u].add(v)
        else:
            self.graph[u] = {v}
        if v in self.graph:
            self.graph[v].add(u)
        else:
            self.graph[v] = {u}

    def adjacent_nodes(self, node):
        return list(self.graph[node])

def test(edges_to_add, nodes_to_test):
    graph = Graph()

    for edge in edges_to_add:
        graph.add_edge(edge[0], edge[1])
        print(f"Added edge: {edge}")
    print()

    for node in nodes_to_test:
        nodes = graph.adjacent_nodes(node)
        print(f"Adjacent nodes of {node}: {sorted(nodes)}")
    print()

def main():
    test(
        [
            (0, 1),
            (1, 2),
            (2, 3),
            (3, 4),
            (4, 5),
        ],
        [0, 1, 2, 3, 4, 5],
    )
    test(
        [
            (0, 1),
            (0, 2),
            (0, 3),
            (1, 2),
            (1, 3),
        ],
        [0, 1, 2, 3],
    )

if __name__ == "__main__":
    main()
