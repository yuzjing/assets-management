<!-- frontend/src/routes/+page.svelte (最终美学优化完整版) -->
<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    // 1. 导入我们需要的图标
    import { Search, RotateCcw, SlidersHorizontal, ChevronsUpDown } from 'svelte-lucide-icons';

    export let data;

    // 筛选和排序的状态
    let filters = {
        gdzc_lb: data.filters.gdzc_lb || '',
        gdzc_userbm: data.filters.gdzc_userbm || '',
        gdzc_zt: data.filters.gdzc_zt || '',
        gdzc_sccj: data.filters.gdzc_sccj || '',
        gdzc_je_gte: data.filters.gdzc_je_gte || '',
        gdzc_je_lte: data.filters.gdzc_je_lte || '',
        gdzc_rzdate_gte: data.filters.gdzc_rzdate_gte || '',
        gdzc_rzdate_lte: data.filters.gdzc_rzdate_lte || '',
    };
    let sortBy = data.filters.sortBy || 'gdzc_id';
    let order = data.filters.order || 'asc';
    let isFilterPanelOpen = false; // 控制筛选面板的显示和隐藏

    // 应用筛选与排序
    function applyFiltersAndSort() {
        const params = new URLSearchParams();
        for (const key in filters) {
            if (filters[key]) params.set(key, filters[key]);
        }
        params.set('sortBy', sortBy);
        params.set('order', order);
        goto(`?${params.toString()}`, { keepFocus: true, noScroll: true, replaceState: true });
    }

    // 清除筛选
    function clearFilters() {
        for (const key in filters) filters[key] = '';
        sortBy = 'gdzc_id';
        order = 'asc';
        goto(`/`, { keepFocus: true, noScroll: true, replaceState: true });
    }
    
    // 点击表头进行排序
    function handleSort(newSortBy) {
        if (sortBy === newSortBy) {
            order = order === 'asc' ? 'desc' : 'asc';
        } else {
            sortBy = newSortBy;
            order = 'asc';
        }
        applyFiltersAndSort();
    }

    // 表格滚动和日期格式化的辅助函数
    let topScrollbarWrapper, mainTableWrapper;
    onMount(() => {
        function syncScrollFromMain() { if (topScrollbarWrapper && mainTableWrapper) { topScrollbarWrapper.scrollLeft = mainTableWrapper.scrollLeft; } }
        function syncScrollFromTop() { if (topScrollbarWrapper && mainTableWrapper) { mainTableWrapper.scrollLeft = topScrollbarWrapper.scrollLeft; } }
        const innerScrollbar = topScrollbarWrapper.querySelector('.top-scrollbar-inner');
        const table = mainTableWrapper.querySelector('table');
        if (innerScrollbar && table) {
            const tableMinWidth = getComputedStyle(table).minWidth;
            innerScrollbar.style.width = tableMinWidth;
        }
        mainTableWrapper.addEventListener('scroll', syncScrollFromMain);
        topScrollbarWrapper.addEventListener('scroll', syncScrollFromTop);
        return () => {
            mainTableWrapper.removeEventListener('scroll', syncScrollFromMain);
            topScrollbarWrapper.removeEventListener('scroll', syncScrollFromTop);
        };
    });
    function dragToScroll(node) {
        let isDown = false, startX, scrollLeft;
        function handleMouseDown(e) { isDown = true; node.classList.add('grabbing'); startX = e.pageX - node.offsetLeft; scrollLeft = node.scrollLeft; }
        function handleMouseLeave() { isDown = false; node.classList.remove('grabbing'); }
        function handleMouseUp() { isDown = false; node.classList.remove('grabbing'); }
        function handleMouseMove(e) { if (!isDown) return; e.preventDefault(); const x = e.pageX - node.offsetLeft; const walk = (x - startX) * 2; node.scrollLeft = scrollLeft - walk; }
        node.addEventListener('mousedown', handleMouseDown);
        node.addEventListener('mouseleave', handleMouseLeave);
        node.addEventListener('mouseup', handleMouseUp);
        node.addEventListener('mousemove', handleMouseMove);
        return { destroy() { /* cleanup */ } };
    }
    function formatDate(dateString) {
        if (!dateString || !isFinite(new Date(dateString))) return '—';
        try { return new Date(dateString).toISOString().split('T')[0]; }
        catch (e) { return '无效日期'; }
    }
