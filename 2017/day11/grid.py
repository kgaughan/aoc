#!/usr/bin/env python3


def expand(n, pos, neg):
    if n > 0:
        return [pos] * n
    if n < 0:
        return [neg] * -n
    return []


def simplify(directions):
    x = 0  # N/S
    y = 0  # NE/SW
    z = 0  # NW/SE

    # Convert to axes:
    for direction in directions:
        if direction == 'n':
            x += 1
        elif direction == 's':
            x -= 1
        elif direction == 'ne':
            y += 1
        elif direction == 'sw':
            y -= 1
        elif direction == 'nw':
            z += 1
        elif direction == 'se':
            z -= 1

    result = expand(x, 'n', 's') + \
        expand(y, 'ne', 'sw') + \
        expand(z, 'nw', 'se')

    return result
