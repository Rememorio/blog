---
name: image
description: "Generate and maintain Rememorio hand-drawn technical article images with image2/imagegen."
argument-hint: "[article-or-figure topic]"
---

# Rememorio Image2 Hand-Drawn Figures

Use this skill for every image2/imagegen asset created for this repository.
The goal is a consistent Rememorio visual language for technical notes and
articles: precise enough for source-code explanation, warm enough for long-form
reading.

## Living Visual Exemplar

The Wine & Chord prompt-cache article is the external visual style and quality
reference:

- Public article: `https://www.wineandchord.com/books/prompt-cache/`.
- Source repository: `https://github.com/WineChord/books`.
- Reference article artifact: `docs/public/prompt-cache/index.html`.
- Reference image pack: `docs/public/prompt-cache/assets/*.png`.

Before producing any new Rememorio article figure, inspect the Wine & Chord
reference article and a contact sheet of the prompt-cache image pack. Use the
reference for visual discipline: warm paper, restrained blue/green palette,
clean sketched arrows, short labels, visible ownership boundaries, and one
mechanism per figure. Do not copy the Wine & Chord topic, prose, or image
content into Rememorio articles.

Do not use a work-in-progress Rememorio article's current images as the style
authority for revising that same article. Existing local images are candidates
to keep, fix, or replace; the Wine & Chord reference is the comparative bar.
If an image generation tool accepts reference images, use the most relevant
`1` to `3` Wine-style references by role. If the tool does not accept reference
images, explicitly translate the selected references' style and layout roles
into the prompt and deterministic post-processing choices.

Adapt the visual system to Rememorio by keeping a restrained bottom-right
`Rememorio` wordmark treatment and the article's source-accurate labels.

## Non-Negotiables

- The page must reference final raster images, usually PNG. Do not leave source
  diagrams as HTML, Mermaid, or SVG-only illustrations in published articles.
- Every final figure must carry one deterministic `Rememorio` raster brand
  treatment. Do not rely on the image model to redraw or reinterpret the logo.
- Every final generated figure must be saved under the relevant article's
  local `assets/` directory before publication. Published article markup should
  use the correct page-relative path for that static page.
- Do not include production notes, prompts, private instructions, model names,
  or process explanations inside public images.
- Do not let generated text invent source facts. If exact identifiers, arrows,
  or labels matter, compose them deterministically into the final raster image
  after generating the hand-drawn base.
- A figure must faithfully teach the mechanism it represents. It may be lively
  and metaphorical, but the node order, ownership boundary, data flow, and
  failure/recovery path must match the article's verified claims.
- Prefer fewer, source-accurate, readable figures over a large set of weak
  decorative images.

## Visual System

Canvas:

- Default aspect ratio: `16:9`.
- Preferred new working size: `2400 x 1350`; acceptable publication size:
  `1672 x 941` or larger.
- Safe margin: at least `7%` on all sides.
- Background: warm off-white paper, not pure white.
- Texture: subtle paper grain; no noisy parchment, stains, or heavy shadows.

Palette:

- Ink navy: `#102a43`.
- Rememorio teal: `#0f766e`.
- Rememorio blue: `#2d5f8b`.
- Codex green: `#1f7a4d` or `#15803d`.
- Slate gray: `#52606d` / `#64748b`.
- Warm amber: `#b7791f`.
- Soft paper: `#fffaf0` / `#fffffb`.
- Use red only for explicit danger or deletion semantics.

Line and shape:

- Hand-drawn ink outlines with slight wobble, but keep geometry legible.
- Rounded boxes, soft shadows, watercolor fills, and sketched arrows.
- Avoid glossy 3D, neon gradients, plastic UI cards, decorative blobs, and
  overly cute cartoons.
- Arrows must be thin-to-medium, directional, and non-crossing.
- Favor lanes, ledgers, stacks, timelines, and left-to-right pipelines.
- Use Rememorio-style structures when they fit the claim: source-reading route
  maps, layered ownership boards, runtime timelines, durable-ledger views,
  lifecycle loops, tool/file/network side-effect paths, and final comparison
  maps.

Brand mark treatment:

- The preferred Rememorio pictorial source is the repository-root
  `Eyjafjalla.png` image. Use the full character image as the brand signature
  whenever space permits, and compose it deterministically with a small clean
  `Rememorio` wordmark instead of asking the image model to redraw it.
- Use a restrained wordmark: dark ink, warm brown, or muted rose tones that fit
  the figure. Do not use the old yellow/green route-line wordmark treatment as
  the default.
- Do not crop the character into a head-only avatar by default. If the full
  character would be too small inside a dense technical figure, keep the figure
  brand treatment quieter rather than switching to a large avatar badge.
- The final published figure should composite exactly one Rememorio brand
  treatment into the image after generation or deterministic drawing.
