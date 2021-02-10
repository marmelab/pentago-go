# Pentago-go 🍻

[Pentago](https://en.wikipedia.org/wiki/Pentago) game implementation in **Go**.

Definition and rules from Wikipedia:

> Pentago is a two-player abstract strategy game invented by Tomas Flodén. The Swedish company Mindtwister has the rights of developing and commercializing the product. The game is played on a 6×6 board divided into four 3×3 sub-boards (or quadrants). Taking turns, the two players place a marble of their color (either black or white) onto an unoccupied space on the board, and then rotate one of the sub-boards by 90 degrees either clockwise or anti-clockwise. A player wins by getting five of their marbles in a vertical, horizontal or diagonal row (either before or after the sub-board rotation in their move). If all 36 spaces on the board are occupied without a row of five being formed then the game is a draw."

Type `make help` to list all commands available

## Prerequisite

- [Docker](https://www.docker.com/) & [docker-compose](https://docs.docker.com/compose/) are installed on your machine.

## Installation

- Clone this repo
- Run `make install`

## Start

### Use A Preconfigured Board

- Run `make FILE="game_run.txt" start` to use one of existing game state and find best moves.
- You can also specify which board to use by typing `make FILE="NAME.txt" start`.
*This file must be present in the `./src/datasets` folder.

### Create Your Own Board

You can also create your own board.

- Copy-paste one file.
- Replace all numbers by which game state you want.

The numbers should follow these rules:

- `0` is used for an empty cell
- `1` is for player 1 (you)
- `2` is for player 2 (the opponent)

## Tests

- Run `make test-verbose`
