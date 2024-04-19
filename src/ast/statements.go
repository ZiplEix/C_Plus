package ast

type BlockStmt struct {
	Body []Stmt
}

func (b *BlockStmt) stmt() {}

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
