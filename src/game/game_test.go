package game

import (
    "testing"
)

func TestDeserializeBoardShouldReturnBoard(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

    board, err := DeserializeBoard(rawBoard)
	numberOfQuadrants := len(board.Quadrants)
	numberOfRows := len(board.Quadrants[0])
	numberOfColumns := len(board.Quadrants[0][0])
	if numberOfQuadrants != 4 || numberOfRows != 3 || numberOfColumns != 3 || err != nil {
		t.Errorf("The board's size are %d*%d*%d 3 d array (4*3*3 expected)", numberOfQuadrants, numberOfRows, numberOfColumns)
	}
}

func TestDeserializeBoardWithMissingLineShouldThrowAnError(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

    _, err := DeserializeBoard(rawBoard)
	if err == nil {
		t.Errorf("The given board should throw an error because second line are missing.")
	}
}

func TestDeserializeBoardWithMissingValueOnLineShouldThrowAnError(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|0|0|0||0|0||
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

    _, err := DeserializeBoard(rawBoard)
	if err == nil {
		t.Errorf("The given board should throw an error because at line 6, there are only 5 digits")
	}
}

func TestGetLineOfBoardForFirstLine(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|1|0|2||1|2|2|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := DeserializeBoard(rawBoard)


	result := GetLineOfBoard(board, 0)
	expected_result := "011021"
	if expected_result != result {
		t.Errorf("GetLineOfBoard returned %v, expected %v", result, expected_result)
	}
}

func TestGetLineOfBoardForFourthLine(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|1|0|2||1|2|2|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := DeserializeBoard(rawBoard)


	result := GetLineOfBoard(board, 3)
	expected_result := "102122"
	if expected_result != result {
		t.Errorf("GetLineOfBoard returned %v, expected %v", result, expected_result)
	}
}

func TestToStringBoard(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|1|0|2||1|2|2|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := DeserializeBoard(rawBoard)

	result := ToStringBoard(board)
	expected_result := "011021000000000000102122000000000000"
	if expected_result != result {
		t.Errorf("ToStringBoard returned %v, expected %v", result, expected_result)
	}
}

var convertQuadrantPositionIntoBoardPoitionDatasets = []struct {
	in  [3]int
	out [2]int
}{
	{[3]int{0, 2, 2}, [2]int{2, 2}},
	{[3]int{1, 2, 2}, [2]int{2, 5}},
	{[3]int{2, 2, 2}, [2]int{5, 2}},
	{[3]int{3, 2, 2}, [2]int{5, 5}},
}
func TestConvertQuadrantPositionIntoBoardPosition(t *testing.T) {
	for _, data := range convertQuadrantPositionIntoBoardPoitionDatasets {
		result, _ := ConvertQuadrantPositionIntoBoardPosition(data.in)

		if result != data.out {
			t.Errorf("ConvertQuadrantPositionIntoBoardPosition returned %v, expected %v", result, data.out)

		}
	}
}

func TestConvertQuadrantPositionIntoBoardPositionThrowAnError(t *testing.T) {
	positions := [3]int{4, 0, 0}

	_, err := ConvertQuadrantPositionIntoBoardPosition(positions)

	if err == nil {
		t.Errorf("ConvertQuadrantPositionIntoBoardPosition should throw an error, quarter 4 doesn't exist")
	}
}


func TestTranspose(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|0|0||1|0|1|
				|0|0|0||0|2|0|
				|────────────|
				|1|0|2||1|2|2|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := DeserializeBoard(rawBoard)
	result := Transpose(board.Quadrants[1])
	expected_result := [3][3]string{{"0", "1", "0"}, {"2", "0", "2"}, {"1", "1", "0"}}
	if result != expected_result {
		t.Errorf("Transpose returned %v, expected %v", result, expected_result)
	}
}


func TestReverse(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|0|0||1|0|1|
				|0|0|0||0|2|0|
				|────────────|
				|1|0|2||1|2|2|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := DeserializeBoard(rawBoard)
	result := Reverse(board.Quadrants[1])
	expected_result := [3][3]string{{"1", "2", "0"}, {"1", "0", "1"}, {"0", "2", "0"}}
	if result != expected_result {
		t.Errorf("Reverse returned %v, expected %v", result, expected_result)
	}
}

func TestRotateClockwise(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|0|0||1|0|1|
				|0|0|0||0|2|0|
				|────────────|
				|1|0|2||1|2|2|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := DeserializeBoard(rawBoard)
	result := Rotate(board.Quadrants[1], "clockwise")
	expected_result := [3][3]string{{"0", "1", "0"}, {"2", "0", "2"}, {"0", "1", "1"}}
	if result != expected_result {
		t.Errorf("Reverse returned %v, expected %v", result, expected_result)
	}
}

func TestRotateCounterClockwise(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|0|0||1|0|1|
				|0|0|0||0|2|0|
				|────────────|
				|1|0|2||1|2|2|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := DeserializeBoard(rawBoard)
	result := Rotate(board.Quadrants[1], "counter clockwise")
	expected_result := [3][3]string{{"1", "1", "0"}, {"2", "0", "2"}, {"0", "1", "0"}}
	if result != expected_result {
		t.Errorf("Reverse returned %v, expected %v", result, expected_result)
	}
}
