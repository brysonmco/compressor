import type { PageServerLoad } from './$types';
import { isAuthenticated } from "$lib/server/auth";

export const load: PageServerLoad = async ({cookies}) => {
    // Check if the user is authenticated
    const authenticated = await isAuthenticated(cookies);
    return {
        authenticated: authenticated,
    };
};