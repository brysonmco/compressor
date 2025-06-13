<script lang="ts">
    let { tokens }: { tokens: number } = $props();
    let showDropdown = $state(false);

    function toggleDropdown() {
        showDropdown = !showDropdown;
    }


</script>

<div class="relative inline-block overflow-visible">
    <button class="text-lg font-medium px-5 py-1.5 border-2 rounded-lg hover:cursor-pointer hover:bg-slate-100"
            onclick={toggleDropdown}>Account
    </button>

    {#if showDropdown}
        <div class="flex flex-col absolute right-0 mt-2 w-48 bg-white p-2 rounded-md shadow-lg z-50 gap-1">
        <span class="flex py-2 px-1 justify-between">
            <span class="text-md font-medium">Tokens:</span>
            <span class="text-md font-medium">{ tokens }</span>
        </span>
            <hr class="text-slate-200">

            <a href="/dashboard"
               class="flex items-center gap-1 hover:cursor-pointer hover:bg-slate-100 w-full py-2 px-1 rounded-lg">
                <i class='bx bx-home text-xl'></i>
                <span class="text-md font-medium">Dashboard</span>
            </a>
            <a href="/account"
               class="flex items-center gap-1 hover:cursor-pointer hover:bg-slate-100 w-full py-2 px-1 rounded-lg">
                <i class='bx bx-user-circle text-xl'></i>
                <span class="text-md font-medium">Account</span>
            </a>
            <hr class="text-slate-200">

            <a href="/logout"
               class="flex items-center gap-1 hover:cursor-pointer hover:bg-slate-100 w-full py-2 px-1 rounded-lg">
                <i class='bx bx-log-out text-xl'></i>
                <span class="text-md font-medium">Log Out</span>
            </a>
        </div>
    {/if}

</div>
