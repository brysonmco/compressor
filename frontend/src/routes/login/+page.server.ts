import type {Actions, PageServerLoad} from './$types';
import {fail, json, redirect} from "@sveltejs/kit";
import {login} from "$lib/server/auth";

export const load: PageServerLoad = async ({cookies}) => {

};

export const actions = {
    default: async ({cookies, request}) => {
        const data = await request.formData();
        const email = data.get('email')!.toString().trim();
        const password = data.get('password')!.toString();

        // Validate the email and password
        const errors: Record<string, string> = {};

        if (!email) {
            errors.email = "Email is required";
        } else if (email.length > 250) {
            errors.email = "Email is too long";
        }

        if (!password) {
            errors.password = "Password is required";
        } else if (password.length > 250) {
            errors.password = "Password is too long";
        }


        if (Object.keys(errors).length > 0) {
            return fail(400, {
                errors,
            });
        }

        // Send the login request to the server
        const res = await login(email, password)
        const resData = await res.json();
        if (!resData.success) {
            switch (resData.error) {
                case "server_error":
                    errors.server = "An error occurred while processing your request. Please try again later.";
                    return fail(500, {
                        errors,
                    });
                case "invalid_credentials":
                    errors.email = "Incorrect email or password";
                    errors.password = "Incorrect email or password";
                    return fail(400, {
                        errors,
                    });
                default:
                    errors.server = "An unexpected error occurred. Please try again later.";
                    return fail(500, {
                        errors,
                    });
            }
        }

        const accessToken = resData.accessToken;
        const refreshToken = resData.refreshToken;

        if (!accessToken || !refreshToken) {
            errors.server = "An error occurred while processing your request. Please try again later.";
            return fail(500, {
                errors,
            });
        }

        cookies.set(
            'accessToken',
            accessToken,
            {
                path: '/',
                httpOnly: true,
                sameSite: 'strict',
            }
        );

        cookies.set(
            'refreshToken',
            refreshToken,
            {
                path: '/',
                httpOnly: true,
                sameSite: 'strict',
            }
        );

        return redirect(300, '/dashboard');
    }
} satisfies Actions;