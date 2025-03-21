package rollback_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fabioods/go-orders/pkg/rollback"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	rb := rollback.New()

	// Assert
	assert.NotNil(t, rb)
}

func TestRollback_Add_Happy(t *testing.T) {
	// Arrange
	rb := rollback.New()

	// Act
	rb.Add("function one", func() error {
		fmt.Println("done function one")
		return nil
	}).Add("function two", func() error {
		fmt.Println("done function two")
		return nil
	}).Add("function tree", func() error {
		fmt.Println("done function two")
		return nil
	})

	// Assert
	assert.NotNil(t, rb)
}

func TestRollback_Do_Happy(t *testing.T) {
	// Arrange
	ctx := context.Background()

	rb := rollback.New().
		Add("function one", func() error {
			fmt.Println("done function one")
			return nil
		}).
		Add("function two", func() error {
			fmt.Println("done function two")
			return nil
		}).
		Add("function tree", func() error {
			fmt.Println("done function two")
			return nil
		})

	// Act
	calls := rb.Do(ctx)

	// Assert
	assert.NotNil(t, rb)
	assert.Equal(t, "function tree", calls[0])
	assert.Equal(t, "function two", calls[1])
	assert.Equal(t, "function one", calls[2])
}
