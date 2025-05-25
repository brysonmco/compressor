<script lang="ts">
    import AccountSelector from "$lib/components/AccountSelector.svelte";
    import Dropzone from "$lib/components/Dropzone.svelte";

    let tokens = 1000;

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
            <div class="flex flex-row gap-4 justify-between">
                <Dropzone></Dropzone>
                <div class="flex flex-col flex-1 bg-slate-50 border-2 border-slate-200 rounded-lg p-3">
                    <form action="/idk" class="flex flex-col gap-3">
                        <label class="flex flex-col gap-1">
                            <span class="font-medium text-gray-700">Output Container</span>
                            <select name="outputContainer" id="outputContainer" class="w-full rounded border border-gray-300 px-3 py-2 bg-white focus:outline-none focus:ring-1 focus:ring-blue-500">
                                <option value="mp4">MP4</option>
                                <option value="mov">MOV</option>
                                <option value="mp4">MKV</option>
                            </select>
                        </label>

                        <label class="flex flex-col gap-1">
                            <span class="font-medium text-gray-700">Output Codec</span>
                            <select name="outputContainer" id="outputContainer" class="w-full rounded border border-gray-300 px-3 py-2 bg-white focus:outline-none focus:ring-1 focus:ring-blue-500">
                                <option value="h264">H.264</option>
                                <option value="h265">H.265</option>
                                <option value="vp9">VP9</option>
                                <option value="av1">AV1</option>
                            </select>
                        </label>

                        <label class="flex flex-col gap-1">
                            <span class="font-medium text-gray-700">Output Resolution</span>
                            <select name="outputContainer" id="outputContainer" class="w-full rounded border border-gray-300 px-3 py-2 bg-white focus:outline-none focus:ring-1 focus:ring-blue-500">
                                <option value="8k">8K (7680x4320)</option>
                                <option value="4k">4K (3840x2160)</option>
                                <option value="2k">2K (2560x1440)</option>
                                <option value="fhd">Full HD (1920x1080)</option>
                                <option value="hd">HD (1280x720)</option>
                                <option value="480">480p (854x480)</option>
                                <option value="360">360p (640x360)</option>
                                <option value="240">240p (320x240)</option>
                                <option value="144">144p (256x144)</option>
                            </select>
                        </label>
                    </form>

                    <hr class="my-3 text-slate-300 border rounded-full">

                    <span class="text-gray-700"><span class="font-semibold">Estimated File Size:</span> 32MB</span>
                    <span class="text-gray-700"><span class="font-semibold">Estimated Compression Time:</span> 1:30:13</span>

                    <hr class="my-3 text-slate-300 border rounded-full">

                    <span class="text-gray-700 mb-2"><span class="font-semibold">Token Cost:</span> 40</span>

                    <button class="bg-purple-500 text-white rounded-lg py-1.5 text-lg font-medium hover:bg-purple-600 cursor-pointer">Compress</button>
                </div>
            </div>
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
                                    on:click={() => alert(`Details for ${job.id}`)}
                            >
                                Details
                            </button>

                            {#if job.status === 'Completed'}
                                <button
                                        class="rounded bg-blue-600 px-3 py-1 text-sm font-medium text-white hover:bg-blue-700 transition"
                                        type="button"
                                        on:click={() => alert(`Downloading ${job.fileName}`)}
                                >
                                    Download
                                </button>
                            {/if}

                            {#if job.status === 'Failed'}
                                <button
                                        class="rounded bg-red-600 px-3 py-1 text-sm font-medium text-white hover:bg-red-700 transition"
                                        type="button"
                                        on:click={() => alert(`Retrying job ${job.id}`)}
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