package models

type Signal struct {
	Signal SignalID `json:"signal" db:"-"`
	Count  int      `json:"count" db:"-"`
}

type SignalID struct {
	Type string `json:"type" db:"-"`
	Name string `json:"name" db:"-"`
}