- Keep it legible but secondary: roughly `10%` to `18%` of image width, with at
  least `3%` canvas padding from the right and bottom edges.
- During image generation, ask for a clean bottom-right paper area reserved for
  the brand treatment. After generation, overlay the selected PNG composition
  deterministically. If the model generated any competing logo or text there,
  cover it with matching paper texture before compositing the brand treatment.
- Do not use a generic badge, alternate logo, image-model imitation, or
  unrelated icon. Legacy brand PNGs in this skill directory are only fallback
  assets when the article explicitly calls for them.

## Reference Image Roles

Use the Wine & Chord prompt-cache image pack as a role library. Pick references
by teaching role, not by filename similarity:

- Cover or overview: broad system route with one clear thesis.
- Provider or API contract: two or three lanes with stable/dynamic boundaries.
- Pressure ladder: stacked mechanisms ordered by escalation.
- Lifecycle: numbered stages with a visible trigger and terminal state.
- Projection or replacement: before/after views with one ownership boundary.
- Persistence or recovery: durable store, runtime view, and replay path.
- Final map: compact comparison of routes without crowding labels.

A new article about memory should use ledger, snapshot, and before/after roles.
A runtime article should use lifecycle, ownership-board, and recovery roles.
A source comparison should use contract lanes and final route maps.

Reference selection workflow:

1. Name the figure's teaching role: cover, overview, contract comparison,
   pressure ladder, lifecycle, persistence path, projection, replacement
   ledger, or final map.
2. Select the closest `1` to `3` Wine-style references for that role.
3. State what each reference contributes: layout, palette, line treatment,
   density, or brand mark.
4. Do not reuse the reference's subject-specific nodes unless the new article
   is explaining the same mechanism.

## Technical Diagram Method

Before generating, write a compact figure brief:

1. Claim: the one source-level idea this figure must teach.
2. Nodes: the exact entities that may appear.
3. Edges: the exact order or relationship between nodes.
4. Uncertainties: anything that must be described as public API contract rather
   than inferred provider internals.
5. Caption role: what the surrounding prose will explain so the image can stay
   visually clean.
6. References: which Wine-style image roles should guide style and layout.

For runtime and protocol figures, prefer lifecycle or before/after layouts over
generic box clusters. A good figure should make one transition visible:
request before vs after projection, UI history vs model-visible history,
durable record vs resume reconstruction, or provider contract vs client
runtime responsibility. If the article already uses JSON examples for exact
fields, the figure should show ownership, order, and consequence rather than
duplicating the whole payload.

Figures should expose the reason for a mechanism, not just its plumbing. When a
diagram depicts a runtime choice, make the protected invariant visually clear:
stable prefix, lossy projection, durable recovery point, ownership boundary, or
failure mode avoided. Avoid decorative complexity that shows more nodes without
making the tradeoff easier to understand.

When a figure depicts source-code behavior, ensure the surrounding article text
or caption links to the corresponding source files, functions, and official
contracts. The image may carry short labels, but it must not be the only place
where a reader can trace a claim back to code or documentation.

Visible figure labels must follow the same abstraction level as the prose. Use
category labels for categories, product labels for products, and source
identifier labels for exact code concepts. Do not let a broad category such as
`coding agent` visually collapse into one vendor product, and do not add
awkward translation glosses such as `代码智能体` unless they make the figure
easier to understand.

Those traceability links must follow the article skill's source-link contract:
canonical GitHub repository, verified path, and line anchor. Never use a local
workspace directory name as the public evidence URL. If the article uses an
immutable source snapshot, keep raw commit hashes out of figure text.

For source-code diagrams, keep visible labels short:

- Prefer code identifiers only when short, for example `cache_control`,
  `prompt_cache_key`, `compact_boundary`, `replacement_history`.
- Prefer concise English labels and source identifiers inside figures, even for
  Chinese articles. This improves generation reliability and preserves a clean
  path to later English localization.
- Avoid long Chinese sentences inside generated images.
- If Chinese labels are required, add them with deterministic post-processing
  and export a single PNG.
- Use font roles deliberately during deterministic post-processing: reserve
  monospace fonts for short code identifiers, function names, flags, and schema
  fields; use a CJK-capable sans font for Chinese prose labels, captions, and
  callouts. Never place Chinese explanation text in a monospace-only font that
  can render missing-glyph boxes.
- Maximum visible labels per figure: `10`; maximum words per label: `4`.

Prompt pattern:

