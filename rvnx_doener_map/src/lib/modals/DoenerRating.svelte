<script>
    import StarRating from "../common/StarRating.svelte";
    import TextRating from "../common/TextRating.svelte";
    import {currentUserStore, modalStore} from "../../stores.js";
    import RatePopup from "./RateShopModal.svelte";
    import {bind} from 'svelte-simple-modal';

    let user;
    currentUserStore.subscribe(value => {
        user = value
    })

    export let shop;

    function showModal(shop) {
        modalStore.set(bind(RatePopup, {
            shop: shop,
        }));
    }
</script>
<div class="flex flex-col mt-4">
    <StarRating unrated={!shop.rating} rating={shop.rating?.score}/>
    {#if shop.rating?.prices }
        <hr class="my-1">
        {#each shop.rating?.prices as p}
            <TextRating title={p[0]} value={p[1]}/>
        {/each}
    {/if}
    {#if user != null }
        <hr class="my-1">
        <button class="border-2 border-blue-400 rounded-full m-1 p-1 hover:bg-gray-100"
                on:click={() => showModal(shop)}>
            <span class="font-semibold">Bewerten</span>
        </button>
        <button class="border-2 border-orange-400 rounded-full m-1 p-1 hover:bg-gray-100">
            <span class="font-semibold">Änderung vorschlagen</span>
        </button>
    {/if}
</div>