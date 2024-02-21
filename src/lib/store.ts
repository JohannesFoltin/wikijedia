import { writable, get } from "svelte/store";
import { BackendObject } from "$lib/structureData";

function currentObjectMethod() {
    const object : null | BackendObject = null;
    const writableObject = writable(object);
    const {set, subscribe} = writableObject;
    return {
        subscribe,
        get: () => get(writableObject),
        set: (value:null|BackendObject) => set(value),
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
        set: (value:string) => set(value),
        reset: () => set("")
    };
}

export const serverURL = currentServerURL();
