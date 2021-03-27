#!/usr/bin/python3

import sys

def copyrights(string, x, y):
    '''
    ' Write your code here
    ' Return solution
    '''
    n = len(string)
    cost = 0
    prev = "?"
    for i in range(n):
        if string[i] == "C" and prev == "J":
            cost += y
            prev = "C"
        elif string[i] == "J" and prev == "C":
            cost += x
            prev = "J"
        elif string[i] == "?":
            continue
        elif string[i] != "?" and prev == "?":
            prev = string[i]

    return cost


if __name__ == '__main__':
    testCases = int(input())

    for testCase in range(1, testCases + 1):
        line = input().strip().split()
        x, y = int(line[0]), int(line[1])
        string = line[2]
        print("Case #" + str(testCase) + ": " + str(copyrights(string, x, y)))

