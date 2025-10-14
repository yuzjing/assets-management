// frontend/src/routes/logs/+page.server.js

/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch }) {
    //  /logs API 获取数据
    const response = await fetch('http://localhost:8123/logs');
    
    if (!response.ok) {
        return { logs: [], error: 'Failed to fetch logs' };
    }

    const logs = await response.json();

    return {
        logs: logs
    };
}