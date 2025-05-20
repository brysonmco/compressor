<script>
    import AccountSelector from "$lib/components/AccountSelector.svelte";
    let tokens = 1000;
</script>

<div class="flex flex-col items-center h-screen">
    <div class="w-1/2 h-full">
        <div class="flex flex-row justify-between py-4 items-center">
            <span class="text-4xl font-medium">Compressor</span>
            <AccountSelector tokens={tokens}></AccountSelector>
        </div>
        <hr class="border rounded-full mb-4">

        <span class="text-2xl font-medium">Compress</span>
        <div class="flex flex-row w-full h-1/3">
            <div class="flex justify-center items-center border-dashed rounded-l-lg border-l-2 border-t-2 border-b-2 w-1/2">
                Hello
            </div>
            <div class="border-2 w-1/2 rounded-r-lg">
                Hello
            </div>
        </div>

        <hr class="border rounded-full mt-6 mb-4">
        <span class="text-2xl font-medium">Past Jobs</span>
        <table class="table-auto border-2 border-collapse border-gray-200 bg-gray-100 rounded-lg w-full">
            <thead>
                <tr>
                    <th class="border-2 border-gray-200">Job ID</th>
                    <th class="border-2 border-gray-200">File Name</th>
                    <th class="border-2 border-gray-200">Status</th>
                    <th class="border-2 border-gray-200">Date</th>
                    <th class="border-2 border-gray-200">Cost</th>
                    <th class="border-2 border-gray-200">Action</th>
                </tr>
            </thead>
            <tbody>
            <tr>
                <td class="border-2 border-gray-200">123456</td>
                <td class="border-2 border-gray-200">example.txt</td>
                <td class="border-2 border-gray-200">Completed</td>
                <td class="border-2 border-gray-200">2023-10-01</td>
                <td class="border-2 border-gray-200">$0.10</td>
                <td class="border-2 border-gray-200">
                    <button class="bg-blue-500 text-white px-4 py-2 rounded-lg">Download</button>
                </td>
            </tr>
            </tbody>
        </table>
    </div>
</div>