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


def combine_blocks(blocks):
    n = len(blocks)
    if n == 1:
        return blocks

    if n == 2:
        str1 = blocks[0]
        str2 = blocks[1]
        if str1[-1] == str2[0]:
            return [str1 + str2]
        elif str2[-1] == str1[0]:
            return [str2 + str1]
        else:
            return blocks

    next = combine_blocks(blocks[1:])
    curr = blocks[0]
    for i in range(len(next)):
        block = next[i]
        if curr[-1] == block[0]:
            next[i] = curr + block
            return combine_blocks(next)
            break
        elif block[-1] == curr[0]:
            next[i] = block + curr
            return combine_blocks(next)
            break
    next.append(curr)
    return next


def letter_blocks(blocks):
    new_blocks = combine_blocks(blocks)

    ans = ""

    for block in new_blocks:
        ans += block
        if not is_valid(ans):
            return "IMPOSSIBLE"

    return ans


if __name__ == '__main__':
    t = int(input())

    for test_case in range(1, t + 1):
        n = int(input())
        blocks = input().split()
        output = letter_blocks(blocks)
        print("Case #" + str(test_case) + ": " + output)

