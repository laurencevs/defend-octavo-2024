import sys

def parse_file(filename):
    with open(filename) as f:
        return [l.rstrip() for l in f.readlines()]

def check_for_word_directed(lines, word, i, j, yOffset, xOffset, reverse):
    for k in range(len(word)):
        if lines[i + yOffset*k][j + xOffset*k] != (word[-k-1] if reverse else word[k]):
            return False
    return True

def check_for_word(lines, word, i, j, yOffset, xOffset):
    return check_for_word_directed(lines, word, i, j, yOffset, xOffset, False) or \
        check_for_word_directed(lines, word, i, j, yOffset, xOffset, True)

def count_word(lines, word):
    count = 0
    for i in range(len(lines)):
        for j in range(len(lines[0])-len(word)+1):
            count += check_for_word(lines, word, i, j, 0, 1)
    for i in range(len(lines)-len(word)+1):
        for j in range(len(lines[0])):
            count += check_for_word(lines, word, i, j, 1, 0)
    for i in range(len(lines)-len(word)+1):
        for j in range(len(lines[0])-len(word)+1):
            count += check_for_word(lines, word, i, j, 1, 1)
            count += check_for_word(lines, word, i+len(word)-1, j, -1, 1)
    return count

def count_word_x(lines, word):
    count = 0
    for i in range(len(lines)-len(word)+1):
        for j in range(len(lines[0])-len(word)+1):
            count += check_for_word(lines, word, i, j, 1, 1) and check_for_word(lines, word, i+len(word)-1, j, -1, 1)
    return count

if __name__ == "__main__":
    lines = parse_file(sys.argv[1])
    print(count_word(lines, "XMAS"))
    print(count_word_x(lines, "MAS"))
