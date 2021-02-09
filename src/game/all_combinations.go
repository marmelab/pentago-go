package game

var ALL_COMBINATIONS_OF_WIN = [32]int64{34905131008, 545392672, 17452565504, 272696336, 8726282752, 136348168, 4363141376, 68174084, 2181570688, 34087042, 1090785344, 17043521, 66571993088, 33285996544, 1040187392, 520093696, 16252928, 8126464, 253952, 126976, 3968, 1984, 62, 31, 34630287488, 270549121, 541098242, 17315143744, 1108378624, 34636832, 2216757248, 17318416}

var FULL_BOARD int64 = 68719476735

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
	// We use bitwise OR operator to combine both player1 & player2 and compare it to
	// The representation of a full board.
	return FULL_BOARD == (player1Int64 | player2Int64)
}
