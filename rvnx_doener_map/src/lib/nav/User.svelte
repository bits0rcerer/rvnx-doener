<script>
    // Due to bugs during build this library is not in use and
    // <Fa size="lg" icon={faRightFromBracket}/> is replaced by the actual svg
    //import Fa from 'svelte-fa'
    //import {faRightFromBracket} from '@fortawesome/free-solid-svg-icons'

    export let currentUser;
</script>

{#if currentUser != null }
    <div class="flex flex-row border-2 border-purple-twitch rounded-full py-2 px-3">
        <img height="32" width="32" class="rounded-full" src="{currentUser.profileImageURL}" alt="profile picture">
        <span class="my-auto mx-auto md:mx-4 font-semibold">{currentUser.name}</span>
        <a rel=external href="/api/twitch/logout" class="cursor-pointer my-auto">
            <!--<Fa size="lg" icon={faRightFromBracket}/>-->
            <svg style="height:1em;line-height:.75em;font-size:1.33333em;vertical-align:-.225em;transform-origin:center;overflow:visible" viewBox="0 0 512 512" aria-hidden="true" role="img" xmlns="http://www.w3.org/2000/svg"><g transform="translate(256 256)" transform-origin="128 0"><g transform="translate(0,0) scale(1,1)"><path d="M96 480h64C177.7 480 192 465.7 192 448S177.7 416 160 416H96c-17.67 0-32-14.33-32-32V128c0-17.67 14.33-32 32-32h64C177.7 96 192 81.67 192 64S177.7 32 160 32H96C42.98 32 0 74.98 0 128v256C0 437 42.98 480 96 480zM504.8 238.5l-144.1-136c-6.975-6.578-17.2-8.375-26-4.594c-8.803 3.797-14.51 12.47-14.51 22.05l-.0918 72l-128-.001c-17.69 0-32.02 14.33-32.02 32v64c0 17.67 14.34 32 32.02 32l128 .001l.0918 71.1c0 9.578 5.707 18.25 14.51 22.05c8.803 3.781 19.03 1.984 26-4.594l144.1-136C514.4 264.4 514.4 247.6 504.8 238.5z" fill="currentColor" transform="translate(-256 -256)"></path></g></g></svg>
        </a>
    </div>
{ :else }
    <a rel=external href="/api/twitch/login">
        <button class="m-auto bg-purple-twitch rounded-xl p-2 flex flex-row items-center">
            <div class="my-auto bg-white rounded-full p-2">
                <img height="18" width="18"
                     src="https://static.twitchcdn.net/assets/favicon-32-e29e246c157142c94346.png" alt="twitch icon">
            </div>
            <span class="text-white font-semibold whitespace-nowrap my-auto mx-2">Login with Twitch</span>
        </button>
    </a>
{/if}