package tests

import (
	"bingbonglang/lexer"
	"bingbonglang/parser"
	"testing"
)

func TestVarStatements(t *testing.T) {
	input := `
var x = 5;
var y = 10;
var foobar = 93939;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},

		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVarStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}
func testVarStatement(t *testing.T, s parser.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var'. got=%q", s.TokenLiteral())
		return false
	}
	varStmt, ok := s.(*parser.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement. got=%T", s)
		return false
	}
	if varStmt.Name.Value != name {
		t.Errorf("VarStmt.Name.Value not '%s'. got=%s", name, varStmt.Name.Value)
		return false
	}
	if varStmt.Name.Token.Literal != name {
		t.Errorf("s.Name not '%s'. got=%s", name, varStmt.Name)
		return false
	}
	return true
}
