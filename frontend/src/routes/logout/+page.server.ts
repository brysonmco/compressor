import type {PageServerLoad} from "../../../.svelte-kit/types/src/routes/login/$types";
import {redirect} from "@sveltejs/kit";

export const load: PageServerLoad = async ({cookies}) => {
    // I'm lazy, so we are not going to revoke the sessions, just delete the cookie
    cookies.delete('accessToken', {path: '/'});
    cookies.delete('refreshToken', {path: '/'});
    redirect(300, '/');
};