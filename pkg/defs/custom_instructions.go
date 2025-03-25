package defs

type InstructionType string

const (
	Type42 InstructionType = "type42"
	Sign   InstructionType = "sign"
)

type CustomInstruction struct {
	Type        InstructionType `json:"type"`
	Instruction string          `json:"instruction"`
}
