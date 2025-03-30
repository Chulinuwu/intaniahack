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

	// Elementary School (Age 0-12)
	age0_12 := GameEvents{
		AgeIndex: 0,
		Events: []Event{
			{
				ID:          "age0_1",
				Type:        "knowledge",
				Description: "You're starting elementary school for the first time!",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 3},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age0_2",
				Type:        "relationship",
				Description: "Join a sports team",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 5},
				},
			},
			{
				ID:          "age0_3",
				Type:        "knowledge",
				Description: "Join the music club",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 3},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age0_4",
				Type:        "knowledge",
				Description: "Join the science club",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 8},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: 2},
				},
			},
			{
				ID:          "age0_5",
				Type:        "happiness",
				Description: "Your family takes you on a memorable vacation",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 7},
					{Stat: "relationship", Value: 5},
				},
			},
			{
				ID:          "age0_6",
				Type:        "knowledge",
				Description: "Spelling Bee",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 7},
					{Stat: "happiness", Value: 3},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age0_7",
				Type:        "relationship",
				Description: "Sports Competition",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 6},
					{Stat: "relationship", Value: 4},
				},
			},
			{
				ID:          "age0_8",
				Type:        "happiness",
				Description: "Art Contest",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 4},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age0_9",
				Type:        "happiness",
				Description: "You've caught a bad flu and need to stay home from school",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: -2},
					{Stat: "happiness", Value: -4},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age0_10",
				Type:        "happiness",
				Description: "Have a big party with all your classmates",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: -3},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 8},
					{Stat: "relationship", Value: 6},
				},
			},
			{
				ID:          "age0_11",
				Type:        "happiness",
				Description: "Quiet celebration with family",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 3},
				},
			},
			{
				ID:          "age0_12",
				Type:        "happiness",
				Description: "Visit a theme park with your best friends",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "money", Value: -5},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 4},
				},
			},
		},
	}

	// Teenage Years (Age 13-18)
	age13_18 := GameEvents{
		AgeIndex: 1,
		Events: []Event{
			{
				ID:          "age1_1",
				Type:        "knowledge",
				Description: "You received recognition for your performance in school!",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 8},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age1_2",
				Type:        "money",
				Description: "Work in retail",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 3},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: 4},
				},
			},
			{
				ID:          "age1_3",
				Type:        "money",
				Description: "Become a tutor",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 2},
					{Stat: "knowledge", Value: 6},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age1_4",
				Type:        "knowledge",
				Description: "Focus on studies",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: -2},
					{Stat: "knowledge", Value: 10},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age1_5",
				Type:        "relationship",
				Description: "You had a disagreement with some friends",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: -4},
					{Stat: "relationship", Value: -3},
				},
			},
			{
				ID:          "age1_6",
				Type:        "relationship",
				Description: "Ask your crush to be your date",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 7},
					{Stat: "relationship", Value: 8},
				},
			},
			{
				ID:          "age1_7",
				Type:        "relationship",
				Description: "Go with a group of friends",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 5},
				},
			},
			{
				ID:          "age1_8",
				Type:        "knowledge",
				Description: "Skip the dance to study",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 7},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: -2},
				},
			},
			{
				ID:          "age1_9",
				Type:        "happiness",
				Description: "You discovered a hobby that you're passionate about",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 4},
					{Stat: "happiness", Value: 8},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age1_10",
				Type:        "relationship",
				Description: "Give in to peer pressure",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 2},
					{Stat: "happiness", Value: -3},
					{Stat: "relationship", Value: 4},
				},
			},
			{
				ID:          "age1_11",
				Type:        "relationship",
				Description: "Refuse and stand your ground",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 3},
					{Stat: "happiness", Value: 2},
					{Stat: "relationship", Value: -3},
				},
			},
			{
				ID:          "age1_12",
				Type:        "knowledge",
				Description: "Suggest a safer alternative",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 4},
					{Stat: "happiness", Value: 3},
					{Stat: "relationship", Value: 2},
				},
			},
		},
	}

	// College Years (Age 19-22)
	age19_22 := GameEvents{
		AgeIndex: 2,
		Events: []Event{
			{
				ID:          "age2_1",
				Type:        "knowledge",
				Description: "Study STEM field",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -10},
					{Stat: "knowledge", Value: 15},
					{Stat: "happiness", Value: -2},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age2_2",
				Type:        "knowledge",
				Description: "Study Business",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -5},
					{Stat: "knowledge", Value: 10},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: 5},
				},
			},
			{
				ID:          "age2_3",
				Type:        "happiness",
				Description: "Study Arts & Humanities",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -8},
					{Stat: "knowledge", Value: 12},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age2_4",
				Type:        "money",
				Description: "Skip college and work full-time",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: 12},
					{Stat: "knowledge", Value: 2},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age2_5",
				Type:        "relationship",
				Description: "You've started your first serious relationship",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 12},
				},
			},
			{
				ID:          "age2_6",
				Type:        "money",
				Description: "College expenses are piling up",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -8},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: -3},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age2_7",
				Type:        "knowledge",
				Description: "Go study abroad",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -12},
					{Stat: "knowledge", Value: 12},
					{Stat: "happiness", Value: 8},
					{Stat: "relationship", Value: 4},
				},
			},
			{
				ID:          "age2_8",
				Type:        "knowledge",
				Description: "Stay at your current university",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -4},
					{Stat: "knowledge", Value: 6},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: 3},
				},
			},
			{
				ID:          "age2_9",
				Type:        "knowledge",
				Description: "Participate in an online international program",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -6},
					{Stat: "knowledge", Value: 8},
					{Stat: "happiness", Value: 4},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age2_10",
				Type:        "knowledge",
				Description: "You've been recognized for your academic accomplishments",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 10},
					{Stat: "happiness", Value: 7},
					{Stat: "relationship", Value: 3},
				},
			},
			{
				ID:          "age2_11",
				Type:        "relationship",
				Description: "Live in the dormitories",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -8},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 8},
				},
			},
			{
				ID:          "age2_12",
				Type:        "relationship",
				Description: "Rent an apartment with roommates",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: -6},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 4},
					{Stat: "relationship", Value: 6},
				},
			},
			{
				ID:          "age2_13",
				Type:        "money",
				Description: "Live with your parents to save money",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "money", Value: 10},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: -3},
					{Stat: "relationship", Value: -2},
				},
			},
		},
	}

	// Young Adult (Age 23-39)
	age23_39 := GameEvents{
		AgeIndex: 3,
		Events: []Event{
			{
				ID:          "age3_1",
				Type:        "money",
				Description: "Corporate job with high salary",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 15},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: -2},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age3_2",
				Type:        "knowledge",
				Description: "Join a startup with potential",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 5},
					{Stat: "knowledge", Value: 10},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age3_3",
				Type:        "happiness",
				Description: "Work for a nonprofit organization",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 3},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 8},
				},
			},
			{
				ID:          "age3_4",
				Type:        "money",
				Description: "You got promoted at work!",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 12},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 8},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age3_5",
				Type:        "relationship",
				Description: "Get married",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: -5},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 15},
				},
			},
			{
				ID:          "age3_6",
				Type:        "money",
				Description: "Wait for the right time",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 5},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: -2},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age3_7",
				Type:        "money",
				Description: "Choose to stay single",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 8},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 3},
					{Stat: "relationship", Value: -5},
				},
			},
			{
				ID:          "age3_8",
				Type:        "money",
				Description: "Accept the relocation offer",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 12},
					{Stat: "knowledge", Value: 8},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: -4},
				},
			},
			{
				ID:          "age3_9",
				Type:        "relationship",
				Description: "Stay in your current location",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 4},
					{Stat: "relationship", Value: 6},
				},
			},
			{
				ID:          "age3_10",
				Type:        "money",
				Description: "Negotiate for remote work options",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 6},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: 3},
				},
			},
			{
				ID:          "age3_11",
				Type:        "money",
				Description: "You've experienced a setback in your career",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: -10},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: -8},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age3_12",
				Type:        "relationship",
				Description: "Have children now",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: -15},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 12},
				},
			},
			{
				ID:          "age3_13",
				Type:        "money",
				Description: "Wait a few more years",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 8},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: -3},
				},
			},
			{
				ID:          "age3_14",
				Type:        "money",
				Description: "Decide not to have children",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: 12},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 3},
					{Stat: "relationship", Value: -6},
				},
			},
		},
	}

	// Middle Age (Age 40-59)
	age40_59 := GameEvents{
		AgeIndex: 4,
		Events: []Event{
			{
				ID:          "age4_1",
				Type:        "money",
				Description: "Invest in stocks",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: 20},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: -3},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age4_2",
				Type:        "money",
				Description: "Buy property",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: 15},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age4_3",
				Type:        "knowledge",
				Description: "Invest in further education",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: -10},
					{Stat: "knowledge", Value: 15},
					{Stat: "happiness", Value: 0},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age4_4",
				Type:        "happiness",
				Description: "You took your family on a wonderful vacation",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: -8},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 12},
					{Stat: "relationship", Value: 10},
				},
			},
			{
				ID:          "age4_5",
				Type:        "happiness",
				Description: "You're experiencing some health problems",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: -5},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: -10},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age4_6",
				Type:        "knowledge",
				Description: "Start a completely new career path",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: -10},
					{Stat: "knowledge", Value: 15},
					{Stat: "happiness", Value: 12},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age4_7",
				Type:        "money",
				Description: "Stay in your current career but seek advancement",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: 10},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 3},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age4_8",
				Type:        "money",
				Description: "Start a side business while keeping your job",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: 8},
					{Stat: "knowledge", Value: 10},
					{Stat: "happiness", Value: 6},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age4_9",
				Type:        "happiness",
				Description: "Your children have achieved significant milestones",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 15},
					{Stat: "relationship", Value: 8},
				},
			},
			{
				ID:          "age4_10",
				Type:        "money",
				Description: "Downsize to a smaller home",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: 15},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age4_11",
				Type:        "happiness",
				Description: "Renovate your current home",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: -10},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age4_12",
				Type:        "happiness",
				Description: "Upgrade to your dream home",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "money", Value: -20},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 15},
					{Stat: "relationship", Value: 5},
				},
			},
		},
	}

	// Senior Years (Age 60-79)
	age60_79 := GameEvents{
		AgeIndex: 5,
		Events: []Event{
			{
				ID:          "age5_1",
				Type:        "happiness",
				Description: "Retire early and enjoy life",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: -15},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 12},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age5_2",
				Type:        "money",
				Description: "Continue working part-time",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: 8},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: -2},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age5_3",
				Type:        "happiness",
				Description: "Retire and volunteer",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: -5},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 8},
				},
			},
			{
				ID:          "age5_4",
				Type:        "relationship",
				Description: "Your children have given you grandchildren",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 15},
					{Stat: "relationship", Value: 12},
				},
			},
			{
				ID:          "age5_5",
				Type:        "happiness",
				Description: "You've discovered a new hobby to fill your time",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 8},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age5_6",
				Type:        "relationship",
				Description: "Move to a retirement community",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: -12},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 8},
					{Stat: "relationship", Value: 10},
				},
			},
			{
				ID:          "age5_7",
				Type:        "happiness",
				Description: "Stay in your home with modifications",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: -5},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age5_8",
				Type:        "relationship",
				Description: "Move closer to family members",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: -8},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 7},
					{Stat: "relationship", Value: 15},
				},
			},
			{
				ID:          "age5_9",
				Type:        "happiness",
				Description: "You're dealing with some age-related health challenges",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: -10},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: -8},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age5_10",
				Type:        "happiness",
				Description: "Travel the world extensively",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: -18},
					{Stat: "knowledge", Value: 12},
					{Stat: "happiness", Value: 15},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age5_11",
				Type:        "happiness",
				Description: "Take occasional trips within the country",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: -8},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: 8},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age5_12",
				Type:        "relationship",
				Description: "Focus on local activities and community",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "money", Value: -3},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 7},
					{Stat: "relationship", Value: 12},
				},
			},
		},
	}

	// Age 80-100 Events (ageIndex: 6)
	age80_100 := GameEvents{
		AgeIndex: 6,
		Events: []Event{
			{
				ID:          "age6_1",
				Type:        "happiness",
				Description: "Donate to charity",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "money", Value: -20},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 15},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age6_2",
				Type:        "relationship",
				Description: "Leave everything to family",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 10},
					{Stat: "relationship", Value: 15},
				},
			},
			{
				ID:          "age6_3",
				Type:        "knowledge",
				Description: "Write your memoirs",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 15},
					{Stat: "happiness", Value: 8},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age6_4",
				Type:        "relationship",
				Description: "Your entire family gathered to celebrate you",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 20},
					{Stat: "relationship", Value: 15},
				},
			},
			{
				ID:          "age6_5",
				Type:        "knowledge",
				Description: "You reflect on your life journey",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 10},
					{Stat: "happiness", Value: 8},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age6_6",
				Type:        "happiness",
				Description: "Seek advanced medical treatments",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "money", Value: -15},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 5},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age6_7",
				Type:        "happiness",
				Description: "Focus on comfort and quality of life",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 0},
					{Stat: "happiness", Value: 12},
					{Stat: "relationship", Value: 8},
				},
			},
			{
				ID:          "age6_8",
				Type:        "knowledge",
				Description: "Try holistic and alternative approaches",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "money", Value: -5},
					{Stat: "knowledge", Value: 8},
					{Stat: "happiness", Value: 7},
					{Stat: "relationship", Value: 0},
				},
			},
			{
				ID:          "age6_9",
				Type:        "happiness",
				Description: "You've lost a close friend from your generation",
				AgeIndex:    6,
				Effects: []Effect{
					{Stat: "money", Value: 0},
					{Stat: "knowledge", Value: 5},
					{Stat: "happiness", Value: -10},
					{Stat: "relationship", Value: -5},
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
		Type:        "happiness",
		Title:       "An Unexpected Day",
		Description: "Something unusual happened today",
		AgeIndex:    ageIndex,
		Effects: []Effect{
			{Stat: "happiness", Value: 2},
		},
	}
}
