<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import CompetiterCard from '../../components/CompetiterCard.svelte';
	import PlayCard from '../../components/PlayCard.svelte';
	import { iconMapColor } from '$lib/utils/iconMapColor';
	import TimeLeft from '../../components/TimeLeft.svelte';
	import { getToken, connectWebSocket } from '$lib/auth';
	import { goto } from '$app/navigation';
	import bg from '$lib/assets/image/bg.gif';
	
	// Lobby variables
	let token = '';
	let ws: WebSocket | null = null;
	let roomId = '';
	let inputRoomId = '';
	let players: string[] = [];
	let host = '';
	let topic = '';
	let gameStatus = '';
	let error = '';
	let currentUsername = '';
	let selectedTopic = '';

	// Game state flag
	let isGameStarted = false;
	let turnCount = 0; // เพิ่มตัวนับรอบการเล่น
	// Game over flag and results
	let isGameOver = false;
	let gameResults: any[] = [];

	// Game variables
	let handCards: any[] = [];
	let currentAgeIndex = 0;
	
	$: currentAge = ageRanges[currentAgeIndex];

	// Current event information
	let currentEvent: { Type: string; Choices: any[]; ID: any; Description: any; } | null = null;

	// Topics for host to choose
	const topics = [
		'GET THE MOST Money',
		'GET THE MOST Happiness',
		'GET THE MOST Knowledge',
		'GET THE MOST Relationship'
	];

	// Initialize age-based card slots
	let droppedCards0_12 = Array(5).fill(null);
	let droppedCards13_18 = Array(5).fill(null);
	let droppedCards19_22 = Array(5).fill(null);
	let droppedCards23_39 = Array(5).fill(null);
	let droppedCards40_59 = Array(5).fill(null);
	let droppedCards60_79 = Array(5).fill(null);
	let droppedCards80_100 = Array(5).fill(null);

	let ageRanges = [
		{ label: '0 - 12', data: droppedCards0_12 },
		{ label: '13 - 18', data: droppedCards13_18 },
		{ label: '19 - 22', data: droppedCards19_22 },
		{ label: '23 - 39', data: droppedCards23_39 },
		{ label: '40 - 59', data: droppedCards40_59 },
		{ label: '60 - 79', data: droppedCards60_79 },
		{ label: '80 - 100', data: droppedCards80_100 }
	];

	// Card selection variables
	let selectedCard: null = null;
	let selectedCardIndex: number | null = null;
	let selectedCardSource: string | null = null;
	let selectedAgeIndex: number | null = null;

	// Variables to track turn state
	let isMyTurn = false;
	let currentEventChoices = [];
	let allPlayerStats: any[] = [];

	// Game notification system
	let notification = '';
	let notificationTimer: string | number | NodeJS.Timeout | null | undefined = null;

	// Game progress
	let currentAgeProgress = 0;
	let totalAges = 7;

	onMount(async () => {
		token = getToken() || '';
		if (!token) {
			goto('/login');
			return;
		}

		const response = await fetch(`http://${import.meta.env.VITE_API_URL}/me`, {
			headers: { Authorization: `Bearer ${token}` }
		});
		const data = await response.json();
		if (response.ok) {
			currentUsername = data.user.username;
		}
	});

	function hostGame() {
		if (!selectedTopic) {
			error = 'Please select a topic';
			return;
		}
		if (ws) ws.close();
		ws = connectWebSocket(token, '/host', '', handleMessage);

		// Send topic with the connection if needed
		if (ws && ws.readyState === WebSocket.OPEN) {
			ws.send(
				JSON.stringify({
					Authorization: `Bearer ${token}`,
					topic: selectedTopic
				})
			);
		}
	}

	function joinGame() {
		if (!inputRoomId) {
			error = 'Please enter a room ID';
			return;
		}
		if (ws) ws.close();
		ws = connectWebSocket(token, '/join', inputRoomId, handleMessage);
	}

	function startGame() {
		if (currentUsername !== host) return;
		if (ws) {
			ws.send(JSON.stringify({ event: 'start_game', room_id: roomId }));
		}
		isGameStarted = true;
	}

	function handleMessage(data: any) {
		console.log('Received:', data);
		if (data.error) {
			error = data.error;
			return;
		}

		// Room setup messages
		if (data.room_id) {
			roomId = data.room_id;
			if (data.topic) topic = data.topic;
		}
		if (data.event === 'player_list') {
			players = data.players;
			host = data.host;
			topic = data.topic;
		}
		if (data.event === 'game_ready') {
			gameStatus = data.message;
		}

		// Game initialization
		if (data.event === 'game_initialized') {
			isGameStarted = true;
			isGameOver = false;
			currentAgeIndex = 0; // เริ่มที่ 0 - 12
			turnCount = 0; // รีเซ็ต turnCount
	

			// Initialize player stats
			if (data.players) {
				updatePlayerStats(data.players);
			}

			showNotification('Game started! Waiting for first player...');
		}

		// Turn management
		if (data.event === 'turn_start') {
			handleTurnStart(data);
		}
		if (data.event === 'waiting_for_turn') {
			handleWaitingForTurn(data);
		}
		if (data.event === 'turn_result') {
			handleTurnResult(data);

			turnCount += 1;
			console.log('Turn Count:', turnCount, 'Players:', players.length);
			// ถ้าครบรอบ (ทุกคนเล่นครบ 1 รอบ)
			if (turnCount >= players.length && currentAgeIndex < ageRanges.length - 1) {
				currentAgeIndex += 1; // เลื่อนไปช่วงอายุถัดไป
				turnCount = 0; // รีเซ็ตรอบ
				ageRanges[currentAgeIndex].data = Array(5).fill(null); // รีเซ็ต dropped cards
				currentAgeProgress = Math.floor((currentAgeIndex / (totalAges - 1)) * 100);
				showNotification(`Age advanced to ${ageRanges[currentAgeIndex].label}`);
			}
		}

		// Age progression
		if (data.event === 'age_advanced') {
			currentAgeIndex = data.age_index;
			currentAgeProgress = Math.floor((currentAgeIndex / (totalAges - 1)) * 100);

			if (data.players) {
				updatePlayerStats(data.players);
			}

			// Reset dropped cards for new age
			ageRanges = ageRanges.map((age, i) => ({
				...age,
				data: i === currentAgeIndex ? Array(5).fill(null) : age.data
			}));

			showNotification(`Age advanced to ${data.age_range.Label}`);
		}

		// Game results
		if (data.event === 'game_results') {
			showGameResults(data.results);
		}

		// Card deck response
		if (data.event === 'card_deck') {
			handCards = data.cards.map((card: { id: any; }) => ({
				...card,
				pic: `../src/lib/imggen/${card.id}.png` // Store the filename as string
			}));
			console.log('Updated handCards:', handCards);
		}
	}

	function showNotification(message: string) {
		notification = message;
		if (notificationTimer) clearTimeout(notificationTimer);
		notificationTimer = setTimeout(() => {
			notification = '';
		}, 5000);
	}

	function showEventResult(event: { description: any; choice_made: any; }, playerName: any) {
		// Show a modal or notification with the event result
		if (!event || !event.description) return;

		let message = `${playerName}: ${event.description}`;
		if (event.choice_made) {
			message += ` (Chose: ${event.choice_made})`;
		}

		showNotification(message);
	}

	function showGameResults(results: any[]) {
		// Show game results modal
		isGameOver = true;
		gameResults = results;
		showNotification('Game Over! The results are in!');
	}

	// Game functions
	function changeAge(direction: number) {
		if (!isMyTurn) return;
		currentAgeIndex = Math.max(0, Math.min(ageRanges.length - 1, currentAgeIndex + direction));
	}

	function getDeck() {
		console.log('Requesting cards...');

		// Request cards from server only if it's the player's turn
		if (isMyTurn && ws) {
			const playerIndex = players.indexOf(currentUsername);
			if (playerIndex === -1) return;

			ws.send(
				JSON.stringify({
					event: 'request_cards',
					room_id: roomId,
					player_index: playerIndex,
					data: {
						age_index: currentAgeIndex,
						count: 6 // Request 6 cards
					}
				})
			);
		} else {
			showNotification("It's not your turn yet. Please wait.");
		}
	}

	// Card selection functions
	function selectCard(card: any, index: number, source: string, ageIndex: number | null = null) {
		// Allow card selection only during player's turn
		if (!isMyTurn) {
			showNotification("It's not your turn. You can't select cards now.");
			return;
		}

		if (
			selectedCardIndex === index &&
			selectedCardSource === source &&
			selectedAgeIndex === ageIndex
		) {
			selectedCard = null;
			selectedCardIndex = null;
			selectedCardSource = null;
			selectedAgeIndex = null;
		} else {
			selectedCard = card;
			selectedCardIndex = index;
			selectedCardSource = source;
			selectedAgeIndex = ageIndex;
		}
	}

	function moveCardToSlot(slotIndex: number) {
		if (!isMyTurn) {
			showNotification("It's not your turn. You can't move cards now.");
			return;
		}

		if (!selectedCard || selectedCardSource === null) return;

		if (selectedCardSource === 'hand' && selectedCardIndex !== null) {
			if (!currentAge.data[slotIndex]) {
				handCards = handCards.filter((_, i) => i !== selectedCardIndex);
				currentAge.data[slotIndex] = selectedCard;
			}
		} else if (
			selectedCardSource === 'dropzone' &&
			selectedCardIndex !== null &&
			selectedAgeIndex !== null
		) {
			if (!currentAge.data[slotIndex]) {
				ageRanges[selectedAgeIndex].data[selectedCardIndex] = null;
				currentAge.data[slotIndex] = selectedCard;
			}
		}

		resetSelection();
	}

	function trashCard() {
		if (!isMyTurn) {
			showNotification("It's not your turn. You can't trash cards now.");
			return;
		}

		if (!selectedCard || selectedCardSource === null) return;

		if (selectedCardSource === 'hand' && selectedCardIndex !== null) {
			handCards = handCards.filter((_, i) => i !== selectedCardIndex);
		} else if (
			selectedCardSource === 'dropzone' &&
			selectedCardIndex !== null &&
			selectedAgeIndex !== null
		) {
			ageRanges[selectedAgeIndex].data[selectedCardIndex] = null;
		}

		resetSelection();
	}

	function resetSelection() {
		selectedCard = null;
		selectedCardIndex = null;
		selectedCardSource = null;
		selectedAgeIndex = null;

		handCards = handCards;
		ageRanges = ageRanges.map((age) => ({
			...age,
			data: [...age.data]
		}));
	}

	function handleSubmit() {
		if (!ws || !isMyTurn) return;

		// For choice events, handle differently than card placement
		if (currentEvent && currentEvent.Type === 'choice' && selectedCardIndex !== null) {
			// For choice events, send the selected choice ID
			const choiceId = currentEvent.Choices?.[selectedCardIndex]?.ID || 'default_choice';

			ws.send(
				JSON.stringify({
					event: 'make_choice',
					room_id: roomId,
					player_index: players.indexOf(currentUsername),
					event_id: currentEvent.ID,
					choice_id: choiceId
				})
			);

			resetSelection();
			return;
		}

		// For card placement, collect all placed cards
		const eventIDs = currentAge.data.filter((card) => card !== null).map((card) => card.id);

		if (eventIDs.length === 0) {
			// If we have a current event but no cards placed, send the event ID directly
			if (currentEvent) {
				ws.send(
					JSON.stringify({
						event: 'make_choice',
						room_id: roomId,
						player_index: players.indexOf(currentUsername),
						event_id: currentEvent.ID,
						choice_id: 'default_choice'
					})
				);
				return;
			}

			// Show error if no cards and no current event
			showNotification('You need to place at least one card before confirming your choice.');
			return;
		}

		// Send choices to backend with multiple event IDs
		ws.send(
			JSON.stringify({
				event: 'make_choice',
				room_id: roomId,
				player_index: players.indexOf(currentUsername),
				choice_id: 'selected_choice',
				data: {
					event_ids: eventIDs
				}
			})
		);

		// Clear placed cards after submission
		ageRanges[currentAgeIndex].data = Array(5).fill(null);
	}

	function handleTurnStart(data: { player_name: string; event_data: any; }) {
		isMyTurn = data.player_name === currentUsername;
		currentEvent = data.event_data;

		// If it's a choice event, prepare choices for display
		if (currentEvent && currentEvent.Type === 'choice' && currentEvent.Choices) {
			currentEventChoices = currentEvent.Choices;

			// If it's a choice event, add choices to hand cards
			if (isMyTurn) {
				handCards = currentEvent.Choices.map((choice: { ID: any; Description: any; Effects: any; }, index: any) => ({
					id: choice.ID,
					type: 'choice',
					description: choice.Description,
					money: getEffectValue(choice.Effects, 'money'),
					happiness: getEffectValue(choice.Effects, 'happiness'),
					knowledge: getEffectValue(choice.Effects, 'knowledge'),
					relationship: getEffectValue(choice.Effects, 'relationship'),
					pic: `../src/lib/imggen/choice_${index}.png` // Placeholder image
				}));
			}
		} else {
			currentEventChoices = [];

			// For regular events, show the event in the slot
			if (isMyTurn) {
				// Clear previous cards
				handCards = [];
				// Request cards for placement
				getDeck();
			}
		}

		if (isMyTurn) {
			// Show turn notification
			showNotification(`It's your turn! ${currentEvent?.Description || ''}`);
		}
	}

	function getEffectValue(effects: any[], stat: string) {
		if (!effects) return 0;
		const effect = effects.find((e: { Stat: any; }) => e.Stat === stat);
		return effect ? effect.Value : 0;
	}

	function handleWaitingForTurn(data: { player_name: any; }) {
		isMyTurn = false;
		showNotification(`Waiting for ${data.player_name}'s turn...`);
	}

	function handleTurnResult(data: { player_name: string; player_stats: any; life_event: any; }) {
		// Update displayed stats for the player
		const playerIndex = players.indexOf(data.player_name);
		if (playerIndex >= 0 && data.player_stats) {
			updatePlayerStat(playerIndex, data.player_stats);
		}

		// Show event result
		showEventResult(data.life_event, data.player_name);

		// Reset current event
		if (data.player_name === currentUsername) {
			currentEvent = null;
			currentEventChoices = [];
			// Clear placed cards
			ageRanges[currentAgeIndex].data = Array(5).fill(null);
		}
	}

	function updatePlayerStats(playersData: any[]) {
		// Initialize the stats array if it doesn't exist
		if (!allPlayerStats || allPlayerStats.length === 0) {
			allPlayerStats = Array(players.length)
				.fill(null)
				.map(() => ({
					money: 50,
					happiness: 50,
					knowledge: 50,
					relationship: 50
				}));
		}

		// Update player statistics from server data
		playersData.forEach((playerData: { money: any; happiness: any; knowledge: any; relationship: any; }, index: number) => {
			if (index < allPlayerStats.length) {
				allPlayerStats[index] = {
					money: playerData.money,
					happiness: playerData.happiness,
					knowledge: playerData.knowledge,
					relationship: playerData.relationship
				};
			}
		});
	}

	function updatePlayerStat(playerIndex: number, stats: { money: any; happiness: any; knowledge: any; relationship: any; }) {
		if (!allPlayerStats) {
			allPlayerStats = Array(players.length)
				.fill(null)
				.map(() => ({
					money: 50,
					happiness: 50,
					knowledge: 50,
					relationship: 50
				}));
		}

		if (playerIndex >= 0 && playerIndex < allPlayerStats.length) {
			allPlayerStats[playerIndex] = {
				money: stats.money,
				happiness: stats.happiness,
				knowledge: stats.knowledge,
				relationship: stats.relationship
			};
		}
	}

	function returnToLobby() {
		isGameStarted = false;
		isGameOver = false;

		// Reset game state
		handCards = [];
		currentAgeIndex = 0;
		turnCount = 0;
		resetAllDroppedCards();
	}

	ageRanges = [
		{ label: '0 - 12', data: droppedCards0_12 },
		{ label: '13 - 18', data: droppedCards13_18 },
		{ label: '19 - 22', data: droppedCards19_22 },
		{ label: '23 - 39', data: droppedCards23_39 },
		{ label: '40 - 59', data: droppedCards40_59 },
		{ label: '60 - 79', data: droppedCards60_79 },
		{ label: '80 - 100', data: droppedCards80_100 }
	];

	function resetAllDroppedCards() {
		droppedCards0_12 = Array(5).fill(null);
		droppedCards13_18 = Array(5).fill(null);
		droppedCards19_22 = Array(5).fill(null);
		droppedCards23_39 = Array(5).fill(null);
		droppedCards40_59 = Array(5).fill(null);
		droppedCards60_79 = Array(5).fill(null);
		droppedCards80_100 = Array(5).fill(null);

		ageRanges = [
			{ label: '0 - 12', data: droppedCards0_12 },
			{ label: '13 - 18', data: droppedCards13_18 },
			{ label: '19 - 22', data: droppedCards19_22 },
			{ label: '23 - 39', data: droppedCards23_39 },
			{ label: '40 - 59', data: droppedCards40_59 },
			{ label: '60 - 79', data: droppedCards60_79 },
			{ label: '80 - 100', data: droppedCards80_100 }
		];
	}

	function playAgain() {
		isGameOver = false;
		if (currentUsername === host) {
			// Host can start a new game
			startGame();
		} else {
			// Non-hosts just wait for host to start
			showNotification('Waiting for the host to start a new game.');
		}
	}

	onDestroy(() => {
		if (ws) ws.close();
		if (notificationTimer) clearTimeout(notificationTimer);
	});
