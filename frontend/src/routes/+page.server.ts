import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({cookies}) => {
    // Check if the user is authenticated

    // Get their subscription
    let profile = null;

    console.log("I EXECUTED");
    return {
        authenticated: false,
        user: profile
    };
};