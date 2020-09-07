package models

//stripe connect

type Charge struct {
	ID *int64
}

type Customer struct {
	ID *int64
}

type Subscription struct {
	ID *int64
}

type Payout struct { //negative charge?

}
