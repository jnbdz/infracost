package hcl

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type VertexOutput struct {
	logger    *logrus.Entry
	evaluator *Evaluator
	block     *Block
}

func (v *VertexOutput) ID() string {
	if v.block.ModuleName() == "" {
		return v.block.Label()
	}

	return fmt.Sprintf("%s.%s", v.block.ModuleAddress(), v.block.Label())
}

func (v *VertexOutput) ModuleAddress() string {
	return v.block.ModuleAddress()
}

func (v *VertexOutput) Evaluator() *Evaluator {
	return v.evaluator
}

func (v *VertexOutput) References() []VertexReference {
	return referencesForBlock(v.block)
}

func (v *VertexOutput) Evaluate() error {
	val, err := v.evaluator.evaluateOutput(v.block)
	if err != nil {
		return fmt.Errorf("could not evaluate output %s: %w", v.ID(), err)
	}

	v.logger.Debugf("adding output %s to the evaluation context", v.ID())
	key := fmt.Sprintf("output.%s", v.block.Label())
	v.evaluator.ctx.SetByDot(val, key)

	return nil
}

func (v *VertexOutput) Expand() ([]*Block, error) {
	return []*Block{v.block}, nil
}