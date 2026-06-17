# Rememorio Blog

Static source-reading notes and engineering essays published at
[rememorio.github.io/blog](https://rememorio.github.io/blog/).

This repository is intentionally small: plain HTML, shared CSS, local raster
figures, and a few project-local writing rules. The goal is to make complex
runtime systems readable without hiding the evidence trail behind a build
pipeline.

## Published Articles

- Language gateway:
  [rememorio.github.io/blog](https://rememorio.github.io/blog/)
- Chinese home:
  [rememorio.github.io/blog/zh/](https://rememorio.github.io/blog/zh/)
- English home:
  [rememorio.github.io/blog/en/](https://rememorio.github.io/blog/en/)

### Codex Source Notes

Codex Source Notes is a bilingual source-reading series on
[openai/codex](https://github.com/openai/codex). It starts by following one
ordinary user request through the runtime path, then drills into context,
tools, permissions, client projection, extensions, hooks, prompt caching, and
recovery as separate mechanisms.

- Language gateway:
  [rememorio.github.io/blog/codex/](https://rememorio.github.io/blog/codex/)
- Chinese overview:
  [Codex 源码阅读（一）：先跟一次请求走完整条运行主线](https://rememorio.github.io/blog/codex/zh/overview/)
- English overview:
  [Codex Source Notes I: Follow One Turn Through the Governed Runtime](https://rememorio.github.io/blog/codex/en/overview/)
- Chinese context management:
  [Codex 源码阅读（二）：上下文不是聊天记录，而是运行时账本](https://rememorio.github.io/blog/codex/zh/context-management/)
- English context management:
  [Codex Source Notes II: Context Is a Runtime Ledger, Not Chat History](https://rememorio.github.io/blog/codex/en/context-management/)
- Chinese protocol and events:
  [Codex 源码阅读（三）：protocol 与事件流，客户端怎样共享同一套事实](https://rememorio.github.io/blog/codex/zh/protocol-events/)
- English protocol and events:
  [Codex Source Notes III: Protocol and Event Stream](https://rememorio.github.io/blog/codex/en/protocol-events/)
- Chinese tools:
  [Codex 源码阅读（四）：工具调用背后的运行时合约](https://rememorio.github.io/blog/codex/zh/tools/)
- English tools:
  [Codex Source Notes IV: Tools Are Runtime Contracts](https://rememorio.github.io/blog/codex/en/tools/)
- Chinese permission and sandboxing:
  [Codex 源码阅读（五）：权限与 sandbox，副作用怎样被运行时收口](https://rememorio.github.io/blog/codex/zh/permissions-sandbox/)
- English permission and sandboxing:
  [Codex Source Notes V: Permission and Sandboxing](https://rememorio.github.io/blog/codex/en/permissions-sandbox/)
- Chinese client projection:
  [Codex 源码阅读（六）：客户端投影，同一套事实怎样变成可见状态](https://rememorio.github.io/blog/codex/zh/client-projection/)
- English client projection:
  [Codex Source Notes VI: Client Projection](https://rememorio.github.io/blog/codex/en/client-projection/)
- Chinese extensions and multi-agent:
  [Codex 源码阅读（七）：扩展与多 agent，额外能力怎样进入一次 turn](https://rememorio.github.io/blog/codex/zh/extensions-agents/)
- English extensions and multi-agent:
  [Codex Source Notes VII: Extensions and Multi-Agent](https://rememorio.github.io/blog/codex/en/extensions-agents/)
- Chinese hooks and lifecycle:
  [Codex 源码阅读（八）：Hooks 与生命周期，边界处还能改变什么](https://rememorio.github.io/blog/codex/zh/hooks-lifecycle/)
- English hooks and lifecycle:
  [Codex Source Notes VIII: Hooks and Lifecycle Slots](https://rememorio.github.io/blog/codex/en/hooks-lifecycle/)
- Chinese performance and prompt cache:
  [Codex 源码阅读（九）：性能与 prompt cache，速度来自请求形状](https://rememorio.github.io/blog/codex/zh/prompt-cache-performance/)
- English performance and prompt cache:
  [Codex Source Notes IX: Performance and Prompt Cache](https://rememorio.github.io/blog/codex/en/prompt-cache-performance/)

### Hermes Agent

Hermes Agent is a source-grounded walkthrough of
[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent). It
explains the runtime loop behind a long-running, self-evolving agent: foreground
turns, model-visible views, durable transcripts, memory, session search, skills,
background review, curator, gateway, cron, and delegation.

- Language gateway:
  [rememorio.github.io/blog/hermes-agent/](https://rememorio.github.io/blog/hermes-agent/)
- Chinese:
  [Hermes Agent: 自进化 Agent 的运行闭环](https://rememorio.github.io/blog/hermes-agent/zh/)
- English:
  [Hermes Agent: the Runtime Loop Behind a Self-Evolving Agent](https://rememorio.github.io/blog/hermes-agent/en/)

## Repository Layout

```text
.
├── index.html                  # Site language gateway
├── zh/index.html               # Chinese blog home page
├── en/index.html               # English blog home page
├── codex/
│   ├── index.html              # Language gateway
│   ├── zh/index.html           # Chinese series index
│   ├── zh/overview/            # Chinese series overview article
│   ├── zh/context-management/  # Chinese context-management article
│   ├── zh/protocol-events/     # Chinese protocol/event-stream article
│   ├── zh/tools/               # Chinese tools-runtime article
│   ├── zh/permissions-sandbox/ # Chinese permission/sandbox article
│   ├── zh/client-projection/   # Chinese client-projection article
│   ├── zh/extensions-agents/   # Chinese extensions/multi-agent article
│   ├── zh/hooks-lifecycle/     # Chinese hooks/lifecycle article
│   ├── zh/prompt-cache-performance/ # Chinese performance/cache article
│   ├── en/index.html           # English series index
│   ├── en/overview/            # English series overview article
│   ├── en/context-management/  # English context-management article
│   ├── en/protocol-events/     # English protocol/event-stream article
│   ├── en/tools/               # English tools-runtime article
│   ├── en/permissions-sandbox/ # English permission/sandbox article
│   ├── en/client-projection/   # English client-projection article
│   ├── en/extensions-agents/   # English extensions/multi-agent article
│   ├── en/hooks-lifecycle/     # English hooks/lifecycle article
│   ├── en/prompt-cache-performance/ # English performance/cache article
│   └── assets/*.png            # Final raster figures
├── hermes-agent/
│   ├── index.html              # Language gateway
│   ├── zh/index.html           # Chinese edition
│   ├── en/index.html           # English edition
│   ├── styles.css              # Shared article styling
│   └── assets/*.png            # Final raster figures
├── .codex/skills/
│   ├── article/SKILL.md        # Article maintenance rules
│   └── image/SKILL.md          # Figure generation and QA rules
├── Eyjafjalla.png              # Rememorio figure brand asset
└── LICENSE
```

## Local Preview

No build step is required.

```bash
python3 -m http.server 4173
```

Then open:

- `http://127.0.0.1:4173/`
- `http://127.0.0.1:4173/zh/`
- `http://127.0.0.1:4173/en/`
- `http://127.0.0.1:4173/codex/`
- `http://127.0.0.1:4173/codex/zh/overview/`
- `http://127.0.0.1:4173/codex/en/overview/`
- `http://127.0.0.1:4173/codex/zh/context-management/`
- `http://127.0.0.1:4173/codex/en/context-management/`
- `http://127.0.0.1:4173/codex/zh/protocol-events/`
- `http://127.0.0.1:4173/codex/en/protocol-events/`
- `http://127.0.0.1:4173/codex/zh/tools/`
- `http://127.0.0.1:4173/codex/en/tools/`
- `http://127.0.0.1:4173/codex/zh/permissions-sandbox/`
- `http://127.0.0.1:4173/codex/en/permissions-sandbox/`
- `http://127.0.0.1:4173/codex/zh/client-projection/`
- `http://127.0.0.1:4173/codex/en/client-projection/`
- `http://127.0.0.1:4173/codex/zh/extensions-agents/`
- `http://127.0.0.1:4173/codex/en/extensions-agents/`
- `http://127.0.0.1:4173/codex/zh/hooks-lifecycle/`
- `http://127.0.0.1:4173/codex/en/hooks-lifecycle/`
- `http://127.0.0.1:4173/codex/zh/prompt-cache-performance/`
- `http://127.0.0.1:4173/codex/en/prompt-cache-performance/`
- `http://127.0.0.1:4173/hermes-agent/`

Use a local server rather than opening files directly when checking relative
URLs, language gateways, anchor redirects, and GitHub Pages behavior.

## Writing Standards

Articles should read as finished technical essays, not source dumps. A good
article in this repository:

- starts with a practical problem before naming internal mechanisms;
- states the evidence boundary early;
- links source claims to verified public source paths or canonical docs;
- explains lifecycle, ownership, and failure boundaries in reader-facing order;
- uses tables for compact comparisons when they improve scanning;
- keeps public prose free of local paths, private notes, and process markers;
- ends with transferable rules the reader can reuse outside the named project.

For bilingual articles, use sibling pages:

- `<article>/zh/`
- `<article>/en/`
- `<article>/` as a language gateway

Each edition should have its own title, metadata, table of contents, language
switch, and localized alt text. Shared figures are preferred when their labels
are concise and language-neutral.

## Figure Standards

Public figures are final PNG assets stored under the article's `assets/`
directory. They should explain mechanisms, not decorate pages.

Figures should:

- use warm paper, restrained colors, and hand-drawn technical diagram styling;
- make ownership boundaries, lifecycles, or state ledgers visible;
- keep labels short and readable at article width;
- avoid misleading arrows, crowded callouts, and overlapping text;
- include exactly one deterministic Rememorio brand treatment;
- be checked at both desktop and mobile widths before publication.

When a prose change alters a lifecycle, side path, state boundary, or comparison,
run a visual impact pass and update the relevant figure if needed.

## Maintenance Checklist

Before publishing changes:

- Verify all relative image and stylesheet paths from each page depth.
- Check desktop and mobile widths for body overflow, table overflow, and image
  overflow.
- Confirm each public image has meaningful `alt` text.
- Scan for local paths, temporary files, private prompts, and tool signatures.
- Sample-check representative GitHub source links.
- For bilingual changes, verify canonical and `hreflang` links, language
  switches, the blog index, and old hash redirects.

## License

This repository is licensed under the [MIT License](./LICENSE).
