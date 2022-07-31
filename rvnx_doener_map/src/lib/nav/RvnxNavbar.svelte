<script>
    import {currentUserStore} from "../../stores.js";
    import User from "./User.svelte";

    let open = false;
    function toggleOpen() {
        open = !open;
    }

    let u;
    currentUserStore.subscribe(value => {
        u = value
        console.log(value)
    })
</script>

<nav class="bg-white border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-900">
    <div class="container flex flex-wrap justify-between items-center mx-auto">
        <a href="/" class="flex items-center">
                <img src="/favicon.png" class="mr-3 h-6 sm:h-9 p-1 rounded bg-white" alt="RVNX Döner Map Logo"/>
            <span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white">RVNX Döner Map <span class="text-red-500">Work-in-progres</span></span>
        </a>
        <button on:click={toggleOpen} type="button"
                class="inline-flex items-center p-2 ml-3 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
                aria-controls="navbar-default" aria-expanded="false">
            <span class="sr-only">Open menu</span>
            <svg class="w-6 h-6" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20"
                 xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                      d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                      clip-rule="evenodd"></path>
            </svg>
        </button>
        <div class:hidden={!open} class="w-full md:block md:w-auto" id="navbar-default">
            <ul class="flex flex-col p-4 mt-4 bg-gray-50 rounded-lg border border-gray-100 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                <li>
                    <User currentUser={u} />
                </li>
            </ul>
        </div>
    </div>
</nav>