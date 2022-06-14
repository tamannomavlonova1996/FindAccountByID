package types

import "errors"

type Money int64
type PaymentCategory string
type PaymentStatus string

const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

var (
	ErrPaymentNotFound = errors.New("err payment not found")
)

type Phone string
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
