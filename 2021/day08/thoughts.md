# Thoughts

* Signals in the data are listed randomly, both within the signal data and
  each signal pattern, but there's no significance to this.
* It's probably about as easy to solve the general problem as the 'simple'
  problem of figuring out what 1, 4, 7, and 8 are.

## Solving the general problem

### Exploitable commonalities?

Let's see if any digits have unique characteristics:

One:
* Only one that uses two segments
Two:
* ?
Three:
* Only five-segment that overlaps with 9
Four:
* Only one that uses four segments
Five:
* Common segments between 6 and 9
* Only five-segment that overlaps with 6
* Only five-segment that overlaps with 4
Six:
* ?
Seven:
* Only one the uses three segments
Eight:
* Only one that uses seven segments
Nine
* Only six-segment number that overlaps with 4
Zero
* ?

This doesn't seem like a practical route.

### Filtering

| Segments | Digits  |
| -------- | ------- |
| 2        | 1       |
| 3        | 7       |
| 4        | 4       |
| 5        | 2, 3, 5 |
| 6        | 0, 6, 9 |
| 7        | 8       |

Layout:

```
 aaaa
b    c
b    c
 dddd
e    f
e    f
 gggg
```

Let's take the example:

```
Input: acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf
Segs:  7       5     5     5     3   6      6      4    6      2
Cands: 8       2,3,5 2,3,5 2,3,5 7   0,6,9  0,6,9  4    0,6,9  1
```

This implies:

```
a -> 
