<script lang="ts">
    import AccountSelector from "$lib/components/AccountSelector.svelte";
    import {enhance} from "$app/forms";

    let tokens = 1000;
    let showCompressModal = $state(false);
    let fileInput: HTMLInputElement;

    // Sample data for past jobs
    let pastJobs = [
        {
            id: "cmp-2a9f",
            fileName: "promo_video.mp4",
            status: "Completed",
            sizeBefore: "120 MB",
            sizeAfter: "32 MB",
            tokensUsed: 94,
            date: "May 25, 2025 · 3:47 PM",
        },
        {
            id: "cmp-3b1x",
            fileName: "holiday_clip.mov",
            status: "Processing",
            sizeBefore: "350 MB",
            sizeAfter: "-",
            tokensUsed: "-",
            date: "May 25, 2025 · 4:01 PM",
        },
        {
            id: "cmp-9z7k",
            fileName: "event_recording.mp4",
            status: "Failed",
            sizeBefore: "200 MB",
            sizeAfter: "-",
            ratio: "-",
            tokensUsed: "-",
            date: "May 24, 2025 · 9:12 AM",
        },
    ];

    // Function to get badge colors based on status
    function statusClasses(status: string): string {
        switch (status) {
            case "Completed":
                return "bg-green-100 text-green-800";
            case "Processing":
                return "bg-yellow-100 text-yellow-800";
            case "Queued":
                return "bg-blue-100 text-blue-800";
            case "Failed":
                return "bg-red-100 text-red-800";
            default:
                return "bg-gray-100 text-gray-800";
        }
    }
</script>

