// frontend/src/routes/+page.server.js (健壮版)

/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch, url }) {
    // 无论成功与否，都先准备好 filters 对象
    const currentFilters = Object.fromEntries(url.searchParams.entries());

    try {
        const response = await fetch(`/api/assets${url.search}`);
        
        if (!response.ok) {
            // 即使失败，也要返回一致的结构
            return { 
                assets: [], 
                filters: currentFilters, 
                error: `Failed to fetch assets: ${response.statusText}` 
            };
        }

        const assets = await response.json();

        // 成功时，返回完整数据
        return {
            assets: assets,
            filters: currentFilters
        };
    } catch (e) {
        // 网络错误等异常情况，也要返回一致的结构
        return {
            assets: [],
            filters: currentFilters,
            error: `Network error: ${e.message}`
        };
    }
}