import { type Handle, redirect } from "@sveltejs/kit";
import { isAuthenticated } from "$lib/server/auth";

export const handle = (async ({ event, resolve }) => {
    if (!shouldProtectRoute(event.route.id!)) {
        return resolve(event);
    }

    // Check if the user is authenticated
    const authenticated = await isAuthenticated(event.cookies)
    if (!authenticated) {
        // If not authenticated, redirect to the login page
        throw redirect(303, '/login');
    }

    return resolve(event);
}) satisfies Handle

function shouldProtectRoute(routeId: string) {
    return routeId.startsWith('/(protected)/')
}