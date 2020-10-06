package models

type Tag struct {
	Key   *string `db:"key" json:"key,omitempty"`
	Value *int64  `db:"value" json:"value,omitempty"`
}

//able to track severity of conditions, also able to use data in ML in future
