<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import type {PageProps} from "./$types";

    let {data}: PageProps = $props();
    console.log(data)
</script>

<Header authenticated={data.authenticated} tokens={data.user?.tokens ?? -1}></Header>
