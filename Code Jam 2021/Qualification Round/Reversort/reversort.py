#!/usr/bin/python3

import sys

def reversort(arr):
    '''
    ' Write your code here
    ' Return solution
    '''
    n = len(arr)
    count = 0
    for i in range(n-1):
        j = arr[i:].index((min(arr[i:])))
        arr[i:(j+i+1)] = reversed(arr[i:(j+i+1)])
        count += j+1
    return count


if __name__ == '__main__':
    testCases = int(input())

    for testCase in range(1, testCases + 1):
        n = int(input())
        arr = list(map(int, input().strip().split()))
        print("Case #" + str(testCase) + ": " + str(reversort(arr)))

