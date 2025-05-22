import {accessToken} from "$lib/stores/auth";
import type {Handle} from "@sveltejs/kit";
import {getProfile} from "$lib/api/users";

export const handle = (async ({ event, resolve }): Promise<Response> => {
    if (!shouldProtectRoute(event.route.id!)) {
        return resolve(event);
    }

    // Check that an accessToken has been set
    if (accessToken === null) {
        // Redirect to login page if access token is not present
        return Response.redirect('/login', 303);
    }

    // Grab their profile and see if it works
    const response = await getProfile();



    return resolve(event);
}) satisfies Handle;

function shouldProtectRoute(routeId: string): boolean {
    return routeId.startsWith('/(protected)/');
}