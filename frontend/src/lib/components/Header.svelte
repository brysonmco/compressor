<script lang="ts">
    import AccountSelector from "$lib/components/AccountSelector.svelte";

    let { authenticated }: { authenticated: boolean } = $props();

    const tokens = 1000;
</script>

<header class="flex flex-row justify-center">
    <div class="flex flex-row justify-between items-center w-full py-2 px-3 md:w-3/4">
        <div class="">
            <span class="text-3xl font-semibold">Compressor</span>
        </div>

        <nav class="hidden text-xl font-medium gap-[2rem] lg:flex">
            <a href="/" class="hover:text-[#0077cc]">Home</a>
            <a href="/features" class="hover:text-[#0077cc]">Features</a>
            <a href="/pricing" class="hover:text-[#0077cc]">Pricing</a>
            <a href="/support" class="hover:text-[#0077cc]">Support</a>
        </nav>

        {#if authenticated}
            <AccountSelector tokens={tokens}></AccountSelector>
        {:else}
            <a href="/signup" class="bg-brand px-6 py-2 rounded-lg text-bg font-medium text-xl hover:bg-brand-dark">Sign
                Up</a>
        {/if}
    </div>
</header>