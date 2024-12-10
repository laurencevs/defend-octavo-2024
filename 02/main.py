import sys

def parse_file(filename):
    with open(filename) as f:
        lines = f.readlines()
    return [[int(x) for x in l.split(" ")] for l in lines]

def is_safe(report):
    inc = report[1] > report[0]
    prev = report[0]
    for v in report[1:]:
        if (v >= prev and not inc) or (v <= prev and inc):
            return False
        if abs(v - prev) > 3:
            return False
        prev = v
    return True

def is_safe_2(report):
    for i in range(len(report)):
        if is_safe(report[:i] + report[i+1:]):
            return True
    return False

if __name__ == "__main__":
    data_filename = sys.argv[1]
    print(sum(is_safe(r) for r in parse_file(data_filename)))
    print(sum(is_safe_2(r) for r in parse_file(data_filename)))
