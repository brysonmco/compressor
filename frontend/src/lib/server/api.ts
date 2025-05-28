import {type Cookies, json} from "@sveltejs/kit";
import { apiBaseUrl } from "$lib/server/config";
import {refresh} from "$lib/server/auth";

export async function requestWithAuth(
    endpoint: string,
    method: "GET" | "POST" | "PUT" | "DELETE",
    cookies: Cookies,
    body: BodyInit
): Promise<Response> {
    let accessToken = cookies.get("accessToken");
    let refreshToken = cookies.get("refreshToken");

    if (!accessToken) {
        if (!refreshToken) {
            cookies.delete('accessToken', {path: '/'});
            cookies.delete('refreshToken', {path: '/'});
            return json({
                success: false,
                message: "No access or refresh token found"
            });
        }

        try {
            const refreshResponse = await fetch(apiBaseUrl + "/auth/refresh", {
                method: "POST",
                credentials: "include",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    refreshToken,
                })
            });
            const refreshData = await refreshResponse.json();

            if (!refreshResponse.ok || !refreshData.data.accessToken) {
                cookies.delete('accessToken', {path: '/'});
                cookies.delete('refreshToken', {path: '/'});
                return json({
                    success: false,
                    status: refreshData.status,
                    message: refreshData.message,
                    error: refreshData.error
                });
            }

            cookies.set('accessToken', refreshData.data.accessToken, {
                httpOnly: true,
                path: '/',
                sameSite: 'strict',
                maxAge: 60 * 60
            });

            accessToken = refreshData.data.accessToken;
        } catch (err) {
            cookies.delete('accessToken', {path: '/'});
            cookies.delete('refreshToken', {path: '/'});
            return json({
                success: false,
                status: 503,
                message: "Could not refresh access token.",
                error: {
                    error: "api_unreachable",
                    message: "Could not reach the API to refresh the access token."
                }
            });
        }
    }

    // At this point, we have an access token, however we are unsure if it is valid.
    try {
        const response = await fetch(apiBaseUrl + endpoint, {
            method: method,
            headers: {
                Authorization: `Bearer ${accessToken}`,
                "Content-Type": "application/json",
            },
            body: body ? JSON.stringify(body) : undefined
        });

        if (!response.ok) {
            const errorData = await response.json();
            return json({
                success: false,
                message: errorData.message || "An error occurred while processing the request."
            });
        }

        return response;
    } catch (err) {
        return json({
            success: false,
            message: "Could not reach API"
        });
    }
}