import itertools
from collections import defaultdict


def is_valid(block):
    m = len(block)
    if m == 1:
        return True

    counter = defaultdict(int)
    counter[block[0]] = 1
    for i in range(1, m):
        char = block[i]
        counter[char] += 1
        if counter[char] > 1 and char != block[i-1]:
            return False
    return True


def letter_blocks(blocks):
    candidates = list(itertools.permutations(blocks))
    candidates = [''.join(candidate) for candidate in candidates]

    for candidate in candidates:
        if is_valid(candidate):
            return candidate
    return "IMPOSSIBLE"


if __name__ == '__main__':
    t = int(input())

    for test_case in range(1, t + 1):
        n = int(input())
        blocks = input().split()
        output = letter_blocks(blocks)
        print("Case #" + str(test_case) + ": " + output)

