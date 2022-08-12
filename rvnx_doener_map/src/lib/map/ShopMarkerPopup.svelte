<script>
    import {onMount} from 'svelte';
    import DoenerRating from "../modals/DoenerRating.svelte";
    import {isApple} from "../common/device-detection.js";

    export let shopID;
    let shop = null;

    function loadShop() {
        fetch("/api/v1/kebabshops/" + shopID)
            .then(resp => resp.json())
            .then(data => data.shop ? data.shop : null)
            .then(shopData => {
                if (shopData.rating?.prices?.length > 0) {
                    shopData.rating.prices.sort((a, b) => a.order_index < b.order_index)
                }

                shop = shopData
            })
    }

    onMount(async () => {
        loadShop()
    });

    function getBaseMapLink() {
        // If it's an iPhone..
        const { isAppleDevice, isIOS } = isApple()

        if (isIOS || isAppleDevice) {
            return "https://maps.apple.com/"
        } else {
            return "https://maps.google.com/"
        }
    }
</script>

<div class="w-52">
    {#if shop == null }
        <span>Loading..</span>
    {:else}
        <h1 class="text-xl">{shop.name}</h1>
        <hr class="mb-1">
        <a href={getBaseMapLink() + "?ll=" + shop.lat + "," + shop.lng +"&q=" + shop.name} target="_blank">🧭
            Navigation</a>
        <DoenerRating reloadShopData={loadShop} shop={shop}/>
    {/if}
</div>