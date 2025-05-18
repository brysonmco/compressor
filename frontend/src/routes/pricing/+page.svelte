<script>
    import Header from "$lib/components/Header.svelte";
    let authenticated = false;
</script>

<Header authenticated={authenticated}></Header>

<div class="flex flex-col items-center h-full mt-8">
    <span class="text-3xl font-medium leading-normal">Pricing</span>
    <span class="text-lg">Get started with a plan that fits your needs.</span>

    <div class="flex flex-row">
        <div class="">
            <span>$4.99/month</span>
        </div>
        <div class="">
            <span>$9.99/month</span>
        </div>
    </div>
</div>
