/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch, url }) {
    const params = new URLSearchParams(url.search);
    
    // 从 URL 中提取所有可能的筛选参数
    const currentFilters = {
        gdzc_bh: params.get('gdzc_bh') || '',
        gdzc_userbm: params.get('gdzc_userbm') || '',
        gdzc_user: params.get('gdzc_user') || '',
        gdzc_lb: params.get('gdzc_lb') || '',
        gdzc_sccj: params.get('gdzc_sccj') || '',
        gdzc_zt: params.get('gdzc_zt') || '', // ✅ 新增：资产状态筛选
        bg_time_gte: params.get('bg_time_gte') || '',
        bg_time_lte: params.get('bg_time_lte') || '',
        sortBy: params.get('sortBy') || 'BGtime',
        order: params.get('order') || 'desc',
    };

    try {
        const response = await fetch(`/api/logs?${params.toString()}`);
        
        if (!response.ok) {
            return { 
                logs: [], 
                filters: currentFilters, 
                error: `加载日志失败: ${response.status} ${response.statusText}` 
            };
        }

        const logs = await response.json();

        return {
            logs,
            filters: currentFilters
        };
    } catch (e) {
        return {
            logs: [],
            filters: currentFilters,
            error: `网络错误或后端服务未响应: ${e.message}`
        };
    }
}