<script lang="ts">
    import { signup } from "$lib/api/auth";

    export let form: any;
    let showErrorModal = false;

    $: if (form?.error) {
        showErrorModal = true;
    }

    let formData = {
        email: '',
        firstName: '',
        lastName: '',
        password: '',
        confirmPassword: ''
    };

    let formErrors = {
        email: '',
        firstName: '',
        lastName: '',
        password: '',
        confirmPassword: ''
    };

    function handleSubmit(event: Event) {
        console.log("This has been called!");
        event.preventDefault();
        signup(
            formData.email,
            formData.firstName,
            formData.lastName,
            formData.password,
            formData.confirmPassword
        );
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
            case 'firstName':
                if (!formData.firstName) {
                    formErrors.firstName = 'First name is required';
                } else {
                    formErrors.firstName = '';
                }
                break;
            case 'lastName':
                if (!formData.lastName) {
                    formErrors.lastName = 'Last name is required';
                } else {
                    formErrors.lastName = '';
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
            case 'confirmPassword':
                if (!formData.confirmPassword) {
                    formErrors.confirmPassword = 'Confirm password is required';
                } else if (formData.confirmPassword !== formData.password) {
                    formErrors.confirmPassword = 'Passwords do not match';
                } else {
                    formErrors.confirmPassword = '';
                }
                break;
        }
    }
</script>


<div class="flex justify-center items-center h-screen">
    <form on:submit|preventDefault={signup}
            class="flex flex-col gap-4 w-1/3 p-6 bg-white rounded-lg items-center" >
        <span class="text-4xl font-semibold">Sign Up</span>

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

        <div class="flex flex-row w-full gap-4">
            <label class="flex flex-col gap-1.5 flex-grow">
                <span class="text-lg font-medium">First Name</span>
                <input
                        type="text"
                        name="firstName"
                        placeholder="Enter your first name..."
                        required
                        bind:value={formData.firstName}
                        on:blur={() => validateField('firstName')}
                        class="rounded-lg bg-bg border-2 text-lg font-medium {formErrors.firstName ? 'border-red-500' : 'border-slate-200' } focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0">
                {#if formErrors.firstName}
                    <span class="text-red-500 text-md font-medium">{formErrors.firstName}</span>
                {/if}
            </label>

            <label class="flex flex-col gap-1.5 flex-grow">
                <span class="text-lg font-medium">Last Name</span>
                <input
                        type="text"
                        name="lastName"
                        placeholder="Enter your last name..."
                        required
                        bind:value={formData.lastName}
                        on:blur={() => validateField('lastName')}
                        class="rounded-lg bg-bg border-2 text-lg font-medium {formErrors.lastName ? 'border-red-500' : 'border-slate-200' }  focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0">
                {#if formErrors.lastName}
                    <span class="text-red-500 text-md font-medium">{formErrors.lastName}</span>
                {/if}
            </label>
        </div>


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

        <label class="flex flex-col gap-1.5 w-full">
            <span class="text-lg font-medium">Confirm Password</span>
            <input
                    type="password"
                    name="confirmPassword"
                    placeholder="Confirm your password..."
                    required
                    bind:value={formData.confirmPassword}
                    on:blur={() => validateField('confirmPassword')}
                    class="rounded-lg bg-bg border-2 text-lg font-medium {formErrors.confirmPassword ? 'border-red-500' : 'border-slate-200' }  focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0">
            {#if formErrors.confirmPassword}
                <span class="text-red-500 text-md font-medium">{formErrors.confirmPassword}</span>
            {/if}
        </label>

        <button type="submit"
                class="bg-brand w-full py-2 text-bg text-xl font-medium rounded-lg hover:bg-brand-dark hover:cursor-pointer">
            Sign Up
        </button>
    </form>
</div>

<!--<div class="fixed inset-0 z-50 flex items-center justify-center bg-gray-900/30">
    <div class="bg-white rounded-2xl shadow-lg max-w-md p-6 text-center">
        <h2 class="text-xl font-semibold mb-4">Error</h2>
        <p class="mb-4 text-red-600">hello</p>
        <button
                class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700"
        >
            Close
        </button>
    </div>
</div>-->
