package chip8

type Chip8 struct {
	stack [16]uint16

	memory [4096]uint8
	V      [16]uint8
	keys   [16]uint8

	opcode uint16
	I      uint16
	PC     uint16
	SP     uint16

	delayTimer uint8
	soundTimer uint8

	display [64 * 32]uint8
}
