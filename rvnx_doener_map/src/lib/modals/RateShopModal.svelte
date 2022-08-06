<script>
    import StarRating from "../common/StarRating.svelte";
    import CurrencyInput from "../common/CurrencyInput.svelte";
    import {modalStore} from "../../stores.js";

    export let shop;
    let userScore = -1;
    let normalKebabPrice = 0;
    let normalKebabCurrency = "EUR";
    let vegiKebabPrice = 0;
    let vegiKebabCurrency = "EUR";
    let opinion = "";

    let failed = "";

    function submitRating(anonymous) {
        let rating = {
            prices: {},
            anonymous: anonymous
        };
        if (userScore > 0)
            rating.userScore = userScore;

        if (normalKebabPrice > 0)
            rating.prices.normalKebab = {
                price: normalKebabPrice.toString(),
                currency: normalKebabCurrency
            };

        if (vegiKebabPrice > 0)
            rating.prices.vegiKebab = {
                price: vegiKebabPrice.toString(),
                currency: vegiKebabCurrency
            };

        if (opinion.trim() !== "")
            rating.opinion = opinion;

        fetch("/api/v1/kebabshops/" + shop.id + "/rate", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                rating: rating,
            })
        }).then((resp) => {
            if (resp.status === 200) {
                modalStore.set(null)
            } else {
                failed = "Absenden fehlgeschlagen. (" + resp.statusText + ")";
            }
        })
    }
</script>

<h1 class="font-semibold text-xl">{shop.name} bewerten</h1>
<hr/>
<h2 class="text-xl my-2">Allgemeine Bewertung</h2>
<StarRating rating={userScore} editable="true" ratingChangedCallback={(r) => userScore = r}/>
<hr class="mt-2"/>
<h2 class="text-xl my-2">Preise</h2>
<div class="grid grid-cols-2">
    <CurrencyInput title="Normaler Döner" value="0" currency="EUR" changedCallback={(price, currency) => {
        normalKebabPrice=price;
        normalKebabCurrency=currency;
    }}/>
    <CurrencyInput title="Vegetarischer Döner" value="0" currency="EUR" changedCallback={(price, currency) => {
        vegiKebabPrice=price;
        vegiKebabCurrency=currency;
    }}/>
</div>
<hr class="mt-2"/>
<h2 class="text-xl my-2">Deine Meinung</h2>
<textarea class="w-full border-2 h-28" placeholder="" on:change={(e) => opinion = e.target.value}></textarea>
<div class="mt-4 flex flex-row justify-center">
    <button on:click={() => submitRating(false)}
            class="border-2 border-green-400 hover:bg-green-50 py-1.5 px-6 rounded-full">
        Absenden
    </button>
    <div class="mx-2"></div>
    <button on:click={() => submitRating(true)}
            class="border-2 border-gray-300 hover:bg-gray-50 py-1.5 px-6 rounded-full">
        Anonym absenden
    </button>
</div>
{#if failed != null}
    <div class="mt-2 flex flex-row justify-center">
        <span class="text-red-500">{failed}</span>
    </div>
{/if}