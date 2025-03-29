<script>
    import { onMount, onDestroy } from 'svelte';
    import { goto } from '$app/navigation';
    import { getToken, connectWebSocket } from '$lib/auth';

    let token = '';
    /**
	 * @type {WebSocket | null}
	 */
    let ws = null;
    let roomId = '';
    let inputRoomId = '';
    /**
	 * @type {string | any[]}
	 */
    let players = [];
    let host = '';
    let topic = ''; // เพิ่มตัวแปรสำหรับ topic
    let gameStatus = '';
    let error = '';
    let currentUsername = '';
    let selectedTopic = ''; // สำหรับ dropdown หรือ input

    // รายการ topic ที่ host สามารถเลือกได้
    const topics = [
        "General Knowledge",
        "Movies",
        "Music",
        "Science",
        "Custom"
    ];

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
        goto(`/testlobby/${roomId}`);
    }

    /**
	 * @param {{ error: string; room_id: string; topic: string; event: string; players: string | any[]; host: string; message: string; }} data
	 */
    function handleMessage(data) {
        console.log('Received:', data);
        if (data.error) {
            error = data.error;
            return;
        }
        if (data.room_id) {
            roomId = data.room_id;
            if (data.topic) topic = data.topic; // รับ topic จาก host
        }
        if (data.event === 'player_list') {
            players = data.players;
            host = data.host;
            topic = data.topic; // อัปเดต topic จาก broadcast
        }
        if (data.event === 'game_ready') {
            gameStatus = data.message;
        }
        if (data.event === 'game_not_ready') {
            gameStatus = data.message;
        }
        if (data.event === 'start_game') {
            goto(`/testlobby/${data.room_id}`);
        }
    }

    onDestroy(() => {
        if (ws) ws.close();
    });
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-xl shadow-lg w-full max-w-md">
        <h1 class="text-3xl font-bold text-gray-800 mb-6 text-center">Test Lobby</h1>

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

<style>
    .space-y-4 > * + * {
        margin-top: 1rem;
    }
    .text-blue-600 {
        color: #2563eb;
        font-weight: bold;
    }
</style>