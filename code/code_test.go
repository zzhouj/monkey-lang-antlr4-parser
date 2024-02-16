package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		op       OpCode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		if len(instruction) != len(tt.expected) {
			t.Errorf("instruction has wrong length. want=%d, got=%d", len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != b {
				t.Errorf("wrong byte at pos %d. want=%d, got=%d", i, b, instruction[i])
			}
		}
	}
}

func TestInstructionsString(t *testing.T) {
	instructions := []Instructions{
		Make(OpConstant, 1),
		Make(OpConstant, 2),
		Make(OpConstant, 65534),
	}

	expected := `0000 OpConstant 1
0003 OpConstant 2
0006 OpConstant 65534
`

	concatted := concatInstructions(instructions)

	actual := concatted.String()
	if expected != actual {
		t.Errorf("wrong string of instructions.\nwant=%q\ngot=%q", expected, actual)
	}
}

func concatInstructions(instructions []Instructions) Instructions {
	out := Instructions{}

	for _, ins := range instructions {
		out = append(out, ins...)
	}

	return out
}
