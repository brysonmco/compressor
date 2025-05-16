import {apiBaseUrl} from "$lib/config";
import {accessToken} from "$lib/stores/auth";
import {get} from "svelte/store";
import {json, redirect} from "@sveltejs/kit";

export async function refreshAccessToken(): Promise<boolean> {
    try {
        const res = await fetch(apiBaseUrl + "/auth/refresh", {
            method: "POST",
            credentials: "include",
        });
        if (!res.ok) {
            throw new Error("Failed to refresh access token");
        }

        const data = await res.json();
        accessToken.set(data.accessToken);
        return true;
    } catch (e) {
        accessToken.set(null)
        return false;
    }
}

export async function fetchWithAuth() {
    let token = get(accessToken);
    if (!token) {
        await refreshAccessToken();
        token = get(accessToken);
    }
}

export async function login(
    email: string, password: string
) {
    try {
        const response = await fetch(apiBaseUrl + "/auth/login", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email,
                password,
            })
        });

        const result = await response.json();

        return json({
            success: true,
            redirect: '/',
            statusCode: 303,
        });
    } catch (err) {
        return json({
            success: false,
            message: "Error occurred while logging in. Please try again later."
        });
    }
}

export async function signup(
    email: string,
    firstName: string,
    lastName: string,
    password: string,
    confirmPassword: string
) {
    try {
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
                confirmPassword,
            })
        });

        const result = await response.json();

        if (!response.ok) {
            switch (result.error) {
                case "missing_fields":
                    return json({
                        success: false,
                        fieldErrors: result.details,
                    });
                case "passwords_mismatch":
                    return json({
                        success: false,
                        fieldErrors: {
                            password: "Passwords do not match",
                            confirmPassword: "Passwords do not match",
                        }
                    });
                case "account_exists":
                    return json({
                        success: false,
                        redirect: '/login?email=' + email + '&error=account_exists',
                        statusCode: 303,
                    });
                default:
                    return json({
                        success: false,
                        message: "Error occurred while signing up. Please try again later."
                    });
            }
        }

        accessToken.set(result.accessToken);
        return json({
            success: true,
            redirect: '/',
            statusCode: 303,
        });
    } catch (err) {
        return json({
            success: false,
            message: "Error occurred while signing up. Please try again later."
        });
    }
}