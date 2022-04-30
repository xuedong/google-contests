def squary(numbers, k):
    if k == 1:
        s1 = 0
        s2 = 0
        for i in range(len(numbers)):
            s1 += numbers[i]
            s2 += numbers[i] * numbers[i]
        if s1 == 0:
            if s2 == 0:
                return str(0)
            return "IMPOSSIBLE"

        x = (s2 - s1 * s1) / (2.0 * s1)
        if x == int(x):
            return str(int(x))
        else:
            return "IMPOSSIBLE"


if __name__ == '__main__':
    t = int(input())

    for test_case in range(1, t + 1):
        n, k = list(map(int, input().split()))
        numbers = list(map(int, input().split()))
        output = squary(numbers, k)
        print("Case #" + str(test_case) + ": " + output)

