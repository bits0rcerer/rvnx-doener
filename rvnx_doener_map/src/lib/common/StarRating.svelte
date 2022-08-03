<script>
    export let rating;
    export let unrated = false;
    export let editable = false;

    export let maxStars = 5;

    let hovering = false;
    let hoverRating = 2;

    export let ratingChangedCallback = null;
    export let starClass = "";
</script>

<div class="flex items-center">
    {#each [...Array(maxStars).keys()] as i}
        <svg aria-hidden="true"
             on:mouseenter={editable ? () => hovering = true : false}
             on:mouseleave={editable ? () => hovering = false: false}
             class={"star " + starClass}
             class:text-yellow-400={(hovering ? hoverRating : rating) >= i+1 && !unrated}
             class:text-gray-300={(hovering ? hoverRating : rating) < i+1 || unrated} fill="currentColor"
             on:mouseenter={editable ? () => hoverRating=i+1 : null}
             on:mouseup={editable ? () => {rating=i+1; if (ratingChangedCallback) ratingChangedCallback(rating)} : null}
             viewBox="0 0 20 20"
             xmlns="http://www.w3.org/2000/svg"><title>{i + 1}. star</title>
            <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
        </svg>
    {/each}

    <span class="ml-auto text-sm font-medium text-gray-500 dark:text-gray-400">{unrated ? "- /" : Math.min(maxStars, Math.max(0, rating)) + " /"} {maxStars}</span>
</div>

<style>
    .star {
        @apply w-6 h-6;
    }
</style>