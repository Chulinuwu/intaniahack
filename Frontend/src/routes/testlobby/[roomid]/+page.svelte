<script>
    import { page } from '$app/stores';

    let roomId = '';

    $: roomId = $page.params.roomid;
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-xl shadow-lg w-full max-w-md">
        <h1 class="text-3xl font-bold text-gray-800 mb-6 text-center">Game Room: {roomId}</h1>
        <p class="text-gray-600 text-center">Game has started! Implement your game logic here.</p>
    </div>
</div>