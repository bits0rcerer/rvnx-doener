<script>
    export let title;
    export let price;
    export let currency;

    export let changedCallback;

    function priceChanged() {
        price = parseFloat(this.value).toFixed(2);
        if (changedCallback) changedCallback(price, currency);
    }

    function clear() {
        price = 0;
        if (changedCallback) changedCallback(price, currency);
    }

    function currencyChanged(e) {
        currency = e.target.value;
        if (changedCallback) changedCallback(price, currency);
    }
</script>

<label class="block text-gray-700 text-sm font-bold my-auto" for="price">
    {title}
</label>
<div class="rounded-md shadow-sm flex flex-row">
    <datalist id="commonPrices">
        {#each [...Array(5).keys()] as i}
            <option value={(i+3.0).toFixed(2)}>
            <option value={(i+3.5).toFixed(2)}>
        {/each}
    </datalist>
    <input
            on:change={priceChanged}
            value={price} type="number" min="0.00" step="0.10" name="normal" id="price"
            list="commonPrices"
            class="focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md px-2 py-1"
            placeholder="0.00">
    <button on:click={clear}>✖</button>
    <select on:change={currencyChanged} id="currency" name="currency"
            class="focus:ring-indigo-500 focus:border-indigo-500 h-full px-2 py-1 border-transparent bg-transparent text-gray-500 sm:text-sm rounded-md">
        <option selected>EUR</option>
        <option>CHF</option>
        <option>JPY</option>
        <option>SEK</option>
        <option>DDK</option>
        <option>USD</option>
        <option>GBP</option>
    </select>
</div>