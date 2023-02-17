package types

// Money = денежная сумма в минимальных единицах.
type Money int64

// Currency = представляет код валюты.
type Currency string

// Status представляет собой статус платежа.
type PaymentStatus string

// Предопределенные статус платежей
const (
	StatusOk PaymentStatus = "OK"
	StatusFail PaymentStatus = "FAIL"
	StatusInProgress PaymentStatus = "INPROGRESS"
)

// Коды валют.
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN = номер карты.
type PAN string

// Category представляет собой категорию, в которой был совершен платеж ("авто", "интернет" и т.д.).
type PaymentCategory string

// Card = информация о платежной карте.
type Card struct{
	ID int
	PAN PAN
	Balance Money
	Currency Currency
	Color string
	Name string
	Active bool
	MinBalance Money
}

// Payment пресдавляет информацию о платеже.
type Payment struct {
	ID string
	AccountID int64
	Amount Money
	Category PaymentCategory
	Status PaymentStatus
}

type Phone string

// Слайс из карт для ДЗ 9.
type PaymentSourse struct{
	Type string // card
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}

type Account struct{
	ID int64
	Phone Phone
	Balance Money
}
