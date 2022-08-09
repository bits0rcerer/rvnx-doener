<script>
    import {getNotificationsContext} from "svelte-notifications";
    import {notificationContextStore} from "../../stores.js";

    const notificationsContext = getNotificationsContext();

    notificationContextStore.set(notificationsContext)
</script>