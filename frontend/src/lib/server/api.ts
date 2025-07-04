import {type Cookies, json} from "@sveltejs/kit";
import {apiBaseUrl} from "$lib/server/config";

export async function requestWithAuth(
    endpoint: string,
    method: "GET" | "POST" | "PUT" | "DELETE",
    cookies: Cookies,
    body: BodyInit | null
) {
    let accessToken = cookies.get("accessToken");
    let refreshToken = cookies.get("refreshToken");

    if (!accessToken) {
        if (!refreshToken) {
            cookies.delete('accessToken', {path: '/'});
            cookies.delete('refreshToken', {path: '/'});
            return json({
                success: false,
                status: 400,
                requestId: null,
                timestamp: Date.now(),
                message: "No access or refresh token found",
                error: {
                    error: "missing_refresh_token",
                    details: null
                }
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
                return json(refreshData);
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
                requestId: null,
                timestamp: Date.now(),
                message: "Could not communicate with API",
                error: {
                    error: "server_error",
                    details: null
                }
            });
        }
    }

    // At this point, we have an access token, however, we are unsure if it is valid.
    try {
        const response = await fetch(apiBaseUrl + endpoint, {
            method: method,
            headers: {
                Authorization: `Bearer ${accessToken}`,
                "Content-Type": "application/json",
            },
            body: body ? JSON.stringify(body) : undefined
        });

        return await response.json();
    } catch (err) {
        return json({
            success: false,
            status: 503,
            requestId: null,
            timestamp: Date.now(),
            message: "Could not communicate with API",
            error: {
                error: "server_error",
                details: null
            }
        });
    }
}