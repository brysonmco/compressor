import { type Handle, redirect } from "@sveltejs/kit";
import {requestWithAuth} from "$lib/server/api";

export const handle = (async ({ event, resolve }) => {
    if (!shouldProtectRoute(event.route.id!)) {
        return resolve(event);
    }

    // Check if the user is authenticated
    const profileResponse = await requestWithAuth("/users/profile", "GET", event.cookies, null);

    return resolve(event);
}) satisfies Handle

function shouldProtectRoute(routeId: string) {
    return routeId.startsWith('/(protected)/')
}