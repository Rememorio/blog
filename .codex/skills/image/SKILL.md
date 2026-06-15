---
name: image
description: "Generate and maintain Rememorio hand-drawn technical article images with imagegen."
argument-hint: "[article-or-figure topic]"
---

# Rememorio Imagegen Hand-Drawn Figures

Use this skill for every imagegen asset created for this repository.
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
into the prompt and later quality checks.

Adapt the visual system to Rememorio by keeping a restrained, integrated
bottom-right `Rememorio` source-notes treatment and the article's
source-accurate labels.

## Non-Negotiables

- The page must reference final raster images, usually PNG. Do not leave source
  diagrams as HTML, Mermaid, or SVG-only illustrations in published articles.
- Use the installed system `imagegen` skill's built-in image generation path by
  default. Do not use local CLI/API workflows unless the user explicitly asks
  for CLI, API, model, or local-key control.
- Every final figure must carry one integrated `Rememorio` source-notes brand
  treatment generated as part of the same visual language as the figure. Do not
  use a pasted-looking logo, generic badge, clover/leaf icon, unrelated mascot,
  or competing generated mark.
- Every final generated figure must be saved under the relevant article's
  local `assets/` directory before publication. Published article markup should
  use the correct page-relative path for that static page.
- Do not include production notes, prompts, private instructions, model names,
  or process explanations inside public images.
- Do not let generated text invent source facts. If exact identifiers, arrows,
  or labels matter, first simplify the label set and regenerate or edit the
  whole figure so text, arrows, and brand remain visually unified.
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

- Generate the Rememorio treatment as part of the whole figure by default,
  usually as a small bottom-right paper note or source-notes card that includes
  a tiny Eyjafjalla-like character and the word `Rememorio`.
- The repository-root `Eyjafjalla.png` image may be used as a visual reference
  for the character's broad look, but the final figure should feel like one
  unified hand-drawn illustration, not a base diagram with a later pasted
  sticker.
- Keep the brand legible but secondary: roughly `10%` to `18%` of image width,
  with at least `3%` canvas padding from the right and bottom edges.
- Ask for the brand area to be naturally integrated into the paper texture,
  line weight, shadows, and palette of the figure. The note should look drawn
  with the same hand as the diagram.
- Do not use a generic badge, alternate logo, image-model imitation of an
  unrelated mark, clover/leaf icon, route-line logo, or large avatar badge.
- Deterministic compositing with the exact `Eyjafjalla.png` asset is an
  exception, not the default. Use it only when the user explicitly requires
  pixel-level brand fidelity or asks to use the exact source image. Before
  choosing that path, note that it can look less integrated than a whole-image
  generation.

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

For side paths such as background review, maintenance jobs, curator passes,
subagents, asynchronous forks, or delayed writes, make sequence semantics
unambiguous:

- draw the foreground path as the main lane;
- draw per-turn background work as a separate lane entered only after the
  trigger point;
- draw periodic or idle maintenance as another lane with its own schedule or
  idle gate;
- do not use a direct arrow when the relationship is accumulation, later
  visibility, or shared storage rather than immediate execution;
- use dotted or muted connectors only for loose dependencies, and label them
  with short phrases such as `after answer`, `future turns`, or
  `skills accumulate`.

If a figure is meant to correct a misconception, design the layout so the wrong
interpretation is hard to read from the arrows. For example, a periodic
curator should not appear as the next step after every single-turn patch.

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
- If Chinese labels are required, keep them short and prefer whole-image
  generation or a built-in edit pass so the lettering stays integrated with the
  figure. Use deterministic text only for small exactness fixes that cannot be
  resolved cleanly through regeneration.
- When deterministic text is truly necessary, use font roles deliberately:
  reserve monospace fonts for short code identifiers, function names, flags,
  and schema fields; use a CJK-capable sans font for Chinese prose labels,
  captions, and callouts. Never place Chinese explanation text in a
  monospace-only font that can render missing-glyph boxes.
