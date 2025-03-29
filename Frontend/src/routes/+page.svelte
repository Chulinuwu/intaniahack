<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { getToken, removeToken } from '$lib/auth';
    import bg from '$lib/assets/image/bg.gif';

    let username = '';
    let email = '';
    let error = '';
    let isAuthenticated = false;

    onMount(async () => {
        const token = getToken();
        if (!token) {
            return; // Stay on landing page
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
            isAuthenticated = true;
        } catch (err) {
            error = err instanceof Error ? err.message : 'An unknown error occurred';
            removeToken();
        }
    });

    function handleLogout() {
        removeToken();
        goto('/');
    }

    function goToLogin() {
        goto('/login');
    }

    function goToSignup() {
        goto('/register');
    }
</script>

<style>
</style>

<div class="min-h-screen flex items-center justify-center relative" style={`background-image: url('${bg}'); background-size: cover; background-position: center; background-color: rgba(0, 0, 0, 0.5); background-blend-mode: darken;`}>
    <!-- Modal box -->
    <div class="relative w-full h-screen flex items-center justify-center">
        <div class=" px-10 py-8 rounded-md shadow-xl relative w-full max-w-md border">
          {#if isAuthenticated}
            <h1 class="text-3xl font-bold text-white mb-4 text-center">
                Welcome, {username}
            </h1>
            <p class="text-gray-300 mb-6 text-center">Email: {email}</p>
            <button 
                on:click={handleLogout} 
                class="w-full bg-red-600 text-white py-3 rounded-md hover:bg-red-700 transition duration-300 font-semibold"
            >
                Logout
            </button>
          {:else}
            <h1 class="text-3xl font-bold text-white mb-6 text-center">
                BANGKOKLIFE
            </h1>
            <div class="flex flex-col gap-4">
                <button 
                    on:click={goToLogin} 
                    class="w-full bg-gray-800 text-white py-3 px-4 rounded-md hover:bg-gray-700 transition duration-300 font-semibold border border-white shadow-md flex items-center justify-center"
                >
                    Login
                </button>
                <button 
                    on:click={goToSignup} 
                    class="w-full bg-gray-800 text-white py-3 px-4 rounded-md hover:bg-gray-700 transition duration-300 font-semibold border border-white shadow-md flex items-center justify-center"
                >
                    Sign Up
                </button>
            </div>
          {/if}

          {#if error}
              <p class="text-red-400 text-sm text-center mt-4">{error}</p>
          {/if}
        </div>
    </div>
</div>