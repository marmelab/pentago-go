
# AI performance results


## Monica Five

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

### Without optimization

- Depth 2 = 135ms
- Depth 3 = 32s
- Depth 4 = Too long

### With Alpha Beta pruning

- Depth 2 = 147ms
- Depth 3 = 7s
- Depth 4 = 9min

### Add Randomize on pruning

- Depth 2 = ~140-150ms
- Depth 3 = ~4-6s
- Depth 4 = ~3-4min

> The game state configuration force players to move on the center of the board. Is it always true ?

See below to check with others configuration !
### Add Go routine for each first move

- Depth 2 = 140ms
- Depth 3 = ~2-3s
- Depth 4 = 1min47
  
## Line straight configuration (To demonstrate randomize)

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

### Without Randomize
- Depth 3 = ~1-2s
### With Randomize

- Depth 3 = ~2-4s

## Close to the end

   0 1 2  3 4 5
  ┌────────────┐
0 |1|0|2||0|2|1|
1 |0|1|0||0|2|0|
2 |2|0|2||1|2|0|
  |────────────|
3 |1|0|1||1|0|2|
4 |1|2|2||2|1|0|
5 |0|0|0||0|0|1|
  └────────────┘

### With AB Pruning, randomize & go routines

- Depth 2 = 19ms
- Depth 3 = ~650ms
- Depth 4 = 9s
- Depth 5 = ~2.30-4min