<div class="flex flex-col items-center h-screen">
    <div class="w-1/2 h-full">
        <div class="flex flex-row justify-between py-4 items-center">
            <span class="text-4xl font-medium">Compressor</span>
            <AccountSelector tokens={tokens}></AccountSelector>
        </div>
        <hr class="border rounded-full mb-4">

        <div class="rounded-xl border bg-white p-6 shadow-sm max-w-full overflow-x-auto">
            <h2 class="text-2xl font-semibold mb-6">Compress</h2>
            <div class="flex flex-col items-center justify-center min-h-96 border-2 hover:border-blue-500 hover:bg-blue-50
        rounded-lg p-3 transition-colors cursor-pointer flex-1 border-dashed border-slate-200 bg-slate-50" onclick={() => fileInput.click()}>
                <i class="bx bx-cloud-upload text-4xl text-gray-600"></i>
                <span class="text-lg font-medium text-gray-600">Drop file or click to upload</span>
            </div>
            <form method="POST" enctype="multipart/form-data" class="hidden" use:enhance>
                <input type="file" name="file" accept="video/*" class="hidden" bind:this={fileInput}>
            </form>
        </div>

        <hr class="border rounded-full mt-6 mb-4">

        <div class="rounded-xl border bg-white p-6 shadow-sm max-w-full overflow-x-auto">
            <h2 class="text-2xl font-semibold mb-6">Past Jobs</h2>

            <table class="min-w-full border-collapse text-left text-sm">
                <thead>
                <tr class="border-b border-gray-300">
                    <th class="py-3 px-4 font-medium text-gray-700">Job ID</th>
                    <th class="py-3 px-4 font-medium text-gray-700">File Name</th>
                    <th class="py-3 px-4 font-medium text-gray-700">Status</th>
                    <th class="py-3 px-4 font-medium text-gray-700">Size Before</th>
                    <th class="py-3 px-4 font-medium text-gray-700">Size After</th>
                    <th class="py-3 px-4 font-medium text-gray-700">Tokens</th>
                    <th class="py-3 px-4 font-medium text-gray-700">Date</th>
                    <th class="py-3 px-4 font-medium text-gray-700">Actions</th>
                </tr>
                </thead>

                <tbody>
                {#each pastJobs as job (job.id)}
                    <tr class="border-b border-gray-200 hover:bg-gray-50">
                        <td class="py-3 px-4 font-mono text-gray-800">{job.id}</td>
                        <td class="py-3 px-4 text-gray-900">{job.fileName}</td>
                        <td class="py-3 px-4">
            <span
                    class="inline-block rounded-full px-3 py-1 text-xs font-semibold"
                    class:bg-green-100={job.status === 'Completed'}
                    class:text-green-800={job.status === 'Completed'}
                    class:bg-yellow-100={job.status === 'Processing'}
                    class:text-yellow-800={job.status === 'Processing'}
                    class:bg-blue-100={job.status === 'Queued'}
                    class:text-blue-800={job.status === 'Queued'}
                    class:bg-red-100={job.status === 'Failed'}
                    class:text-red-800={job.status === 'Failed'}
            >
              {job.status}
            </span>
                        </td>
                        <td class="py-3 px-4 text-gray-800">{job.sizeBefore}</td>
                        <td class="py-3 px-4 text-gray-800">{job.sizeAfter}</td>
                        <td class="py-3 px-4 text-gray-800">{job.tokensUsed}</td>
                        <td class="py-3 px-4 text-gray-800">{job.date}</td>
                        <td class="py-3 px-4 space-x-2">
                            <button
                                    class="rounded bg-gray-200 px-3 py-1 text-sm font-medium text-gray-700 hover:bg-gray-300 transition"
                                    type="button"
                                    onclick={() => alert(`Details for ${job.id}`)}
                            >
                                Details
                            </button>

                            {#if job.status === 'Completed'}
                                <button
                                        class="rounded bg-blue-600 px-3 py-1 text-sm font-medium text-white hover:bg-blue-700 transition"
                                        type="button"
                                        onclick={() => alert(`Downloading ${job.fileName}`)}
                                >
                                    Download
                                </button>
                            {/if}

                            {#if job.status === 'Failed'}
                                <button
                                        class="rounded bg-red-600 px-3 py-1 text-sm font-medium text-white hover:bg-red-700 transition"
                                        type="button"
                                        onclick={() => alert(`Retrying job ${job.id}`)}
                                >
                                    Retry
                                </button>
                            {/if}
                        </td>
                    </tr>
                {/each}
                </tbody>
            </table>
        </div>
    </div>
</div>

{#if showCompressModal}
    <div class="fixed top-0 left-0 w-full h-full flex items-center justify-center bg-slate-800/30 z-50">
        <div class="relative flex flex-col bg-white p-6 rounded-xl shadow-sm w-1/3 gap-2">
            <button class="absolute top-6 right-6" onclick={() => showCompressModal = false} aria-label="close modal">
                <i class="bx bx-x bx-md text-slate-700 hover:bg-slate-100 hover:cursor-pointer rounded-full"></i>
            </button>

            <span class="mt-8 text-3xl text-slate-700 font-semibold">Compress</span>

            <hr class="border-slate-600 border rounded-full w-full my-2">

            <div class="relative flex flex-row gap-2 bg-slate-100 rounded-xl p-3">
                <img src="/thumbnail.jpg" alt="thumbnail" class="w-48 rounded-sm">
                <div class="flex flex-col">
                    <span class="text-xl text-slate-700 font-medium">Jim's Pie Prank.mp4</span>
                    <span class="text-md text-slate-600">120 MB • 02:03</span>
                </div>
                <button class="absolute top-2 right-2" aria-label="close modal">
                    <i class="bx bx-x bx-sm text-slate-700 hover:bg-slate-200 hover:cursor-pointer rounded-full"></i>
                </button>
            </div>

            <hr class="border-slate-600 border rounded-full w-full my-2">

            <div class="flex flex-col gap-2">
                <label class="flex flex-col gap-1">
                    <span class="text-lg text-slate-700 font-medium">Compression Level</span>
                    <select name="compressionLevel"
                            class="bg-slate-50 rounded-xl text-lg border-2 focus:ring-2 focus:border-slate-100 ring-indigo-800">
                        <option value="low">Low (Fastest)</option>
                        <option value="medium" selected>Medium (Balanced)</option>
                        <option value="high">High (Best Quality)</option>
                    </select>
                </label>

                <label class="flex flex-col gap-1">
                    <span class="text-lg text-slate-700 font-medium">Output Format</span>
                    <select name="outputFormat"
                            class="bg-slate-50 rounded-xl text-lg border-2 focus:ring-2 focus:border-slate-100 ring-indigo-800">
                        <option value="mp4" selected>MP4</option>
                        <option value="webm">WebM</option>
                        <option value="avi">AVI</option>
                    </select>
                </label>

                <button class="bg-indigo-800 hover:bg-indigo-900 text-white py-2 rounded-xl text-lg font-medium mt-4">
                    Start Compression
                </button>
            </div>

        </div>
    </div>
{/if}