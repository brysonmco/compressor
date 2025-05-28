import type { PageServerLoad } from './$types';
import { isAuthenticated } from "$lib/server/auth";
import { getProfile } from "$lib/server/users";

export const load: PageServerLoad = async ({cookies}) => {
    // Check if the user is authenticated
    const authenticated = await isAuthenticated(cookies);

    // Get their subscription
    let profile = null;
    if (authenticated) {
        profile = await getProfile(cookies.get('accessToken')!);
    }

    return {
        authenticated: authenticated,
        user: profile,
    };
};