```text
High-resolution 16:9 hand-drawn technical article infographic for Rememorio.
Use the provided Wine-style article images as style references, especially
[reference role filenames].
Warm off-white paper, precise ink sketch lines, subtle watercolor fills.
Palette: deep navy, Rememorio blue, muted forest green, slate gray, small
amber highlights.
Subject: [one precise source-level idea].
Composition: [nodes and arrows in exact order].
Leave a clean warm-paper safe area in the bottom-right for the Rememorio
wordmark/avatar treatment that will be composited after generation.
No process notes, no prompt references, no meta text.
Use only these short English labels: [label list].
Keep arrows clean and non-crossing; leave generous margins.
Professional editorial diagram, not cartoonish, not glossy 3D.
```

If exact labels matter, ask the image model for the unlabeled or lightly labeled
base, then overlay final text with deterministic composition. Keep final labels
aligned with article prose and source links.

## Post-Processing

Use deterministic composition when accuracy requires it:

- Keep the generated or hand-drawn base as the background or illustration layer.
- Overlay exact labels, node titles, arrows, or callouts with a script or design
  renderer.
- Always overlay one selected Rememorio brand treatment as the final
  bottom-right mark.
- Export one final PNG for publication.
- Keep source files near the article only when they are useful for regeneration.

Mandatory brand overlay protocol:

1. Generate the figure body with a clean bottom-right warm-paper safe area.
2. Build or select one Rememorio brand treatment based on the figure brief,
   article voice, and visual fit. Prefer a deterministic composition from the
   full `Eyjafjalla.png` character plus a small clean wordmark. Record the
   source asset and placement choice so later revisions can reproduce the same
   final figure.
3. Load the selected brand treatment and scale it proportionally to `10%` to
   `18%` of canvas width.
4. Place it in the bottom-right safe area with at least `3%` right and bottom
   padding.
5. If the selected brand treatment background differs from the target paper,
   softly key out the light paper background or blend it into a small
   matching-paper patch. Do not distort the avatar or wordmark.
6. Save the composited result as the only image referenced by the article.

Recommended final asset layout:

```text
<article>/assets/<figure-name>.png
```

When mirroring a standalone source article into generated HTML, keep the same
filename in the article's `assets/` directory and update both Markdown and HTML
references.

Static blog publication protocol:

1. Finish the image generation or deterministic drawing pass, label fixes, and
   selected `Rememorio` brand overlay first.
2. Save the final PNG locally under the article's assets directory so the source
   package remains reproducible.
3. Reference the final PNG with the correct page-relative path: usually
   `./assets/<figure-name>.png` from an article root page and
   `../assets/<figure-name>.png` from chapter pages.
4. If the article source is Markdown and generated HTML is checked in, update
   the source first, then render or mirror the generated page.
5. Update captions, alt text, cover references, and social metadata when the
   image's role changes.
6. Verify the rendered page locally for broken image paths, overflow, and
   mobile readability before considering the figure published.

## Figure Density

Plan figures from the article argument, not from existing assets:

- Every substantial source-heavy article should have a cover plus an early
  overview figure.
- Add figures for provider contracts, runtime ownership boundaries, lifecycle
  transitions, tool/file/network side effects, compaction/recovery paths, and
  final comparison maps.
- A long article should rarely go more than `700` to `1100` Chinese characters
  without a visual anchor unless code/table layout already carries the visual
  load.
- Prefer replacing weak legacy diagrams over preserving them for convenience.
- Keep diagrams varied: do not repeat the same lane layout for every figure.

## Quality Gate

Inspect every generated figure before publishing:

- Text is readable and not misspelled.
- No missing-glyph boxes, clipped words, accidental line breaks inside code
  identifiers, or forced wraps such as splitting `observation` across lines.
- Labels do not escape boxes.
- Arrows do not cross in confusing ways.
- The figure matches the article's source claims.
- The Rememorio mark is one deterministic composited raster treatment,
  proportionally scaled and placed in the bottom-right. No image-model
  imitation, alternate badge, generic icon, distorted crop, recolor, or
  competing generated logo is visible.
- The published article references the correct local relative URL for each
  generated figure, and the referenced PNG exists in the repository.
- No private instruction, prompt, unfinished editorial marker, or process note
  is visible.
- Mobile rendering keeps the image legible at article width.
- The figure fits the Rememorio visual family when viewed beside the
  reference assets.
- The figure is vivid enough to explain the concept visually, but not so
  decorative that it obscures the real mechanism.

Regenerate or post-process if any item fails.

For article batches, create a temporary contact sheet of new figures beside the
Wine-style reference images and inspect the set as one family. The batch should
look systematic, not like unrelated one-off diagrams.

For any figure with deterministic text overlays, also inspect at least the
highest-risk images at full size. Contact sheets are useful for family style,
but they can hide missing glyphs, cramped labels, and bad wrapping.

## Evolution

Update this skill when a repeated image correction becomes a reusable rule.
Keep changes small, general, and independent of any one conversation. Do not
record private prompts, one-off preferences, or unpublished source details.
