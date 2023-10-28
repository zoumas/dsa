class Graph:
    def __init__(self, n):
        self.graph = []
        for i in range(n):
            self.graph.append([False] * n)

    def add_edge(self, u, v):
        self.graph[u][v] = True
        self.graph[v][u] = True

    def edge_exists(self, u, v):
        n = len(self.graph)
        if n == 0:
            return False
        if u < 0 or u >= n:
            return False
        if v < 0 or v >= n:
            return False
        return self.graph[u][v]

def test(n, edges_to_add, edges_to_test):
    graph = Graph(n)
    for edge in edges_to_add:
        graph.add_edge(edge[0], edge[1])
        print(f"Added edge: {edge}")
    print()

    for edge in edges_to_test:
        exists = graph.edge_exists(edge[0], edge[1])
        print(f"{edge} exists: {exists}")
    print()

def main():
    test(3, [(0, 1), (2, 0)], [(1, 0), (1, 2), (2, 0)])
    test(
        6, 
        [(0, 1), (1, 2), (2, 3), (3, 4), (4, 5)],
        [(0, 1), (1, 2), (0, 4), (2, 5), (5, 0)],
    )

if __name__ == "__main__":
    main()
