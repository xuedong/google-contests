#!/usr/bin/env python3


if __name__ == "__main__":
    for test_case in range(1, int(input())+1):
        n = int(input())
        nums = sorted(map(int, input().split()), reverse=True)

        for idx, value in enumerate(nums):
            if value + idx < n:
                n = value + idx
                nums = nums[:n]

        print(f'Case #{test_case}: {len(nums)}')

