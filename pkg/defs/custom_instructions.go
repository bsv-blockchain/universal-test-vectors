package defs

// InstructionType is a type that represents the type of a custom instruction.
type InstructionType string

// Possible values for InstructionType
const (
	Type42 InstructionType = "type42"
	Sign   InstructionType = "sign"
)

// CustomInstruction is a struct that represents a custom instruction, used for locking and unlocking scripts.
type CustomInstruction struct {
	Type        InstructionType `json:"type"`
	Instruction string          `json:"instruction"`
}
