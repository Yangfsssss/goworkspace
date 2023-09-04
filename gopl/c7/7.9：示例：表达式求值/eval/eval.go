package eval

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() (bool, string)
}

type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (v Var) String() (bool, string) {
	return true, string(v)
}

type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) String() (bool, string) {
	return true, fmt.Sprintf("%g", l)
}

type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}

	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unsupported unary operator: %q", u.op)
	}

	return u.x.Check(vars)
}

func (u unary) String() (bool, string) {
	return true, fmt.Sprintf("(%c %s)", u.op, u.x)
}

type binary struct {
	op   rune
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}

	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unsupported binary operator: %q", b.op)
	}

	if err := b.x.Check(vars); err != nil {
		return err
	}

	return b.y.Check(vars)
}

func (b binary) String() (bool, string) {
	return true, fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

type call struct {
	fn   string
	args []Expr
	ast  string
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}

	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

var numParams = map[string]int{
	"pow":  2,
	"sin":  1,
	"sqrt": 1,
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function call: %s", c.fn)
	}

	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}

	for _, args := range c.args {
		if err := args.Check(vars); err != nil {
			return err
		}
	}

	return nil
}

func (c *call) String() (bool, string) {
	var buf bytes.Buffer
	buf.WriteString(c.fn)
	buf.WriteByte('(')
	for i, arg := range c.args {
		if i > 0 {
			buf.WriteByte(',')
		}

		_, s := arg.String()

		buf.WriteString(s)
	}

	buf.WriteByte(')')

	fmt.Println(buf.String())
	fmt.Println(c.ast)

	if c.ast != buf.String() {
		c.ast = buf.String()
		return false, buf.String()
	}

	return true, buf.String()
}
