#!/usr/bin/env python3


def find_cycle(banks):
    banks = list(banks)
    n = 0

    seen = {tuple(banks): 0}
    while True:
        i = 0
        for i_next, value in enumerate(banks):
            if value > banks[i]:
                i = i_next

        # Copy out the block count and reset the current bank.
        blocks_remaining, banks[i] = banks[i], 0
        while blocks_remaining > 0:
            i += 1
            if i == len(banks):
                i = 0
            banks[i] += 1
            blocks_remaining -= 1

        n += 1
        current_state = tuple(banks)
        if current_state in seen:
            return n, n - seen[current_state]
        seen[current_state] = n


def main():
    banks = None
    with open('input.txt') as fh:
        for line in fh:
            banks = [int(bank) for bank in line.rstrip().split('\t')]
            break

    print(find_cycle(banks))


if __name__ == '__main__':
    main()
