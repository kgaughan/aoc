# Notes on AoC 2024

## Day 20

An alternative method: store all the points in an array, ordered by their
distance from the start, then scan only the points afterwards, checking their
Manhattan distance. That's O(n*n), but the aggregate real cose might be lower
than what I'm doing here.

Apparently [quadtrees](https://en.wikipedia.org/wiki/Quadtree) or [k-d
trees](https://en.wikipedia.org/wiki/K-d_tree) would also work here, but I'm
not familiar with either, because I've never done anything with them. Two more
things on the learning pile, I guess...

A whole bunch of people look to have used Dijkstra's algorithm to trace the
path, but that's frankly overkill for what's we already know is a straight
line!

Another thing is that I'm effectively scanning pairs of points twice! However,
I'm not super sure of how to remove points from consideration later. This is
something I'd need to have a think about and would likely half the runtime.

Finally, rather than using IntPairMap, I could use Hashtbl.
