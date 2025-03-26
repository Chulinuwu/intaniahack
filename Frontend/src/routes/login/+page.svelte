<script>
    import { goto } from '$app/navigation';
    import { setToken } from '$lib/auth';

    let email = '';
    let password = '';
    let error = '';

    async function handleLogin() {
        try {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ email, password })
            });

            const data = await response.json();
            
            if (!response.ok) {
                throw new Error(data.error || 'Login failed');
            }

            setToken(data.token);
            goto('/');
        } catch (err) {
            error = err instanceof Error ? err.message : 'An unknown error occurred';
        }
    }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-xl shadow-lg w-full max-w-md">
        <h1 class="text-3xl font-bold text-gray-800 mb-6 text-center">Login</h1>
        
        {#if error}
            <p class="text-red-500 text-sm mb-4 text-center">{error}</p>
        {/if}

        <form on:submit|preventDefault={handleLogin} class="space-y-6">
            <div>
                <input 
                    type="email" 
                    bind:value={email} 
                    placeholder="Email" 
                    required 
                    class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition"
                />
            </div>
            <div>
                <input 
                    type="password" 
                    bind:value={password} 
                    placeholder="Password" 
                    required 
                    class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition"
                />
            </div>
            <button 
                type="submit" 
                class="w-full bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition duration-300 font-semibold"
            >
                Login
            </button>
        </form>

        <p class="mt-4 text-center text-gray-600">
            Don't have an account? 
            <a href="/register" class="text-blue-600 hover:underline">Register</a>
        </p>
    </div>
</div>