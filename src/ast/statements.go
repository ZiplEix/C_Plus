package ast

type VarType int

// Enum for the different types of variables that can be declared.
const (
	VOID VarType = iota
	INT
	FLOAT
	DOUBLE
	CHAR
	SHORT
	LONG
)

// Enum for the different types of comments that can be declared.
const (
	SINGLE_LINE_COMMENT VarType = iota
	MULTI_LINE_COMMENT
)

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

type ExprStmt struct {
	Expr Expr
}

func (e ExprStmt) stmt() {}

type VarDeclarationStmt struct {
	Name         string
	IsConst      bool
	IsSigned     bool
	PointerLevel int
	Type         VarType
	AssignedExpr Expr
}

func (v VarDeclarationStmt) stmt() {}

type ReturnStmt struct {
	Expr Expr
}

func (r ReturnStmt) stmt() {}

type CommentStmt struct {
	Value string
	Type  VarType
}

func (c CommentStmt) stmt() {}

type IncluderStmt struct {
	Value string
}

func (i IncluderStmt) stmt() {}

type Parameter struct {
	Name         string
	IsConst      bool
	IsSigned     bool
	PointerLevel int
	Type         VarType
}

type FunctionDeclarationStmt struct {
	Parameters   []Parameter
	Name         string
	Body         []Stmt
	ReturnType   VarType
	IsConst      bool
	IsSigned     bool
	PointerLevel int
}

func (f FunctionDeclarationStmt) stmt() {}
