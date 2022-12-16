package query

type Query struct {
	ID       int
	Priority int
}

func NewQuery(id, priority int) *Query {
	return &Query{ID: id, Priority: priority}
}
