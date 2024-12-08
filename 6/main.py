import sys
from collections import defaultdict
from tqdm import tqdm

def parse_file(filename):
    with open(filename) as f:
        return [l.rstrip() for l in f.readlines()]

def gallivant(map, extra_obstruction=None):
    map_array = []
    guard_pos = None
    for y, line in enumerate(map):
        map_array.append(list(line))
        x = line.find("^")
        if x != -1:
            guard_pos = (y, x)
    if extra_obstruction:
        map_array[extra_obstruction[0]][extra_obstruction[1]] = "#"
    guard_direction = (-1, 0)
    height, width = len(map_array), len(map_array[0])
    visited_directed = defaultdict(bool)
    loop = False
    while 0 <= guard_pos[0] < height and 0 <= guard_pos[1] < width:
        if map_array[guard_pos[0]][guard_pos[1]] == "#":
            guard_pos = (guard_pos[0] - guard_direction[0], guard_pos[1] - guard_direction[1])
            guard_direction = (guard_direction[1], -guard_direction[0])
            guard_pos = (guard_pos[0] + guard_direction[0], guard_pos[1] + guard_direction[1])
            continue
        if visited_directed[(guard_pos, guard_direction)]:
            loop = True
            break
        map_array[guard_pos[0]][guard_pos[1]] = "X"
        visited_directed[(guard_pos, guard_direction)] = True
        guard_pos = (guard_pos[0] + guard_direction[0], guard_pos[1] + guard_direction[1])
    return sum(c == "X" for line in map_array for c in line), loop

def coverage(map):
    cov, _ = gallivant(map)
    return cov

def obstruction_options(map):
    count = 0
    for y in tqdm(range(len(map))):
        for x in range(len(map[y])):
            if map[y][x] == ".":
                _, loop = gallivant(map, (y, x))
                if loop:
                    count += 1
    return count

if __name__ == "__main__":
    map = parse_file(sys.argv[1])
    print(coverage(map))
    print(obstruction_options(map))
