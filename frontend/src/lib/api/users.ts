import {apiBaseUrl} from "$lib/config";
import {get} from "svelte/store";
import {accessToken} from "$lib/stores/auth";

type User = {
    id: bigint,
    email: string,
}

export type UserProfile = {
    id: bigint,
    firstName: string,
    lastName: string,
    avatarUrl: string,
}

export async function getUserProfile(id: bigint): Promise<UserProfile> {
    const res = await fetch(apiBaseUrl + "/users/" + id, {
        method: "GET",
        credentials: "include",
        headers: {
            "Authorization": "Bearer " + get(accessToken)
        }
    });

    if (!res.ok) {
        throw new Error("Failed to get user profile");
    }

    return res.json();
}