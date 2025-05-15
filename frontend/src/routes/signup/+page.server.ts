import type {Actions} from "./$types";
import {apiBaseUrl} from "$lib/config";
import {fail, redirect} from "@sveltejs/kit";

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
                case "bad_json":
                    return fail(400, result.message);
                case "missing_fields":
                    // TODO: implement this later, it seems like a headache
                case "email_already_exists":
                    throw redirect(303, '/login?email=' + email);

            }
        }
    }
} satisfies Actions;