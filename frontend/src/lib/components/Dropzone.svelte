<script lang="ts">
    let fileInput: HTMLInputElement;
    let form: HTMLFormElement;
    let dragActive = false;

    let file: File | null = null;

    let inputMetadata = {
        name: '',
        container: '',
        codec: '',
        resolution: '',
        sizeMB: '',
        duration: ''
    };

    function handleFileInput(selected: FileList | null) {
        if (!selected || selected.length === 0) return;
        file = selected[0];

        form.submit();

        inputMetadata = {
            name: file.name,
            container: 'mp4',
            codec: 'h.264',
            resolution: '1920x1080',
            sizeMB: (file.size / 1024 / 1024).toFixed(1),
            duration: '02:03'
        };
    }

    function handleDrop(event: DragEvent) {
        dragActive = false;
        if (event.dataTransfer?.files?.length) {
            handleFileInput(event.dataTransfer.files);
        }
    }

    function removeFile() {
        file = null;
    }
</script>

<form method="POST" enctype="multipart/form-data" class="hidden" action="?/upload" bind:this={form}>
    <input type="file" class="hidden" name="file" accept="video/*" bind:this={fileInput} on:change={() => handleFileInput(fileInput.files)}>
</form>

{#if file}
    <div class={`flex flex-col min-h-96 border-2 hover:border-red-500 hover:bg-red-50
        rounded-lg p-3 transition-colors cursor-pointer flex-1 ${dragActive ? 'border-blue-500 bg-blue-50' : 'border-slate-200 bg-slate-50'}`}
         on:click={removeFile}>
        <div class="flex flex-row gap-2 mb-2">
            <i class="bx bx-file bx-lg"></i>
            <div class="flex flex-col flex-grow">
                <span class="text-xl font-medium">{ inputMetadata.name }</span>
                <span>{ inputMetadata.sizeMB } MB • 02:03</span>
            </div>
        </div>
        <span><span class="font-semibold">Resolution:</span> 1920x1080</span>
        <span><span class="font-semibold">Container:</span> MP4</span>
        <span><span class="font-semibold">Codec:</span> H.264</span>

    </div>
{:else }
    <div
            class={`flex flex-col items-center justify-center min-h-96 border-2 hover:border-blue-500 hover:bg-blue-50
        rounded-lg p-3 transition-colors cursor-pointer flex-1 ${dragActive ? 'border-blue-500 bg-blue-50' : 'border-dashed border-slate-200 bg-slate-50'}`}
            on:click={() => fileInput.click()}
            on:drop|preventDefault={handleDrop}
            on:dragover|preventDefault={() => dragActive = true}
            on:dragleave={() => dragActive = false}>
        <i class="bx bx-cloud-upload text-4xl text-gray-600"></i>
        <span class="text-lg font-medium text-gray-600">Drop file or click to upload</span>
    </div>
{/if}
