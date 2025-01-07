import { writable } from 'svelte/store';

// Store to track the current stage and user data
export const formStage = writable(1);
export const userData = writable({
    id: "",
    names: [],
    guests: []
});