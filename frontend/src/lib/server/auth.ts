import {type Cookies, json} from "@sveltejs/kit";
import { apiBaseUrl } from "$lib/server/config";
import {getProfile} from "$lib/server/users";

export async function refresh(
    refreshToken: string
): Promise<Response> {
    try {
        const response = await fetch(apiBaseUrl + "/auth/refresh", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                refreshToken,
            })
        });

        const data = await response.json();
        if (!response.ok) {
            return json({
                success: false,
                message: "Error occurred while refreshing token. Please try again later."
            });
        }

        return json({
            success: true,
            accessToken: data.accessToken,
        });
    } catch (e) {
        return json({
            success: false,
            message: "Error occurred while refreshing token. Please try again later."
        });
    }
}

export async function login(
    email: string,
    password: string
): Promise<Response> {
    try {
        const response = await fetch(apiBaseUrl + "/auth/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email,
                password,
            })
        });

        const data = await response.json();
        if (!response.ok) {
            return json({
                success: false,
                error: data.error.error
            });
        }

        return json({
            success: true,
            accessToken: data.data.accessToken,
            refreshToken: data.data.refreshToken,
        });
    } catch (e) {
        return json({
            success: false,
            error: "server_error",
        });
    }
}

export async function isAuthenticated(
    cookies: Cookies
): Promise<boolean> {
    let accessToken = cookies.get("accessToken");
    const refreshToken = cookies.get("refreshToken");

    if (!accessToken) {
        if (!refreshToken) {
            cookies.delete('accessToken', {path: '/'});
            cookies.delete('refreshToken', {path: '/'});
            return false;
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
                return false;
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
            return false;
        }
    }

    try {
        const profileResponse = await fetch(apiBaseUrl + "/users/profile", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${accessToken}`
            }
        })
        if (!profileResponse.ok) {
            cookies.delete('accessToken', {path: '/'});
            cookies.delete('refreshToken', {path: '/'});
            return false;
        }
        const profileData = await profileResponse.json();
        if (!profileData.success) {
            cookies.delete('accessToken', {path: '/'});
            cookies.delete('refreshToken', {path: '/'});
            return false;
        }
        return true;
    } catch (err) {
        cookies.delete('accessToken', {path: '/'});
        cookies.delete('refreshToken', {path: '/'});
        return false;
    }
}

export async function signup(
    email: string,
    password: string,
    firstName: string,
    lastName: string
): Promise<Response> {
    try {
        const response = await fetch(apiBaseUrl + "/auth/signup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email,
                password,
                firstName,
                lastName,
            })
        });

        const data = await response.json();
        if (!response.ok) {
            switch (data.error) {
                case "missing_fields":
                    return json({
                        success: false,
                        fieldErrors: data.details,
                    });
                case "email_already_exists":
                    return json({
                        success: false,
                        fieldErrors: {
                            email: "Email already exists.",
                        }
                    });
                default:
                    return json({
                        success: false,
                        message: "Error occurred while registering. Please try again later."
                    });
            }
        }

        return json({
            success: true,
            accessToken: data.accessToken,
            refreshToken: data.refreshToken,
        });
    } catch (e) {
        return json({
            success: false,
            message: "Error occurred while registering. Please try again later."
        });
    }
}