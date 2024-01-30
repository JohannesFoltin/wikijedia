// @ts-nocheck
import { writable } from "svelte/store";

function currentObjectID() {
    const { subscribe, set } = writable("");
    return {
        subscribe,
        set: (value) => set(value),
        reset: () => set("")
    };
}

export const objectID = currentObjectID();
