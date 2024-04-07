import { writable, get } from "svelte/store";
import type { BackendFolder, BackendObject } from "./types";

function currentObjectMethod() {
    const object: BackendObject | null = null;
    const writableObject = writable(object);
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
        set: (value:string) => set(value),
        reset: () => set("")
    };
}

export const serverURL = currentServerURL();


function updateNotifiyer() {
    const writableObject = writable(false);
    const {set, subscribe} = writableObject;
    return {
        subscribe,
        get: () => get(writableObject),
        set: () => set(true),
        reset: () => set(false)
    };
}

export const updateStructure  = updateNotifiyer();

function showMoveDialog2() {
    const object: BackendObject | BackendFolder | null = null;
    const writableObject = writable(object);
    const {set, subscribe} = writableObject;
    return {
        subscribe,
        get: () => get(writableObject),
        set: (object) => set(object),
        reset: () => set(null)
    };
}

export const showMoveDialog  = showMoveDialog2();