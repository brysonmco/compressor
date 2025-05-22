import {apiBaseUrl} from "$lib/config";
import {get} from "svelte/store";
import {accessToken} from "$lib/stores/auth";
import {refreshAccessToken} from "$lib/api/auth";
import {json} from "@sveltejs/kit";

type User = {
    id: bigint,
    email: string,
}

export type UserProfile = {
    id: bigint,
    firstName: string,
    lastName: string,
}

export async function getProfile(): Promise<Response> {
    const response = await fetch(apiBaseUrl + "/users/profile", {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${get(accessToken)}`,
            "Content-Type": "application/json"
        }});

    const data = await response.json();
    if (!response.ok) {
        switch (data.error.error) {
            case "invalid_token":
                accessToken.set(null);
                return json({
                    success: false,
                    redirect: '/login',
                    statusCode: 303,
                });
            case "expired_token":
                const ref = await refreshAccessToken();

                let refData = await ref.json();
                if (refData.success) {
                    return getProfile();
                } else {
                    return json({
                        success: false,
                        redirect: '/login',
                        statusCode: 303,
                    });
                }
            case "user_not_found":
                // Somehow the user has a bad token that the backend accepts as valid, not good
                accessToken.set(null);
                return json({
                    success: false,
                    redirect: '/login',
                    statusCode: 303,
                });
            case "internal_error":
                return json({
                    success: false,
                    message: "Error occurred while communicating with api. Please try again later."
                });
            default:
                return json({
                    success: false,
                    redirect: '/login',
                    statusCode: 303,
                });
        }
    }

    return json({
        success: true,
        profile: {
            id: data.id,
            firstName: data.first_name,
            lastName: data.last_name,
        }
    });
}