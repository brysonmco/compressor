import { type Handle, redirect } from "@sveltejs/kit";
import { isAuthenticated, refresh } from "$lib/server/auth";

export const handle = (async ({ event, resolve }) => {
    if (!shouldProtectRoute(event.url.pathname)) {
        return resolve(event);
    }

    // Get tokens from cookies
    let accessToken = event.cookies.get('accessToken');
    let refreshToken = event.cookies.get('refreshToken');

    if (!accessToken) {
        if (!refreshToken) {
            event.cookies.delete('accessToken', { path: '/' });
            event.cookies.delete('refreshToken', { path: '/' });
            throw redirect(303, '/signup');
        }

        const refreshResponse = await refresh(refreshToken);
        const refreshData = await refreshResponse.json();

        if (!refreshData.success || !refreshData.accessToken) {
            event.cookies.delete('accessToken', { path: '/' });
            event.cookies.delete('refreshToken', { path: '/' });
            throw redirect(303, '/login');
        }

        event.cookies.set('accessToken', refreshData.accessToken, {
            httpOnly: true,
            path: '/',
            sameSite: 'strict',
            maxAge: 60 * 60
        });

        accessToken = refreshData.accessToken;
    }

    // Check if the user is authenticated
    const authResponse = await isAuthenticated(accessToken!);
    const authData = await authResponse.json();

    if (!authData.success) {
        event.cookies.delete('accessToken', { path: '/' });
        event.cookies.delete('refreshToken', { path: '/' });
        throw redirect(303, '/login');
    }

    return resolve(event);
}) satisfies Handle

function shouldProtectRoute(routeId: string) {
    return routeId.startsWith('/(protected)/')
}