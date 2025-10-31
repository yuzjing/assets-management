import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    kit: {
        // 将 adapter 从 adapter-auto() 更换为 adapter-static()
        adapter: adapter({
            // SvelteKit 构建输出的默认目录是 'build'
            // 这两行可以不写，因为它们是默认值
            pages: 'build',
            assets: 'build',

            // 这是单页应用 (SPA) 的关键设置!
            // 它会创建一个 index.html 作为所有未匹配路由的回退页面。
            // 这样，当你在浏览器直接访问 /some/route 时，你的 Go 服务器
            // 仍然可以返回 index.html，然后 Svelte 的前端路由接管。
            fallback: 'index.html',

            precompress: false,
            strict: true
        })
    }
};

export default config;