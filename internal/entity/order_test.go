package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder_valid(t *testing.T) {
	order := NewOrder("cb834259-7b73-4510-9963-730fab8c6511")
	orderItem := NewOrderItem("Product 1", 10.0, 1, order.ID)
	order.AddItem(orderItem)
	order.CalculateTotal()
	assert.Nil(t, order.Validate())
	assert.Nil(t, orderItem.Validate())
}

func TestOrder_invalid(t *testing.T) {
	order := NewOrder("cb834259-7b73-4510-9963-730fab8c6511")
	orderItem := NewOrderItem("Product 1", 0, 1, order.ID)
	order.AddItem(orderItem)
	assert.NotNil(t, order.Validate())
	assert.NotNil(t, orderItem.Validate())
}

func TestOrder_changeStatusToCanceld(t *testing.T) {
	order := NewOrder("cb834259-7b73-4510-9963-730fab8c6511")
	orderItem := NewOrderItem("Product 1", 10.0, 1, order.ID)
	order.AddItem(orderItem)
	order.CalculateTotal()
	assert.Equal(t, Created, order.Status)
	assert.Nil(t, order.Validate())
	order.Cancel()
	assert.Equal(t, Canceled, order.Status)
}

func TestOrder_changeStatusToPaid(t *testing.T) {
	order := NewOrder("cb834259-7b73-4510-9963-730fab8c6511")
	orderItem := NewOrderItem("Product 1", 10.0, 1, order.ID)
	order.AddItem(orderItem)
	order.CalculateTotal()
	assert.Equal(t, Created, order.Status)
	assert.Nil(t, order.Validate())
	order.Pay()
	assert.Equal(t, Paid, order.Status)
}
