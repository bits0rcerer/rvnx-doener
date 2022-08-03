<!-- Content.svelte -->
<script>
    export let Popup;

    import { getContext } from 'svelte';
    const { open } = getContext('simple-modals');
    const showSurprise = () => open(Popup, { message: "It's a modals!" });
</script>

<p><button on:click={showSurprise}>Show me a surprise!</button></p>