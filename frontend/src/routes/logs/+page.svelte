<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { Search, RotateCcw, SlidersHorizontal, ChevronsUpDown } from 'svelte-lucide-icons';

    export let data;

    // 状态管理
    let filters = {
        gdzc_bh: data.filters.gdzc_bh || '',
        gdzc_userbm: data.filters.gdzc_userbm || '',
        gdzc_user: data.filters.gdzc_user || '',
        gdzc_lb: data.filters.gdzc_lb || '',
        gdzc_sccj: data.filters.gdzc_sccj || '',
        gdzc_zt: data.filters.gdzc_zt || '', // ✅ 新增：资产状态筛选
        bg_time_gte: data.filters.bg_time_gte || '',
        bg_time_lte: data.filters.bg_time_lte || '',
    };
    let sortBy = data.filters.sortBy || 'BGtime';
    let order = data.filters.order || 'desc';
    let isFilterPanelOpen = false;

    // 交互函数 (无需改动)
    function applyFiltersAndSort() {
        const params = new URLSearchParams();
        for (const key in filters) {
            if (filters[key]) params.set(key, filters[key]);
        }
        params.set('sortBy', sortBy);
        params.set('order', order);
        goto(`/logs?${params.toString()}`, { keepFocus: true, noScroll: true, replaceState: true });
        isFilterPanelOpen = false;
    }

    function clearFilters() {
        for (const key in filters) filters[key] = '';
        sortBy = 'BGtime';
        order = 'desc';
        goto(`/logs`, { keepFocus: true, noScroll: true, replaceState: true });
    }
    
    function handleSort(newSortBy) {
        if (sortBy === newSortBy) {
            order = order === 'asc' ? 'desc' : 'asc';
        } else {
            sortBy = newSortBy;
            order = 'desc';
        }
        applyFiltersAndSort();
    }

    // 辅助函数 (无需改动)
    function formatDateTime(dateString) {
        if (!dateString) return '—';
        try {
            const date = new Date(dateString);
            if (isNaN(date.getTime())) return '无效日期';
            return date.toLocaleString('zh-CN', { hour12: false }).replace(/\//g, '-');
        } catch (e) { return '格式化错误'; }
    }
    function formatDate(dateString) {
        if (!dateString) return '—';
        try { return new Date(dateString).toISOString().split('T')[0]; }
        catch (e) { return '无效日期'; }
    }

    // 表格滚动同步逻辑 (无需改动)
    let topScrollbarWrapper, mainTableWrapper;
    onMount(() => {
        if (!topScrollbarWrapper || !mainTableWrapper) return;
        function syncScrollFromMain() { if(topScrollbarWrapper) topScrollbarWrapper.scrollLeft = mainTableWrapper.scrollLeft; }
        function syncScrollFromTop() { if(mainTableWrapper) mainTableWrapper.scrollLeft = topScrollbarWrapper.scrollLeft; }
        const innerScrollbar = topScrollbarWrapper.querySelector('.top-scrollbar-inner');
        const table = mainTableWrapper.querySelector('table');
        if (innerScrollbar && table) {
            innerScrollbar.style.width = getComputedStyle(table).minWidth;
        }
        mainTableWrapper.addEventListener('scroll', syncScrollFromMain);
        topScrollbarWrapper.addEventListener('scroll', syncScrollFromTop);
        return () => {
            if (mainTableWrapper) mainTableWrapper.removeEventListener('scroll', syncScrollFromMain);
            if (topScrollbarWrapper) topScrollbarWrapper.removeEventListener('scroll', syncScrollFromTop);
        };
    });
</script>

<div class="bg-slate-50 min-h-screen font-sans text-slate-800">
    <main class="max-w-screen-2xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        {#if data.error}
            <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg relative mb-4" role="alert">
                <strong class="font-bold">加载错误:</strong> <span class="block sm:inline">{data.error}</span>
            </div>
        {/if}

        <header class="flex justify-between items-center mb-6">
            <div>
                <h1 class="text-3xl font-bold text-slate-900">资产变更日志</h1>
                <p class="mt-1 text-md text-slate-600">所有资产的历史记录快照</p>
            </div>
            <a href="/" class="btn-primary">
                返回资产列表
            </a>
        </header>

        <div class="bg-white p-3 rounded-lg border border-slate-200 shadow-sm mb-1 flex justify-between items-center">
             <div>
                <button on:click={() => isFilterPanelOpen = !isFilterPanelOpen} class="inline-flex items-center gap-2 px-4 py-2 text-sm font-medium text-slate-700 bg-slate-100 hover:bg-slate-200 rounded-md transition">
                    <SlidersHorizontal size={16} /> 筛选 / 排序
                </button>
             </div>
             <div class="text-sm text-slate-500">
                共 <span class="font-semibold text-slate-700">{data.logs?.length || 0}</span> 条记录
             </div>
        </div>

        {#if isFilterPanelOpen}
            <div class="bg-white p-4 rounded-lg border border-slate-200 shadow-sm mb-4">
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                    <!-- 文本筛选 -->
                    <div>
                        <label for="filter_bh" class="block text-sm font-medium text-gray-700">资产编号</label>
                        <input type="text" id="filter_bh" bind:value={filters.gdzc_bh} class="input-filter" placeholder="e.g., SX-BJB-001">
                    </div>
                    <div>
                        <label for="filter_lb" class="block text-sm font-medium text-gray-700">资产类别</label>
                        <input type="text" id="filter_lb" bind:value={filters.gdzc_lb} class="input-filter" placeholder="e.g., 笔记本">
                    </div>
                    <div>
                        <label for="filter_sccj" class="block text-sm font-medium text-gray-700">生产厂家</label>
                        <input type="text" id="filter_sccj" bind:value={filters.gdzc_sccj} class="input-filter" placeholder="e.g., 联想">
                    </div>
                     <div>
                        <label for="filter_userbm" class="block text-sm font-medium text-gray-700">使用部门</label>
                        <input type="text" id="filter_userbm" bind:value={filters.gdzc_userbm} class="input-filter" placeholder="e.g., 开发部">
                    </div>
                    <!-- ✅ 新增: 资产状态筛选 -->
                    <div>
                        <label for="filter_zt" class="block text-sm font-medium text-gray-700">资产状态</label>
                        <input type="text" id="filter_zt" bind:value={filters.gdzc_zt} class="input-filter" placeholder="e.g., 使用中">
                    </div>
                    <div>
                        <label for="filter_user" class="block text-sm font-medium text-gray-700">使用人</label>
                        <input type="text" id="filter_user" bind:value={filters.gdzc_user} class="input-filter" placeholder="e.g., XXX">
                    </div>
                    
                    <!-- 范围筛选：变更时间 -->
                    <div class="md:col-span-2">
                        <label class="block text-sm font-medium text-gray-700">变更时间范围</label>
                        <div class="mt-1 flex items-center space-x-2">
                            <input type="datetime-local" bind:value={filters.bg_time_gte} class="input-filter">
                            <span>-</span>
                            <input type="datetime-local" bind:value={filters.bg_time_lte} class="input-filter">
                        </div>
                    </div>
                    
                    <!-- 排序 -->
                    <div class="lg:col-span-2">
                        <label class="block text-sm font-medium text-gray-700">排序方式</label>
                        <div class="mt-1 flex items-center space-x-2">
                            <select bind:value={sortBy} class="input-filter">
                                <option value="BGtime">变更时间</option>
                                <option value="gdzc_ID">日志ID</option>
                                <option value="gdzc_je">金额</option>
                                <option value="gdzc_rzdate">入账日期</option>
                                <option value="gdzc_djdate">登记日期</option>
                            </select>
                            <select bind:value={order} class="input-filter">
                                <option value="desc">降序</option>
                                <option value="asc">升序</option>
                            </select>
                        </div>
                    </div>
                </div>
                <div class="mt-4 pt-4 border-t border-slate-200 flex justify-end gap-3">
                    <button on:click={clearFilters} class="btn-secondary"><RotateCcw size={16} /> 清除</button>
                    <button on:click={applyFiltersAndSort} class="btn-primary"><Search size={16} /> 应用</button>
                </div>
            </div>
        {/if}

        <!-- 表格部分 (无需改动) -->
        <div class="top-scrollbar-wrapper" bind:this={topScrollbarWrapper}><div class="top-scrollbar-inner"></div></div>
        <div class="max-h-[70vh] overflow-auto shadow-md rounded-lg border border-slate-200" bind:this={mainTableWrapper}>
            <!-- ... 表格 HTML ... -->
            <table class="min-w-full divide-y divide-slate-200 relative">
                <thead class="bg-slate-100">
                     <tr>
                        <th class="th-sticky w-48"><button on:click={() => handleSort('BGtime')} class="th-btn group">变更时间<ChevronsUpDown size={14} class="icon" /></button></th>
                        <th class="th-sticky w-24"><button on:click={() => handleSort('gdzc_ID')} class="th-btn group">日志ID<ChevronsUpDown size={14} class="icon" /></button></th>
                        <th class="th-sticky left-0 z-30 w-40">资产编号</th>
                        <th class="th-sticky w-32">资产类别</th><th class="th-sticky w-32">规格型号</th><th class="th-sticky w-40">生产厂家</th>
                        <th class="th-sticky w-24">增加方式</th><th class="th-sticky w-28">使用部门</th><th class="th-sticky w-24">资产状态</th>
                        <th class="th-sticky min-w-[250px]">关键信息</th><th class="th-sticky w-28">使用人</th>
                        <th class="th-sticky w-32"><button on:click={() => handleSort('gdzc_rzdate')} class="th-btn group">入账日期<ChevronsUpDown size={14} class="icon" /></button></th>
                        <th class="th-sticky w-32">领用日期</th><th class="th-sticky w-28">公司名称</th><th class="th-sticky w-20">单位</th>
                        <th class="th-sticky w-24"><button on:click={() => handleSort('gdzc_je')} class="th-btn group">金额<ChevronsUpDown size={14} class="icon" /></button></th>
                        <th class="th-sticky w-28">年限</th><th class="th-sticky min-w-[100px]">备注</th>
                        <th class="th-sticky w-32"><button on:click={() => handleSort('gdzc_djdate')} class="th-btn group">登记日期<ChevronsUpDown size={14} class="icon" /></button></th>
                        <th class="th-sticky w-28">登记人</th>
                    </tr>
                </thead>
                <tbody class="bg-white divide-y divide-slate-200">
                    {#if data.logs && data.logs.length > 0}
                        {#each data.logs as log (log.log_id)}
                            <tr class="hover:bg-slate-50">
                                <td class="px-3 py-3 text-sm text-slate-700 whitespace-nowrap">{formatDateTime(log.bg_time)}</td>
                                <td class="px-3 py-3 text-sm text-slate-700">{log.log_id}</td>
                                <td class="sticky left-0 z-20 bg-white hover:bg-slate-50 px-4 py-3 whitespace-nowrap text-sm font-medium text-slate-900 border-r border-slate-200">{log.gdzc_bh}</td>
                                <td class="td-cell">{log.gdzc_lb}</td><td class="td-cell">{log.gdzc_ggxh}</td><td class="td-cell">{log.gdzc_sccj}</td>
                                <td class="td-cell">{log.gdzc_zjfs}</td><td class="td-cell">{log.gdzc_userbm}</td><td class="td-cell">{log.gdzc_zt}</td>
                                <td class="px-4 py-3 text-sm text-slate-700 break-words">{log.gdzc_gjxx}</td><td class="td-cell">{log.gdzc_user}</td>
                                <td class="td-cell">{formatDate(log.gdzc_rzdate)}</td><td class="td-cell">{formatDate(log.gdzc_lydate)}</td>
                                <td class="td-cell">{log.gdzc_gsname}</td><td class="px-3 py-3 text-sm text-slate-700">{log.gdzc_dw}</td>
                                <td class="px-3 py-3 text-sm text-slate-700 text-right">{log.gdzc_je}</td><td class="px-3 py-3 text-sm text-slate-700 text-right">{log.gdzc_nx}</td>
                                <td class="px-4 py-3 text-sm text-slate-700 break-words">{log.gdzc_beizhu}</td>
                                <td class="td-cell">{formatDate(log.gdzc_djdate)}</td><td class="td-cell">{log.gdzc_djuser}</td>
                            </tr>
                        {/each}
                    {:else if !data.error}
                        <tr><td colspan="20" class="text-center py-10 text-slate-500">没有找到符合条件的日志记录。</td></tr>
                    {/if}
                </tbody>
            </table>
        </div>
    </main>
</div>

<style>
    .input-filter { @apply mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm; }
    .btn-primary { @apply inline-flex items-center gap-2 px-4 py-2 text-sm bg-indigo-600 text-white rounded-md hover:bg-indigo-700 font-semibold shadow-sm; }
    .btn-secondary { @apply inline-flex items-center gap-2 px-4 py-2 text-sm bg-slate-200 text-slate-700 rounded-md hover:bg-slate-300 font-semibold; }
    .th-sticky { @apply sticky top-0 z-10 bg-slate-100 px-3 py-3 text-left text-xs font-semibold text-slate-600 uppercase tracking-wider whitespace-nowrap; }
    .th-sticky.left-0 { @apply z-30; }
    .th-btn { @apply w-full h-full flex items-center; }
    .icon { @apply ml-auto text-slate-400 group-hover:text-slate-600 transition-colors; }
    .td-cell { @apply px-4 py-3 whitespace-nowrap text-sm text-slate-700; }
    .top-scrollbar-wrapper { overflow-x: auto; overflow-y: hidden; }
    .top-scrollbar-wrapper::-webkit-scrollbar { display: none; }
    .top-scrollbar-inner { width: 3200px; height: 1px; }
    .overflow-auto::-webkit-scrollbar { width: 12px; height: 12px; }
    .overflow-auto::-webkit-scrollbar-track { background-color: #f1f5f9; }
    .overflow-auto::-webkit-scrollbar-thumb { background-color: #cbd5e1; border-radius: 6px; border: 2px solid #f1f5f9; }
</style>