<script>
    import StarRating from "../common/StarRating.svelte";
    import TextRating from "../common/TextRating.svelte";
    import {currentUserStore, modalStore, notificationContextStore} from "../../stores.js";
    import RatePopup from "./RateShopModal.svelte";
    import {bind} from 'svelte-simple-modal';

    let notificationContext;
    notificationContextStore.subscribe((value) => {
        notificationContext = value
    })

    let user;
    currentUserStore.subscribe(value => {
        user = value
    })

    export let shop;
    export let reloadShopData;

    function showModal(shop) {
        modalStore.set(bind(RatePopup, {
            shop: shop,
            reloadShopData: reloadShopData,
        }));
    }

    function formatPrice(price, currency) {
        return parseFloat(price).toFixed(2) + currencyDisplay(currency)
    }

    function currencyDisplay(currency) {
        switch (currency) {
            case "EUR":
                return " €"
            case "USD":
                return " $"
            default:
                return " " + currency
        }
    }

    function priceTypeDisplay(priceKey) {
        switch (priceKey) {
            case "normalKebab":
                return "Döner"
            case "vegiKebab":
                return "Vegitarischer Döner"
            case "normalYufka":
                return "Yufka"
            case "vegiYufka":
                return "Vegitarischer Yufka"
            case "doenerBox":
                return "Dönerbox"
            default:
                return priceKey
        }
    }

    function notifySubsOnly() {
        notificationContext.addNotification({
            id: "notActivatedWarning",
            text: 'Sorry, diese Funktion ist nur für Subs von RvNxSoul, RvNxMango oder Mahluna verfügbar. :/',
            position: 'bottom-right',
            type: 'danger',
            removeAfter: 8000,
        })
    }
</script>
<div class="flex flex-col mt-4">
    <StarRating unrated={!shop.rating.score} rating={shop.rating?.score}/>
    {#if shop.rating.prices }
        <hr class="my-1">
        {#each shop.rating.prices as p}
            <TextRating title={priceTypeDisplay(p.type)} value={formatPrice(p.price, p.currency)}/>
        {/each}
    {/if}
    {#if shop.rating.reviews }
        <hr class="my-1">
        {#each shop.rating.reviews as r}
            <div class="flex flex-col mt-1.5">
                <div class="flex flex-row">
                    <h4 class="text-xs font-semibold">{r.author ? r.author : "Anonym"}</h4>
                    <span class="text-xs font-light mx-auto"> </span>
                    <span class="text-xs font-light">{ new Date(Date.parse(r.date)).toLocaleDateString()}</span>
                </div>
                <span class="mt-0.5">{r.review}</span>
            </div>
        {/each}
    {/if}
    {#if user != null }
        <hr class="my-1">
        <button class={user.activated ? 'rate-button-enabled': 'button-disabled'}
                on:click={() => {
                    if (user.activated) {
                        showModal(shop);
                    } else {
                        notifySubsOnly();
                    }
                }}>
            <span class="font-semibold">Bewerten</span>
        </button>
        <!--<button class="border-2 border-orange-400 rounded-full m-1 p-1 hover:bg-gray-100">
            <span class="font-semibold">Änderung vorschlagen</span>
        </button>-->
    {/if}
</div>

<style>
    button {
        @apply border-2 rounded-full m-1 p-1;
    }

    button:hover {
        @apply bg-gray-100;
    }

    .rate-button-enabled {
        @apply border-blue-400 ;
    }

    .button-disabled {
        @apply border-gray-400 text-gray-400;
    }
</style>