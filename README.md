# Rememorio Blog

Source for [rememorio.github.io/blog](https://rememorio.github.io/blog/), a
static technical essay site about AI coding agents, runtime design, context
management, tool execution, permissions, and prompt-cache behavior.

The repository is intentionally plain: hand-authored HTML, shared CSS, local
PNG figures, and a small set of writing rules. There is no application
framework, bundler, package manager, or build pipeline between the source files
and GitHub Pages.

## Published Site

| Area | Description | Entry points |
| --- | --- | --- |
| Blog home | Language gateway and bilingual home pages. | [Gateway](https://rememorio.github.io/blog/) / [中文](https://rememorio.github.io/blog/zh/) / [English](https://rememorio.github.io/blog/en/) |
| Claude Code Source Notes | Bilingual source-reading series on the public [Rememorio/claude-code](https://github.com/Rememorio/claude-code) mirror. | [Gateway](https://rememorio.github.io/blog/claude-code/) / [中文](https://rememorio.github.io/blog/claude-code/zh/) / [English](https://rememorio.github.io/blog/claude-code/en/) |
| Codex Source Notes | Bilingual source-reading series on [openai/codex](https://github.com/openai/codex). | [Gateway](https://rememorio.github.io/blog/codex/) / [中文](https://rememorio.github.io/blog/codex/zh/) / [English](https://rememorio.github.io/blog/codex/en/) |
| Hermes Agent | Source-grounded walkthrough of [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent). | [Gateway](https://rememorio.github.io/blog/hermes-agent/) / [中文](https://rememorio.github.io/blog/hermes-agent/zh/) / [English](https://rememorio.github.io/blog/hermes-agent/en/) |
| Mem0 Memory | Bilingual source-grounded article on [mem0ai/mem0](https://github.com/mem0ai/mem0), algorithm evolution, temporal reasoning, and memory decay. | [Gateway](https://rememorio.github.io/blog/mem0-memory/) / [中文](https://rememorio.github.io/blog/mem0-memory/zh/) / [English](https://rememorio.github.io/blog/mem0-memory/en/) |

Use the public site indexes as the canonical article catalog. The README should
stay stable as new essays are added.

## Repository Layout

```text
.
├── index.html                  # Site language gateway
├── zh/                         # Chinese blog home
├── en/                         # English blog home
├── claude-code/                # Claude Code series
│   ├── index.html              # Series language gateway
│   ├── zh/                     # Chinese articles and figures
│   └── en/                     # English articles and figures
├── codex/                      # Codex series
│   ├── index.html              # Series language gateway
│   ├── zh/                     # Chinese articles and figures
│   └── en/                     # English articles and figures
├── hermes-agent/               # Hermes Agent article
│   ├── index.html              # Article language gateway
│   ├── zh/
│   ├── en/
│   ├── styles.css
│   └── assets/
├── mem0-memory/                # Mem0 memory article
│   ├── index.html              # Article language gateway
│   ├── zh/
│   └── en/
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

## Figure Conventions

Public figures are final PNG assets stored under the relevant article's
`assets/` directory. They should explain mechanisms rather than decorate pages.

Figures should:

- use warm paper, restrained colors, and hand-drawn technical diagram styling;
- make ownership boundaries, lifecycles, or state ledgers visible;
- keep labels short and readable at article width;
- avoid misleading arrows, crowded callouts, and overlapping text;
- include exactly one deterministic Rememorio brand treatment;
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

## Acknowledgements

Thanks to [WineChord/books](https://github.com/WineChord/books) for inspiring
the long-form technical article presentation and source-reading discipline.

## License

This repository is licensed under the [MIT License](./LICENSE).
