<script>
    import "../app.css";
    import RvnxNavbar from "../lib/nav/RvnxNavbar.svelte";
    import {onMount} from "svelte";
    import {currentUserStore, modalStore} from "../stores.js";
    import {Modal} from "svelte-simple-modal";
    import Notifications from 'svelte-notifications';
    import NotificationContextGetter from "../lib/common/NotifcationContextGetter.svelte";
    import CustomNotification from "../lib/common/CustomNotification.svelte";

    onMount(() => {
        fetch("/api/twitch/me")
            .then(resp => resp.json())
            .then(data => {
                if (data.user) {
                    currentUserStore.update(data.user.id, data.user.name, data.user.profile_image_url, data.user.activated)
                }
            })
    })
</script>

<Notifications zIndex="1000" item={CustomNotification}>
    <NotificationContextGetter/>

    <Modal show={$modalStore}>
        <div class="flex flex-col h-screen">
            <RvnxNavbar/>
            <slot/>
        </div>
    </Modal>
</Notifications>