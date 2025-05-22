import type { Actions } from './$types';
import {login} from "$lib/server/auth";
import {json, redirect} from "@sveltejs/kit";

export const actions = {
    default: async ({ cookies, request }) => {
        const data = await request.formData();
        const email = data.get('email')!.toString();
        const password = data.get('password')!.toString();

        let req = await login(email, password)

        let res = await req.json();
        if (!res.success) {
            return {
                "success": false,
                "message": res.message,
                "formErrors": res.fieldErrors,
            }
        }

        cookies.set('accessToken', res.accessToken, {
            httpOnly: true,
            path: '/',
        });
        cookies.set('refreshToken', res.refreshToken, {
            httpOnly: true,
            path: '/',
        });

        redirect(303, '/dashboard');
    }
} satisfies Actions;