<script>
    import StarRating from "./StarRating.svelte";
    import TextRating from "./TextRating.svelte";
    import {currentUserStore} from "../../stores.js";

    let user;
    currentUserStore.subscribe(value => {
        user = value
    })

    export let doenerRating = {
        score: 4,
        prices: [
            ["Normaler Döner", "4,50€"],
            ["Vegitarischer Döner", "4,00€"],
        ],
    }
</script>

<div class="flex flex-col mt-4">
    <StarRating unrated={!doenerRating} rating={doenerRating?.score}/>
    {#if doenerRating?.prices }
        <hr class="my-1">
        {#each doenerRating?.prices as p}
            <TextRating title={p[0]} value={p[1]}/>
        {/each}
    {/if}
    {#if user != null }
        <hr class="my-1">
        <button class="border-2 border-blue-400 rounded-full m-1 p-1 hover:bg-gray-100">
            <span class="font-semibold">Bewerten</span>
        </button>
        <button class="border-2 border-orange-400 rounded-full m-1 p-1 hover:bg-gray-100">
            <span class="font-semibold">Änderung vorschlagen</span>
        </button>
    {/if}
</div>