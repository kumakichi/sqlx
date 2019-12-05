package builder

import (
	"context"
)

type CombinationAddition struct {
}

func (CombinationAddition) weight() additionWeight {
	return combinationStmt
}

func Union() *combination {
	return &combination{
		operator: "UNION",
	}
}

func Intersect() *combination {
	return &combination{
		operator: "INTERSECT",
	}
}

func Expect() *combination {
	return &combination{
		operator: "EXCEPT",
	}
}

var _ Addition = (*combination)(nil)

type combination struct {
	CombinationAddition
	operator   string // UNION | INTERSECT | EXCEPT
	method     string // ALL | DISTINCT
	stmtSelect *StmtSelect
}

func (c *combination) IsNil() bool {
	return c == nil || IsNilExpr(c.stmtSelect)
}

func (c combination) All(stmtSelect *StmtSelect) *combination {
	c.method = "ALL"
	c.stmtSelect = stmtSelect
	return &c
}

func (c combination) Distinct(stmtSelect *StmtSelect) *combination {
	c.method = "DISTINCT"
	c.stmtSelect = stmtSelect
	return &c
}

func (c *combination) Ex(ctx context.Context) *Ex {
	e := Expr(c.operator)
	e.WriteByte(' ')

	if c.method != "" {
		e.WriteString(c.method)
		e.WriteByte(' ')
	}

	e.WriteExpr(c.stmtSelect)

	return e.Ex(ctx)
}