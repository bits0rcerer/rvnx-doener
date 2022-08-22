<script>
	import { currentUserStore } from '../../stores.js';
	import User from './User.svelte';
	import Fa from 'svelte-fa/src/fa.svelte';
	import { faCirclePlus } from '@fortawesome/free-solid-svg-icons/index.es';
	import { faGithub } from '@fortawesome/free-brands-svg-icons/index.es';

	let open = false;

	function toggleOpen() {
		open = !open;
	}

	let u;
	currentUserStore.subscribe((value) => {
		u = value;
	});
</script>

<nav class="bg-white border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-900">
	<div class="container flex flex-wrap justify-between items-center mx-auto">
		<a href="/" class="flex items-center">
			<img
				src="/favicon.png"
				class="mr-3 h-6 sm:h-9 p-1 rounded bg-white"
				alt="RVNX Döner Map Logo"
			/>
			<span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
				><span class="text-red-500">[WIP]</span> RVNX Döner Map</span
			>
		</a>
		<button
			on:click={toggleOpen}
			type="button"
			class="inline-flex items-center p-2 ml-3 text-sm text-gray-500 rounded-lg lg:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
			aria-controls="navbar-default"
			aria-expanded="false"
		>
			<span class="sr-only">Open menu</span>
			<svg
				class="w-6 h-6"
				aria-hidden="true"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
			>
				<path
					fill-rule="evenodd"
					d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
					clip-rule="evenodd"
				/>
			</svg>
		</button>
		<div class="w-full lg:block lg:w-auto" class:hidden={!open} id="navbar-default">
			<ul
				class="flex flex-col p-4 mt-4 bg-gray-50 rounded-lg border border-gray-100 md:flex-row lg:space-x-8 lg:mt-0 lg:text-sm lg:font-medium md:border-0 lg:bg-white dark:bg-gray-800 lg:dark:bg-gray-900 dark:border-gray-700"
			>
				{#if u != null}
					<li class="m-auto py-2">
						<a
							class="text-xs font-semibold hover:text-blue-500 cursor-pointer dark:text-white"
							href="/new"
						>
							<Fa class="inline px-1" size="lg" icon={faCirclePlus} />
							<span class="hover:underline underline-offset-1">Dönerladen hinzufügen</span></a
						>
					</li>
				{/if}
				<li class="m-auto py-2">
					<a
						class="text-xs font-semibold hover:text-blue-500 cursor-pointer dark:text-white"
						target="_blank"
						href="https://github.com/bits0rcerer/rvnx-doener"
					>
						<Fa class="inline px-1" size="lg" icon={faGithub} />
						<span class="hover:underline underline-offset-1">GitHub</span></a
					>
				</li>
				<li>
					<User currentUser={u} />
				</li>
			</ul>
		</div>
	</div>
</nav>
