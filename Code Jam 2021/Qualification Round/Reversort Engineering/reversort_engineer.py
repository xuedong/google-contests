#!/usr/bin/python3

import sys

def reversort_c(n, c):
    if c < n-1:
        return None
    if c > (n+2)*(n-1)//2:
        return None

    def aux(start, end, cost):
        if start == end:
            return [start]
        if start + 1 == end:
            if cost == 1:
                return [start, end]
            elif cost == 2:
                return [end, start]

        low = end - start
        if cost >= 2 * low:
            rest = aux(start+1, end, cost-low-1)
            rest = rest[::-1]
            rest.append(start)
            return rest
        elif cost <= low:
            rest = aux(start+1, end, cost-1)
            rest.insert(0, start)
            return rest
        else:
            idx = cost - low
            left = [start+i for i in reversed(range(1, idx+1))]
            right = [start+idx+i for i in range(end-idx)]
            return left + [start] + right

    ans = aux(1, n, c)
    return ans


if __name__ == '__main__':
    testCases = int(input())

    for testCase in range(1, testCases + 1):
        n, c = list(map(int, input().strip().split()))
        ans = reversort_c(n, c)
        if ans == None:
            print("Case #" + str(testCase) + ": IMPOSSIBLE")
        else:
            print("Case #" + str(testCase) + ":", end=" ")
            for i in ans:
                print(i, end=" ")
            print()


