import {apiBaseUrl} from "$lib/config";
import {json} from "@sveltejs/kit";

export async function getProfile(
    accessToken: string
): Promise<Response> {
    try {
        const response = await fetch(apiBaseUrl + "/users/profile", {
            method: "GET",
            headers: {
                Authorization: `Bearer ${accessToken}`,
                "Content-Type": "application/json",
            }
        });

        const data = await response.json();

        return json({});
    } catch (err) {
        return json({
            success: false,
            message: "Could not fetch user profile. Please try again later."
        });
    }
}