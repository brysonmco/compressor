<script>
    import AccountSelector from "$lib/components/AccountSelector.svelte";

    let tokens = -1;
</script>

<div class="flex flex-row p-4 h-screen gap-4">
    <div class="flex flex-col p-4 bg-white border rounded-xl h-full min-w-[240px] max-w-xs shadow-sm">
        <!-- Product Name -->
        <div class="mb-6 px-2">
            <h1 class="text-2xl font-bold text-slate-800">Compressor</h1>
        </div>

        <!-- Navigation -->
        <nav class="flex flex-col gap-1 text-base">
            <a href="/admin" class="flex items-center gap-2 px-3 py-2 rounded-lg font-medium text-slate-900 bg-slate-100">
                <i class='bx bx-home-alt'></i>
                Dashboard
            </a>
            <a href="/admin/users" class="flex items-center gap-2 px-3 py-2 rounded-lg font-medium text-slate-700 hover:bg-slate-50">
                <i class='bx bx-user'></i>
                Users
            </a>
            <a href="/admin/subscriptions" class="flex items-center gap-2 px-3 py-2 rounded-lg font-medium text-slate-700 hover:bg-slate-50">
                <i class='bx bx-receipt'></i>
                Subscriptions
            </a>
            <a href="/admin/revenue" class="flex items-center gap-2 px-3 py-2 rounded-lg font-medium text-slate-700 hover:bg-slate-50">
                <i class='bx bx-dollar-circle'></i>
                Revenue
            </a>
            <a href="/admin/storage" class="flex items-center gap-2 px-3 py-2 rounded-lg font-medium text-slate-700 hover:bg-slate-50">
                <i class='bx bx-hdd'></i>
                Storage
            </a>
            <a href="/admin/settings" class="flex items-center gap-2 px-3 py-2 rounded-lg font-medium text-slate-700 hover:bg-slate-50">
                <i class='bx bx-cog'></i>
                Settings
            </a>
        </nav>

        <!-- Footer Info -->
        <div class="mt-auto text-xs text-slate-400 px-2 pt-6">
            v1.0 · © 2025 Compressor
        </div>
    </div>
    <slot/>
</div>