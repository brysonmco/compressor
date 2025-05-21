<script>
    import Header from "$lib/components/Header.svelte";
    let authenticated = false;
</script>

<Header authenticated={authenticated}></Header>

<div class="flex flex-col items-center h-full mt-8">
    <span class="text-3xl font-medium leading-normal">Pricing</span>
    <span class="text-lg mb-6">Choose a plan that fits your needs</span>

    <div class="flex flex-col 2xl:flex-row w-full px-4 sm:px-0 sm:w-3/4 md:w-1/2 2xl:w-3/4 justify-between gap-6">
        <div class="flex flex-col flex-grow bg-white shadow-md rounded-2xl p-6 gap-4 w-full mx-auto">
            <div class="text-center">
                <h2 class="text-2xl font-bold text-gray-800">Free</h2>
                <p class="text-4xl font-extrabold text-purple-600 mt-1">$0<span class="text-base font-medium text-gray-500">/forever</span></p>
                <p class="text-lg font-medium text-gray-600 mt-1">Explore Compressor</p>
            </div>

            <a href="/account"
               class="w-full text-center bg-purple-600 rounded-lg text-white py-2 text-lg font-semibold hover:bg-purple-700 transition">
                Sign Up
            </a>

            <hr class="border-t border-slate-300 my-2" />

            <ul class="flex flex-col gap-2 text-gray-700 text-base font-medium">
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 100 Tokens / month
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Standard Job Priority
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 1 Concurrent Job
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Full HD Max Resolution
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 100MB Max File Size
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Codecs: H.264
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 1 Hour File Retention
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Watermark on Output
                </li>
            </ul>
        </div>

        <div class="flex flex-col flex-grow bg-white shadow-md rounded-2xl p-6 gap-4 w-full mx-auto">
            <div class="text-center">
                <h2 class="text-2xl font-bold text-gray-800">Basic</h2>
                <p class="text-4xl font-extrabold text-purple-600 mt-1">$4.99<span class="text-base font-medium text-gray-500">/month</span></p>
                <p class="text-lg font-medium text-gray-600 mt-1">Essential features</p>
            </div>

            <a href="/account"
               class="w-full text-center bg-purple-600 rounded-lg text-white py-2 text-lg font-semibold hover:bg-purple-700 transition">
                Subscribe
            </a>

            <hr class="border-t border-slate-300 my-2" />

            <ul class="flex flex-col gap-2 text-gray-700 text-base font-medium">
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 1000 Tokens / month
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Standard Job Priority
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 5 Concurrent Jobs
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Full HD Max Resolution
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 1GB Max File Size
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Codecs: H.264, H.265, VP9
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 24 Hour File Retention
                </li>
            </ul>
        </div>
        <div class="flex flex-col flex-grow bg-white shadow-md rounded-2xl p-6 gap-4 w-full mx-auto">
            <div class="text-center">
                <h2 class="text-2xl font-bold text-gray-800">Pro</h2>
                <p class="text-4xl font-extrabold text-purple-600 mt-1">$9.99<span class="text-base font-medium text-gray-500">/month</span></p>
                <p class="text-lg font-medium text-gray-600 mt-1">Creators and media experts</p>
            </div>

            <a href="/account"
               class="w-full text-center bg-purple-600 rounded-lg text-white py-2 text-lg font-semibold hover:bg-purple-700 transition">
                Subscribe
            </a>

            <hr class="border-t border-slate-300 my-2" />

            <ul class="flex flex-col gap-2 text-gray-700 text-base font-medium">
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Unlimited tokens*
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Express Job Priority
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Unlimited Concurrent Jobs
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 4K Max Resolution
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 10GB Max File Size
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Codecs: H.264, H.265, VP9, AV1, HEVC
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 48 Hour File Retention
                </li>
            </ul>
        </div>
        <div class="flex flex-col flex-grow bg-white shadow-md rounded-2xl p-6 gap-4 w-full mx-auto">
            <div class="text-center">
                <h2 class="text-2xl font-bold text-gray-800">Ultimate</h2>
                <p class="text-4xl font-extrabold text-purple-600 mt-1">$19.99<span class="text-base font-medium text-gray-500">/month</span></p>
                <p class="text-lg font-medium text-gray-600 mt-1">Teams handling large workloads</p>
            </div>

            <a href="/account"
               class="w-full text-center bg-purple-600 rounded-lg text-white py-2 text-lg font-semibold hover:bg-purple-700 transition">
                Subscribe
            </a>

            <hr class="border-t border-slate-300 my-2" />

            <ul class="flex flex-col gap-2 text-gray-700 text-base font-medium">
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Unlimited tokens*
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Express Job Priority
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Unlimited Concurrent Jobs
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 8K Max Resolution
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 100GB Max File Size
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> Codecs: H.264, H.265, VP9, AV1, HEVC
                </li>
                <li class="flex items-center">
                    <i class='bx bx-check text-purple-600 mr-2 text-xl'></i> 7 Day File Retention
                </li>
            </ul>
        </div>
    </div>
</div>
