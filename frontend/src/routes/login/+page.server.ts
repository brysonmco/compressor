import type { Actions } from './$types';
import {isAuthenticated, login} from "$lib/server/auth";
import {json, redirect} from "@sveltejs/kit";
import type {PageServerLoad} from "../../../.svelte-kit/types/src/routes/pricing/$types";

export const load: PageServerLoad = async ({cookies}) => {
    // Check if the user is authenticated
    const authenticated = await isAuthenticated(cookies);

    if (authenticated) {
        // If authenticated, redirect to the home page
        return redirect(303, '/dashboard');
    }
};

export const actions = {
    default: async ({ cookies, request }) => {
        const data = await request.formData();
        const email = data.get('email')!.toString();
        const password = data.get('password')!.toString();

        let req = await login(email, password)

        let res = await req.json();
        if (!res.success) {
            return {
                "success": false,
                "message": res.message,
                "formErrors": res.fieldErrors,
            }
        }

        cookies.set('accessToken', res.accessToken, {
            httpOnly: true,
            path: '/',
        });
        cookies.set('refreshToken', res.refreshToken, {
            httpOnly: true,
            path: '/',
        });

        redirect(303, '/dashboard');
    }
} satisfies Actions;