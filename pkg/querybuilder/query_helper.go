package querybuilder

import "github.com/nikunjratanpara/cmms/pkg/linq"

func GetOneComponent[T Clause](clauses []Clause, isClauses ...IsClause) *T {
	clause := linq.FirstOrDefault[Clause](
		clauses,
		func(clause Clause) bool {
			return linq.AllMatch[IsClause](
				isClauses,
				func(isClause IsClause) bool { return isClause(clause) })
		},
		nil)

	if clause != nil {
		tClause, ok := clause.(T)
		if ok {
			return &tClause
		}
	}
	return nil
}
