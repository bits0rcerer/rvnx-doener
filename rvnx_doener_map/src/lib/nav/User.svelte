<script>
	import Fa from 'svelte-fa';
	import { faRightFromBracket } from '@fortawesome/free-solid-svg-icons';

	export let currentUser;
</script>

{#if currentUser != null}
	<div class="flex flex-row border-2 border-purple-twitch rounded-full py-2 px-3">
		<img
			height="32"
			width="32"
			class="rounded-full"
			src={currentUser.profileImageURL}
			alt="profile"
		/>
		<span class="my-auto mx-auto md:mx-4 font-semibold">{currentUser.name}</span>
		<a rel="external" href="/api/twitch/logout" class="cursor-pointer my-auto">
			<Fa size="lg" icon={faRightFromBracket} />
		</a>
	</div>
{:else}
	<a rel="external" href="/api/twitch/login">
		<div class="m-auto bg-purple-twitch rounded-xl p-2 flex flex-row items-center">
			<div class="my-auto bg-white rounded-full p-2">
				<img
					height="18"
					width="18"
					src="https://static.twitchcdn.net/assets/favicon-32-e29e246c157142c94346.png"
					alt="twitch icon"
				/>
			</div>
			<span class="text-white font-semibold whitespace-nowrap my-auto mx-2">Login with Twitch</span>
		</div>
	</a>
{/if}
