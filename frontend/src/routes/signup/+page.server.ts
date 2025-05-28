import type {Actions} from "./$types";
import {apiBaseUrl} from "$lib/config";
import {fail, redirect} from "@sveltejs/kit";
import {accessToken} from "$lib/stores/auth";
import type {PageServerLoad} from "../../../.svelte-kit/types/src/routes/pricing/$types";
import {isAuthenticated} from "$lib/server/auth";

export const load: PageServerLoad = async ({cookies}) => {
    // Check if the user is authenticated
    const authenticated = await isAuthenticated(cookies);

    if (authenticated) {
        // If authenticated, redirect to the home page
        return redirect(303, '/dashboard');
    }
};

export const actions = {
    default: async ({cookies, request}) => {
        const data = await request.formData();
        const email = data.get('email');
        const firstName = data.get('firstName');
        const lastName = data.get('lastName');
        const password = data.get('password');
        const passwordConfirm = data.get('passwordConfirm');

        const response = await fetch(apiBaseUrl + "/auth/signup", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email,
                firstName,
                lastName,
                password,
                passwordConfirm,
            })
        });

        const result = await response.json();

        if (!response.ok) {
            switch (result.error) {
                case "missing_fields":
                    // TODO: implement this later
                case "passwords_mismatch":
                // TODO: implement this later
                case "account_exists":
                    redirect(303, '/login?email=' + email);
                    break;
                default:
                    // TODO: implement this later
            }
        }

        accessToken.set(result.accessToken);
        return redirect(303, '/');
    }
} satisfies Actions;