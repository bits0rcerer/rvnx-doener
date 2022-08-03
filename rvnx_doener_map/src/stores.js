import {writable} from "svelte/store";

function createUserStore() {
    const {subscribe, set, update} = writable(null);

    return {
        subscribe,
        update: (userID, userName, profileImageURL) => update(() => {
            return {
                userID: userID,
                name: userName,
                profileImageURL: profileImageURL
            }
        }),
        clear: () => set(null)
    };
}

export const currentUserStore = createUserStore();

export const modalStore = writable(null)