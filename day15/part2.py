def find_min_distance(ready_vertices):
    closest = None
    closest_distance = None
    for vertex,distance in ready_vertices.items():
        # all ready_vertices have a preliminary distance set. it's what makes them 'ready'
        if closest is None or distance < closest_distance:
            closest = vertex
            closest_distance = distance
    return closest

def find_neighbours_of(source, ready_vertices, unchecked_vertices):
    surroundings = [
        (source[0],   source[1]+1),
        (source[0],   source[1]-1),
        (source[0]+1, source[1]  ),
        (source[0]-1, source[1]  )
    ]
    neighbours = {}
    for potential_neighbour in surroundings:
        if potential_neighbour in ready_vertices:
            neighbours[potential_neighbour] = ready_vertices[potential_neighbour]
        elif potential_neighbour in unchecked_vertices:
            neighbours[potential_neighbour] = None
    return neighbours

def dijkstra(vertices, origin, destination):
    unchecked_vertices = set()
    ready_vertices = {}
    final_distances = {}
    previous_vertices = {}
    for vertex in vertices:  # (x, y): single-hop distance to enter this vertex
        final_distances[vertex] = None
        previous_vertices[vertex] = None
        unchecked_vertices.add(vertex)
    ready_vertices[origin] = 0
    while len(ready_vertices) > 0:
        # print(f"Ready vertices: {len(ready_vertices)} | Unchecked vertices: {len(unchecked_vertices)}")
        current_vertex = find_min_distance(ready_vertices)
        if current_vertex == destination:
            # done!
            print(f"Shortest path found is {ready_vertices[destination]}")
            break
        else:
            current_distance = ready_vertices.pop(current_vertex)
            final_distances[current_vertex] = current_distance
            for neighbour,distance in find_neighbours_of(current_vertex, ready_vertices, unchecked_vertices).items():
                new_distance = current_distance + vertices[neighbour]
                if distance is None or new_distance < distance:
                    ready_vertices[neighbour] = new_distance
                    unchecked_vertices.discard(neighbour)
                    previous_vertices[neighbour] = current_vertex

if __name__ == '__main__':
    with open('input') as input_file:
        input_data = [line.rstrip('\n') for line in input_file]
    vertices = {}
    origin = None
    destination = None
    last_found = None
    line_length = len(input_data[0])
    line_count = len(input_data)
    for line_index,line in enumerate(input_data):
        for char_index,char in enumerate(line):
            for line_dupe in range(5):
                for char_dupe in range(5):
                    vertex = (char_index + (char_dupe * line_length), line_index + (line_dupe * line_count))
                    vertices[vertex] = ((int(char) + char_dupe + line_dupe) % 9) or 9
                    if origin is None: origin = vertex
                    last_found = vertex
    destination = last_found
    dijkstra(vertices, origin, destination)
