package models

import (
	"sync"
	"time"
)

// GameState holds the current state of a game
type GameState struct {
	RoomID      string
	GameID      string
	Started     bool
	CurrentTurn int // Index of player whose turn it is
	CurrentAge  int // Current age index (0-6)
	Players     []Player
	Events      []Event
	Mutex       sync.Mutex
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Player represents a player's state in the game
type Player struct {
	Username     string
	ProfileImage string
	Money        int
	Happiness    int
	Knowledge    int
	Relationship int
	Events       map[int][]Event // Map of age index to events
	CurrentAge   int
	Turn         bool // Whether it's this player's turn
}

// AgeRange represents a range of ages in the game
type AgeRange struct {
	Label       string
	MinAge      int
	MaxAge      int
	Description string
}

// Event represents a life event in the game
type Event struct {
	ID          string
	Type        string // "positive", "negative", "neutral", "choice"
	Description string
	AgeIndex    int
	Effects     []Effect
	Choices     []Choice
}

// Effect represents a stat change from an event
type Effect struct {
	Stat  string // "money", "happiness", "knowledge", "relationship"
	Value int    // Can be positive or negative
}

// Choice represents a decision a player can make
type Choice struct {
	ID          string
	Description string
	Effects     []Effect
}

// GetAgeRanges returns all age ranges in the game
func GetAgeRanges() []AgeRange {
	return []AgeRange{
		{Label: "0 - 12", MinAge: 0, MaxAge: 12, Description: "Childhood"},
		{Label: "13 - 18", MinAge: 13, MaxAge: 18, Description: "Adolescence"},
		{Label: "19 - 22", MinAge: 19, MaxAge: 22, Description: "Early Adulthood"},
		{Label: "23 - 39", MinAge: 23, MaxAge: 39, Description: "Adulthood"},
		{Label: "40 - 59", MinAge: 40, MaxAge: 59, Description: "Middle Age"},
		{Label: "60 - 79", MinAge: 60, MaxAge: 79, Description: "Senior Years"},
		{Label: "80 - 100", MinAge: 80, MaxAge: 100, Description: "Golden Years"},
	}
}

// GameEvents stores predefined events for each age range
type GameEvents struct {
	AgeIndex int
	Events   []Event
}
