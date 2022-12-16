package buffer

import (
	"fmt"
	"sort"
	"sync"

	"SimulationModelingCourseWork/internal"
	"SimulationModelingCourseWork/internal/query"
)

type Buffer struct {
	mu      *sync.Mutex
	content []query.Query
	size    int
}

func NewBuffer() *Buffer {
	return &Buffer{
		mu:      &sync.Mutex{},
		content: []query.Query{},
		size:    internal.L,
	}
}

func (b *Buffer) AddQuery(q query.Query) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.content) < b.size {
		fmt.Printf("Added new query: ID = [%d], Priority = [%d]\n", q.ID, q.Priority)
		b.content = append(b.content, q)
	} else {
		fmt.Printf("Buffer is full, can't add new query: ID = [%d], Priority = [%d]\n", q.ID, q.Priority)
	}
}

func (b *Buffer) GetQuery() (query.Query, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.content) == 0 {
		return query.Query{}, fmt.Errorf("Empty buffer")
	}

	sort.Slice(b.content, func(i, j int) bool {
		return b.content[i].Priority < b.content[j].Priority
	})

	result := b.content[0]

	b.content = b.content[1:]

	return result, nil
}
