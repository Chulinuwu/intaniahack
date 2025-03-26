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


<div class="container">
    <h1>Login</h1>
    {#if error}
        <p class="error">{error}</p>
    {/if}
    <form on:submit|preventDefault={handleLogin}>
        <input type="email" bind:value={email} placeholder="Email" required />
        <input type="password" bind:value={password} placeholder="Password" required />
        <button type="submit">Login</button>
    </form>
    <p>Don't have an account? <a href="/register">Register</a></p>
</div>

<style>
    .container {
        max-width: 400px;
        margin: 0 auto;
        padding: 2rem;
    }
    input {
        display: block;
        width: 100%;
        margin: 1rem 0;
        padding: 0.5rem;
    }
    .error {
        color: red;
    }
</style>