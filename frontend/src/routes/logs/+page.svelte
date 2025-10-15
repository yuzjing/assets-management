<!-- frontend/src/routes/logs/+page.svelte -->
<script>
    export let data;

    function formatDateTime(dateString) {
        if (!dateString || !isFinite(new Date(dateString))) return '—';
        try {
            // 使用 toLocaleString 显示本地化的日期和时间
            return new Date(dateString).toLocaleString();
        } catch (e) {
            return '无效日期';
        }
    }
</script>

<div class="bg-gray-50 min-h-screen">
    <main class="max-w-screen-2xl mx-auto py-4 px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center mb-4">
            <div class="text-left">
                <h1 class="text-2xl font-bold text-gray-900">资产变更日志</h1>
                <p class="mt-1 text-sm text-gray-600">所有资产的变更记录</p>
            </div>
            <!-- 返回主表的按钮 -->
            <a href="/" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition">
                返回资产列表
            </a>
        </div>
        
        <div class="overflow-x-auto shadow-sm rounded-lg border border-gray-200">
            <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-100">
                    <tr>
                        <!-- 日志表的表头 -->
                        <th class="px-4 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">日志ID</th>
                        <th class="px-4 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">固定资产编号</th>
                        <th class="px-4 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">资产类别</th>
                        <th class="px-4 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">使用部门</th>
                        <th class="px-4 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">使用人</th>
                        <th class="px-4 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">变更时间</th>
                        <th class="px-4 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">登记人</th>
                    </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                    {#each data.logs as log}
                        <tr class="hover:bg-gray-50">
                            <!-- 日志表的数据 -->
                            <td class="px-4 py-4 whitespace-nowrap text-sm">{log.log_id}</td>
                            <td class="px-4 py-4 whitespace-nowrap text-sm">{log.gdzc_bh}</td>
                            <td class="px-4 py-4 whitespace-nowrap text-sm">{log.gdzc_lb}</td>
                            <td class="px-4 py-4 whitespace-nowrap text-sm">{log.ggxh}</td>
                            <td class="px-4 py-4 whitespace-nowrap text-sm">{log.gdzc_user}</td>
                            <td class="px-4 py-4 whitespace-nowrap text-sm">{formatDateTime(log.bg_time)}</td>
                            <td class="px-4 py-4 whitespace-nowrap text-sm">{log.gdzc_djuser}</td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    </main>
</div>