package parser_test

import (
	"testing"

	"github.com/skippyza/monkey-interpreter/ast"
	"github.com/skippyza/monkey-interpreter/lexer"
	"github.com/skippyza/monkey-interpreter/parser"
	"github.com/stretchr/testify/assert"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	assert.NotNil(t, program)
	assert.Len(t, program.Statements, 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	assert.Equal(t, "let", s.TokenLiteral())

	assert.IsType(t, &ast.LetStatement{}, s)
	letStmt := s.(*ast.LetStatement)

	assert.Equal(t, letStmt.Name.Value, name)
	assert.Equal(t, letStmt.Name.TokenLiteral(), name)
}
