<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { getToken, removeToken } from '$lib/auth';

    let username = '';
    let email = '';
    let error = '';

    onMount(async () => {
        const token = getToken();
        if (!token) {
            goto('/login');
            return;
        }

        try {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/me`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            const data = await response.json();
            
            if (!response.ok) {
                throw new Error(data.error || 'Failed to fetch user data');
            }

        
            username = data.user.username;
            email = data.user.email;      
        } catch (err) {
            error = err instanceof Error ? err.message : 'An unknown error occurred';
            removeToken();
            goto('/login');
        }
    });

    function handleLogout() {
        removeToken();
        goto('/login');
    }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-xl shadow-lg w-full max-w-md">
        {#if username}
            <h1 class="text-3xl font-bold text-gray-800 mb-4 text-center">
                Welcome, {username}
            </h1>
            <p class="text-gray-600 mb-6 text-center">Email: {email}</p>
            <button 
                on:click={handleLogout} 
                class="w-full bg-red-600 text-white py-3 rounded-lg hover:bg-red-700 transition duration-300 font-semibold"
            >
                Logout
            </button>
        {/if}
        {#if error}
            <p class="text-red-500 text-sm mt-4 text-center">{error}</p>
        {/if}
    </div>
</div>