<!-- frontend/src/routes/+layout.svelte -->
<script>
    // 我们需要导入 app.css 来让 Tailwind 生效
    import '../app.css';
</script>

<!-- svelte:head 是一个特殊的 Svelte 元素 -->
<!-- 放在这里面的所有标签，都会被 SvelteKit 智能地渲染到最终 HTML 的 <head> 部分 -->
<svelte:head>
    <title>资产管理系统</title>
    <!-- 你还可以在这里添加其他的 meta 标签, 比如 description -->
    <meta name="description" content="一个使用 Gin 和 SvelteKit 构建的固定资产管理系统" />
</svelte:head>

<!-- <slot /> 是页面的占位符，代表了当前路由的页面内容会被渲染在这里 -->
<slot />