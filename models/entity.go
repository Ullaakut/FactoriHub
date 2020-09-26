package models

import (
	"fmt"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/pkg/errors"
)

type Entity struct {
	// Index of the entity, 1-based.
	Index int `json:"entity_number" db:"-"`
	// Prototype name of the entity (e.g. "offshore-pump").
	Name string `json:"name" db:"-"`
	Type string `json:"type" db:"-"`

	// The following fields are not stored in the database at the moment.

	Position  Position `json:"position" db:"-"`
	Direction int      `json:"direction,omitempty" db:"-"`
	// Orientation of cargo wagon or locomotive, value 0 to 1 (optional).
	Orientation     float64         `json:"orientation" db:"-"`
	Connections     Connections     `json:"connections" db:"-"`
	ControlBehavior ControlBehavior `json:"control_behavior,omitempty" db:"-"`
	Items           ItemRequest     `json:"items" db:"-"`
	Recipe          string          `json:"recipe" db:"-"`
	// Used by Prototype/Container, optional. The index of the first inaccessible item slot due to limiting with the red "bar". 0-based Types/ItemStackIndex.
	// See https://wiki.factorio.com/Types/ItemStackIndex
	Bar                uint16                 `json:"bar" db:"-"`
	Inventory          Inventory              `json:"inventory" db:"-"`
	InfinitySettings   InfinitySettings       `json:"infinity_settings" db:"-"`
	InputPriority      string                 `json:"input_priority" db:"-"`
	OutputPriority     string                 `json:"output_priority" db:"-"`
	Filter             string                 `json:"filter" db:"-"`
	Filters            []ItemFilter           `json:"filters" db:"-"`
	FilterMode         string                 `json:"filter_mode" db:"-"`
	OverrideStackSize  uint8                  `json:"override_stack_size" db:"-"`
	DropPosition       Position               `json:"drop_position" db:"-"`
	PickupPosition     Position               `json:"pickup_position" db:"-"`
	RequestFilters     []LogisticFilter       `json:"request_filters" db:"-"`
	RequestFromBuffers bool                   `json:"request_from_buffers" db:"-"`
	Parameters         SpeakerParameters      `json:"parameters" db:"-"`
	AlertParameters    SpeakerAlertParameters `json:"alert_parameters" db:"-"`
	AutoLaunch         bool                   `json:"auto_launch" db:"-"`
	// See https://wiki.factorio.com/Types/GraphicsVariation
	Variation uint8  `json:"variation" db:"-"`
	Color     Color  `json:"color" db:"-"`
	Station   string `json:"station" db:"-"`
}

type Entities []Entity

// ValidateCreate decodes a blueprint's info from its raw string and verifies whether it is valid.
func (e Entity) ValidateCreate(_ *pop.Connection) (*validate.Errors, error) {
	fmt.Println("Validating entity")

	if e.Name == "" {
		return validate.NewErrors(), errors.New("missing entity name")
	}

	fmt.Println("Validated entity")

	return validate.NewErrors(), nil
}

// ValidateUpdate decodes a blueprint's info from its raw string and verifies whether it is valid.
func (e Entity) ValidateUpdate(_ *pop.Connection) (*validate.Errors, error) {
	if e.Name == "" {
		return validate.NewErrors(), errors.New("missing entity name")
	}

	return validate.NewErrors(), nil
}

type Inventory struct {
	Filters []ItemFilter `json:"filters"`
	// Used by Prototype/Container, optional. The index of the first inaccessible item slot due to limiting with the red "bar". 0-based Types/ItemStackIndex.
	// See https://wiki.factorio.com/Types/ItemStackIndex
	Bar uint16 `json:"bar"`
}

type Connections struct {
	First  ConnectionPoint `json:"1"`
	Second ConnectionPoint `json:"2"`
}

type ConnectionPoint struct {
	Red   []ConnectionData `json:"red"`
	Green []ConnectionData `json:"green"`
}

type ConnectionData struct {
	EntityID  int `json:"entity_id"`
	CircuitID int `json:"circuit_id"`
}

type ItemRequest map[string]uint32

type ItemFilter struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type InfinitySettings struct {
	RemoveUnfilteredItems bool             `json:"remove_unfiltered_items"`
	Filters               []InfinityFilter `json:"filters"`
}

type InfinityFilter struct {
	Name  string `json:"name"`
	Count uint32 `json:"count"`
	Mode  string `json:"mode"`
	Index int    `json:"index"`
}

type LogisticFilter struct {
	Name  string `json:"name"`
	Count uint32 `json:"count"`
	Index int    `json:"index"`
}

type SpeakerParameters struct {
	PlaybackVolume   float64 `json:"playback_volume"`
	PlaybackGlobally bool    `json:"playback_globally"`
	AllowPolyphony   bool    `json:"allow_polyphony"`
}

type SpeakerAlertParameters struct {
	ShowAlert      bool     `json:"show_alert"`
	ShowAlertOnMap bool     `json:"show_alert_on_map"`
	IconSignalID   SignalID `json:"icon_signal_id"`
	AlertMessage   string   `json:"alert_message"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ControlBehavior struct {
	Filters []Filters `json:"filters"`
}

type Filters struct {
	Signal Signal `json:"signal"`
	Count  int    `json:"count"`
	Index  int    `json:"index"`
}

type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}

type Tile struct {
	Name     string   `json:"name"`
	Position Position `json:"position"`
}

type Schedule struct {
	Schedule    []ScheduleRecord `json:"schedule"`
	Locomotives []int            `json:"locomotives"`
}

type ScheduleRecord struct {
	Station        string          `json:"station"`
	WaitConditions []WaitCondition `json:"wait_conditions"`
}

type WaitCondition struct {
	// One of "time", "inactivity", "full", "empty", "item_count", "circuit", "robots_inactive", "fluid_count", "passenger_present", "passenger_not_present".
	Type string `json:"type"`

	// Either "and", or "or". Tells how this condition is to be compared with the preceding conditions in the corresponding wait_conditions array.
	CompareType string `json:"compare_type"`

	// Number of ticks to wait or of inactivity. Only present when type is "time" or "inactivity". Optional.
	Ticks uint `json:"ticks"`

	// CircuitCondition Object, only present when type is "item_count", "circuit" or "fluid_count".
	Condition CircuitCondition `json:"condition"`
}

type CircuitCondition struct {
	Comparator   string `json:"comparator"`
	FirstSignal  Signal `json:"first_signal"`
	SecondSignal Signal `json:"second_signal"`
	Constant     int    `json:"constant"`
}
