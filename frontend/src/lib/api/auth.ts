import {apiBaseUrl} from "$lib/config";
import {accessToken} from "$lib/stores/auth";
import {get} from "svelte/store";
import {resolveRoute} from "$app/paths";
import {redirect} from "@sveltejs/kit";

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

export async function signup(
    email: string,
    firstName: string,
    lastName: string,
    password: string,
    passwordConfirm: string
) {
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