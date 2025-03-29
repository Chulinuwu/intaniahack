<script>
    import { goto } from '$app/navigation';
    import { setToken } from '$lib/auth';
    import bg from '$lib/assets/image/bg.gif';
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

<style>
</style>

<div class="min-h-screen flex items-center justify-center relative" style={`background-image: url('${bg}'); background-size: cover; background-position: center; background-color: rgba(0, 0, 0, 0.5); background-blend-mode: darken;`}>
    <!-- Modal box -->
    <div class="relative w-full h-screen flex items-center justify-center">
        <div class="absolute h-[20rem] w-[26rem] bg-gray-800 rounded-md shadow-xl max-w-md border border-white z-10"></div>
        <div class="bg-gray-800 text-white px-10 py-8 rounded-md shadow-xl relative w-full max-w-md border border-white z-20">
          
          <!-- Close button -->
          <button 
            class="absolute top-4 right-4 text-white text-2xl font-bold hover:text-red-400 transition"
            on:click={() => goto('/')}
          >
            ×
          </button>
      
          <form on:submit|preventDefault={handleLogin} class="space-y-6">
            <div>
              <label for="email" class="block text-lg font-semibold tracking-wide mb-1">Email</label>
              <input 
                id="email"
                type="email" 
                bind:value={email} 
                required 
                class="w-full px-3 py-2 rounded-md bg-gray-200 text-black focus:outline-none focus:ring-2 focus:ring-blue-500 transition"
              />
            </div>
      
            <div>
              <label for="password" class="block text-lg font-semibold tracking-wide mb-1">Password</label>
              <input 
                id="password"
                type="password" 
                bind:value={password} 
                required 
                class="w-full px-3 py-2 rounded-md bg-gray-200 text-black focus:outline-none focus:ring-2 focus:ring-blue-500 transition"
              />
            </div>
      
            {#if error}
              <p class="text-red-400 text-sm text-center">{error}</p>
            {/if}
      
            <button 
              type="submit" 
              class="w-full bg-white text-blue-800 py-3 rounded-md font-bold hover:bg-gray-200 transition duration-200"
            >
              Login
            </button>
          </form>
        </div>
      </div>
      
</div>
