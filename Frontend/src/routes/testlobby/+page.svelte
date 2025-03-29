<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
    import { goto } from '$app/navigation';
    import { getToken, connectWebSocket } from '$lib/auth';

    let token = '';
    let ws: WebSocket | null = null;
    let roomId = '';
    let inputRoomId = '';
    let players: string[] = [];
    let gameStatus = '';
    let error = '';

    onMount(() => {
        token = getToken() ?? '';
        if (!token) {
            goto('/login');
        }
        console.log(token);
    });

    function hostGame() {
        if (ws) ws.close(); // ปิดการเชื่อมต่อเก่าถ้ามี
        ws = connectWebSocket(token, '/host', '', handleMessage);
    }

    function joinGame() {
        if (!inputRoomId) {
            error = 'Please enter a room ID';
            return;
        }
        if (ws) ws.close(); // ปิดการเชื่อมต่อเก่าถ้ามี
        ws = connectWebSocket(token, '/join', inputRoomId, handleMessage);
    }

    function handleMessage(data: { error?: string; room_id?: string; event?: string; players?: string[]; message?: string }) {
        console.log('Received:', data);
        if (data.error) {
            error = data.error;
            return;
        }
        if (data.room_id) {
            roomId = data.room_id;
        }
        if (data.event === 'player_list') {
            players = data.players ?? [];
        }
        if (data.event === 'game_ready') {
            gameStatus = data.message ?? '';
        }
        if (data.event === 'game_not_ready') {
            gameStatus = data.message ?? '';
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
            <button 
                on:click={hostGame} 
                class="w-full bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition duration-300 font-semibold"
            >
                Host Game
            </button>

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

        {#if players.length > 0}
            <h3 class="mt-4 text-lg font-semibold text-gray-800">Players:</h3>
            <ul class="list-disc pl-5">
                {#each players as player}
                    <li>{player}</li>
                {/each}
            </ul>
        {/if}

        {#if gameStatus}
            <p class="mt-4 text-gray-600 text-center">{gameStatus}</p>
        {/if}
    </div>
</div>

<style>
    .space-y-4 > * + * {
        margin-top: 1rem;
    }
</style>