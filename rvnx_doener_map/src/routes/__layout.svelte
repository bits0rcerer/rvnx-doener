<script>
    import "../app.css";
    import RvnxNavbar from "../lib/nav/RvnxNavbar.svelte";
    import {onMount} from "svelte";
    import { currentUserStore } from "../stores.js";

    onMount(() => {
        fetch("/api/twitch/me")
            .then(resp => resp.json())
            .then(data => {
                if (data.user) {
                    console.log(data)
                    currentUserStore.update(data.user.id, data.user.name, data.user.profile_image_url)
                }
            })
    })
</script>

<div class="flex flex-col h-screen">
    <RvnxNavbar/>
    <slot/>
</div>

