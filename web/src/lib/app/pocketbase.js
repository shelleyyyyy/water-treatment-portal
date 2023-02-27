import PocketBase from 'pocketbase';
import { writable } from 'svelte/store';
import { PUBLIC_POCKETBASE, PUBLIC_POCKETBASE_PORT } from '$env/static/public'

console.log(PUBLIC_POCKETBASE)
export const pb = new PocketBase(`http://${PUBLIC_POCKETBASE}:${PUBLIC_POCKETBASE_PORT}`); 

export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange((auth) => {
    console.log('authStore changed', auth);
    currentUser.set(pb.authStore.model);
});