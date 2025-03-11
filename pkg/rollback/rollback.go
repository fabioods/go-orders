package rollback

import (
	"context"
	"log"
)

type Func struct {
	Name string
	Func func() error
}

type IRollback interface {
	Add(name string, function func() error) IRollback
	Do(ctx context.Context) []string
}

type Rollback struct {
	functions []Func
}

func (s *Rollback) Add(name string, function func() error) IRollback {
	s.functions = append(s.functions, Func{
		Name: name,
		Func: function,
	})
	return s
}

func (s *Rollback) Do(ctx context.Context) []string {
	callFuncName := make([]string, 0)
	for i := len(s.functions) - 1; i >= 0; i-- {
		item := s.functions[i]
		log.Println("Rollback: ", item.Name)
		if err := item.Func(); err != nil {
			log.Printf("Erro ao executar rollback %s: %v\n", item.Name, err)
		}
		callFuncName = append(callFuncName, item.Name)
	}
	return callFuncName
}

func New() *Rollback {
	return &Rollback{
		functions: make([]Func, 0),
	}
}
