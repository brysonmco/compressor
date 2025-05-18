<script lang="ts">
    let activeSubscription = false;
</script>

<div class="max-w-3xl mx-auto p-6 space-y-10">
    <!-- Profile Information -->
    <section>
        <h2 class="text-4xl font-semibold mb-4">Profile Information</h2>
        <form class="space-y-4">
            <label class="flex flex-col gap-1.5">
                <span class="text-md font-medium">Email</span>
                <input
                        type="email"
                        name="email"
                        placeholder="users@email"
                        class="rounded-lg bg-bg border-2 text-md font-medium border-slate-200 focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0"/>
            </label>
            <label class="flex flex-col gap-1.5">
                <span class="text-md font-medium">First Name</span>
                <input
                        type="text"
                        name="firstName"
                        placeholder="User's First Name"
                        class="rounded-lg bg-bg border-2 text-md font-medium border-slate-200 focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0"/>
            </label>
            <label class="flex flex-col gap-1.5">
                <span class="text-md font-medium">Last Name</span>
                <input
                        type="text"
                        name="lastName"
                        placeholder="User's Last Name"
                        class="rounded-lg bg-bg border-2 text-md font-medium border-slate-200 focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0"/>
            </label>
            <button type="submit"
                    class="bg-brand text-bg text-lg font-medium px-4 py-2 rounded-lg hover:bg-brand-dark hover:cursor-pointer">
                Save Changes
            </button>
        </form>
    </section>

    <!-- Change Password -->
    <section>
        <h2 class="text-4xl font-semibold mb-4">Change Password</h2>
        <form class="space-y-4">
            <label class="flex flex-col gap-1.5">
                <span class="text-md font-medium">Current Password</span>
                <input
                        type="password"
                        name="currentPassword"
                        placeholder="Enter your password..."
                        class="rounded-lg bg-bg border-2 text-md font-medium border-slate-200 focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0"/>
            </label>
            <label class="flex flex-col gap-1.5">
                <span class="text-md font-medium">New Password</span>
                <input
                        type="password"
                        name="newPassword"
                        placeholder="Enter new password..."
                        class="rounded-lg bg-bg border-2 text-md font-medium border-slate-200 focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0"/>
            </label>
            <label class="flex flex-col gap-1.5">
                <span class="text-md font-medium">Current Password</span>
                <input
                        type="password"
                        name="newPasswordConfirm"
                        placeholder="Confirm new password..."
                        class="rounded-lg bg-bg border-2 text-md font-medium border-slate-200 focus:border-brand active:outline-none focus:outline-none  ring-0 focus:ring-0"/>
            </label>
            <button type="submit"
                    class="bg-brand text-bg text-lg font-medium px-4 py-2 rounded-lg hover:bg-brand-dark hover:cursor-pointer">
                Update Password
            </button>
        </form>
    </section>

    <!-- Subscription Management -->
    <section>
        <h2 class="text-4xl font-semibold mb-4">Subscription</h2>
        {#if activeSubscription}
            <div class="border p-4 rounded-lg space-y-2">
                <p class="font-medium">Plan: <span class="text-blue-700">Pro</span></p>
                <p class="text-sm text-gray-600">Next billing date: June 1, 2025</p>
                <div class="flex space-x-4 mt-2">
                    <button class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700">
                        Manage Subscription
                    </button>
                    <button class="bg-gray-200 text-gray-800 px-4 py-2 rounded-lg hover:bg-gray-300">
                        Cancel Plan
                    </button>
                </div>
            </div>
        {:else}
            <div class="border border-blue-200 bg-blue-50 p-4 rounded-lg space-y-2 md:w-2/3">
                <p class="font-medium text-blue-800">You don't have an active subscription.</p>
                <p class="text-sm text-blue-700">Subscribe now to unlock unlimited compression and advanced
                    features.</p>
                <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700">
                    Choose a Plan
                </button>
            </div>
        {/if}
    </section>

    <!-- Danger Zone -->
    <section>
        <h2 class="text-4xl font-semibold mb-4">Danger Zone</h2>
        <div class="border border-red-300 p-4 rounded-lg bg-red-50 space-y-4">
            <p class="text-red-700">
                Deleting your account will permanently remove all your data and cannot be undone.
            </p>
            <button class="bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700">
                Delete Account
            </button>
        </div>
    </section>
</div>
