import {type Actions} from "@sveltejs/kit";
import type {PageServerLoad} from "../../../../.svelte-kit/types/src/routes/login/$types";

export const load: PageServerLoad = async ({cookies}) => {
    // Get # of tokens, past jobs, and subscription tier
};

export const actions = {
    upload: async ({request, cookies}) => {
        let data = await request.formData();
        let file = data.get("file") as File;

        if (!file) {
            return {
                "success": false,
                "message": "File is required.",
                "formErrors": "file",
            }
        }
        return
    }
} satisfies Actions;