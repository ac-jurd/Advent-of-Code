from typing import List

def is_linear(l: List[int]) -> bool:
    increasing = l[0] < l[1]

    for i in range(1, len(l)-1):
        if (l[i] < l[i+1]) != increasing:
            return False
    return True

def is_valid_adj(l: List[int]) -> bool:
    for i in range(len(l)-1):
        diff = abs(l[i] - l[i+1])
        if diff < 1 or diff > 3:
            return False
    return True

def is_line_valid(l: List[int]) -> bool:
    return is_linear(l) and is_valid_adj(l)

def process_without_levels(l: List[int]) -> bool:
    if is_line_valid(l):
        return True

    for i in range(len(l)):
        copy = l[:i] + l[i+1:]
        if is_line_valid(copy):
            return True
    return False

def main():
    c = 0
    with open('input.txt', 'r') as f:
        while True:
            try:
                l = f.readline()
                spl = l.split(' ')
                spl = [int(i) for i in spl]
                if process_without_levels(spl):
                    c += 1
            except:
                break

    print(c)

main()
