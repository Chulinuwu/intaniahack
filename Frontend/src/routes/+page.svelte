<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { getToken, removeToken } from '$lib/auth';

    let username = '';
    let error = '';

    onMount(async () => {
        try {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/me`, {
                headers: {
                    'Authorization': `Bearer ${getToken()}`
                }
            });

            const data = await response.json();
            
            if (!response.ok) {
                throw new Error('Failed to fetch user data');
            }

            username = data.email;
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

<div class="container">
    {#if username}
        <h1>Welcome, {username}</h1>
        <button on:click={handleLogout}>Logout</button>
    {/if}
    {#if error}
        <p class="error">{error}</p>
    {/if}
</div>

<style>
    .container {
        max-width: 400px;
        margin: 0 auto;
        padding: 2rem;
    }
    .error {
        color: red;
    }
</style>