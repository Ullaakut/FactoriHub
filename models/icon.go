package models

// Icon is used by pop to map your icons database table to your go code.
type Icon struct {
	Signal SignalID `json:"signal" db:"-"`
	Index  int      `json:"index" db:"-"`
}

type Icons []Icon
