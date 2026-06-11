# Hermes Agent 源码阅读

这个目录是一个零构建的静态页面。

- 入口：`index.html`
- 样式：`styles.css`
- 视觉资产：`assets/hermes-banner.png`
- 章节源稿：`chapters/*.md`
- 章节页面：`chapters/*.html`
- 官方文档：https://hermes-agent.nousresearch.com/docs/
- 源码版本：https://github.com/NousResearch/hermes-agent/tree/3edd09a46

本组内容基于 `NousResearch/hermes-agent@3edd09a46`，主题是从自进化闭环切入，拆解 Hermes Agent 的 memory、skills、session_search、background review、curator 和运行时触发链路。

章节顺序：

1. `chapters/00-series-map.md`：源码阅读路线
2. `chapters/01-learning-loop.md`：自进化闭环
3. `chapters/02-memory.md`：Memory
4. `chapters/03-skills.md`：Skills
5. `chapters/04-session-search.md`：Session Search
6. `chapters/05-runtime-triggers.md`：运行时触发点
7. `chapters/06-summary-and-questions.md`：总结与追问

本地预览可以直接打开 `index.html`，或在仓库根目录启动任意静态文件服务后访问 `/hermes-agent/`。

修改章节源稿后，运行 `python3 hermes-agent/render_chapters.py` 重新生成 `chapters/*.html`。
