package interpreter

import (
	"fmt"

	"github.com/4chain-ag/universal-test-vectors/pkg/defs"
)

// Interpreter is a struct that is used to interpret custom instructions.
type Interpreter[R Resolver[TKey], TKey InputKeys] struct {
	resolver R
}

// NewInterpreter creates a new interpreter for custom instructions with a given resolver.
func NewInterpreter[R Resolver[TKey], TKey InputKeys](resolver R) *Interpreter[R, TKey] {
	return &Interpreter[R, TKey]{
		resolver: resolver,
	}
}

// Process processes custom instructions for a given key.
func (p *Interpreter[I, TKey]) Process(key *TKey, instructions []defs.CustomInstruction) (I, error) {
	var err error
	var proceed bool

	err = p.resolver.Initialize(key)
	if err != nil {
		return p.resolver, fmt.Errorf("failed to initialize resolver: %w", err)
	}

	for _, instruction := range instructions {
		switch instruction.Type {
		case defs.Type42:
			proceed, err = p.resolver.Type42(instruction.Instruction)
		case defs.Sign:
			proceed, err = p.resolver.Sign(instruction.Instruction)
		default:
			return p.resolver, fmt.Errorf("unsupported instruction type: %s", instruction.Type)
		}
		if err != nil {
			return p.resolver, fmt.Errorf("failed to process instruction: %w", err)
		}
		if !proceed {
			break
		}
	}

	err = p.resolver.Finalize()
	if err != nil {
		return p.resolver, fmt.Errorf("failed to finalize resolver: %w", err)
	}

	return p.resolver, nil
}
