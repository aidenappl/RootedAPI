package service

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/aidenappl/rootedapi/structs"
	"github.com/aidenappl/rootedapi/tools"
)

func ApplyPagination(q sq.SelectBuilder, limit *int, offset *int) sq.SelectBuilder {
	if limit == nil {
		limit = tools.IntP(structs.DefaultListLimit)
	}
	if *limit > structs.MaximumListLimit {
		*limit = structs.MaximumListLimit
	}

	q = q.Limit(uint64(*limit))

	if offset != nil {
		q = q.Offset(uint64(*offset))
	} else {
		q = q.Offset(0)
	}

	return q
}