</script>

{#if !isGameStarted}
	<!-- Lobby Interface -->
	<div
		class="relative flex min-h-screen items-center justify-center p-4"
		style={`background-image: url('${bg}'); background-size: cover; background-position: center; background-color: rgba(0, 0, 0, 0.5); background-blend-mode: darken;`}
	>
		<div class="w-full max-w-md rounded-md border px-8 py-4 shadow-xl">
			<h1 class="mb-2 text-center text-3xl font-bold text-white">Lobby</h1>

			{#if error}
				<p class="mb-4 text-center text-sm text-red-500">{error}</p>
			{/if}

			<div class="space-y-4">
				<div>
					<select
						bind:value={selectedTopic}
						class="w-full rounded-lg border border-gray-300 px-4 py-3 transition focus:border-transparent focus:outline-none focus:ring-2 focus:ring-blue-500"
					>
						<option value="" disabled selected>Select a topic</option>
						{#each topics as t}
							<option value={t}>{t}</option>
						{/each}
					</select>
					<button
						on:click={hostGame}
						class="mt-2 w-full rounded-lg bg-blue-600 py-3 font-medium text-white transition duration-300 hover:bg-blue-700"
					>
						Host Game
					</button>
				</div>

				<div>
					<input
						type="text"
						bind:value={inputRoomId}
						placeholder="Enter Room ID"
						class="w-full rounded-lg border border-gray-300 px-4 py-3 transition focus:border-transparent focus:outline-none focus:ring-2 focus:ring-blue-500"
					/>
					<button
						on:click={joinGame}
						class="mt-2 w-full rounded-lg bg-green-600 py-3 font-medium text-white transition duration-300 hover:bg-green-700"
					>
						Join Game
					</button>
				</div>
			</div>

			{#if roomId}
				<h2 class="mt-2 text-xl font-medium text-white">Room ID: {roomId}</h2>
			{/if}

			{#if topic}
				<p class="mt-2 text-white">Topic: {topic}</p>
			{/if}

			{#if players.length > 0}
				<h3 class="mt-1 text-lg font-medium text-white">Players:</h3>
				<ul class="list-disc pl-5 text-white">
					{#each players as player}
						<li class:text-blue-600={player === host}>
							{player}
							{player === host ? '(Host)' : ''}
						</li>
					{/each}
				</ul>
			{/if}

			{#if gameStatus}
				<p class="mt-2 text-center text-white">{gameStatus}</p>
			{/if}

			{#if gameStatus && currentUsername === host}
				<button
					on:click={startGame}
					class="mt-2 w-full rounded-lg bg-purple-600 py-3 font-semibold text-white transition duration-300 hover:bg-purple-700"
				>
					Start Game
				</button>
			{/if}
		</div>
	</div>
{:else if isGameOver}
	<!-- Game Results Interface -->
	<div
		class="flex h-full w-full flex-col items-center justify-center px-5 py-2"
		style="background-image: url('../src/lib/assets/image/play-bg.png'); background-size: cover; background-position: center; height: 100vh;"
	>
		<div class="w-full max-w-3xl rounded-xl bg-white bg-opacity-90 p-6 shadow-lg">
			<h1 class="mb-4 text-center text-3xl font-bold text-gray-800">Game Results</h1>

			<div class="mb-6">
				<h2 class="mb-2 text-center text-xl font-semibold">{topic || 'Game Results'}</h2>
				<p class="text-center text-gray-600">See who made the most of their life!</p>
			</div>

			<div class="space-y-4">
				{#each gameResults as result, index}
					<div
						class="flex items-center justify-between rounded-lg bg-gray-100 p-4 {index === 0
							? 'border-2 border-yellow-400 bg-yellow-100'
							: ''}"
					>
						<div class="flex items-center">
							<span class="mr-2 font-bold">{index + 1}.</span>
							<div class="ml-2">
								<p class="font-bold">{result.username} {index === 0 ? '👑' : ''}</p>
								<div class="mt-1 flex flex-wrap gap-2">
									<div class="flex items-center">
										<img src={iconMapColor['money']} alt="money icon" class="mr-1 h-4" />
										<span>{result.stats.money}</span>
									</div>
									<div class="flex items-center">
										<img src={iconMapColor['happiness']} alt="happiness icon" class="mr-1 h-4" />
										<span>{result.stats.happiness}</span>
									</div>
									<div class="flex items-center">
										<img src={iconMapColor['knowledge']} alt="knowledge icon" class="mr-1 h-4" />
										<span>{result.stats.knowledge}</span>
									</div>
									<div class="flex items-center">
										<img
											src={iconMapColor['relationship']}
											alt="relationship icon"
											class="mr-1 h-4"
										/>
										<span>{result.stats.relationship}</span>
									</div>
								</div>
							</div>
						</div>
						<div class="text-right">
							<p class="text-xl font-bold">{result.total_score}</p>
							<p class="text-sm text-gray-600">Total Score</p>
						</div>
					</div>
				{/each}
			</div>

			<div class="mt-6 flex justify-center gap-4">
				<button
					on:click={returnToLobby}
					class="rounded-lg bg-blue-600 px-6 py-3 font-semibold text-white transition duration-300 hover:bg-blue-700"
				>
					Return to Lobby
				</button>

				{#if currentUsername === host}
					<button
						on:click={playAgain}
						class="rounded-lg bg-green-600 px-6 py-3 font-semibold text-white transition duration-300 hover:bg-green-700"
					>
						Play Again
					</button>
				{/if}
			</div>
		</div>
	</div>
{:else}
	<!-- Game Interface -->
	<div
		class="flex h-full w-full flex-col items-center justify-center justify-between px-5 py-2"
		style="background-image: url('../src/lib/assets/image/play-bg.png'); background-size: cover; background-position: center; height: 100vh;"
	>
		<div class="flex w-full justify-between">
			<!-- Display player cards based on actual players -->
			{#if players.length > 0 && allPlayerStats && allPlayerStats.length > 0}
				<!-- First player card (can be host or first player) -->
				<CompetiterCard
					profileImage={`https://api.dicebear.com/7.x/bottts/svg?seed=${players[0]}`}
					playerName={players[0]}
					money={allPlayerStats[0]?.money || 50}
					happiness={allPlayerStats[0]?.happiness || 50}
					knowledge={allPlayerStats[0]?.knowledge || 50}
					relationship={allPlayerStats[0]?.relationship || 50}
					isCurrentPlayer={players[0] === currentUsername}
					isCurrentTurn={isMyTurn && players[0] === currentUsername}
				/>
			{:else}
				<!-- Placeholder if no player data yet -->
				<CompetiterCard
					profileImage="../src/lib/assets/image/profile.jpg"
					playerName="Player 1"
					money={50}
					happiness={50}
					knowledge={50}
					relationship={50}
				/>
			{/if}

			<button
				class="mt-10 hover:scale-110 disabled:opacity-50"
				on:click={() => changeAge(-1)}
				disabled={currentAgeIndex === 0 || !isMyTurn}
			>
				<img src="../src/lib/assets/icon/left.svg" alt="left icon" class="h-5" />
			</button>

			<div class="flex flex-col text-center text-white">
				<div class="text-sm">{topic || 'GET THE MOST MONEY'}</div>
				<TimeLeft />
				<div class="text-xs">AGE {currentAge.label}</div>

				<!-- Current event display -->
				{#if currentEvent}
					<div class="mb-2 mt-1 text-sm">
						{currentEvent.Description}
					</div>
				{/if}

				<!-- Turn indicator -->
				<div class="mb-2 text-sm font-bold">
					{#if isMyTurn}
						<span class="text-green-400">It's your turn!</span>
					{:else}
						<span class="text-yellow-300">Waiting for other players...</span>
					{/if}
				</div>

				<!-- Drop Zones for current age -->
				<div class="mt-2 flex gap-3">
					{#each currentAge.data as card, i}
						<div
							class="flex h-32 w-24 cursor-pointer items-center justify-center rounded-md border-2 border-dashed border-white {!isMyTurn
								? 'opacity-80'
								: ''}"
							on:click={() => {
								if (card) {
									selectCard(card, i, 'dropzone', currentAgeIndex);
								} else if (selectedCard) {
									moveCardToSlot(i);
								}
							}}
							on:keydown={(e) => {
								if (e.key === 'Enter' || e.key === ' ') {
									if (card) {
										selectCard(card, i, 'dropzone', currentAgeIndex);
									} else if (selectedCard) {
										moveCardToSlot(i);
									}
								}
							}}
							role="button"
							tabindex="0"
							class:bg-[#474848]={selectedCard && !card && isMyTurn}
						>
							{#if card}
								<div
									on:click|stopPropagation={() => selectCard(card, i, 'dropzone', currentAgeIndex)}
									on:keydown|stopPropagation={(e) => {
										if (e.key === 'Enter' || e.key === ' ') {
											selectCard(card, i, 'dropzone', currentAgeIndex);
										}
									}}
									role="button"
									tabindex="0"
									class:scale-110={selectedCardIndex === i &&
										selectedCardSource === 'dropzone' &&
										selectedAgeIndex === currentAgeIndex}
									class="transition-transform duration-200"
								>
									<PlayCard {...card} />
								</div>
							{/if}
						</div>
					{/each}
				</div>
			</div>

			<button
				class="mt-10 hover:scale-110 disabled:opacity-50"
				on:click={() => changeAge(1)}
				disabled={currentAgeIndex === ageRanges.length - 1 || !isMyTurn}
			>
				<img src="../src/lib/assets/icon/right.svg" alt="right icon" class="h-5" />
			</button>

			{#if players.length > 1 && allPlayerStats && allPlayerStats.length > 1}
				<!-- Second player card (if exists) -->
				<CompetiterCard
					profileImage={`https://api.dicebear.com/7.x/bottts/svg?seed=${players[1]}`}
					playerName={players[1]}
					money={allPlayerStats[1]?.money || 50}
					happiness={allPlayerStats[1]?.happiness || 50}
					knowledge={allPlayerStats[1]?.knowledge || 50}
					relationship={allPlayerStats[1]?.relationship || 50}
					isCurrentPlayer={players[1] === currentUsername}
					isCurrentTurn={isMyTurn && players[1] === currentUsername}
				/>
			{:else}
				<!-- Placeholder if no second player yet -->
				<CompetiterCard
					profileImage="../src/lib/assets/image/profile.jpg"
					playerName="Player 2"
					money={50}
					happiness={50}
					knowledge={50}
					relationship={50}
				/>
			{/if}
		</div>

		<div class="flex w-full justify-between text-white">
			<button
				on:click={() => getDeck()}
				class={!isMyTurn ? 'cursor-not-allowed opacity-60' : 'hover:scale-105'}
			>
				<img
					src="../src/lib/assets/image/play/random-deck-button.svg"
					alt="desk"
					class="h-[165px] w-[169px]"
				/>
			</button>

			<div class="flex flex-col gap-1">
				<div class="flex justify-between">
					<div class="flex flex-col">
						<div class="text-2xl">GAME PROGRESS</div>
						<div class="flex gap-2">
							<div class="flex items-center px-1">
								<img src={iconMapColor['money']} alt={`money icon`} class="mr-0.5 h-4" />
								<span class="text-sm font-medium">
									{allPlayerStats && players.indexOf(currentUsername) !== -1
										? allPlayerStats[players.indexOf(currentUsername)]?.money || 50
										: 50}
								</span>
							</div>
							<div class="flex items-center px-1">
								<img src={iconMapColor['happiness']} alt={`happiness icon`} class="mr-0.5 h-4" />
								<span class="text-sm font-medium">
									{allPlayerStats && players.indexOf(currentUsername) !== -1
										? allPlayerStats[players.indexOf(currentUsername)]?.happiness || 50
										: 50}
								</span>
							</div>
							<div class="flex items-center px-1">
								<img src={iconMapColor['knowledge']} alt={`knowledge icon`} class="mr-0.5 h-4" />
								<span class="text-sm font-medium">
									{allPlayerStats && players.indexOf(currentUsername) !== -1
										? allPlayerStats[players.indexOf(currentUsername)]?.knowledge || 50
										: 50}
								</span>
							</div>
							<div class="flex items-center px-1">
								<img
									src={iconMapColor['relationship']}
									alt={`relationship icon`}
									class="mr-0.5 h-4"
								/>
								<span class="text-sm font-medium">
									{allPlayerStats && players.indexOf(currentUsername) !== -1
										? allPlayerStats[players.indexOf(currentUsername)]?.relationship || 50
										: 50}
								</span>
							</div>
						</div>
					</div>

					<div class="flex flex-col gap-2">
						{#if isMyTurn}
							<button
								on:click={() => handleSubmit()}
								class="flex cursor-pointer items-center justify-center rounded-md bg-red-500 p-2 transition duration-200 hover:bg-red-600"
							>
								Confirm your choice
							</button>
						{:else}
							<button
								disabled
								class="flex cursor-not-allowed items-center justify-center rounded-md bg-gray-500 p-2 opacity-50"
							>
								Waiting for your turn...
							</button>
						{/if}

						<button
							on:click={returnToLobby}
							class="flex cursor-pointer items-center justify-center rounded-md bg-blue-500 p-2 transition duration-200 hover:bg-blue-600"
						>
							Return to Lobby
						</button>
					</div>

					<div class="flex flex-col items-center justify-center">
						<div class="flex items-center gap-2">
							<img
								src="../src/lib/assets/icon/PlayCard/heart.svg"
								alt="heart icon"
								class="mr-0.5 h-4"
							/>
							<div class="relative h-4 w-32 overflow-hidden rounded-full bg-[#D9D9D9]">
								<div
									class="absolute left-0 top-0 h-full bg-[#FF8787]"
									style="width: calc({(currentAgeIndex / 6) * 100}%)"
								></div>
							</div>
							<span class="text-sm font-medium">{currentAgeIndex + 1}</span>
						</div>

						<div class="flex items-center gap-2">
							<img
								src="../src/lib/assets/icon/PlayCard/hourglass.svg"
								alt="hourglass icon"
								class="mr-0.5 h-4 w-4"
							/>
							<div class="relative h-4 w-32 overflow-hidden rounded-full bg-[#D9D9D9]">
								<div
									class="absolute left-0 top-0 h-full bg-[#99624E]"
									style="width: calc({((totalAges - currentAgeIndex - 1) / totalAges) * 100}%)"
								></div>
							</div>
							<span class="text-sm font-medium">{totalAges - currentAgeIndex - 1}</span>
						</div>
					</div>
				</div>

				<div
					class="relative flex h-[114px] w-[530px] cursor-pointer items-center justify-center gap-2 rounded-md border border-white bg-[#474848] {!isMyTurn
						? 'opacity-80'
						: ''}"
					on:click={(e) => {
						if (!isMyTurn) return;

						const target = e.target as HTMLElement;
						if (selectedCard && !target.closest('.play-card-container')) {
							if (
								selectedCardSource === 'dropzone' &&
								selectedCardIndex !== null &&
								selectedAgeIndex !== null
							) {
								handCards = [...handCards, selectedCard];
								ageRanges[selectedAgeIndex].data[selectedCardIndex] = null;
							}
							resetSelection();
						}
					}}
					class:bg-[#5a5b5b]={selectedCard && selectedCardSource === 'dropzone' && isMyTurn}
				>
					{#each handCards as card, index}
						<div
							class="play-card-container cursor-pointer transition-transform duration-200"
							on:click|stopPropagation={() => selectCard(card, index, 'hand')}
							class:scale-110={selectedCardIndex === index && selectedCardSource === 'hand'}
						>
							<PlayCard {...card} />
						</div>
					{/each}

					{#if handCards.length === 0 && selectedCard && selectedCardSource === 'dropzone'}
						<div class="pointer-events-none absolute inset-0 flex items-center justify-center">
							<span class="text-sm text-white opacity-70">. . .</span>
						</div>
					{/if}

					{#if handCards.length === 0 && isMyTurn}
						<div class="text-sm text-white opacity-70">
							{currentEvent?.Type === 'choice'
								? 'Click to select a choice...'
								: 'Click the deck to get cards...'}
						</div>
					{/if}
				</div>
			</div>

			<div class="flex flex-col items-center justify-center gap-2">
				<div class="flex gap-2">
					<div class="relative w-[75px]">
						<img src="../src/lib/assets/image/play/job-bg.svg" alt="job" class="h-auto w-full" />
						<div
							class="absolute inset-0 flex items-center justify-center text-sm font-medium text-white"
						>
							job
						</div>
					</div>
					<div class="relative w-[75px]">
						<img
							src="../src/lib/assets/image/play/investment-bg.svg"
							alt="investment"
							class="h-auto w-full"
						/>
						<div
							class="absolute inset-0 flex items-center justify-center text-sm font-medium text-white"
						>
							investment
						</div>
					</div>
				</div>
				<div class="relative flex gap-2">
					<div class="relative w-[75px]">
						<img src="../src/lib/assets/image/play/love-bg.svg" alt="love" class="h-auto w-full" />
						<div
							class="absolute inset-0 flex items-center justify-center text-sm font-medium text-white"
						>
							love
						</div>
					</div>
					<div
						class="relative w-[75px] cursor-pointer transition-transform duration-200 {!isMyTurn
							? 'cursor-not-allowed opacity-60'
							: ''}"
						on:click={trashCard}
						class:scale-110={selectedCard !== null && isMyTurn}
					>
						<img src="../src/lib/assets/image/play/bin.svg" alt="bin" class="h-[73px] w-full" />
						{#if selectedCard}
							<div
								class="pointer-events-none absolute inset-0 flex items-center justify-center"
							></div>
						{/if}
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}

{#if notification}
	<div
		class="fixed left-1/2 top-4 z-50 -translate-x-1/2 transform rounded-md bg-black bg-opacity-70 px-4 py-2 text-white"
	>
		{notification}
	</div>
{/if}

<style>
	.space-y-4 > * + * {
		margin-top: 1rem;
	}
	.text-blue-600 {
		color: #2563eb;
		font-weight: bold;
	}
</style>
