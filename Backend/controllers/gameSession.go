package controllers

import (
	"backend-go/models"
	"backend-go/utils"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var gameSessionsMutex sync.Mutex
var gameSessions = make(map[string]*models.GameState)
var ctx = context.Background()

// InitGameSession initializes a new game session
func InitGameSession(roomID string, playerNames []string) *models.GameState {
	gameSessionsMutex.Lock()
	defer gameSessionsMutex.Unlock()

	gameID := utils.GenerateGameID()

	// Create players based on player names
	players := make([]models.Player, len(playerNames))
	for i, name := range playerNames {
		players[i] = models.Player{
			Username:     name,
			ProfileImage: fmt.Sprintf("https://api.dicebear.com/7.x/bottts/svg?seed=%s", name),
			Money:        50,
			Happiness:    50,
			Knowledge:    50,
			Relationship: 50,
			Events:       make(map[int][]models.Event),
			CurrentAge:   0,
			Turn:         i == 0, // First player starts
		}
	}

	// Create game state
	gameState := &models.GameState{
		RoomID:      roomID,
		GameID:      gameID,
		Started:     true,
		CurrentTurn: 0,
		CurrentAge:  0,
		Players:     players,
		Events:      []models.Event{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Store in memory
	gameSessions[gameID] = gameState

	// Store in Redis for persistence
	// jsonData, _ := json.Marshal(gameState)
	// config.RedisClient.Set(ctx, fmt.Sprintf("game:%s", gameID), jsonData, 24*time.Hour)

	return gameState
}

// HandleStartGame handles the start game request
func HandleStartGame(roomID string, room *models.Room) {
	roomsMutex.Lock()
	playerNames := getPlayersList(room)
	fmt.Println("Player names:", playerNames)
	roomsMutex.Unlock()

	// Initialize game events if not already initialized
	if len(models.EventsDB) == 0 {
		models.InitEvents()
	}

	// Initialize game session
	gameState := InitGameSession(roomID, playerNames)

	// Update room with game state
	room.Mutex.Lock()
	room.GameState = gameState
	room.Mutex.Unlock()

	// Notify all players that game has started
	message := gin.H{
		"event":        "game_initialized",
		"game_id":      gameState.GameID,
		"players":      gameStateToPlayerInfo(gameState),
		"current_age":  gameState.CurrentAge,
		"current_turn": gameState.CurrentTurn,
		"age_ranges":   models.GetAgeRanges(),
	}

	broadcastToRoom(room, message)

	// Start first player's turn
	startPlayerTurn(room, gameState.CurrentTurn)
}

// startPlayerTurn sends an event to indicate it's a player's turn
func startPlayerTurn(room *models.Room, playerIndex int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC in startPlayerTurn:", r)
		}
	}()

	fmt.Println("ENTERING startPlayerTurn for player:", playerIndex)

	if room == nil {
		fmt.Println("Error: Room is nil in startPlayerTurn")
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	fmt.Println("Starting turn for player index:", playerIndex)

	// Safety check for total players
	totalPlayers := len(room.Players) + 1 // +1 for host
	fmt.Println("Total players:", totalPlayers)

	if playerIndex >= totalPlayers {
		fmt.Println("Error: playerIndex out of bounds, resetting to 0")
		playerIndex = 0
	}

	gameState := room.GameState
	if gameState == nil {
		fmt.Println("Error: GameState is nil in startPlayerTurn")
		return
	}

	if len(gameState.Players) <= playerIndex {
		fmt.Println("Error: Not enough players in gameState")
		return
	}

	gameState.Mutex.Lock()
	gameState.CurrentTurn = playerIndex

	// Update player turns with safety check
	for i := range gameState.Players {
		if i < len(gameState.Players) {
			gameState.Players[i].Turn = (i == playerIndex)
		}
	}

	// Generate random event for the current player
	event := models.GetRandomEvent(gameState.CurrentAge)
	gameState.Mutex.Unlock()

	// Determine which connection to send to
	var targetConn *websocket.Conn
	if playerIndex == 0 {
		targetConn = room.Host
		if targetConn == nil {
			fmt.Println("Error: Host connection is nil")
			return
		}
	} else {
		if playerIndex-1 >= len(room.Players) || playerIndex-1 < 0 {
			fmt.Println("Error: Player index out of bounds:", playerIndex-1)
			return
		}
		targetConn = room.Players[playerIndex-1]
		if targetConn == nil {
			fmt.Println("Error: Player connection is nil for index:", playerIndex-1)
			return
		}
	}

	// Send turn start event to the player
	message := gin.H{
		"event":        "turn_start",
		"player_name":  gameState.Players[playerIndex].Username,
		"player_index": playerIndex,
		"event_data":   event,
	}

	if err := targetConn.WriteJSON(message); err != nil {
		fmt.Println("Error sending turn_start message:", err)
	}

	// Broadcast to other players that it's this player's turn
	otherPlayerMessage := gin.H{
		"event":        "waiting_for_turn",
		"player_name":  gameState.Players[playerIndex].Username,
		"player_index": playerIndex,
	}

	// Send to host if not their turn
	if playerIndex != 0 {
		if err := room.Host.WriteJSON(otherPlayerMessage); err != nil {
			fmt.Println("Error sending waiting_for_turn message to host:", err)
		}
	}

	// Send to other players
	for i, conn := range room.Players {
		if i != playerIndex-1 {
			if err := conn.WriteJSON(otherPlayerMessage); err != nil {
				fmt.Println("Error sending waiting_for_turn message to player:", err)
			}
		}
	}
}

// HandlePlayerChoice processes a player's choice
func HandlePlayerChoice(room *models.Room, playerIndex int, choiceID string, eventID string) {
	if room == nil {
		fmt.Println("Error: Room is nil in HandlePlayerChoice")
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	fmt.Println("Handling player choice:", choiceID, eventID)
	fmt.Println("Player index:", playerIndex)

	gameState := room.GameState
	if gameState == nil {
		fmt.Println("Error: GameState is nil in HandlePlayerChoice")
		return
	}

	gameState.Mutex.Lock()
	defer gameState.Mutex.Unlock()

	// Verify it's the player's turn
	if gameState.CurrentTurn != playerIndex {
		return
	}

	// Find the event that matches the event ID
	var currentEvent models.Event
	var found bool

	for _, ageEvents := range models.EventsDB {
		if ageEvents.AgeIndex == gameState.CurrentAge {
			for _, event := range ageEvents.Events {
				if event.ID == eventID {
					currentEvent = event
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	fmt.Println("Current Event:", currentEvent)
	fmt.Println("Found:", found)

	if !found {
		// Event not found, use a default effect
		ApplyEffectToPlayer(&gameState.Players[playerIndex], models.Effect{Stat: "happiness", Value: 2})
	} else {
		// If this is a choice event, find the selected choice
		if currentEvent.Type == "choice" {
			for _, choice := range currentEvent.Choices {
				if choice.ID == choiceID {
					// Apply all effects from this choice
					for _, effect := range choice.Effects {
						ApplyEffectToPlayer(&gameState.Players[playerIndex], effect)
					}
					break
				}
			}
		} else {
			// For non-choice events, apply the default effects
			for _, effect := range currentEvent.Effects {
				ApplyEffectToPlayer(&gameState.Players[playerIndex], effect)
			}
		}

		// Add this event to the player's history
		if gameState.Players[playerIndex].Events == nil {
			gameState.Players[playerIndex].Events = make(map[int][]models.Event)
		}
		gameState.Players[playerIndex].Events[gameState.CurrentAge] = append(
			gameState.Players[playerIndex].Events[gameState.CurrentAge],
			currentEvent,
		)
	}

	// Update game state in Redis
	// jsonData, _ := json.Marshal(gameState)
	// config.RedisClient.Set(ctx, fmt.Sprintf("game:%s", gameState.GameID), jsonData, 24*time.Hour)

	// Broadcast the turn result to all players
	broadcastTurnResult(room, gameState, playerIndex, currentEvent, choiceID)
}

func scheduleNextTurn(room *models.Room, nextPlayerIndex int) {
	fmt.Println("Scheduling next turn for player index:", nextPlayerIndex)
	gameState := room.GameState
	if nextPlayerIndex == 0 {
		gameState.CurrentAge++
	}
	fmt.Println("nextPlayerIndex:", nextPlayerIndex)
	fmt.Println("gameState.CurrentAge:", gameState.CurrentAge)
	if gameState.CurrentAge == 7 && nextPlayerIndex == 0 {
		fmt.Println("lingan guliguli")
		finalizeGame(room)
		return
	}
	startPlayerTurn(room, nextPlayerIndex)
}

// ApplyEffectToPlayer applies effects to a player's stats
func ApplyEffectToPlayer(player *models.Player, effect models.Effect) {
	switch effect.Stat {
	case "money":
		player.Money += effect.Value
		if player.Money < 0 {
			player.Money = 0
		} else if player.Money > 100 {
			player.Money = 100
		}
	case "happiness":
		player.Happiness += effect.Value
		if player.Happiness < 0 {
			player.Happiness = 0
		} else if player.Happiness > 100 {
			player.Happiness = 100
		}
	case "knowledge":
		player.Knowledge += effect.Value
		if player.Knowledge < 0 {
			player.Knowledge = 0
		} else if player.Knowledge > 100 {
			player.Knowledge = 100
		}
	case "relationship":
		player.Relationship += effect.Value
		if player.Relationship < 0 {
			player.Relationship = 0
		} else if player.Relationship > 100 {
			player.Relationship = 100
		}
	}
}

// broadcastTurnResult sends turn results to all players
func broadcastTurnResult(room *models.Room, gameState *models.GameState, playerIndex int, event models.Event, choiceID string) {
	player := gameState.Players[playerIndex]

	// Find the choice description if applicable
	var choiceDescription string
	if event.Type == "choice" {
		for _, choice := range event.Choices {
			if choice.ID == choiceID {
				choiceDescription = choice.Description
				break
			}
		}
	}

	message := gin.H{
		"event":        "turn_result",
		"player_name":  player.Username,
		"player_index": playerIndex,
		"player_stats": gin.H{
			"money":        player.Money,
			"happiness":    player.Happiness,
			"knowledge":    player.Knowledge,
			"relationship": player.Relationship,
		},
		"life_event": gin.H{
			"title":       event.Title,
			"description": event.Description,
			"type":        event.Type,
			"choice_made": choiceDescription,
		},
	}

	broadcastToRoom(room, message)
}

// checkAgeProgression checks if all players have completed their turns and advances age if needed
func checkAgeProgression(room *models.Room, playerIndex int) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	gameState := room.GameState
	if gameState == nil {
		return
	}

	gameState.Mutex.Lock()
	defer gameState.Mutex.Unlock()

	// totalPlayers := len(room.Players) + 1

	// If the current turn is back to player 0, we've completed a round
	if gameState.CurrentAge == 0 && gameState.CurrentTurn == 1 {
		// Reset all players' turns

	}
	// if gameState.CurrentTurn == 0 && playerIndex == 0 {
	// 	gameState.CurrentAge++

	// }

	if gameState.CurrentTurn == 0 && gameState.CurrentAge != 0 {
		// Move to the next age
		fmt.Println("Blah Balh Balh")
		fmt.Println("Current Age:", gameState.CurrentAge)
		fmt.Println("Current Turn:", gameState.CurrentTurn)
		if gameState.CurrentAge < 6 { // 0-6 are our 7 age ranges
			fmt.Println("Bleh Belh Belh")
			// Broadcast age advancement
			message := gin.H{
				"event":     "age_advanced",
				"age_index": gameState.CurrentAge,
				"age_range": models.GetAgeRanges()[gameState.CurrentAge],
				"players":   gameStateToPlayerInfo(gameState),
			}

			broadcastToRoom(room, message)

			// Update game state in Redis
			// jsonData, _ := json.Marshal(gameState)
			// config.RedisClient.Set(ctx, fmt.Sprintf("game:%s", gameState.GameID), jsonData, 24*time.Hour)
		}
		// else {
		// 	// Game is complete if we've gone through all ages
		// 	finalizeGame(room)
		// }
	}
}

// finalizeGame calculates final scores and ends the game
func finalizeGame(room *models.Room) {
	// TODO : room mutex is not working because it already lock in checkAgeProgression

	gameState := room.GameState
	if gameState == nil {
		return
	}

	// Calculate final scores for each player
	type PlayerResult struct {
		Username     string                    `json:"username"`
		ProfileImage string                    `json:"profile_image"`
		TotalScore   int                       `json:"total_score"`
		Stats        map[string]int            `json:"stats"`
		EventsByAge  map[string][]models.Event `json:"events_by_age"`
	}

	results := make([]PlayerResult, len(gameState.Players))
	for i, player := range gameState.Players {
		totalScore := player.Money + player.Happiness + player.Knowledge + player.Relationship

		// Convert events map to use string keys for JSON
		eventsByAge := make(map[string][]models.Event)
		for age, events := range player.Events {
			ageStr := fmt.Sprintf("%d", age)
			eventsByAge[ageStr] = events
		}

		results[i] = PlayerResult{
			Username:     player.Username,
			ProfileImage: player.ProfileImage,
			TotalScore:   totalScore,
			Stats: map[string]int{
				"money":        player.Money,
				"happiness":    player.Happiness,
				"knowledge":    player.Knowledge,
				"relationship": player.Relationship,
			},
			EventsByAge: eventsByAge,
		}

	}

	// Sort results by total score (highest first)
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].TotalScore > results[i].TotalScore {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	// Send game results to all players
	message := gin.H{
		"event":   "game_results",
		"results": results,
	}

	broadcastToRoom(room, message)

	// Clean up game resources
	gameSessionsMutex.Lock()
	delete(gameSessions, gameState.GameID)
	gameSessionsMutex.Unlock()
}

// gameStateToPlayerInfo converts game state to player info for the client
func gameStateToPlayerInfo(gameState *models.GameState) []gin.H {
	players := make([]gin.H, len(gameState.Players))
	for i, player := range gameState.Players {
		players[i] = gin.H{
			"username":      player.Username,
			"profile_image": player.ProfileImage,
			"money":         player.Money,
			"happiness":     player.Happiness,
			"knowledge":     player.Knowledge,
			"relationship":  player.Relationship,
			"current_age":   player.CurrentAge,
			"is_turn":       player.Turn,
		}
	}
	return players
}

// broadcastToRoom sends a message to all players in a room
func broadcastToRoom(room *models.Room, message gin.H) {
	fmt.Println("WoW:", message)
	jsonMessage, _ := json.Marshal(message)

	// Send to host
	room.Host.WriteMessage(websocket.TextMessage, jsonMessage)

	// Send to all players
	for _, conn := range room.Players {
		conn.WriteMessage(websocket.TextMessage, jsonMessage)
	}
}
