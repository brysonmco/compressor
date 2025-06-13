import type {Actions, PageServerLoad} from './$types';
import {fail, json, redirect} from "@sveltejs/kit";
import {login, signup} from "$lib/server/auth";

export const load: PageServerLoad = async ({cookies}) => {

};

export const actions = {
    default: async ({cookies, request}) => {
        const data = await request.formData();
        const email = data.get('email')!.toString().trim();
        const firstName = data.get('firstName')!.toString().trim();
        const lastName = data.get('lastName')!.toString().trim();
        const password = data.get('password')!.toString();
        const confirmPassword = data.get('confirmPassword')!.toString();

        // Validate the fields
        const errors: Record<string, string> = {};

        if (!email) {
            errors.email = "Email is required";
        } else if (email.length > 250) {
            errors.email = "Email is too long";
        }

        if (!firstName) {
            errors.firstName = "First name is required";
        } else if (firstName.length > 100) {
            errors.firstName = "First name is too long";
        }

        if (!lastName) {
            errors.lastName = "Last name is required";
        } else if (lastName.length > 100) {
            errors.lastName = "Last name is too long";
        }

        if (!password) {
            errors.password = "Password is required";
        } else if (password.length < 8) {
            errors.password = "Password must be at least 8 characters long";
        } else if (password.length > 32) {
            errors.password = "Password is too long";
        }

        if (!confirmPassword) {
            errors.confirmPassword = "Confirm password is required";
        } else if (confirmPassword !== password) {
            errors.password = "Passwords do not match";
            errors.confirmPassword = "Passwords do not match";
        }


        if (Object.keys(errors).length > 0) {
            return fail(400, {
                errors,
            });
        }

        // Send the login request to the server
        const res = await signup(email, firstName, lastName, password);
        const resData = await res.json();
        if (!resData.success) {
            switch (resData.error) {
                case "server_error":
                    errors.server = "An error occurred while processing your request. Please try again later.";
                    return fail(500, {
                        errors,
                    });
                case "internal_error":
                    errors.server = "An error occurred while processing your request. Please try again later.";
                    return fail(500, {
                        errors,
                    });
                case "account_exists":
                    return redirect(300, '/login');
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

        return redirect(300, '/dashboard'); // hooks.server.ts will kick them to verify-email
    }
} satisfies Actions;