def find_min_distance(distances, unvisited):
    closest = None
    for vertex in unvisited:
        if closest is None:
            closest = vertex
        elif distances[vertex] is not None and distances[vertex] < distances[closest]:
            closest = vertex
    return closest

def find_neighbours_of(source, unvisited):
    surroundings = [
        (source[0], source[1]+1),
        (source[0], source[1]-1),
        (source[0]+1, source[1]),
        (source[0]-1, source[1])
    ]
    return [neighbour for neighbour in surroundings if neighbour in unvisited]

def dijkstra(vertices, origin, destination):
    unvisited = []
    distances = {}
    previous_vertices = {}
    for vertex in vertices:  # (x, y): single-hop distance to enter this vertex
        distances[vertex] = None
        previous_vertices[vertex] = None
        unvisited.append(vertex)
    distances[origin] = 0
    while len(unvisited) > 0:
        current_vertex = find_min_distance(distances, unvisited)
        if current_vertex == destination:
            # done!
            print(f"Shortest path found is {distances[destination]}")
            break
        else:
            unvisited.remove(current_vertex)
            current_distance = distances[current_vertex]
            for neighbour in find_neighbours_of(current_vertex, unvisited):
                new_distance = current_distance + vertices[neighbour]
                if distances[neighbour] is None or new_distance < distances[neighbour]:
                    distances[neighbour] = new_distance
                    previous_vertices[neighbour] = current_vertex

if __name__ == '__main__':
    with open('input') as input_file:
        input_data = [line.rstrip('\n') for line in input_file]
    vertices = {}
    origin = None
    destination = None
    last_found = None
    for line_index,line in enumerate(input_data):
        for char_index,char in enumerate(line):
            vertex = (char_index, line_index)
            vertices[vertex] = int(char)
            if origin is None: origin = vertex
            last_found = vertex
    destination = last_found
    dijkstra(vertices, origin, destination)