</script>

<div class="bg-slate-50 min-h-screen font-sans text-slate-800">
    <main class="max-w-screen-2xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <!-- 页面头部 -->

        <!-- ================= 新增的错误提示 ================= -->
        {#if data.error}
            <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg relative mb-4" role="alert">
                <strong class="font-bold">加载错误:</strong>
                <span class="block sm:inline">{data.error}</span>
            </div>
        {/if}
        <!-- ================================================== -->
        <header class="flex justify-between items-center mb-6">
            <div>
                <h1 class="text-3xl font-bold text-slate-900">固定资产管理系统</h1>
                <p class="mt-1 text-md text-slate-600">资产列表</p>
            </div>
            <a href="/logs" class="inline-flex items-center gap-2 px-4 py-2 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition font-semibold shadow-sm">
                查看变更日志
            </a>
        </header>

        <!-- 操作与筛选栏 -->
        <div class="bg-white p-3 rounded-lg border border-slate-200 shadow-sm mb-1 flex justify-between items-center">
             <div class="flex items-center gap-4">
                <button on:click={() => isFilterPanelOpen = !isFilterPanelOpen} class="inline-flex items-center gap-2 px-4 py-2 text-sm font-medium text-slate-700 bg-slate-100 hover:bg-slate-200 rounded-md transition">
                    <SlidersHorizontal size={16} />
                    筛选 / 排序
                </button>
             </div>
             <div class="text-sm text-slate-500">
                共 <span class="font-semibold text-slate-700">{data.assets?.length || 0}</span> 条记录
             </div>
        </div>

        <!-- 可伸缩的筛选面板 -->
        {#if isFilterPanelOpen}
            <div class="bg-white p-4 rounded-lg border border-slate-200 shadow-sm mb-4 transition-all">
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                    <!-- 文本筛选 -->
                    <div>
                        <label for="gdzc_lb" class="block text-sm font-medium text-gray-700">资产类别</label>
                        <input type="text" id="gdzc_lb" bind:value={filters.gdzc_lb} class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" placeholder="e.g., 笔记本">
                    </div>
                    <div>
                        <label for="gdzc_userbm" class="block text-sm font-medium text-gray-700">使用部门</label>
                        <input type="text" id="gdzc_userbm" bind:value={filters.gdzc_userbm} class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" placeholder="e.g., 开发部">
                    </div>
                    <div>
                        <label for="gdzc_zt" class="block text-sm font-medium text-gray-700">资产状态</label>
                        <input type="text" id="gdzc_zt" bind:value={filters.gdzc_zt} class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" placeholder="e.g., 使用中">
                    </div>
                    <div>
                        <label for="gdzc_sccj" class="block text-sm font-medium text-gray-700">生产厂家</label>
                        <input type="text" id="gdzc_sccj" bind:value={filters.gdzc_sccj} class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" placeholder="e.g., 联想">
                    </div>

                    <!-- 范围筛选：金额 -->
                    <div class="md:col-span-2">
                        <label class="block text-sm font-medium text-gray-700">金额范围</label>
                        <div class="mt-1 flex items-center space-x-2">
                            <input type="number" bind:value={filters.gdzc_je_gte} class="block w-full rounded-md border-gray-300 shadow-sm sm:text-sm" placeholder="最小金额">
                            <span>-</span>
                            <input type="number" bind:value={filters.gdzc_je_lte} class="block w-full rounded-md border-gray-300 shadow-sm sm:text-sm" placeholder="最大金额">
                        </div>
                    </div>

                    <!-- 范围筛选：入账日期 -->
                    <div class="md:col-span-2">
                        <label class="block text-sm font-medium text-gray-700">入账日期范围</label>
                        <div class="mt-1 flex items-center space-x-2">
                            <input type="date" bind:value={filters.gdzc_rzdate_gte} class="block w-full rounded-md border-gray-300 shadow-sm sm:text-sm">
                            <span>-</span>
                            <input type="date" bind:value={filters.gdzc_rzdate_lte} class="block w-full rounded-md border-gray-300 shadow-sm sm:text-sm">
                        </div>
                    </div>
                    
                    <!-- 排序 -->
                    <div class="lg:col-span-2">
                        <label class="block text-sm font-medium text-gray-700">排序方式</label>
                        <div class="mt-1 flex items-center space-x-2">
                            <select bind:value={sortBy} class="block w-full rounded-md border-gray-300 shadow-sm sm:text-sm">
                                <option value="gdzc_id">系统编号</option>
                                <option value="gdzc_je">金额</option>
                                <option value="gdzc_nx">使用年限</option>
                                <option value="gdzc_rzdate">入账日期</option>
                                <option value="gdzc_lydate">领用日期</option>
                                <option value="gdzc_djdate">登记日期</option>
                            </select>
                            <select bind:value={order} class="block w-full rounded-md border-gray-300 shadow-sm sm:text-sm">
                                <option value="asc">升序</option>
                                <option value="desc">降序</option>
                            </select>
                        </div>
                    </div>
                </div>
                <div class="mt-4 pt-4 border-t border-slate-200 flex justify-end gap-3">
                    <button on:click={clearFilters} class="inline-flex items-center gap-2 px-4 py-2 text-sm bg-slate-200 text-slate-700 rounded-md hover:bg-slate-300 font-semibold">
                        <RotateCcw size={16} />
                        清除
                    </button>
                    <button on:click={applyFiltersAndSort} class="inline-flex items-center gap-2 px-4 py-2 text-sm bg-indigo-600 text-white rounded-md hover:bg-indigo-700 font-semibold shadow-sm">
                        <Search size={16} />
                        应用
                    </button>
                </div>
            </div>
        {/if}

        <!-- 顶部滚动条 -->
        <div class="top-scrollbar-wrapper" bind:this={topScrollbarWrapper}>
            <div class="top-scrollbar-inner"></div>
        </div>

        <!-- 主表格容器 -->
        <div class="max-h-[70vh] overflow-auto shadow-md rounded-lg border border-slate-200" bind:this={mainTableWrapper} use:dragToScroll>
            <table class="min-w-full divide-y divide-slate-200 relative">
                <thead class="bg-slate-100">
                    <tr>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-24">
                            <button on:click={() => handleSort('gdzc_id')} class="w-full h-full flex items-center justify-start px-3 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group">
                                <span>ID</span>
                                <ChevronsUpDown size={14} class="ml-auto text-slate-400 group-hover:text-slate-600 transition-colors" />
                            </button>
                        </th>
                        <th scope="col" class="sticky left-0 top-0 z-30 bg-slate-100 w-40">
                             <button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap">
                                <span>资产编号</span>
                            </button>
                        </th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-32"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap"><span>资产类别</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-32"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap"><span>规格型号</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-40"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap"><span>生产厂家</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-24"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap"><span>增加方式</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-28"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap"><span>使用部门</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-24"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap"><span>资产状态</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 min-w-[250px]"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group"><span>关键信息</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-28"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap"><span>使用人</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-32">
                            <button on:click={() => handleSort('gdzc_rzdate')} class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap">
                                <span>入账日期</span>
                                <ChevronsUpDown size={14} class="ml-auto text-slate-400 group-hover:text-slate-600 transition-colors" />
                            </button>
                        </th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-32">
                            <button on:click={() => handleSort('gdzc_lydate')} class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap">
                                <span>领用日期</span>
                                <ChevronsUpDown size={14} class="ml-auto text-slate-400 group-hover:text-slate-600 transition-colors" />
                            </button>
                        </th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-28"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap"><span>公司名称</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-20"><button class="w-full h-full flex items-center justify-start px-3 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group"><span>单位</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-24">
                            <button on:click={() => handleSort('gdzc_je')} class="w-full h-full flex items-center justify-start px-3 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group">
                                <span>金额</span>
                                <ChevronsUpDown size={14} class="ml-auto text-slate-400 group-hover:text-slate-600 transition-colors" />
                            </button>
                        </th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-28">
                            <button on:click={() => handleSort('gdzc_nx')} class="w-full h-full flex items-center justify-start px-3 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group">
                                <span>年限</span>
                                <ChevronsUpDown size={14} class="ml-auto text-slate-400 group-hover:text-slate-600 transition-colors" />
                            </button>
                        </th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 min-w-[100px]"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group"><span>备注</span></button></th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-32">
                            <button on:click={() => handleSort('gdzc_djdate')} class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap">
                                <span>登记日期</span>
                                <ChevronsUpDown size={14} class="ml-auto text-slate-400 group-hover:text-slate-600 transition-colors" />
                            </button>
                        </th>
                        <th scope="col" class="sticky top-0 z-10 bg-slate-100 w-28"><button class="w-full h-full flex items-center justify-start px-4 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider group whitespace-nowrap"><span>登记人</span></button></th>
                    </tr>
                </thead>
                <tbody class="bg-white divide-y divide-slate-200">
                    {#each data.assets as asset}
                        <tr class="hover:bg-slate-50">
                            <td class="px-3 py-3 text-sm text-slate-700 text-left">{asset.gdzc_id}</td>
                            <td class="sticky left-0 z-20 bg-white px-4 py-3 whitespace-nowrap text-sm font-medium text-slate-900 text-left border-r border-slate-200">{asset.gdzc_bh}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{asset.gdzc_lb}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{asset.gdzc_ggxh}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{asset.gdzc_sccj}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{asset.gdzc_zjfs}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{asset.gdzc_userbm}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{asset.gdzc_zt}</td>
                            <td class="px-4 py-3 text-sm text-slate-700 text-left break-words">{asset.gdzc_gjxx}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{asset.gdzc_user}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{formatDate(asset.gdzc_rzdate)}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{formatDate(asset.gdzc_lydate)}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{asset.gdzc_gsname}</td>
                            <td class="px-3 py-3 text-sm text-slate-700 text-left">{asset.gdzc_dw}</td>
                            <td class="px-3 py-3 text-sm text-slate-700 text-left">{asset.gdzc_je}</td>
                            <td class="px-3 py-3 text-sm text-slate-700 text-left">{asset.gdzc_nx}</td>
                            <td class="px-4 py-3 text-sm text-slate-700 text-left break-words">{asset.gdzc_beizhu}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{formatDate(asset.gdzc_djdate)}</td>
                            <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-left">{asset.gdzc_djuser}</td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    </main>
</div>

<style>
    .grabbing { cursor: grabbing; }
    tr:hover .sticky { background-color: #f8fafc; /* Tailwind's slate-50 */ }
    .top-scrollbar-wrapper { overflow-x: auto; overflow-y: hidden; margin-bottom: -1px; }
    .top-scrollbar-wrapper::-webkit-scrollbar { height: 12px; }
    .top-scrollbar-wrapper::-webkit-scrollbar-track { background-color: #f1f5f9; }
    .top-scrollbar-wrapper::-webkit-scrollbar-thumb { background-color: #cbd5e1; border-radius: 6px; }
    .top-scrollbar-inner { width: 3200px; height: 1px; } /* Increased width for safety */
    .overflow-auto::-webkit-scrollbar { width: 12px; height: 12px; }
    .overflow-auto::-webkit-scrollbar-track { background-color: #f1f5f9; }
    .overflow-auto::-webkit-scrollbar-thumb { background-color: #cbd5e1; border-radius: 6px; border: 2px solid #f1f5f9; }
</style>