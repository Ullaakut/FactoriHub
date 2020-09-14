package models

import (
	"fmt"
	"strconv"
)

type BlueprintData struct {
	BlueprintBook BlueprintBook `json:"blueprint_book"`
}

type BlueprintBook struct {
	Author       User     `json:"-"`
	Favorites    uint     `json:"-"`
	Labels       []string `json:"-"`
	Image        string   `json:"-"`
	RawString    string   `json:"-"`
	RedditThread string   `json:"-"`

	// 	String, the name of the item that was saved ("blueprint-book" in vanilla).
	Item string `json:"item"`
	// 	String, the name of the blueprint set by the user.
	Label      string `json:"label"`
	LabelColor Color  `json:"label_color"`
	// The actual content of the blueprint book, array of objects containing an "index" key and 0-based value and a "blueprint" key with a #Blueprint object as the value.
	Blueprints []BlueprintReference `json:"blueprints"`
	// Index of the currently selected blueprint, 0-based.
	ActiveIndex int           `json:"active_index"`
	Version     VersionString `json:"version"`
}

type BlueprintReference struct {
	Blueprint Blueprint `json:"blueprint"`
	Index     int       `json:"index"`
}

// See https://wiki.factorio.com/Version_string_format
type VersionString struct {
	Major  uint16
	Minor  uint16
	Hotfix uint16
	Dev    uint16
}

// Bit shift to the right + AND with a full uint16 to extract the binary info from the uint64.
func (v *VersionString) UnmarshalJSON(data []byte) error {
	ver, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}

	v.Major = uint16(ver >> 48 & 65535)
	v.Minor = uint16(ver >> 32 & 65535)
	v.Hotfix = uint16(ver >> 16 & 65535)
	v.Dev = uint16(ver & 65535)

	return nil
}

func (v VersionString) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Hotfix)
}

type Blueprint struct {
	// String, the name of the item that was saved ("blueprint" in vanilla).
	Item       string     `json:"item"`
	Label      string     `json:"label"`
	LabelColor Color      `json:"label_color"`
	Entities   []Entities `json:"entities"`
	Tiles      []Tile     `json:"tiles"`
	Icons      []Icon     `json:"icons"`
	Schedules  []Schedule `json:"schedules"`
	// The map version of the map the blueprint was created in.
	Version VersionString `json:"version"`
}

func (b Blueprint) TotalEntities() map[string]uint {
	var total = make(map[string]uint)
	for _, e := range b.Entities {
		typ := e.Type
		switch typ {
		// No type specified.
		case "":
			typ = "item"

		// Legacy type naming.
		case "input", "output":
			typ = "item"
		}

		total[typ+"/"+e.Name] += 1
	}
	return total
}

type Entities struct {
	// Index of the entity, 1-based.
	EntityNumber int `json:"entity_number"`
	// Prototype name of the entity (e.g. "offshore-pump").
	Name      string   `json:"name"`
	Position  Position `json:"position"`
	Direction int      `json:"direction,omitempty"`
	// Orientation of cargo wagon or locomotive, value 0 to 1 (optional).
	Orientation     float64         `json:"orientation"`
	Connections     Connections     `json:"connections"`
	ControlBehavior ControlBehavior `json:"control_behavior,omitempty"`
	Items           ItemRequest     `json:"items"`
	Recipe          string          `json:"recipe"`
	// Used by Prototype/Container, optional. The index of the first inaccessible item slot due to limiting with the red "bar". 0-based Types/ItemStackIndex.
	// See https://wiki.factorio.com/Types/ItemStackIndex
	Bar                uint16                 `json:"bar"`
	Inventory          Inventory              `json:"inventory"`
	InfinitySettings   InfinitySettings       `json:"infinity_settings"`
	Type               string                 `json:"type"`
	InputPriority      string                 `json:"input_priority"`
	OutputPriority     string                 `json:"output_priority"`
	Filter             string                 `json:"filter"`
	Filters            []ItemFilter           `json:"filters"`
	FilterMode         string                 `json:"filter_mode"`
	OverrideStackSize  uint8                  `json:"override_stack_size"`
	DropPosition       Position               `json:"drop_position"`
	PickupPosition     Position               `json:"pickup_position"`
	RequestFilters     LogisticFilter         `json:"request_filters"`
	RequestFromBuffers bool                   `json:"request_from_buffers"`
	Parameters         SpeakerParameters      `json:"parameters"`
	AlertParameters    SpeakerAlertParameters `json:"alert_parameters"`
	AutoLaunch         bool                   `json:"auto_launch"`
	// See https://wiki.factorio.com/Types/GraphicsVariation
	Variation uint8  `json:"variation"`
	Color     Color  `json:"color"`
	Station   string `json:"station"`
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

type Signal struct {
	Signal SignalID `json:"signal"`
	Count  int      `json:"count"`
}

type SignalID struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type Icon struct {
	Signal SignalID `json:"signal"`
	Index  int      `json:"index"`
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
