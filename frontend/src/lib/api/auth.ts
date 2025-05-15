import { apiBaseUrl } from "$lib/config";
import { accessToken } from "$lib/stores/auth";
import { get } from "svelte/store";

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