<!-- 1. เพิ่มตัวแปรสำหรับแสดงผลลัพธ์เกม -->
<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import CompetiterCard from '../../components/CompetiterCard.svelte';
	import PlayCard from '../../components/PlayCard.svelte';
	import { iconMapColor } from '$lib/utils/iconMapColor';
	import TimeLeft from '../../components/TimeLeft.svelte';
	import { getToken, connectWebSocket } from '$lib/auth';
	import { goto } from '$app/navigation';

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
	
	// Game over flag and results
	let isGameOver = false;
	let gameResults = [];

	// Game variables
	let handCards = [];
	let currentAgeIndex = 0;
	$: currentAge = ageRanges[currentAgeIndex];

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
	let selectedCard = null;
	let selectedCardIndex = null;
	let selectedCardSource = null;
	let selectedAgeIndex = null;

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
		// Instead of navigating to a new page, we just switch the view
		isGameStarted = true;
		// Initialize game by getting initial cards
		getDeck();
	}

	function handleMessage(data) {
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
			isGameOver = false; // Reset game over state
			currentAgeIndex = data.current_age;
			ageRanges = data.age_ranges.map((range, i) => ({
				label: range,
				data: ageRanges[i].data // Keep existing data structure
			}));

			// Update player data
			updatePlayerData(data.players);
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
		}

		// Age progression
		if (data.event === 'age_advanced') {
			currentAgeIndex = data.age_index;
			updatePlayerData(data.players);
		}

		// Game results
		if (data.event === 'game_results') {
			showGameResults(data.results);
		}

		// Card deck response
		if (data.event === 'card_deck') {
			handCards = data.cards;
		}
	}
	// Add these to your script section
	let notification = '';
	let notificationTimer = null;
	let playerStats = [];

	function showNotification(message) {
		notification = message;
		if (notificationTimer) clearTimeout(notificationTimer);
		notificationTimer = setTimeout(() => {
			notification = '';
		}, 3000);
	}

	function showEventResult(event, playerName) {
		// Show a modal or notification with the event result
		showNotification(
			`${playerName}: ${event.description} ${event.choice_made ? `(Chose: ${event.choice_made})` : ''}`
		);
	}

	function showGameResults(results) {
		// Show game results modal
		isGameOver = true;
		gameResults = results;
		showNotification('Game Over! The results are in!');
	}

	// Game functions
	function changeAge(direction) {
		currentAgeIndex = Math.max(0, Math.min(ageRanges.length - 1, currentAgeIndex + direction));
	}

	function getDeck() {
		console.log('Deck clicked!');

		// Request cards from server only if it's the player's turn
		if (isMyTurn) {
			ws.send(
				JSON.stringify({
					event: 'request_cards',
					room_id: roomId,
					player_index: players.findIndex((p) => p === currentUsername),
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
	function selectCard(card, index, source, ageIndex = null) {
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

	function moveCardToSlot(slotIndex) {
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

		// Collect all placed cards for the current age
		const eventIDs = currentAge.data.filter((card) => card !== null).map((card) => card.id);

		if (eventIDs.length === 0) {
			// Show error or warning if no cards placed
			showNotification("You need to place at least one card before confirming your choice.");
			return;
		}

		// Send choices to backend
		ws.send(
			JSON.stringify({
				event: 'make_choice',
				room_id: roomId,
				player_index: players.findIndex((p) => p === currentUsername),
				choice_id: 'selected_choice', // This would be for choice events
				data: {
					event_ids: eventIDs
				}
			})
		);
	}
	// Variables to track turn state
	let isMyTurn = false;
	let currentEvent = null;

	function handleTurnStart(data) {
		isMyTurn = data.player_name === currentUsername;
		currentEvent = data.event_data;

		if (isMyTurn) {
			// Show turn notification
			showNotification(`It's your turn!`);

			// Request cards if we don't have any
			if (handCards.length === 0) {
				getDeck();
			}
		}
	}

	function handleWaitingForTurn(data) {
		isMyTurn = false;
		showNotification(`Waiting for ${data.player_name}'s turn...`);
	}

	function handleTurnResult(data) {
		// Update player stats
		const playerIndex = players.findIndex((p) => p === data.player_name);
		if (playerIndex >= 0) {
			playerStats[playerIndex] = data.player_stats;
		}

		// Show event result
		showEventResult(data.life_event, data.player_name);
	}

	function updatePlayerData(playersData) {
		// Update player statistics from server data
		playerStats = playersData.map((p) => ({
			money: p.money,
			happiness: p.happiness,
			knowledge: p.knowledge,
			relationship: p.relationship
		}));
	}

	function returnToLobby() {
		isGameStarted = false;
		isGameOver = false;
	}
	
	function playAgain() {
		isGameOver = false;
		if (currentUsername === host) {
			// Host can start a new game
			startGame();
		} else {
			// Non-hosts just wait for host to start
			showNotification("Waiting for the host to start a new game.");
		}
	}

	onDestroy(() => {
		if (ws) ws.close();
	});
</script>

{#if !isGameStarted}
	<!-- Lobby Interface -->
	<div class="flex min-h-screen items-center justify-center bg-gray-100">
		<div class="w-full max-w-md rounded-xl bg-white p-8 shadow-lg">
			<h1 class="mb-6 text-center text-3xl font-bold text-gray-800">Lobby</h1>

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
						class="mt-2 w-full rounded-lg bg-blue-600 py-3 font-semibold text-white transition duration-300 hover:bg-blue-700"
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
						class="mt-2 w-full rounded-lg bg-green-600 py-3 font-semibold text-white transition duration-300 hover:bg-green-700"
					>
						Join Game
					</button>
				</div>
			</div>

			{#if roomId}
				<h2 class="mt-6 text-xl font-semibold text-gray-800">Room ID: {roomId}</h2>
			{/if}

			{#if topic}
				<p class="mt-2 text-gray-600">Topic: {topic}</p>
			{/if}

			{#if players.length > 0}
				<h3 class="mt-4 text-lg font-semibold text-gray-800">Players:</h3>
				<ul class="list-disc pl-5">
					{#each players as player}
						<li class:text-blue-600={player === host}>
							{player}
							{player === host ? '(Host)' : ''}
						</li>
					{/each}
				</ul>
			{/if}

			{#if gameStatus}
				<p class="mt-4 text-center text-gray-600">{gameStatus}</p>
			{/if}

			{#if gameStatus && currentUsername === host}
				<button
					on:click={startGame}
					class="mt-4 w-full rounded-lg bg-purple-600 py-3 font-semibold text-white transition duration-300 hover:bg-purple-700"
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
				<h2 class="mb-2 text-xl font-semibold text-center">{topic || 'Game Results'}</h2>
				<p class="text-center text-gray-600">See who made the most of their life!</p>
			</div>
			
			<div class="space-y-4">
				{#each gameResults as result, index}
					<div class="flex items-center justify-between rounded-lg bg-gray-100 p-4 {index === 0 ? 'bg-yellow-100 border-2 border-yellow-400' : ''}">
						<div class="flex items-center">
							<span class="mr-2 font-bold">{index + 1}.</span>
							<div class="ml-2">
								<p class="font-bold">{result.username} {index === 0 ? '👑' : ''}</p>
								<div class="flex flex-wrap gap-2 mt-1">
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
										<img src={iconMapColor['relationship']} alt="relationship icon" class="mr-1 h-4" />
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
			<CompetiterCard
				profileImage="../src/lib/assets/image/profile.jpg"
				playerName="John Doe"
				money={1000}
				happiness={80}
				knowledge={90}
				relationship={70}
				age0_12={[
					{ type: 'money', description: 'Learning and growing.' },
					{ type: 'happiness', description: 'Exploring and discovering.' },
					{ type: 'relationship', description: 'Building a career and relationships.' },
					{ type: 'knowledge', description: 'Learning and growing.' }
				]}
				age13_18={[
					{ type: 'money', description: 'Learning and growing.' },
					{ type: 'knowledge', description: 'Exploring and discovering.' },
					{ type: 'relationship', description: 'Building a career and relationships.' }
				]}
			/>
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
				<div class="mb-2 mt-1 text-sm font-bold">
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
							class="flex h-32 w-24 cursor-pointer items-center justify-center rounded-md border-2 border-dashed border-white {!isMyTurn ? 'opacity-80' : ''}"
							on:click={() => {
								if (card) {
									selectCard(card, i, 'dropzone', currentAgeIndex);
								} else if (selectedCard) {
									moveCardToSlot(i);
								}
							}}
							class:bg-[#474848]={selectedCard && !card && isMyTurn}
						>
							{#if card}
								<div
									on:click|stopPropagation={() => selectCard(card, i, 'dropzone', currentAgeIndex)}
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
			<CompetiterCard
				profileImage="../src/lib/assets/image/profile.jpg"
				playerName="John Doe"
				money={1000}
				happiness={80}
				knowledge={90}
				relationship={70}
				age0_12={[
					{ type: 'money', description: 'Learning and growing.' },
					{ type: 'happiness', description: 'Exploring and discovering.' },
					{ type: 'relationship', description: 'Building a career and relationships.' },
					{ type: 'knowledge', description: 'Learning and growing.' }
				]}
				age13_18={[
					{ type: 'money', description: 'Learning and growing.' },
					{ type: 'knowledge', description: 'Exploring and discovering.' },
					{ type: 'relationship', description: 'Building a career and relationships.' }
				]}
			/>
		</div>
		<div class="flex w-full justify-between text-white">
			<button on:click={() => getDeck()} class="{!isMyTurn ? 'opacity-60 cursor-not-allowed' : 'hover:scale-105'}">
				<img
					src="../src/lib/assets/image/play/random-deck-button.svg"
					alt="desk"
					class="h-[165px] w-[169px]"
				/>
			</button>
			<div class="flex flex-col gap-1">
				<div class="flex justify-between">
					<div class="flex flex-col">
						<div class="text-2xl">REMAINING COST</div>
						<div class="flex gap-2">
							<div class="flex items-center px-1">
								<img src={iconMapColor['money']} alt={`money icon`} class="mr-0.5 h-4" />
								<span class="text-sm font-medium">10</span>
							</div>
							<div class="flex items-center px-1">
								<img src={iconMapColor['happiness']} alt={`happiness icon`} class="mr-0.5 h-4" />
								<span class="text-sm font-medium">10</span>
							</div>
							<div class="flex items-center px-1">
								<img src={iconMapColor['knowledge']} alt={`knowledge icon`} class="mr-0.5 h-4" />
								<span class="text-sm font-medium">10</span>
							</div>
							<div class="flex items-center px-1">
								<img
									src={iconMapColor['relationship']}
									alt={`relationship icon`}
									class="mr-0.5 h-4"
								/>
								<span class="text-sm font-medium">10</span>
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
									style="width: calc((1 / 8) * 100%)"
								></div>
							</div>
							<span class="text-sm font-medium">1</span>
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
									style="width: calc((7 / 8) * 100%)"
								></div>
							</div>
							<span class="text-sm font-medium">7</span>
						</div>
					</div>
				</div>
				<div
					class="relative flex h-[114px] w-[530px] cursor-pointer items-center justify-center gap-2 rounded-md border border-white bg-[#474848] {!isMyTurn ? 'opacity-80' : ''}"
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
						class="relative w-[75px] cursor-pointer transition-transform duration-200 {!isMyTurn ? 'opacity-60 cursor-not-allowed' : ''}"
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