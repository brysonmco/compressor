<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import type { PageProps } from "./$types";

    let { data }: PageProps = $props();
</script>

<Header authenticated={data.authenticated}></Header>
