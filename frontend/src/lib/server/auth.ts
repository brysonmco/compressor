import { json } from "@sveltejs/kit";
import { apiBaseUrl } from "$lib/server/config";

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

export async function isAuthenticated(
    accessToken: string
): Promise<Response> {
    try {
        const response = await fetch(apiBaseUrl + "/users/profile", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${accessToken}`
            }
        });

        if (!response.ok) {
            return json({
                success: false,
                message: "Authentication failed. Please log in again."
            });
        }

        return json({
            success: true,
        });
    } catch (e) {
        return json({
            success: false,
            message: "Error occurred while checking authentication. Please try again later."
        });
    }
}