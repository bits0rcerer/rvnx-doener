<script>
    import {onMount} from 'svelte';
    import DoenerRating from "../modals/DoenerRating.svelte";

    export let shopID;
    let shop = null;

    onMount(async () => {
        fetch("/api/v1/kebabshops/" + shopID)
            .then(resp => resp.json())
            .then(data => data.shop ? data.shop : null)
            .then(data => {
                // TODO: REMOVE MOCK DATA
                console.warn("TODO: REMOVE MOCK DATA")
                    data.rating = {
                        score: 4,
                        prices: [
                            ["Normaler Döner", "4,50€"],
                            ["Vegitarischer Döner", "4,00€"],
                        ],
                    }
                    shop = data
                }
            )
    });

    function getBaseMapLink() {
        // If it's an iPhone..
        const isIOS = /iPad|iPhone|iPod/.test(navigator.userAgent);
        const isAppleDevice = navigator.userAgent.includes('Macintosh');

        if (isIOS || isAppleDevice) {
            return "https://maps.apple.com/"
        } else {
            return "https://maps.google.com/"
        }
    }
</script>

<div class="w-48">
    {#if shop == null }
        <span>Loading..</span>
    {:else}
        <h1 class="text-xl">{shop.name}</h1>
        <hr class="mb-1">
        <a href={getBaseMapLink() + "?ll=" + shop.lat + "," + shop.lng +"&q=" + shop.name} target="_blank">🧭
            Navigation</a>
        <DoenerRating shop={shop}/>
    {/if}
</div>