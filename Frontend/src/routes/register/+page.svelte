<script>
    import { goto } from '$app/navigation';

    let email = '';
    let username = '';
    let password = '';
    let error = '';

    async function handleRegister() {
        try {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ email, username, password })
            });

            const data = await response.json();
            
            if (!response.ok) {
                throw new Error(data.error || 'Registration failed');
            }

            goto('/login');
        } catch (err) {
            error = err instanceof Error ? err.message : 'An unknown error occurred';
        }
    }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-xl shadow-lg w-full max-w-md">
        <h1 class="text-3xl font-bold text-gray-800 mb-6 text-center">Register</h1>
        
        {#if error}
            <p class="text-red-500 text-sm mb-4 text-center">{error}</p>
        {/if}

        <form on:submit|preventDefault={handleRegister} class="space-y-6">
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
                    type="text" 
                    bind:value={username} 
                    placeholder="Username" 
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
                Register
            </button>
        </form>

        <p class="mt-4 text-center text-gray-600">
            Already have an account? 
            <a href="/login" class="text-blue-600 hover:underline">Login</a>
        </p>
    </div>
</div>