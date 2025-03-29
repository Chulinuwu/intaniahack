<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { browser } from '$app/environment';
    import { token } from '$lib/auth';
    import '../app.css';
    import RotatePage from '../components/Rotate/RotatePage.svelte';

    let isPortrait = false;

    function checkOrientation() {
        isPortrait = window.matchMedia('(orientation: portrait)').matches;
    }

    onMount(() => {
        if (browser) {
            const pathname = window.location.pathname;
            if (!$token && pathname !== '/login' && pathname !== '/register') {
                goto('/login');
            }

            // check orientation on mount
            checkOrientation();
            window.addEventListener('resize', checkOrientation);

            return () => {
                window.removeEventListener('resize', checkOrientation);
            };
        }
    });

    // reactive statement สำหรับ client-side changes
    $: if (browser && !$token && window.location.pathname !== '/login' && window.location.pathname !== '/register') {
        goto('/login');
    }
</script>

{#if isPortrait}
    <RotatePage />
{:else}
    <slot />
{/if}