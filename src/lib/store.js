// @ts-nocheck
import { writable, get } from "svelte/store";

function currentObjectMethod() {
    const writableObject = writable(null);
    const {set, subscribe} = writableObject;
    return {
        subscribe,
        get: () => get(writableObject),
        set: (value) => set(value),
        reset: () => set(null)
    };
}

export const currentObject = currentObjectMethod();

function currentServerURL() {
    const writableObject = writable("http://localhost:8080/");
    const {set, subscribe} = writableObject;
    return {
        subscribe,
        get: () => get(writableObject),
        set: (value) => set(value),
        reset: () => set("")
    };
}

export const serverURL = currentServerURL();