- Maximum visible labels per figure: `10`; maximum words per label: `4`.
- When a diagram is likely to be reused across language editions, prefer short
  English labels and move localized explanation into alt text, captions, and
  surrounding prose. Do not regenerate a figure solely to translate labels if
  the existing labels are readable and language-neutral.

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
Integrate a small bottom-right Rememorio source-notes paper card into the same
hand-drawn style, with a tiny Eyjafjalla-like character and the exact word
`Rememorio`.
No process notes, no prompt references, no meta text.
Use only these short English labels: [label list].
Avoid generic logos, clovers, leaves, pasted stickers, competing badges, and
any brand mark that looks detached from the illustration.
Keep arrows clean and non-crossing; leave generous margins.
Professional editorial diagram, not cartoonish, not glossy 3D.
```

If exact labels matter, simplify the label list first and iterate with a
targeted whole-image generation or built-in edit. Use deterministic overlays
only for small exactness fixes where the result still looks like a single
finished raster image.

## Post-Processing

Use post-processing sparingly. The preferred final figure is one coherent image
whose text, arrows, and Rememorio treatment were generated or edited together.
Local processing is appropriate for format conversion, downscaling, light
cropping, or a small exactness correction that does not make the figure look
assembled from layers.

- Keep the generated figure as the primary visual source.
- Prefer regenerating or built-in editing when labels, arrows, or the brand
  treatment drift.
- Use deterministic labels, arrows, or callouts only when exact source
  identifiers require it and the correction remains visually integrated.
- Use deterministic brand compositing only when the user explicitly requires
  the exact source image or pixel-level brand fidelity.
- Export one final PNG for publication and keep source files near the article
  only when they are useful for regeneration.

When using deterministic text overlays, fit labels by changing the wording and
layout before shrinking text aggressively. Prefer shorter labels, wider boxes,
separate lanes, or fewer callouts over tiny typography. After overlaying text,
inspect the final raster at full size, not only as a contact sheet.

Integrated brand protocol:

1. Include the Rememorio treatment in the figure brief before generation.
2. Place it as a small bottom-right source-notes card, roughly `10%` to `18%`
   of canvas width, with at least `3%` right and bottom padding.
3. Keep the character, wordmark, card paper, line weight, shadows, and palette
   consistent with the rest of the figure.
4. If the brand area looks like a pasted sticker, generic icon, clover/leaf
   badge, unrelated mascot, or detached logo, regenerate or use a built-in edit
   for the whole figure area rather than layering a patch on top.
5. Save the final integrated image as the only PNG referenced by the article.

Recommended final asset layout:

```text
<article>/assets/<figure-name>.png
```

When mirroring a standalone source article into generated HTML, keep the same
filename in the article's `assets/` directory and update both Markdown and HTML
references.

Static blog publication protocol:

1. Finish the built-in image generation or edit pass, label fixes, and
   integrated `Rememorio` treatment first.
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
- Labels, arrows, lane captions, and brand marks do not overlap.
- Arrows do not cross in confusing ways, point to the wrong node, or imply a
  sequence that the article says is only periodic, optional, or delayed.
- Separate foreground, background, and maintenance paths visually when their
  triggers and permissions differ.
- The figure matches the article's source claims.
- The Rememorio treatment is integrated into the bottom-right of the figure and
  looks drawn with the same paper, line weight, palette, and shadows. No pasted
  sticker, alternate badge, generic icon, clover/leaf mark, distorted crop,
  recolor, or competing generated logo is visible.
- The published article references the correct local relative URL for each
  generated figure, and the referenced PNG exists in the repository.
- No private instruction, prompt, unfinished editorial marker, or process note
  is visible.
- Mobile rendering keeps the image legible at article width.
- The figure fits the Rememorio visual family when viewed beside the
  reference assets.
- The figure is vivid enough to explain the concept visually, but not so
  decorative that it obscures the real mechanism.

If any item fails, prefer regeneration or a built-in edit that keeps the whole
figure visually unified. Use local post-processing only for small exactness or
format fixes.

For figures that use hand-drawn or script-like fonts, perform an extra
readability pass. If the font makes letters ambiguous, pushes text into header
bands, or creates awkward baselines, switch to a clearer hand-drawn font or
reduce the label. A hand-drawn feel is secondary to legibility and correct
mechanism flow.

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
