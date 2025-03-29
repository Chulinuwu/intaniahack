package models

import (
	"math/rand"
	"time"
)

// EventsDB defines all available events in the game
var EventsDB []GameEvents

// InitEvents initializes all predefined events in the game
func InitEvents() {
	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Age 0-12 (Childhood)
	age0_12 := GameEvents{
		AgeIndex: 0,
		Events: []Event{
			{
				ID:          "age0_1",
				Type:        "positive",
				Title:       "First Day of School",
				Description: "You're starting elementary school for the first time!",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 3},
				},
			},
			{
				ID:          "age0_2",
				Type:        "choice",
				Title:       "Join After-School Activity",
				Description: "Your school offers after-school activities. Which one will you join?",
				AgeIndex:    0,
				Choices: []Choice{
					{
						ID:          "sports",
						Description: "Join a sports team",
						Effects: []Effect{
							{Stat: "happiness", Value: 5},
							{Stat: "relationship", Value: 5},
						},
					},
					{
						ID:          "music",
						Description: "Join the music club",
						Effects: []Effect{
							{Stat: "knowledge", Value: 5},
							{Stat: "happiness", Value: 3},
						},
					},
					{
						ID:          "science",
						Description: "Join the science club",
						Effects: []Effect{
							{Stat: "knowledge", Value: 8},
							{Stat: "relationship", Value: 2},
						},
					},
				},
			},
			{
				ID:          "age0_3",
				Type:        "neutral",
				Title:       "Family Vacation",
				Description: "Your family takes you on a memorable vacation",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "happiness", Value: 7},
					{Stat: "relationship", Value: 5},
				},
			},
		},
	}

	// Age 13-18 (Adolescence)
	age13_18 := GameEvents{
		AgeIndex: 1,
		Events: []Event{
			{
				ID:          "age1_1",
				Type:        "positive",
				Title:       "First High School Award",
				Description: "You received recognition for your performance in school!",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "knowledge", Value: 8},
					{Stat: "happiness", Value: 5},
				},
			},
			{
				ID:          "age1_2",
				Type:        "choice",
				Title:       "Part-time Job Opportunity",
				Description: "You have a chance to get a part-time job. What will you choose?",
				AgeIndex:    1,
				Choices: []Choice{
					{
						ID:          "retail",
						Description: "Work in retail",
						Effects: []Effect{
							{Stat: "money", Value: 3},
							{Stat: "relationship", Value: 4},
						},
					},
					{
						ID:          "tutor",
						Description: "Become a tutor",
						Effects: []Effect{
							{Stat: "money", Value: 2},
							{Stat: "knowledge", Value: 6},
						},
					},
					{
						ID:          "focus",
						Description: "Focus on studies",
						Effects: []Effect{
							{Stat: "knowledge", Value: 10},
							{Stat: "money", Value: -2},
						},
					},
				},
			},
			{
				ID:          "age1_3",
				Type:        "negative",
				Title:       "Friend Conflict",
				Description: "You had a disagreement with some friends",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "happiness", Value: -4},
					{Stat: "relationship", Value: -3},
				},
			},
		},
	}

	// Age 19-22 (Early Adulthood)
	age19_22 := GameEvents{
		AgeIndex: 2,
		Events: []Event{
			{
				ID:          "age2_1",
				Type:        "choice",
				Title:       "College Decision",
				Description: "You need to decide on your higher education path",
				AgeIndex:    2,
				Choices: []Choice{
					{
						ID:          "stem",
						Description: "Study STEM field",
						Effects: []Effect{
							{Stat: "knowledge", Value: 15},
							{Stat: "money", Value: -10},
							{Stat: "happiness", Value: -2},
						},
					},
					{
						ID:          "business",
						Description: "Study Business",
						Effects: []Effect{
							{Stat: "knowledge", Value: 10},
							{Stat: "money", Value: -5},
							{Stat: "relationship", Value: 5},
						},
					},
					{
						ID:          "arts",
						Description: "Study Arts & Humanities",
						Effects: []Effect{
							{Stat: "knowledge", Value: 12},
							{Stat: "money", Value: -8},
							{Stat: "happiness", Value: 5},
						},
					},
					{
						ID:          "work",
						Description: "Skip college and work full-time",
						Effects: []Effect{
							{Stat: "money", Value: 12},
							{Stat: "knowledge", Value: 2},
						},
					},
				},
			},
			{
				ID:          "age2_2",
				Type:        "positive",
				Title:       "First Relationship",
				Description: "You've started your first serious relationship",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 12},
				},
			},
			{
				ID:          "age2_3",
				Type:        "negative",
				Title:       "Financial Struggle",
				Description: "College expenses are piling up",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -8},
					{Stat: "happiness", Value: -3},
				},
			},
		},
	}

	// Age 23-39 (Adulthood)
	age23_39 := GameEvents{
		AgeIndex: 3,
		Events: []Event{
			{
				ID:          "age3_1",
				Type:        "choice",
				Title:       "Career Opportunity",
				Description: "You have different job offers. Which one will you take?",
				AgeIndex:    3,
				Choices: []Choice{
					{
						ID:          "corporate",
						Description: "Corporate job with high salary",
						Effects: []Effect{
							{Stat: "money", Value: 15},
							{Stat: "happiness", Value: -2},
						},
					},
					{
						ID:          "startup",
						Description: "Join a startup with potential",
						Effects: []Effect{
							{Stat: "money", Value: 5},
							{Stat: "knowledge", Value: 10},
							{Stat: "happiness", Value: 5},
						},
					},
					{
						ID:          "nonprofit",
						Description: "Work for a nonprofit organization",
						Effects: []Effect{
							{Stat: "money", Value: 3},
							{Stat: "happiness", Value: 10},
							{Stat: "relationship", Value: 8},
						},
					},
				},
			},
			{
				ID:          "age3_2",
				Type:        "positive",
				Title:       "Promotion",
				Description: "You got promoted at work!",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 12},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 8},
				},
			},
			{
				ID:          "age3_3",
				Type:        "choice",
				Title:       "Life Partner",
				Description: "You're considering settling down with a partner",
				AgeIndex:    3,
				Choices: []Choice{
					{
						ID:          "marry",
						Description: "Get married",
						Effects: []Effect{
							{Stat: "relationship", Value: 15},
							{Stat: "happiness", Value: 10},
							{Stat: "money", Value: -5},
						},
					},
					{
						ID:          "wait",
						Description: "Wait for the right time",
						Effects: []Effect{
							{Stat: "money", Value: 5},
							{Stat: "happiness", Value: -2},
						},
					},
					{
						ID:          "single",
						Description: "Choose to stay single",
						Effects: []Effect{
							{Stat: "money", Value: 8},
							{Stat: "relationship", Value: -5},
							{Stat: "happiness", Value: 3},
						},
					},
				},
			},
		},
	}

	// Age 40-59 (Middle Age)
	age40_59 := GameEvents{
		AgeIndex: 4,
		Events: []Event{
			{
				ID:          "age4_1",
				Type:        "choice",
				Title:       "Mid-life Investment",
				Description: "You have some savings to invest",
				AgeIndex:    4,
				Choices: []Choice{
					{
						ID:          "stocks",
						Description: "Invest in stocks",
						Effects: []Effect{
							{Stat: "money", Value: 20},
							{Stat: "happiness", Value: -3},
						},
					},
					{
						ID:          "property",
						Description: "Buy property",
						Effects: []Effect{
							{Stat: "money", Value: 15},
							{Stat: "happiness", Value: 5},
						},
					},
					{
						ID:          "education",
						Description: "Invest in further education",
						Effects: []Effect{
							{Stat: "knowledge", Value: 15},
							{Stat: "money", Value: -10},
						},
					},
				},
			},
			{
				ID:          "age4_2",
				Type:        "positive",
				Title:       "Family Vacation",
				Description: "You took your family on a wonderful vacation",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "happiness", Value: 12},
					{Stat: "relationship", Value: 10},
					{Stat: "money", Value: -8},
				},
			},
			{
				ID:          "age4_3",
				Type:        "negative",
				Title:       "Health Issue",
				Description: "You're experiencing some health problems",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "happiness", Value: -10},
					{Stat: "money", Value: -5},
				},
			},
		},
	}

	// Age 60-79 (Senior Years)
	age60_79 := GameEvents{
		AgeIndex: 5,
		Events: []Event{
			{
				ID:          "age5_1",
				Type:        "choice",
				Title:       "Retirement Planning",
				Description: "You need to decide about your retirement",
				AgeIndex:    5,
				Choices: []Choice{
					{
						ID:          "early",
						Description: "Retire early and enjoy life",
						Effects: []Effect{
							{Stat: "happiness", Value: 12},
							{Stat: "money", Value: -15},
						},
					},
					{
						ID:          "continue",
						Description: "Continue working part-time",
						Effects: []Effect{
							{Stat: "money", Value: 8},
							{Stat: "knowledge", Value: 5},
							{Stat: "happiness", Value: -2},
						},
					},
					{
						ID:          "volunteer",
						Description: "Retire and volunteer",
						Effects: []Effect{
							{Stat: "happiness", Value: 10},
							{Stat: "relationship", Value: 8},
							{Stat: "money", Value: -5},
						},
					},
				},
			},
			{
				ID:          "age5_2",
				Type:        "positive",
				Title:       "Grandchildren",
				Description: "Your children have given you grandchildren",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "happiness", Value: 15},
					{Stat: "relationship", Value: 12},
				},
			},
			{
				ID:          "age5_3",
				Type:        "neutral",
				Title:       "New Hobby",
				Description: "You've discovered a new hobby to fill your time",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "happiness", Value: 8},
					{Stat: "knowledge", Value: 5},
				},
			},
		},
	}

	// Age 80-100 (Golden Years)
	age80_100 := GameEvents{
		AgeIndex: 6,
		Events: []Event{
			{
				ID:          "age6_1",
				Type:        "choice",
				Title:       "Legacy Planning",
				Description: "You're thinking about your legacy",
				AgeIndex:    6,
				Choices: []Choice{
					{
						ID:          "charity",
						Description: "Donate to charity",
						Effects: []Effect{
							{Stat: "happiness", Value: 15},
							{Stat: "money", Value: -20},
						},
					},
					{
						ID:          "family",
						Description: "Leave everything to family",
						Effects: []Effect{
							{Stat: "relationship", Value: 15},
							{Stat: "happiness", Value: 10},
						},
					},
					{
						ID:          "memoir",
						Description: "Write your memoirs",
						Effects: []Effect{
							{Stat: "knowledge", Value: 15},
							{Stat: "happiness", Value: 8},
						},
					},
				},
			},
			{
				ID:          "age6_2",
				Type:        "positive",
				Title:       "Family Reunion",
				Description: "Your entire family gathered to celebrate you",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "happiness", Value: 20},
					{Stat: "relationship", Value: 15},
				},
			},
			{
				ID:          "age6_3",
				Type:        "neutral",
				Title:       "Reflection",
				Description: "You reflect on your life journey",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "knowledge", Value: 10},
					{Stat: "happiness", Value: 8},
				},
			},
		},
	}

	// Add all age ranges to the database
	EventsDB = []GameEvents{
		age0_12,
		age13_18,
		age19_22,
		age23_39,
		age40_59,
		age60_79,
		age80_100,
	}
}

// GetRandomEvent returns a random event for the given age index
func GetRandomEvent(ageIndex int) Event {
	for _, ageEvents := range EventsDB {
		if ageEvents.AgeIndex == ageIndex {
			return ageEvents.Events[rand.Intn(len(ageEvents.Events))]
		}
	}

	// Return a default event if none found (should never happen)
	return Event{
		ID:          "default",
		Type:        "neutral",
		Title:       "An Unexpected Day",
		Description: "Something unusual happened today",
		AgeIndex:    ageIndex,
		Effects: []Effect{
			{Stat: "happiness", Value: 2},
		},
	}
}
