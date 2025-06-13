<script lang="ts">
    import AccountSelector from "$lib/components/AccountSelector.svelte";

    let { authenticated, tokens }: { authenticated: boolean, tokens: number } = $props();
</script>

<header class="flex flex-row justify-center">
    <div class="flex flex-row justify-between items-center w-full py-2 px-3 md:w-3/4">
        <div class="">
            <span class="text-3xl font-semibold">Compressor</span>
        </div>

        <nav class="hidden text-xl font-medium gap-[2rem] lg:flex">
            <a href="/" class="hover:text-indigo-800">Home</a>
            <a href="/features" class="hover:text-indigo-800">Features</a>
            <a href="/pricing" class="hover:text-indigo-800">Pricing</a>
            <a href="/support" class="hover:text-indigo-800">Support</a>
        </nav>

        {#if !authenticated}
            <AccountSelector tokens={tokens}></AccountSelector>
        {:else}
            <a href="/signup" class="bg-indigo-800 px-5 py-1.5 rounded-lg text-bg font-medium text-lg hover:bg-indigo-900 transition-colors">Sign
                Up</a>
        {/if}
    </div>
</header>