<svelte:head>
    <title>Log in | Compressor</title>
    <meta name="description"
          content="Log in to your Compressor account to start compressing videos and managing your files.">
</svelte:head>

<script lang="ts">
    import {applyAction, enhance} from "$app/forms";
    import type {PageProps} from "./$types";
    import type {SubmitFunction} from "@sveltejs/kit";
    import {goto} from "$app/navigation";

    let {form}: PageProps = $props();
    let loading = $state(false);
    let showModal = $state(false);

    const submit: SubmitFunction = () => {
        loading = true;

        return async ({result}) => {
            loading = false;
            if (result.type === 'redirect') {
                await goto(result.location);
            } else {
                await applyAction(result);
                if (form?.errors.server) {
                    showModal = true;
                }
            }
        }
    };
</script>


<div class="flex flex-row h-screen">
    <div class="flex flex-1 justify-center items-center bg-indigo-800 md:bg-slate-100 max-w-full">
        <form method="POST" use:enhance={submit}
              class="flex flex-col bg-white rounded-xl border p-6 shadow-sm gap-3 2xl:w-1/2 w-full max-w-[80%]">
            <span class="text-4xl font-semibold mb-2">Login</span>
            <label class="flex flex-col gap-2">
                <span class="text-2xl font-medium">Email</span>
                <input type="email" name="email"
                       class={`bg-slate-50 rounded-xl text-xl border-2 focus:ring-2 focus:border-slate-100 ring-indigo-800 ${form?.errors.email ? 'border-red-600' : 'border-slate-100'}`}
                       placeholder="Enter your email">
                {#if form?.errors.email}
                    <span class="text-red-600 font-medium">{form?.errors.email}</span>
                {/if}
            </label>
            <label class="flex flex-col gap-2 mb-2">
                <span class="text-2xl font-medium">Password</span>
                <input type="password" name="password"
                       class={`bg-slate-50 rounded-xl text-xl border-2 focus:ring-2 focus:border-slate-100 ring-indigo-800 ${form?.errors.password ? 'border-red-600' : 'border-slate-100'}`}
                       placeholder="Enter your password">
                {#if form?.errors.password}
                    <span class="text-red-600 font-medium">{form?.errors.password}</span>
                {/if}
            </label>

            {#if !loading}
                <button class="bg-indigo-800 hover:bg-indigo-900 hover:cursor-pointer text-white py-2 rounded-xl text-xl font-medium">
                    Login
                </button>
            {:else }
                <button disabled
                        class="flex flex-row items-center justify-center gap-2 bg-indigo-800 hover:bg-indigo-900 hover:cursor-not-allowed text-white py-2 rounded-xl text-xl font-medium">
                    <svg width="24" height="24" fill="#fff" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                        <style>.spinner_z9k8 {
                            transform-origin: center;
                            animation: spinner_StKS .75s infinite linear
                        }

                        @keyframes spinner_StKS {
                            100% {
                                transform: rotate(360deg)
                            }
                        }
                        </style>
                        <path d="M12,1A11,11,0,1,0,23,12,11,11,0,0,0,12,1Zm0,19a8,8,0,1,1,8-8A8,8,0,0,1,12,20Z"
                              opacity=".25"/>
                        <path d="M12,4a8,8,0,0,1,7.89,6.7A1.53,1.53,0,0,0,21.38,12h0a1.5,1.5,0,0,0,1.48-1.75,11,11,0,0,0-21.72,0A1.5,1.5,0,0,0,2.62,12h0a1.53,1.53,0,0,0,1.49-1.3A8,8,0,0,1,12,4Z"
                              class="spinner_z9k8"/>
                    </svg>
                    Logging in
                </button>
            {/if}


            <span>Don't have an account? <a href="/signup" class="text-blue-500 hover:underline font-medium">Sign Up</a></span>
            <a href="/forgot-password" class="text-blue-500 hover:underline font-medium">Forgot Password?</a>
        </form>
    </div>
    <div class="bg-indigo-800 h-full flex-1 hidden md:flex"></div>
</div>

{#if showModal}
    <div class="fixed top-0 left-0 w-full h-full flex items-center justify-center bg-slate-800/30 z-50">
        <div class="relative flex flex-col items-center justify-center bg-white p-6 rounded-xl shadow-sm w-1/3 h-1/3 gap-2">
            <button class="absolute top-2 right-2" onclick={() => showModal = false} aria-label="close modal">
                <i class="bx bx-x bx-md hover:bg-slate-100 hover:cursor-pointer rounded-full"></i>
            </button>
            <i class="bx bx-error-circle bx-lg text-red-500"></i>
            <span class="text-2xl text-slate-700 font-medium">There was an error processing your request.</span>
            <span class="text-md text-slate-600">Please try again later, if the issue persists contact support</span>

        </div>
    </div>
{/if}