package ast

type Type int

// Enum for the different types of variables that can be declared.
const (
	VOID Type = iota
	INT
	FLOAT
	DOUBLE
	CHAR
	SHORT
	LONG
)

// Enum for the different types of comments that can be declared.
const (
	SINGLE_LINE_COMMENT Type = iota
	MULTI_LINE_COMMENT
)

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
	PointerLevel int
	Type         Type
	AssignedExpr Expr
}

func (v *VarDeclarationStmt) stmt() {}

type ReturnStmt struct {
	Expr Expr
}

func (r *ReturnStmt) stmt() {}

type CommentStmt struct {
	Value string
	Type  Type
}

func (c *CommentStmt) stmt() {}

type IncluderStmt struct {
	Value string
}

func (i *IncluderStmt) stmt() {}
