package game

var ALL_COMBINATIONS_OF_WIN = [32]int64{
	0b100000_100000_100000_100000_100000_000000,
	0b000000_100000_100000_100000_100000_100000,
	0b010000_010000_010000_010000_010000_000000,
	0b000000_010000_010000_010000_010000_010000,
	0b001000_001000_001000_001000_001000_000000,
	0b000000_001000_001000_001000_001000_001000,
	0b000100_000100_000100_000100_000100_000000,
	0b000000_000100_000100_000100_000100_000100,
	0b000010_000010_000010_000010_000010_000000,
	0b000000_000010_000010_000010_000010_000010,
	0b000001_000001_000001_000001_000001_000000,
	0b000000_000001_000001_000001_000001_000001,
	0b111110_000000_000000_000000_000000_000000,
	0b011111_000000_000000_000000_000000_000000,
	0b000000_111110_000000_000000_000000_000000,
	0b000000_011111_000000_000000_000000_000000,
	0b000000_000000_111110_000000_000000_000000,
	0b000000_000000_011111_000000_000000_000000,
	0b000000_000000_000000_111110_000000_000000,
	0b000000_000000_000000_011111_000000_000000,
	0b000000_000000_000000_000000_111110_000000,
	0b000000_000000_000000_000000_011111_000000,
	0b000000_000000_000000_000000_000000_111110,
	0b000000_000000_000000_000000_000000_011111,
	0b100000_010000_001000_000100_000010_000000,
	0b000000_010000_001000_000100_000010_000001,
	0b000000_100000_010000_001000_000100_000010,
	0b010000_001000_000100_000010_000001_000000,
	0b000001_000010_000100_001000_010000_000000,
	0b000000_000010_000100_001000_010000_100000,
	0b000010_000100_001000_010000_100000_000000,
	0b000000_000001_000010_000100_001000_010000}
	
var FULL_BOARD int64 = 0b111111_111111_111111_111111_111111_111111

func GetAllCombinations()[32]int64 {
	return ALL_COMBINATIONS_OF_WIN
}

func BinaryCompareInt64(combination int64, playerBoard int64) bool {
	// bitwise AND operator between the given combination and the playerBoard.
	// If the bitwise and operator = the combination
	// It means that the player have at least every bits of the combination
	// (He has played every cell on the board to win by this combination)
	return combination & playerBoard == combination
}

func IsBoardFull(player1Int64 int64, player2Int64 int64) bool {
	// We use bitwise OR operator to combine both player1 0 player2 and compare it to
	// The representation of a full board.
	return FULL_BOARD == (player1Int64 | player2Int64)
}