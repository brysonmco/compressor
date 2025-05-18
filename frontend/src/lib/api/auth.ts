import {apiBaseUrl} from "$lib/config";
import {accessToken} from "$lib/stores/auth";
import {json} from "@sveltejs/kit";

// TODO: I'm not sure if I want the refreshAccessToken function to redirect the user, I'm still not sure about how I want to handle refreshing.
export async function refreshAccessToken(): Promise<Response> {
    // Clear the access token
    accessToken.set(null)

    try {
        const response = await fetch(apiBaseUrl + "/auth/refresh", {
            method: "POST",
            credentials: "include",
        });

        const data = await response.json();
        if (!response.ok) {
            // TODO: This switch is a little redundant at the moment, but I intend to change how this is handled in the future.
            switch (data.error) {
                case "invalid_token":
                    return json({
                        success: false,
                        redirect: '/login',
                        statusCode: 303,
                    });
                case "expired_token":
                    return json({
                        success: false,
                        redirect: '/login',
                        statusCode: 303,
                    });
                case "revoked_token":
                    return json({
                        success: false,
                        redirect: '/login',
                        statusCode: 303,
                    });
                default:
                    return json({
                        success: false,
                        redirect: '/login',
                        statusCode: 303,
                    });
            }
        }

        accessToken.set(data.accessToken);
        return json({
            success: true,
        });
    } catch (e) {
        return json({
            success: false,
            redirect: '/login',
            statusCode: 303,
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

        accessToken.set(data.accessToken);
        return json({
            success: true,
            redirect: '/',
            statusCode: 303,
        });
    } catch (err) {
        return json({
            success: false,
            message: "Error occurred while logging in. Please try again later."
        });
    }
}

export async function signup(
    email: string,
    firstName: string,
    lastName: string,
    password: string,
    confirmPassword: string
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
                firstName,
                lastName,
                password,
                confirmPassword,
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
                case "password_mismatch":
                    return json({
                        success: false,
                        fieldErrors: {
                            password: "Passwords do not match",
                            confirmPassword: "Passwords do not match",
                        }
                    });
                case "account_exists":
                    return json({
                        success: false,
                        redirect: '/login?email=' + email,
                        statusCode: 303,
                    });
                default:
                    return json({
                        success: false,
                        message: "Error occurred while signing up. Please try again later."
                    });
            }
        }

        accessToken.set(data.accessToken);
        return json({
            success: true,
            redirect: '/',
            statusCode: 303,
        });
    } catch (err) {
        return json({
            success: false,
            message: "Error occurred while signing up. Please try again later."
        });
    }
}