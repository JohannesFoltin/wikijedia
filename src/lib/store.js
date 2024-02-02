// @ts-nocheck
import { writable, get } from "svelte/store";

function currentObjectID() {
    const writableObject = writable("");
    const {set, subscribe} = writableObject;
    return {
        subscribe,
        get: () => get(writableObject),
        set: (value) => set(value),
        reset: () => set("")
    };
}

export const objectID = currentObjectID();

function currentServerURL() {
    const { subscribe, set } = writable("http://localhost:8080/");
    return {
        subscribe,
        set: (value) => set(value),
        reset: () => set("")
    };
}

export const serverURL = currentServerURL();
