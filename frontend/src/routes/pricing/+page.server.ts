import type { PageServerLoad } from './$types';
import { getProfile } from "$lib/server/users";

export const load: PageServerLoad = async ({cookies}) => {
    // Check if the user is authenticated

    // Get their subscription
    let profile = null;

    return {
        authenticated: false,
        user: profile,
    };
};