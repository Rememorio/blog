# Rememorio Blog

Source for [rememorio.github.io/blog](https://rememorio.github.io/blog/), a
static technical essay site about AI coding agents, runtime design, context
management, tool execution, permissions, long-term memory, and prompt-cache
behavior.

The repository is intentionally plain: hand-authored HTML, shared CSS, local
PNG figures, and a small set of writing rules. There is no application
framework, bundler, package manager, or build pipeline between the source files
and GitHub Pages.

## Published Site

| Area | Description | Entry points |
| --- | --- | --- |
| Blog home | Language gateway and bilingual home pages. | [Gateway](https://rememorio.github.io/blog/) / [中文](https://rememorio.github.io/blog/zh/) / [English](https://rememorio.github.io/blog/en/) |
| AI Engineering | Bilingual systems series on prompts, context, harnesses, loops, and evals. | [Gateway](https://rememorio.github.io/blog/ai-engineering/) / [中文](https://rememorio.github.io/blog/ai-engineering/zh/) / [English](https://rememorio.github.io/blog/ai-engineering/en/) |
| Agent Framework Source Notes | Bilingual source-reading series comparing agent frameworks by runtime ownership, including [agentscope-ai/agentscope](https://github.com/agentscope-ai/agentscope), [earendil-works/pi](https://github.com/earendil-works/pi), [google/adk-python](https://github.com/google/adk-python), [agno-agi/agno](https://github.com/agno-agi/agno), [microsoft/autogen](https://github.com/microsoft/autogen), [crewAIInc/crewAI](https://github.com/crewAIInc/crewAI), [cloudwego/eino](https://github.com/cloudwego/eino), and [trpc-group/trpc-agent-go](https://github.com/trpc-group/trpc-agent-go). | [Gateway](https://rememorio.github.io/blog/agent-framework/) / [中文](https://rememorio.github.io/blog/agent-framework/zh/) / [English](https://rememorio.github.io/blog/agent-framework/en/) |
| Claude Code Source Notes | Bilingual source-reading series on the public [Rememorio/claude-code](https://github.com/Rememorio/claude-code) mirror. | [Gateway](https://rememorio.github.io/blog/claude-code/) / [中文](https://rememorio.github.io/blog/claude-code/zh/) / [English](https://rememorio.github.io/blog/claude-code/en/) |
| Codex Source Notes | Bilingual source-reading series on [openai/codex](https://github.com/openai/codex). | [Gateway](https://rememorio.github.io/blog/codex/) / [中文](https://rememorio.github.io/blog/codex/zh/) / [English](https://rememorio.github.io/blog/codex/en/) |
| Hermes Agent Source Notes | Bilingual two-part series on the runtime boundaries in [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) and evaluated skill evolution in [NousResearch/hermes-agent-self-evolution](https://github.com/NousResearch/hermes-agent-self-evolution) with [GEPA](https://arxiv.org/abs/2507.19457v2). | [Series](https://rememorio.github.io/blog/hermes-agent/) / [Runtime 中文](https://rememorio.github.io/blog/hermes-agent/zh/) / [Runtime English](https://rememorio.github.io/blog/hermes-agent/en/) / [Self-Evolution 中文](https://rememorio.github.io/blog/hermes-agent/zh/self-evolution/) / [Self-Evolution English](https://rememorio.github.io/blog/hermes-agent/en/self-evolution/) |
| Temporal | Bilingual source-grounded walkthrough of [temporalio/temporal](https://github.com/temporalio/temporal), durable execution, and its boundary with [trpc-group/trpc-agent-go](https://github.com/trpc-group/trpc-agent-go). | [Gateway](https://rememorio.github.io/blog/temporal/) / [中文](https://rememorio.github.io/blog/temporal/zh/) / [English](https://rememorio.github.io/blog/temporal/en/) |
| Agent Memory | Bilingual route map and source notes for memory systems including [mem0ai/mem0](https://github.com/mem0ai/mem0), [letta-ai/letta](https://github.com/letta-ai/letta), [getzep/graphiti](https://github.com/getzep/graphiti), [langchain-ai/langmem](https://github.com/langchain-ai/langmem), [TencentCloud/TencentDB-Agent-Memory](https://github.com/TencentCloud/TencentDB-Agent-Memory), [topoteretes/cognee](https://github.com/topoteretes/cognee), and [supermemoryai/supermemory](https://github.com/supermemoryai/supermemory). | [Gateway](https://rememorio.github.io/blog/agent-memory/) / [中文](https://rememorio.github.io/blog/agent-memory/zh/) / [English](https://rememorio.github.io/blog/agent-memory/en/) / [Mem0 Gateway](https://rememorio.github.io/blog/agent-memory/mem0-memory/) / [Mem0 中文](https://rememorio.github.io/blog/agent-memory/mem0-memory/zh/) / [Mem0 English](https://rememorio.github.io/blog/agent-memory/mem0-memory/en/) / [Letta 中文](https://rememorio.github.io/blog/agent-memory/letta/zh/) / [Letta English](https://rememorio.github.io/blog/agent-memory/letta/en/) / [Graphiti 中文](https://rememorio.github.io/blog/agent-memory/graphiti/zh/) / [Graphiti English](https://rememorio.github.io/blog/agent-memory/graphiti/en/) / [LangMem 中文](https://rememorio.github.io/blog/agent-memory/langmem/zh/) / [LangMem English](https://rememorio.github.io/blog/agent-memory/langmem/en/) / [TencentDB 中文](https://rememorio.github.io/blog/agent-memory/tencentdb-agent-memory/zh/) / [TencentDB English](https://rememorio.github.io/blog/agent-memory/tencentdb-agent-memory/en/) / [Cognee/Supermemory 中文](https://rememorio.github.io/blog/agent-memory/cognee-supermemory/zh/) / [Cognee/Supermemory English](https://rememorio.github.io/blog/agent-memory/cognee-supermemory/en/) |

Use the public site indexes as the canonical article catalog. The README should
stay stable as new essays are added.

## Repository Layout

```text
.
├── index.html                  # Site language gateway
├── zh/                         # Chinese blog home
├── en/                         # English blog home
├── ai-engineering/             # AI Engineering systems series
│   ├── index.html              # Series language gateway
│   ├── assets/                 # Shared language-neutral figures and styles
│   ├── zh/                     # Chinese articles
│   └── en/                     # English articles
├── agent-framework/             # Agent framework source-reading series
│   ├── index.html              # Series language gateway
│   ├── zh/                     # Chinese articles and figures
│   └── en/                     # English articles and figures
├── claude-code/                # Claude Code series
│   ├── index.html              # Series language gateway
│   ├── zh/                     # Chinese articles and figures
│   └── en/                     # English articles and figures
├── codex/                      # Codex series
│   ├── index.html              # Series language gateway
│   ├── zh/                     # Chinese articles and figures
│   └── en/                     # English articles and figures
├── hermes-agent/               # Hermes Agent source-reading series
│   ├── index.html              # Series gateway and route map
│   ├── zh/                     # Chinese runtime article + self-evolution chapter
│   ├── en/                     # English runtime article + self-evolution chapter
│   └── styles.css
├── temporal/                   # Temporal durable execution article
│   ├── index.html              # Article language gateway
│   ├── zh/
│   └── en/
├── agent-memory/               # Agent memory series map
│   ├── index.html              # Series language gateway
│   ├── zh/
│   ├── en/
│   ├── mem0-memory/            # Mem0 chapter used by the Agent Memory series
│   ├── letta/                  # Letta stateful agent memory chapter
│   ├── graphiti/               # Graphiti temporal context graph chapter
│   ├── langmem/                # LangMem and LangGraph workflow memory chapter
│   ├── tencentdb-agent-memory/ # TencentDB context offload and L0-L3 memory chapter
│   └── cognee-supermemory/     # Cognee and Supermemory platform memory chapter
├── .codex/skills/              # Local writing and image-generation rules
├── Eyjafjalla.png              # Rememorio figure brand asset
└── LICENSE
```

Every published page is an `index.html` file. Article URLs are therefore stable
directory URLs such as `/codex/zh/tools/` rather than filename URLs.

## Local Preview

No install or build step is required. Run any static file server from the
repository root; Python's built-in server is the lowest-dependency option:

```bash
python3 -m http.server 4173
```

Then open [http://127.0.0.1:4173/](http://127.0.0.1:4173/). Use a local server
rather than opening files directly when checking relative links, language
gateways, canonical URLs, redirects, stylesheets, and images.

## Content Conventions

Articles should read as finished technical essays, not source dumps. A good
article in this repository:

- starts with a practical problem before naming internal mechanisms;
- states the evidence boundary early;
- links source claims to verified public source paths or canonical docs;
- explains lifecycle, ownership, and failure boundaries in reader-facing order;
- uses tables only when they make comparison or scanning easier;
- keeps public prose free of local paths, private notes, prompts, and process
  markers;
- ends with transferable rules the reader can reuse outside the named project.

For bilingual work, keep each edition as a sibling page:

- `<topic>/zh/`
- `<topic>/en/`
- `<topic>/` as the language gateway

Each edition should have its own title, metadata, table of contents, language
switch, canonical link, `hreflang` links, and localized alt text. Shared figures
are preferred when labels are concise and language-neutral.
For series pages, both languages should expose the same public chapter set,
route order, overview image, and navigation graph.

## Figure Conventions

Public figures are final PNG assets stored under the relevant article's
`assets/` directory. They should explain mechanisms rather than decorate pages.

Figures should:

- use warm paper, restrained colors, and hand-drawn technical diagram styling;
- make ownership boundaries, lifecycles, or state ledgers visible;
- keep labels short and readable at article width;
- avoid misleading arrows, crowded callouts, and overlapping text;
- include one integrated Rememorio treatment; established series should preserve
  the in-scene Eyjafjalla guide rather than replacing it with a generic card;
- keep series overview or route-map figures aligned with the published chapter
  list, ordering, labels, and alt text;
- be checked at both desktop and mobile widths before publication.

When prose changes alter a lifecycle, side path, state boundary, or comparison,
run a visual impact pass and update the relevant figure if needed.

## Maintenance Checklist

Before publishing changes:

- Preview the affected pages through a local static server.
- Verify relative image and stylesheet paths from each page depth.
- Check desktop and mobile widths for body, table, and image overflow.
- Confirm public images have meaningful `alt` text.
- Scan for local paths, temporary files, private prompts, and tool signatures.
- Sample-check representative GitHub source links.
- For bilingual changes, verify indexes, language switches, canonical links,
  `hreflang` links, and old hash redirects.
- For series changes, verify chapter navigation, overview images, route-map alt
  text, home cards, README entries, and language gateways all describe the same
  sequence.

## Acknowledgements

Thanks to [WineChord/books](https://github.com/WineChord/books) for inspiring
the long-form technical article presentation and source-reading discipline.

## License

This repository is licensed under the [MIT License](./LICENSE).
