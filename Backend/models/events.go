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
			{
				ID:          "age0_4",
				Type:        "choice",
				Title:       "School Competition",
				Description: "There's a competition at school. What would you like to participate in?",
				AgeIndex:    0,
				Choices: []Choice{
					{
						ID:          "spelling",
						Description: "Spelling Bee",
						Effects: []Effect{
							{Stat: "knowledge", Value: 7},
							{Stat: "happiness", Value: 3},
						},
					},
					{
						ID:          "sports_competition",
						Description: "Sports Competition",
						Effects: []Effect{
							{Stat: "happiness", Value: 6},
							{Stat: "relationship", Value: 4},
						},
					},
					{
						ID:          "art",
						Description: "Art Contest",
						Effects: []Effect{
							{Stat: "knowledge", Value: 4},
							{Stat: "happiness", Value: 5},
						},
					},
				},
			},
			{
				ID:          "age0_5",
				Type:        "negative",
				Title:       "Childhood Illness",
				Description: "You've caught a bad flu and need to stay home from school",
				AgeIndex:    0,
				Effects: []Effect{
					{Stat: "happiness", Value: -4},
					{Stat: "knowledge", Value: -2},
				},
			},
			{
				ID:          "age0_6",
				Type:        "choice",
				Title:       "Birthday Party",
				Description: "It's your birthday! How would you like to celebrate?",
				AgeIndex:    0,
				Choices: []Choice{
					{
						ID:          "big_party",
						Description: "Have a big party with all your classmates",
						Effects: []Effect{
							{Stat: "happiness", Value: 8},
							{Stat: "relationship", Value: 6},
							{Stat: "money", Value: -3},
						},
					},
					{
						ID:          "family_only",
						Description: "Quiet celebration with family",
						Effects: []Effect{
							{Stat: "happiness", Value: 5},
							{Stat: "relationship", Value: 3},
						},
					},
					{
						ID:          "theme_park",
						Description: "Visit a theme park with your best friends",
						Effects: []Effect{
							{Stat: "happiness", Value: 10},
							{Stat: "relationship", Value: 4},
							{Stat: "money", Value: -5},
						},
					},
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
			{
				ID:          "age1_4",
				Type:        "choice",
				Title:       "School Dance",
				Description: "The school dance is coming up. What will you do?",
				AgeIndex:    1,
				Choices: []Choice{
					{
						ID:          "ask_crush",
						Description: "Ask your crush to be your date",
						Effects: []Effect{
							{Stat: "relationship", Value: 8},
							{Stat: "happiness", Value: 7},
						},
					},
					{
						ID:          "go_friends",
						Description: "Go with a group of friends",
						Effects: []Effect{
							{Stat: "relationship", Value: 5},
							{Stat: "happiness", Value: 5},
						},
					},
					{
						ID:          "skip_dance",
						Description: "Skip the dance to study",
						Effects: []Effect{
							{Stat: "knowledge", Value: 7},
							{Stat: "relationship", Value: -2},
						},
					},
				},
			},
			{
				ID:          "age1_5",
				Type:        "positive",
				Title:       "Found a Hobby",
				Description: "You discovered a hobby that you're passionate about",
				AgeIndex:    1,
				Effects: []Effect{
					{Stat: "happiness", Value: 8},
					{Stat: "knowledge", Value: 4},
				},
			},
			{
				ID:          "age1_6",
				Type:        "choice",
				Title:       "Peer Pressure",
				Description: "Your friends are pressuring you to try something risky. What will you do?",
				AgeIndex:    1,
				Choices: []Choice{
					{
						ID:          "give_in",
						Description: "Give in to peer pressure",
						Effects: []Effect{
							{Stat: "relationship", Value: 4},
							{Stat: "happiness", Value: -3},
							{Stat: "knowledge", Value: 2},
						},
					},
					{
						ID:          "refuse",
						Description: "Refuse and stand your ground",
						Effects: []Effect{
							{Stat: "relationship", Value: -3},
							{Stat: "happiness", Value: 2},
							{Stat: "knowledge", Value: 3},
						},
					},
					{
						ID:          "compromise",
						Description: "Suggest a safer alternative",
						Effects: []Effect{
							{Stat: "relationship", Value: 2},
							{Stat: "happiness", Value: 3},
							{Stat: "knowledge", Value: 4},
						},
					},
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
			{
				ID:          "age2_4",
				Type:        "choice",
				Title:       "Study Abroad Opportunity",
				Description: "You have a chance to study abroad for a semester. Will you go?",
				AgeIndex:    2,
				Choices: []Choice{
					{
						ID:          "go_abroad",
						Description: "Go study abroad",
						Effects: []Effect{
							{Stat: "knowledge", Value: 12},
							{Stat: "relationship", Value: 4},
							{Stat: "money", Value: -12},
							{Stat: "happiness", Value: 8},
						},
					},
					{
						ID:          "stay_home",
						Description: "Stay at your current university",
						Effects: []Effect{
							{Stat: "knowledge", Value: 6},
							{Stat: "relationship", Value: 3},
							{Stat: "money", Value: -4},
						},
					},
					{
						ID:          "online_program",
						Description: "Participate in an online international program",
						Effects: []Effect{
							{Stat: "knowledge", Value: 8},
							{Stat: "money", Value: -6},
							{Stat: "happiness", Value: 4},
						},
					},
				},
			},
			{
				ID:          "age2_5",
				Type:        "positive",
				Title:       "Academic Achievement",
				Description: "You've been recognized for your academic accomplishments",
				AgeIndex:    2,
				Effects: []Effect{
					{Stat: "knowledge", Value: 10},
					{Stat: "happiness", Value: 7},
					{Stat: "relationship", Value: 3},
				},
			},
			{
				ID:          "age2_6",
				Type:        "choice",
				Title:       "Living Situation",
				Description: "You need to decide where to live during your college years",
				AgeIndex:    2,
				Choices: []Choice{
					{
						ID:          "dorms",
						Description: "Live in the dormitories",
						Effects: []Effect{
							{Stat: "relationship", Value: 8},
							{Stat: "money", Value: -8},
							{Stat: "happiness", Value: 5},
						},
					},
					{
						ID:          "apartment",
						Description: "Rent an apartment with roommates",
						Effects: []Effect{
							{Stat: "relationship", Value: 6},
							{Stat: "money", Value: -6},
							{Stat: "happiness", Value: 4},
						},
					},
					{
						ID:          "parents",
						Description: "Live with your parents to save money",
						Effects: []Effect{
							{Stat: "money", Value: 10},
							{Stat: "relationship", Value: -2},
							{Stat: "happiness", Value: -3},
						},
					},
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
			{
				ID:          "age3_4",
				Type:        "choice",
				Title:       "Relocation Opportunity",
				Description: "You have an opportunity to relocate for work",
				AgeIndex:    3,
				Choices: []Choice{
					{
						ID:          "move",
						Description: "Accept the relocation offer",
						Effects: []Effect{
							{Stat: "money", Value: 12},
							{Stat: "knowledge", Value: 8},
							{Stat: "relationship", Value: -4},
						},
					},
					{
						ID:          "stay",
						Description: "Stay in your current location",
						Effects: []Effect{
							{Stat: "relationship", Value: 6},
							{Stat: "happiness", Value: 4},
						},
					},
					{
						ID:          "negotiate",
						Description: "Negotiate for remote work options",
						Effects: []Effect{
							{Stat: "money", Value: 6},
							{Stat: "knowledge", Value: 5},
							{Stat: "relationship", Value: 3},
						},
					},
				},
			},
			{
				ID:          "age3_5",
				Type:        "negative",
				Title:       "Career Setback",
				Description: "You've experienced a setback in your career",
				AgeIndex:    3,
				Effects: []Effect{
					{Stat: "money", Value: -10},
					{Stat: "happiness", Value: -8},
					{Stat: "knowledge", Value: 5},
				},
			},
			{
				ID:          "age3_6",
				Type:        "choice",
				Title:       "Family Planning",
				Description: "You and your partner are discussing starting a family",
				AgeIndex:    3,
				Choices: []Choice{
					{
						ID:          "have_children",
						Description: "Have children now",
						Effects: []Effect{
							{Stat: "relationship", Value: 12},
							{Stat: "happiness", Value: 10},
							{Stat: "money", Value: -15},
						},
					},
					{
						ID:          "wait_children",
						Description: "Wait a few more years",
						Effects: []Effect{
							{Stat: "money", Value: 8},
							{Stat: "knowledge", Value: 5},
							{Stat: "relationship", Value: -3},
						},
					},
					{
						ID:          "no_children",
						Description: "Decide not to have children",
						Effects: []Effect{
							{Stat: "money", Value: 12},
							{Stat: "relationship", Value: -6},
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
			{
				ID:          "age4_4",
				Type:        "choice",
				Title:       "Career Change",
				Description: "You're considering a mid-life career change",
				AgeIndex:    4,
				Choices: []Choice{
					{
						ID:          "new_career",
						Description: "Start a completely new career path",
						Effects: []Effect{
							{Stat: "knowledge", Value: 15},
							{Stat: "money", Value: -10},
							{Stat: "happiness", Value: 12},
						},
					},
					{
						ID:          "stay_career",
						Description: "Stay in your current career but seek advancement",
						Effects: []Effect{
							{Stat: "money", Value: 10},
							{Stat: "knowledge", Value: 5},
							{Stat: "happiness", Value: 3},
						},
					},
					{
						ID:          "side_business",
						Description: "Start a side business while keeping your job",
						Effects: []Effect{
							{Stat: "money", Value: 8},
							{Stat: "knowledge", Value: 10},
							{Stat: "happiness", Value: 6},
						},
					},
				},
			},
			{
				ID:          "age4_5",
				Type:        "positive",
				Title:       "Children's Success",
				Description: "Your children have achieved significant milestones",
				AgeIndex:    4,
				Effects: []Effect{
					{Stat: "happiness", Value: 15},
					{Stat: "relationship", Value: 8},
				},
			},
			{
				ID:          "age4_6",
				Type:        "choice",
				Title:       "Living Arrangements",
				Description: "You're considering changing your living situation",
				AgeIndex:    4,
				Choices: []Choice{
					{
						ID:          "downsize",
						Description: "Downsize to a smaller home",
						Effects: []Effect{
							{Stat: "money", Value: 15},
							{Stat: "happiness", Value: 5},
						},
					},
					{
						ID:          "renovate",
						Description: "Renovate your current home",
						Effects: []Effect{
							{Stat: "money", Value: -10},
							{Stat: "happiness", Value: 10},
						},
					},
					{
						ID:          "dream_home",
						Description: "Upgrade to your dream home",
						Effects: []Effect{
							{Stat: "money", Value: -20},
							{Stat: "happiness", Value: 15},
							{Stat: "relationship", Value: 5},
						},
					},
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
			{
				ID:          "age5_4",
				Type:        "choice",
				Title:       "Senior Living Options",
				Description: "You're considering your living arrangements as you age",
				AgeIndex:    5,
				Choices: []Choice{
					{
						ID:          "retirement_community",
						Description: "Move to a retirement community",
						Effects: []Effect{
							{Stat: "relationship", Value: 10},
							{Stat: "happiness", Value: 8},
							{Stat: "money", Value: -12},
						},
					},
					{
						ID:          "stay_home",
						Description: "Stay in your home with modifications",
						Effects: []Effect{
							{Stat: "money", Value: -5},
							{Stat: "happiness", Value: 10},
						},
					},
					{
						ID:          "move_family",
						Description: "Move closer to family members",
						Effects: []Effect{
							{Stat: "relationship", Value: 15},
							{Stat: "money", Value: -8},
							{Stat: "happiness", Value: 7},
						},
					},
				},
			},
			{
				ID:          "age5_5",
				Type:        "negative",
				Title:       "Health Concerns",
				Description: "You're dealing with some age-related health challenges",
				AgeIndex:    5,
				Effects: []Effect{
					{Stat: "happiness", Value: -8},
					{Stat: "money", Value: -10},
				},
			},
			{
				ID:          "age5_6",
				Type:        "choice",
				Title:       "Travel Plans",
				Description: "You're planning how to spend your retirement leisure time",
				AgeIndex:    5,
				Choices: []Choice{
					{
						ID:          "world_travel",
						Description: "Travel the world extensively",
						Effects: []Effect{
							{Stat: "happiness", Value: 15},
							{Stat: "knowledge", Value: 12},
							{Stat: "money", Value: -18},
						},
					},
					{
						ID:          "local_travel",
						Description: "Take occasional trips within the country",
						Effects: []Effect{
							{Stat: "happiness", Value: 8},
							{Stat: "knowledge", Value: 5},
							{Stat: "money", Value: -8},
						},
					},
					{
						ID:          "stay_active",
						Description: "Focus on local activities and community",
						Effects: []Effect{
							{Stat: "relationship", Value: 12},
							{Stat: "happiness", Value: 7},
							{Stat: "money", Value: -3},
						},
					},
				},
			},
		},
	}

	// Age 80-100 (Golden Years)
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
		{
			ID:          "age6_4",
			Type:        "choice",
			Title:       "Healthcare Decisions",
			Description: "You need to make important healthcare decisions",
			AgeIndex:    6,
			Choices: []Choice{
				{
					ID:          "advanced_treatment",
					Description: "Seek advanced medical treatments",
					Effects: []Effect{
						{Stat: "happiness", Value: 5},
						{Stat: "money", Value: -15},
					},
				},
				{
					ID:          "comfort_care",
					Description: "Focus on comfort and quality of life",
					Effects: []Effect{
						{Stat: "happiness", Value: 12},
						{Stat: "relationship", Value: 8},
					},
				},
				{
					ID:          "holistic",
					Description: "Try holistic and alternative approaches",
					Effects: []Effect{
						{Stat: "knowledge", Value: 8},
						{Stat: "happiness", Value: 7},
						{Stat: "money", Value: -5},
					},
				},
			},
		},
		{
			ID:          "age6_5",
			Type:        "negative",
			Title:       "Loss of a Friend",
			Description: "You've lost a close friend from your generation",
			AgeIndex:    6,
			Effects: []Effect{
				{Stat: "happiness", Value: -10},
				{Stat: "relationship", Value: -5},
				{Stat: "knowledge", Value: 5},
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
