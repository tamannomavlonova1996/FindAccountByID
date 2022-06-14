package payments

import (
	"FindAccountById/types"
)

type Payment struct {
	ID       string
	Amount   types.Money
	Category types.PaymentCategory
	Status   types.PaymentStatus
}

var payments = []Payment{
	{
		ID:       "1",
		Amount:   5,
		Category: "auto",
		Status:   types.PaymentStatusInProgress,
	},
	{
		ID:       "2",
		Amount:   10,
		Category: "food",
		Status:   types.PaymentStatusOk,
	},
}

func (p *Payment) Reject(paymentID string) (*Payment, error) {
	payment, err := p.FindPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}
	payment.Status = types.PaymentStatusFail

	return payment, nil
}

func (p *Payment) FindPaymentByID(paymentID string) (*Payment, error) {
	for _, v := range payments {
		if v.ID == paymentID {
			return &v, nil
		}
	}

	return nil, types.ErrPaymentNotFound
}
