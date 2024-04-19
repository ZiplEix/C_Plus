package ast

type BlockAst struct {
	Body []Stmt
}

func (b *BlockAst) stmt() {}

type ExprStmt struct {
	Expr Expr
}

func (e *ExprStmt) stmt() {}

type VarDeclarationStmt struct {
	Name         string
	IsConst      bool
	IsSigned     bool
	AssignedExpr Expr
}

func (v *VarDeclarationStmt) stmt() {}
