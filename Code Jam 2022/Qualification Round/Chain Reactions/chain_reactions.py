#!/usr/bin/env python3

from collections import defaultdict


def min_factors(n, funs, pointers):
    funs = [-1] + funs
    pointers = [-1] + pointers

    nodes = range(n+1)
    nodes = sorted(nodes, key=lambda x: pointers[x], reverse=True)

    nums_in = defaultdict(int)
    for pointer in pointers:
        nums_in[pointer] += 1

    ans = 0

    i = 0
    while i < len(nodes):
        relevant = nodes[i:i+nums_in[pointers[nodes[i]]]]

        if pointers[nodes[i]] == 0:
            return ans + sum([funs[i] for i in relevant])

        if nums_in[pointers[nodes[i]]] == 1:
            funs[pointers[nodes[i]]] = max(funs[nodes[i]], funs[pointers[nodes[i]]])
            i += 1
            continue

        minimum = min([funs[i] for i in relevant])
        ans += sum([funs[i] for i in relevant]) - minimum
        funs[pointers[nodes[i]]] = max(minimum, funs[pointers[nodes[i]]])
        i += len(relevant)

    return ans


if __name__ == '__main__':
    for t in range(1, int(input())+1):
        n = int(input())
        funs = list(map(int, input().split()))
        pointers = list(map(int, input().split()))

        ans = min_factors(n, funs, pointers)
        print(f'Case #{t}: {str(ans)}')

