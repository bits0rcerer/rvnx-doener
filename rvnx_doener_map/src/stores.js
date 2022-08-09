import {writable} from "svelte/store";

function createUserStore() {
    const {subscribe, set, update} = writable(null);

    return {
        subscribe,
        update: (userID, userName, profileImageURL, activated) => update(() => {
            return {
                userID: userID,
                name: userName,
                profileImageURL: profileImageURL,
                activated: activated,
            }
        }),
        clear: () => set(null)
    };
}

export const currentUserStore = createUserStore();

export const modalStore = writable(null)

export const notificationContextStore = writable(null)