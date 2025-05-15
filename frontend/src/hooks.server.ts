import {accessToken} from "$lib/stores/auth";

export async function handle({ event, resolve }) {
    if (accessToken) {

    }

    return resolve(event);
}