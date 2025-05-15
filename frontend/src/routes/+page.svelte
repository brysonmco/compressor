<script>
    import Header from "$lib/components/Header.svelte";
    let authenticated = false;
    let avatarUrl = "";
</script>

<Header authenticated={authenticated} avatarUrl={avatarUrl}></Header>
