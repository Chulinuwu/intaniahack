<script>
    import { goto } from '$app/navigation';

    let email = '';
    let password = '';
    let error = '';

    async function handleRegister() {
        try {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ email, password })
            });

            const data = await response.json();
            
            if (!response.ok) {
                throw new Error(data.message || 'Registration failed');
            }

            goto('/login');
        } catch (err) {
            error = err instanceof Error ? err.message : 'An unknown error occurred';
        }
    }
</script>

<div class="container">
    <h1>Register</h1>
    {#if error}
        <p class="error">{error}</p>
    {/if}
    <form on:submit|preventDefault={handleRegister}>
        <input type="email" bind:value={email} placeholder="Email" required />
        <input type="password" bind:value={password} placeholder="Password" required />
        <button type="submit">Register</button>
    </form>
    <p>Already have an account? <a href="/login">Login</a></p>
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