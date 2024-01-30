// @ts-nocheck
import { writable } from "svelte/store";

function currentText() {
    const { subscribe, set } = writable({});
    return {
        subscribe,
        set: (value) => set(value),
        reset: () => set({})
    };
}
export const textObjekt = currentText();
