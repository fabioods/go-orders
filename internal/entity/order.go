package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Status string

const Created Status = "CREATED"
const Canceled Status = "CANCELED"
const Paid Status = "PAID"

type Order struct {
	ID        string       `json:"id" gorm:"type:uuid;primary_key" validate:"required,uuid4"`
	UserID    string       `json:"user_id" gorm:"type:uuid" validate:"required,uuid4"`
	Status    Status       `json:"status" gorm:"type:varchar(20)" validate:"required,oneof=CREATED CANCELED PAID"`
	Total     float64      `json:"total" gorm:"type:float" validate:"required,gt=1"`
	Items     []*OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime" validate:"required"`
	UpdatedAt time.Time    `json:"updated_at"  gorm:"autoUpdateTime" validate:"required"`
}

type OrderItem struct {
	ID       string  `json:"id" gorm:"type:uuid;primary_key" validate:"required,uuid4"`
	OrderID  string  `json:"order_id" gorm:"type:uuid" validate:"required,uuid4"`
	Name     string  `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Price    float64 `json:"price" gorm:"type:float" validate:"required,numeric,gt=1"`
	Quantity int64   `json:"quantity" gorm:"type:int" validate:"required,numeric,min=1,gte=1"`
}

func NewOrder(userID string) *Order {
	return &Order{
		ID:        uuid.New().String(),
		UserID:    userID,
		Status:    Created,
		Total:     0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (o *Order) CalculateTotal() {
	var total float64
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	o.Total = total
}

func (o *Order) Cancel() {
	o.Status = Canceled
}

func (o *Order) Pay() {
	o.Status = Paid
}

func (o *Order) AddItem(item *OrderItem) {
	item.OrderID = o.ID
	o.Items = append(o.Items, item)
}

func (o *Order) Validate() error {
	validate := validator.New()
	err := validate.Struct(o)
	if err != nil {
		var errMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf("Field '%s' invalid: %s", err.Field(), err.Tag()))
		}
		return fmt.Errorf("%s", strings.Join(errMsgs, "; "))
	}
	return nil
}

func NewOrderItem(name string, price float64, quantity int64, orderId string) *OrderItem {
	return &OrderItem{
		ID:       uuid.New().String(),
		Name:     name,
		Price:    price,
		Quantity: quantity,
		OrderID:  orderId,
	}
}

func (o *OrderItem) Validate() error {
	validate := validator.New()
	err := validate.Struct(o)
	if err != nil {
		var errMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf("Field '%s' invalid: %s", err.Field(), err.Tag()))
		}
		return fmt.Errorf("%s", strings.Join(errMsgs, "; "))
	}
	return nil
}
