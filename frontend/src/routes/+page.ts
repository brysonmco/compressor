import type { PageLoad } from './$types';
import {accessToken} from "$lib/stores/auth";


export const load: PageLoad = ({ params }) => {
    // Check if an access token is present
}