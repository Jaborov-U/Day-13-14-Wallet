package types

// Money = денежная сумма в минимальных единицах.
type Money int64

// Currency = представляет код валюты.
type Currency string

// Status представляет собой статус платежа.
type Status string

// Предопределенные статус платежей
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
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
type Category string

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
	ID int
	Amount Money
	Category Category
	Status Status
}

// Слайс из карт для ДЗ 9.
type PaymentSourse struct{
	Type string // card
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}

