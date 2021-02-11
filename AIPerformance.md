
# AI performance results

All of this AI benchmarks used the same datasets `monica_five.txt`

```
   0 1 2  3 4 5
  ┌────────────┐
0 |1|0|0||0|0|0|
1 |0|1|0||0|2|0|
2 |0|0|0||0|2|0|
  |────────────|
3 |0|0|0||0|0|0|
4 |0|2|0||0|1|0|
5 |0|0|0||0|0|0|
  └────────────┘
```

## Without optimization

- Depth 2 = 135ms
- Depth 3 = 32s
- Depth 4 = Too long

## With Alpha Beta pruning

- Depth 2 = 147ms
- Depth 3 = 7s
- Depth 4 = 9min

## With Randomize + Alpha Beta Pruning

- Depth 2 = ~140-150ms
- Depth 3 = ~4-6s
- Depth 4 = ~3-4min

> The game state configuration force players to move on the center of the board. Is it always true ?

Let's go to check with others configurations !

### Line straight configuration

```
   0 1 2  3 4 5
  ┌────────────┐
0 |0|1|0||0|0|0|
1 |0|0|0||2|2|0|
2 |0|0|2||0|0|0|
  |────────────|
3 |0|0|0||0|0|0|
4 |0|2|0||0|0|0|
5 |0|1|1||0|0|0|
  └────────────┘
```

### With Alpha Beta pruning
- Depth 3 = 4s

## With Randomize + Alpha Beta Pruning
- Depth 3 = ~4-6s

## With Go routine

TODO

## With both Alpha Beta pruning + Go routine

TODO