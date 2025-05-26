import type { Actions } from "@sveltejs/kit";

export const actions = {
    upload: async ({ request, cookies }) => {
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