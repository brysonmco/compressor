<script lang="ts">
    import { login } from "$lib/api/auth";

    let showErrorModal = false;
    let errorModalMessage = '';


    let formData = {
        email: '',
        password: '',
    };

    let formErrors = {
        email: '',
        password: '',
    };

    async function handleSubmit(event: Event | SubmitEvent) {
        event.preventDefault();
        const res = await login(
            formData.email,
            formData.password,
        );

        const data = await res.json();

        if (data.fieldErrors) {
            formErrors = data.fieldErrors;
        }

        if (!data.success && data.message) {
            showErrorModal = true;
            errorModalMessage = data.message;
        }

        if (data.redirect) {
            window.location.href = data.redirect;
        }
    }

    function validateField(field: string) {
        switch (field) {
            case 'email':
                if (!formData.email) {
                    formErrors.email = 'Email is required';
                } else if (formData.email.includes(' ')) {
                    formErrors.email = 'Email is invalid';
                } else {
                    formErrors.email = '';
                }
                break;
            case 'password':
                if (!formData.password) {
                    formErrors.password = 'Password is required';
                } else if (formData.password.length < 8) {
                    formErrors.password = 'Password must be at least 8 characters long';
                } else {
                    formErrors.password = '';
                }
                break;
        }
    }
</script>


<div class="flex justify-center items-center h-screen">
    <form on:submit|preventDefault={handleSubmit}
          class="flex flex-col gap-4 2xl:w-1/3 md:w-2/3 md:m-0 m-4 p-6 bg-white rounded-lg items-center">
        <span class="text-4xl font-semibold">Login</span>

        <label class="flex flex-col gap-1.5 w-full">
            <span class="text-lg font-medium">Email</span>
            <input
                    type="email"
                    name="email"
                    placeholder="Enter your email..."
                    required
                    bind:value={formData.email}
                    on:blur={() => validateField('email')}
                    class="rounded-lg bg-bg border-2 text-lg font-medium {formErrors.email ? 'border-red-500' : 'border-slate-200' } focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0">
            {#if formErrors.email}
                <span class="text-red-500 text-md font-medium">{formErrors.email}</span>
            {/if}
        </label>

        <label class="flex flex-col gap-1 w-full">
            <span class="text-lg font-medium">Password</span>
            <input
                    type="password"
                    name="password"
                    placeholder="Enter a password..."
                    required
                    bind:value={formData.password}
                    on:blur={() => validateField('password')}
                    class="rounded-lg bg-bg border-2 text-lg font-medium {formErrors.password ? 'border-red-500' : 'border-slate-200' }  focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0">
            {#if formErrors.password}
                <span class="text-red-500 text-md font-medium">{formErrors.password}</span>
            {/if}
        </label>

        <button type="submit"
                class="bg-brand w-full py-2 text-bg text-xl font-medium rounded-lg hover:bg-brand-dark hover:cursor-pointer">
            Login
        </button>
        <span class="text-lg font-medium">Don't have an account? <a href="/signup" class="text-brand hover:underline">Sign Up</a></span>
    </form>
</div>

{#if showErrorModal}
    <div class="fixed inset-0 z-50 flex items-center justify-center bg-slate-900/30">
        <div class="flex flex-col bg-white rounded-lg p-6 text-center gap-4">
            <h2 class="text-3xl font-medium">Error on Login</h2>
            <p class="text-lg text-red-600">{ errorModalMessage }</p>
            <button
                    class="bg-brand w-full py-2 text-bg text-lg rounded-lg hover:bg-brand-dark hover:cursor-pointer"
                    on:click={() => (showErrorModal = false)}>
                Close
            </button>
        </div>
    </div>
{/if}