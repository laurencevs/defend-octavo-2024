from sys import argv

def parse_blocks(s):
    out = []
    for i, c in enumerate(s):
        out += [i//2 if i%2 == 0 else -1] * int(c)
    return out

def checksum_blocks(blocks):
    i1, i2 = 0, len(blocks)-1
    total = 0
    while i1 <= i2:
        if blocks[i1] == -1:
            total += blocks[i2] * i1
            i2 -= 1
            while i2 >= 0 and blocks[i2] == -1:
                i2 -= 1
        else:
            total += blocks[i1] * i1
        i1 += 1
    return total

def parse_files(s):
    return [[i, int(c), int(s[2*i+1]) if 2*i+1 < len(s) else 0] for i, c in enumerate(s[::2])]

def checksum_files(files):
    i = len(files)-1
    while i >= 0:
        for j in range(i):
            if files[j][2] >= files[i][1]:
                file_to_move = files.pop(i)
                files.insert(j+1, [file_to_move[0], file_to_move[1], files[j][2]-file_to_move[1]])
                files[j][2] = 0
                files[i][2] += file_to_move[1] + file_to_move[2]
                break
        else:
            i -= 1
    i = 0
    total = 0
    for file in files:
        total += file[0] * file[1] * (2*i + file[1] - 1)//2
        i += file[1] + file[2]
    return total

if __name__ == "__main__":
    with open(argv[1]) as f:
        data = f.read()
    print(checksum_blocks(parse_blocks(data)))
    print(checksum_files(parse_files(data)))
