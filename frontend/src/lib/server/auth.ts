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
            credentials: "include",
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
            switch (data.error) {
                case "missing_fields":
                    return json({
                        success: false,
                        fieldErrors: data.details,
                    });
                case "invalid_credentials":
                    return json({
                        success: false,
                        fieldErrors: {
                            email: "Incorrect email or password.",
                            password: "Incorrect email or password.",
                        }
                    });
                default:
                    return json({
                        success: false,
                        message: "Error occurred while logging in. Please try again later."
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
            message: "Error occurred while logging in. Please try again later."
        });
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
            credentials: "include",
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

export async function isAuthenticated(cookies: Cookies): Promise<boolean> {
    let accessToken = cookies.get('accessToken');
    let refreshToken = cookies.get('refreshToken');

    if (!accessToken) {
        if (!refreshToken) {
            cookies.delete('accessToken', {path: '/'});
            cookies.delete('refreshToken', {path: '/'});
            return false;
        }

        const refreshResponse = await refresh(refreshToken);
        const refreshData = await refreshResponse.json();

        if (!refreshData.success || !refreshData.accessToken) {
            cookies.delete('accessToken', {path: '/'});
            cookies.delete('refreshToken', {path: '/'});
            return false;
        }

        cookies.set('accessToken', refreshData.accessToken, {
            httpOnly: true,
            path: '/',
            sameSite: 'strict',
            maxAge: 60 * 60
        });

        accessToken = refreshData.accessToken;
    }

    // Check if the user is authenticated
    try {
        const profileResponse = await getProfile(accessToken!);
        const authData = await profileResponse.json();

        if (!profileResponse.ok || !authData.success) {
            cookies.delete('accessToken', {path: '/'});
            cookies.delete('refreshToken', {path: '/'});
            return false;
        }
    } catch (err) {
        cookies.delete('accessToken', {path: '/'});
        cookies.delete('refreshToken', {path: '/'});
        return false;
    }

    return true;
}