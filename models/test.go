package models

type Test struct {
	ID     *int64  `json:"id,omitempty" db:"id"`
	Tester *string `json:"tester,omitempty" db:"tester"`
	Opt    *bool   `json:"opt,omitempty" db:"opt"`
	//Created *time.Time `json:"created,omitempty"`
	Amount *int64 `json:"amount,omitempty" db:"amount"`
}
