from collections import defaultdict, deque

CAVE_MAP = defaultdict(list)

def find_paths_from(cave_name, paths, visit_counter, second_visit_available):
    paths.append(cave_name)
    visit_counter[cave_name] += 1
    for adjacent_cave_name in CAVE_MAP[cave_name]:
        if adjacent_cave_name == 'start':
            continue
        if adjacent_cave_name == 'end':
            visit_counter[adjacent_cave_name] += 1
        elif adjacent_cave_name.isupper() or visit_counter[adjacent_cave_name] < 1:
            find_paths_from(adjacent_cave_name, paths, visit_counter, second_visit_available)
        elif second_visit_available:
            find_paths_from(adjacent_cave_name, paths, visit_counter, False)
    paths.pop()
    visit_counter[cave_name] -= 1
    return visit_counter['end']

if __name__ == "__main__":
    with open('input', 'r') as input_file:
        for line in input_file:
            cave1_name, cave2_name = line.rstrip().split('-')
            CAVE_MAP[cave1_name].append(cave2_name)
            CAVE_MAP[cave2_name].append(cave1_name)
    paths = deque()
    visit_counter = defaultdict(int)  # cave_name: visit_count
    paths_found = find_paths_from('start', paths, visit_counter, True)
    print(f"Paths found to get us out of here: {paths_found}")
