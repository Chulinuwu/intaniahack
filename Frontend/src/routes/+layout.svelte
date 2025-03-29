<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { browser } from '$app/environment';
    import { token } from '$lib/auth';
	import '../app.css'

    onMount(() => {
        if (browser) {
            const pathname = window.location.pathname;
            if (!$token && pathname !== '/login' && pathname !== '/register') {
                goto('/home');
            }
        }
    });

    // reactive statement สำหรับ client-side changes
    $: if (browser && !$token && window.location.pathname !== '/login' && window.location.pathname !== '/register') {
        goto('/home');
    }
</script>

<slot />