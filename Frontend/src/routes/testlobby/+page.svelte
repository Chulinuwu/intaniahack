<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
    import CompetiterCard from "../../components/CompetiterCard.svelte";
    import PlayCard from "../../components/PlayCard.svelte";
    import { iconMapColor } from '$lib/utils/iconMapColor';
    import TimeLeft from "../../components/TimeLeft.svelte";
    import { getToken, connectWebSocket } from '$lib/auth';
    import { goto } from '$app/navigation';

    // Lobby variables
    let token = '';
    let ws = null;
    let roomId = '';
    let inputRoomId = '';
    let players = [];
    let host = '';
    let topic = '';
    let gameStatus = '';
    let error = '';
    let currentUsername = '';
    let selectedTopic = '';
    
    // Game state flag
    let isGameStarted = false;

    // Game variables
    let handCards = [];
    let currentAgeIndex = 0;
    $: currentAge = ageRanges[currentAgeIndex];

    // Topics for host to choose
    const topics = [
        "GET THE MOST Money",
        "GET THE MOST Happiness",
        "GET THE MOST Knowledge",
        "GET THE MOST Relationship"
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
        { label: "0 - 12", data: droppedCards0_12 },
        { label: "13 - 18", data: droppedCards13_18 },
        { label: "19 - 22", data: droppedCards19_22 },
        { label: "23 - 39", data: droppedCards23_39 },
        { label: "40 - 59", data: droppedCards40_59 },
        { label: "60 - 79", data: droppedCards60_79 },
        { label: "80 - 100", data: droppedCards80_100 }
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
            headers: { 'Authorization': `Bearer ${token}` }
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
            ws.send(JSON.stringify({ event: "start_game", room_id: roomId }));
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
        if (data.event === 'game_not_ready') {
            gameStatus = data.message;
        }
        if (data.event === 'start_game') {
            // Start the game interface instead of redirecting
            isGameStarted = true;
            // Initialize game by getting initial cards
            getDeck();
        }
    }

    // Game functions
    function changeAge(direction) {
        currentAgeIndex = Math.max(0, Math.min(ageRanges.length - 1, currentAgeIndex + direction));
    }

    function getDeck() {
        console.log("Deck clicked!");

        // Simulate fetch from events-data.json
        // In production, replace with actual fetch
        const sampleCards = [
            {
                type: "money",
                pic: '../src/lib/assets/image/play/money.png',
                description: "Saved money from part-time job",
                money: 5,
                happiness: 2,
                knowledge: 0,
                relationship: 0
            },
            {
                type: "happiness",
                pic: '../src/lib/assets/image/play/money.png',
                description: "Went to a fun summer camp",
                money: -2,
                happiness: 5,
                knowledge: 3,
                relationship: 4
            },
            {
                type: "knowledge",
                pic: '../src/lib/assets/image/play/money.png',
                description: "Studied hard for exams",
                money: 0,
                happiness: -1,
                knowledge: 6,
                relationship: 0
            },
            {
                type: "relationship",
                pic: '../src/lib/assets/image/play/money.png',
                description: "Made new friends at school",
                money: 0,
                happiness: 4,
                knowledge: 2,
                relationship: 5
            },
            {
                type: "money",
                pic: '../src/lib/assets/image/play/money.png',
                description: "Birthday money from relatives",
                money: 3,
                happiness: 3,
                knowledge: 0,
                relationship: 1
            },
            {
                type: "knowledge",
                pic: '../src/lib/assets/image/play/money.png',
                description: "Read many books from library",
                money: 0,
                happiness: 2,
                knowledge: 5,
                relationship: 0
            }
        ];

        handCards = sampleCards;
    }

    // Card selection functions
    function selectCard(card, index, source, ageIndex = null) {
        if (selectedCardIndex === index && selectedCardSource === source && selectedAgeIndex === ageIndex) {
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
        if (!selectedCard || selectedCardSource === null) return;

        if (selectedCardSource === 'hand' && selectedCardIndex !== null) {
            if (!currentAge.data[slotIndex]) {
                handCards = handCards.filter((_, i) => i !== selectedCardIndex);
                currentAge.data[slotIndex] = selectedCard;
            }
        }
        else if (selectedCardSource === 'dropzone' && selectedCardIndex !== null && selectedAgeIndex !== null) {
            if (!currentAge.data[slotIndex]) {
                ageRanges[selectedAgeIndex].data[selectedCardIndex] = null;
                currentAge.data[slotIndex] = selectedCard;
            }
        }

        resetSelection();
    }

    function trashCard() {
        if (!selectedCard || selectedCardSource === null) return;

        if (selectedCardSource === 'hand' && selectedCardIndex !== null) {
            handCards = handCards.filter((_, i) => i !== selectedCardIndex);
        }
        else if (selectedCardSource === 'dropzone' && selectedCardIndex !== null && selectedAgeIndex !== null) {
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
        ageRanges = ageRanges.map(age => ({
            ...age,
            data: [...age.data]
        }));
    }

    function handleSubmit() {
        console.log("Submitting final choices:", ageRanges);
        // Here you would send the final choices to the server
    }

    function returnToLobby() {
        isGameStarted = false;
    }

    onDestroy(() => {
        if (ws) ws.close();
    });
</script>

{#if !isGameStarted}
<!-- Lobby Interface -->
<div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-xl shadow-lg w-full max-w-md">
        <h1 class="text-3xl font-bold text-gray-800 mb-6 text-center">Lobby</h1>

        {#if error}
            <p class="text-red-500 text-sm mb-4 text-center">{error}</p>
        {/if}

        <div class="space-y-4">
            <div>
                <select 
                    bind:value={selectedTopic} 
                    class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition"
                >
                    <option value="" disabled selected>Select a topic</option>
                    {#each topics as t}
                        <option value={t}>{t}</option>
                    {/each}
                </select>
                <button 
                    on:click={hostGame} 
                    class="w-full mt-2 bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition duration-300 font-semibold"
                >
                    Host Game
                </button>
            </div>

            <div>
                <input 
                    type="text" 
                    bind:value={inputRoomId} 
                    placeholder="Enter Room ID" 
                    class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition"
                />
                <button 
                    on:click={joinGame} 
                    class="w-full mt-2 bg-green-600 text-white py-3 rounded-lg hover:bg-green-700 transition duration-300 font-semibold"
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
                    <li class:text-blue-600={player === host}>{player} {player === host ? '(Host)' : ''}</li>
                {/each}
            </ul>
        {/if}

        {#if gameStatus}
            <p class="mt-4 text-gray-600 text-center">{gameStatus}</p>
        {/if}

        {#if gameStatus && currentUsername === host}
            <button 
                on:click={startGame} 
                class="w-full mt-4 bg-purple-600 text-white py-3 rounded-lg hover:bg-purple-700 transition duration-300 font-semibold"
            >
                Start Game
            </button>
        {/if}
    </div>
</div>
{:else}
<!-- Game Interface -->
<div
  class="flex flex-col w-full h-full justify-between py-2 px-5 items-center justify-center"
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
                { type: "money", description: "Learning and growing." },
                { type: "happiness", description: "Exploring and discovering." },
                { type: "relationship", description: "Building a career and relationships." },
                { type: "knowledge", description: "Learning and growing." }, 
            ]}
            age13_18={[
                { type: "money", description: "Learning and growing." },
                { type: "knowledge", description: "Exploring and discovering." },
                { type: "relationship", description: "Building a career and relationships." }
            ]}
        />
        <button 
            class="mt-10 hover:scale-110 disabled:opacity-50" 
            on:click={() => changeAge(-1)} 
            disabled={currentAgeIndex === 0}
        >
            <img src='../src/lib/assets/icon/left.svg' alt="left icon" class="h-5" />
        </button>
        <div class="flex flex-col text-white text-center">
            <div class="text-sm">{topic || "GET THE MOST MONEY"}</div>
            <TimeLeft />
            <div class="text-xs">AGE {currentAge.label}</div>
            <!-- Drop Zones for current age -->
            <div class="flex mt-2 gap-3">
                {#each currentAge.data as card, i}
                  <div 
                      class="w-24 h-32 border-2 border-dashed border-white rounded-md flex items-center justify-center cursor-pointer"
                      on:click={() => {
                          if (card) {
                              selectCard(card, i, 'dropzone', currentAgeIndex);
                          } else if (selectedCard) {
                              moveCardToSlot(i);
                          }
                      }}
                      class:bg-[#474848]={selectedCard && !card}
                  >
                      {#if card}
                          <div
                              on:click|stopPropagation={() => selectCard(card, i, 'dropzone', currentAgeIndex)}
                              class:scale-110={selectedCardIndex === i && selectedCardSource === 'dropzone' && selectedAgeIndex === currentAgeIndex}
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
            disabled={currentAgeIndex === ageRanges.length - 1}
        >
          <img src='../src/lib/assets/icon/right.svg' alt="right icon" class="h-5" />
        </button>
        <CompetiterCard
            profileImage="../src/lib/assets/image/profile.jpg"
            playerName="John Doe"
            money={1000}
            happiness={80}
            knowledge={90}
            relationship={70}
            age0_12={[
                { type: "money", description: "Learning and growing." },
                { type: "happiness", description: "Exploring and discovering." },
                { type: "relationship", description: "Building a career and relationships." },
                { type: "knowledge", description: "Learning and growing." },
            ]}
            age13_18={[
                { type: "money", description: "Learning and growing." },
                { type: "knowledge", description: "Exploring and discovering." },
                { type: "relationship", description: "Building a career and relationships." }
            ]}
        />
    </div>
    <div class="flex w-full text-white justify-between">
        <button on:click={() => getDeck()}>
            <img src="../src/lib/assets/image/play/random-deck-button.svg" alt="desk" class="w-[169px] h-[165px]" />
        </button>
        <div class="flex flex-col gap-1">
            <div class="flex justify-between">
                <div class="flex flex-col">
                    <div class="text-2xl">
                        REMAINING COST
                    </div>
                    <div class="flex gap-2">
                        <div class="flex items-center px-1">
                            <img src={iconMapColor['money']} alt={`money icon`} class="h-4 mr-0.5" />
                            <span class="font-medium text-sm">10</span>
                        </div>
                        <div class="flex items-center px-1">
                            <img src={iconMapColor['happiness']} alt={`happiness icon`} class="h-4 mr-0.5" />
                            <span class="font-medium text-sm">10</span>
                        </div>
                        <div class="flex items-center px-1">
                            <img src={iconMapColor['knowledge']} alt={`knowledge icon`} class="h-4 mr-0.5" />
                            <span class="font-medium text-sm">10</span>
                        </div>
                        <div class="flex items-center px-1">
                            <img src={iconMapColor['relationship']} alt={`relationship icon`} class="h-4 mr-0.5" />
                            <span class="font-medium text-sm">10</span>
                        </div>
                    </div>
                </div>
                <div class="flex flex-col gap-2">
                    <button
                      on:click={() => handleSubmit()}
                      class="flex p-2 bg-red-500 rounded-md items-center justify-center cursor-pointer hover:bg-red-600 transition duration-200"
                    >
                        Confirm your choice
                    </button>
                    <button
                      on:click={returnToLobby}
                      class="flex p-2 bg-blue-500 rounded-md items-center justify-center cursor-pointer hover:bg-blue-600 transition duration-200"
                    >
                        Return to Lobby
                    </button>
                </div>
                <div class="flex flex-col items-center justify-center">
                    <div class="flex items-center gap-2">
                        <img src='../src/lib/assets/icon/PlayCard/heart.svg' alt="heart icon" class="h-4 mr-0.5" />
                        <div class="relative w-32 h-4 bg-[#D9D9D9] rounded-full overflow-hidden">
                            <div
                                class="absolute top-0 left-0 h-full bg-[#FF8787]"
                                style="width: calc((1 / 8) * 100%)"
                            ></div>
                        </div>
                        <span class="font-medium text-sm">1</span>
                    </div>
                    <div class="flex items-center gap-2">
                        <img src='../src/lib/assets/icon/PlayCard/hourglass.svg' alt="hourglass icon" class="h-4 w-4 mr-0.5" />
                        <div class="relative w-32 h-4 bg-[#D9D9D9] rounded-full overflow-hidden">
                            <div
                                class="absolute top-0 left-0 h-full bg-[#99624E]"
                                style="width: calc((7 / 8) * 100%)"
                            ></div>
                        </div>
                        <span class="font-medium text-sm">7</span>
                    </div>
                </div>
            </div>
            <div 
              class="flex gap-2 w-[530px] h-[114px] bg-[#474848] border border-white rounded-md items-center justify-center cursor-pointer relative"
              on:click={(e) => {
                  const target = e.target as HTMLElement;
                  if (selectedCard && !target.closest('.play-card-container')) {
                      if (selectedCardSource === 'dropzone' && selectedCardIndex !== null && selectedAgeIndex !== null) {
                          handCards = [...handCards, selectedCard];
                          ageRanges[selectedAgeIndex].data[selectedCardIndex] = null;
                      }
                      resetSelection();
                  }
              }}
              class:bg-[#5a5b5b]={selectedCard && selectedCardSource === 'dropzone'}
          >
              {#each handCards as card, index}
                  <div 
                      class="play-card-container transition-transform duration-200 cursor-pointer" 
                      on:click|stopPropagation={() => selectCard(card, index, 'hand')}
                      class:scale-110={selectedCardIndex === index && selectedCardSource === 'hand'}
                  >
                      <PlayCard {...card} />
                  </div>
              {/each}
              
              {#if handCards.length === 0 && selectedCard && selectedCardSource === 'dropzone'}
                  <div class="absolute inset-0 flex items-center justify-center pointer-events-none">
                      <span class="text-white text-sm opacity-70">. . .</span>
                  </div>
              {/if}
          </div>
        </div>
        <div class="flex flex-col gap-2 items-center justify-center">
            <div class="flex gap-2">
                <div class="w-[75px] relative">
                    <img src="../src/lib/assets/image/play/job-bg.svg" alt="job" class="w-full h-auto" />
                    <div class="absolute inset-0 flex items-center justify-center text-white font-medium text-sm">
                        job
                    </div>
                </div>
                <div class="w-[75px] relative">
                    <img src="../src/lib/assets/image/play/investment-bg.svg" alt="investment" class="w-full h-auto" />
                    <div class="absolute inset-0 flex items-center justify-center text-white font-medium text-sm">
                        investment
                    </div>
                </div>
            </div>
            <div class="flex gap-2 relative">
                <div class="w-[75px] relative">
                    <img src="../src/lib/assets/image/play/love-bg.svg" alt="love" class="w-full h-auto" />
                    <div class="absolute inset-0 flex items-center justify-center text-white font-medium text-sm">
                        love
                    </div>
                </div>
                <div 
                      class="w-[75px] relative cursor-pointer transition-transform duration-200"
                      on:click={trashCard}
                      class:scale-110={selectedCard !== null}
                  >
                      <img 
                          src="../src/lib/assets/image/play/bin.svg" 
                          alt="bin" 
                          class="w-full h-[73px]"
                      />
                      {#if selectedCard}
                          <div class="absolute inset-0 flex items-center justify-center pointer-events-none">
                          </div>
                      {/if}
                  </div>
            </div>
        </div>
    </div>
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