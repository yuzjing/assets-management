// frontend/vite.config.js
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		// --- 添加或修改这个 proxy 配置 ---
		proxy: {
			// 匹配所有以 '/api' 开头的请求
			'/api': {
				// 将请求转发到你的 Go 后端服务器地址
				target: 'http://localhost:8123',
				
				// 必须设置为 true, 否则后端可能会因为来源(Origin)不对而拒绝请求
				changeOrigin: true,

                // 注意: 在我们的设置中，不需要 rewrite。
                // 因为我们的 Go 后端已经在 '/api' 路由组下监听 '/assets'，
                // 所以 Go 服务监听的完整路径就是 '/api/assets'。
                // 代理直接将 '/api/assets' 请求转发到目标服务器的 '/api/assets' 即可。
			}
		}
	},
	preview: {
        // 让 preview 服务器也使用同样的代理规则
        proxy: {
			'/api': {
				target: 'http://localhost:8123',
				changeOrigin: true
			}
        }
    }
});