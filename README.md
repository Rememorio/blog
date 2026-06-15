# Rememorio Blog

Static source-reading notes and engineering essays published at
[rememorio.github.io/blog](https://rememorio.github.io/blog/).

This repository is intentionally small: plain HTML, shared CSS, local raster
figures, and a few project-local writing rules. The goal is to make complex
runtime systems readable without hiding the evidence trail behind a build
pipeline.

## Published Articles

### Codex Source Notes

Codex Source Notes is a bilingual source-reading series on
[openai/codex](https://github.com/openai/codex). It starts by following one
ordinary user turn through the governed runtime, then drills into context,
tools, permissions, client projection, and recovery as separate mechanisms.

- Language gateway:
  [rememorio.github.io/blog/codex/](https://rememorio.github.io/blog/codex/)
- Chinese overview:
  [Codex 源码阅读（一）：从一次 turn 看懂受控运行时](https://rememorio.github.io/blog/codex/zh/overview/)
- English overview:
  [Codex Source Notes I: Follow One Turn Through the Governed Runtime](https://rememorio.github.io/blog/codex/en/overview/)
- Chinese context management:
  [Codex 源码阅读（二）：上下文不是聊天记录，而是运行时账本](https://rememorio.github.io/blog/codex/zh/context-management/)
- English context management:
  [Codex Source Notes II: Context Is a Runtime Ledger, Not Chat History](https://rememorio.github.io/blog/codex/en/context-management/)

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
├── index.html                  # Blog home page
├── codex/
│   ├── index.html              # Language gateway
│   ├── cn/index.html           # Chinese series index
│   ├── cn/overview/            # Chinese series overview article
│   ├── cn/context-management/  # Chinese context-management article
│   ├── en/index.html           # English series index
│   ├── en/overview/            # English series overview article
│   ├── en/context-management/  # English context-management article
│   └── assets/*.png            # Final raster figures
├── hermes-agent/
│   ├── index.html              # Language gateway
│   ├── cn/index.html           # Chinese edition
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
- `http://127.0.0.1:4173/codex/`
- `http://127.0.0.1:4173/codex/zh/overview/`
- `http://127.0.0.1:4173/codex/en/overview/`
- `http://127.0.0.1:4173/codex/zh/context-management/`
- `http://127.0.0.1:4173/codex/en/context-management/`
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
