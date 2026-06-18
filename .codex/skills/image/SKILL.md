---
name: image
description: "Generate and maintain Rememorio hand-drawn technical article images with imagegen."
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
`Rememorio` source-notes treatment and the article's source-accurate labels.
The default treatment is a small bottom-right paper card, but a series may
establish a stronger in-scene guide character or source-notes motif. Once that
treatment has been accepted, carry it forward instead of forcing every later
figure back into the default card layout.

For Rememorio brand fidelity, treat the repository-root `Eyjafjalla.png` as
the canonical character reference for every Rememorio character treatment:
bottom-right card, in-scene guide, or small source-notes motif. Accepted
in-repo figures may be used to calibrate placement, paper-card or in-scene
integration, line weight, and `Rememorio / source notes` feel; they must not
replace the character reference or the Wine & Chord reference as the overall
article-figure style authority.
Before generating or repairing a figure, do a short character/style pass:
inspect `Eyjafjalla.png` for identity and inspect one or two accepted in-repo
figures for how the character is integrated into the paper, labels, arrows, and
mechanism scene. The finished treatment should read as a faithful stylized
redraw of the same character, not as a new logo concept.

## Non-Negotiables

- The page must reference final raster images, usually PNG. Do not leave source
  diagrams as HTML, Mermaid, or SVG-only illustrations in published articles.
- Use the installed system `imagegen` skill's built-in image generation path by
  default. Do not use local CLI/API workflows unless the user explicitly asks
  for CLI, API, model, or local-key control.
- Every final standalone figure needs one integrated `Rememorio` source-notes
  treatment generated as part of the same visual language as the figure. Use
  the default bottom-right card for new one-off figures. For an established
  series, preserve the accepted series treatment, including an Eyjafjalla-like
  guide character drawn into the mechanism scene. Do not use a pasted-looking
  logo, generic badge, clover/leaf icon, unrelated mascot, or competing
  generated mark.
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

- For new standalone figures, generate the Rememorio treatment as part of the
  whole figure, usually as a small bottom-right paper note or source-notes card
  that includes a tiny Eyjafjalla-like character and the word `Rememorio`.
  For a series with accepted in-scene character treatment, keep that treatment
  instead of adding a second card or moving the character to the corner.
- Resolve the canonical character reference from the current project, commonly
  repository-root `Eyjafjalla.png` in this blog. If the file is not available,
  ask for the project-provided character reference rather than inventing a new
  logo. Accepted in-repo figures are style calibration for placement and line
  integration, not replacements for the character reference.
- Preserve Eyjafjalla's identity anchors even after stylization: small chibi
  girl proportions, warm chestnut-brown hair with rounded volume and loose
  side locks, segmented pale curled horns, red eyes, small red hair beads or
  ties where visible, white/red outfit cues, and darker glove or boot accents.
  At small sizes, prioritize the silhouette, horns, hair mass, and red eyes.
- Redraw those traits into the figure's hand-drawn paper style; do not paste,
  crop, trace, recolor, or mechanically composite the source image unless the
  user explicitly asks for exact pixel-level fidelity.
- Keep the brand legible but secondary. A default card is roughly `10%` to
  `18%` of image width with at least `3%` canvas padding from the right and
  bottom edges. An in-scene guide character may be larger when it teaches the
  mechanism, but it must not crowd labels, arrows, or the article's evidence.
- Ask for the brand area or guide character to be naturally integrated into the
  paper texture, line weight, shadows, and palette of the figure. It should
  look drawn with the same hand as the diagram.
- When revising a figure whose mechanism is otherwise acceptable but the brand
  card drifted, use the current figure as the edit target and regenerate or
  edit the whole image area so the brand remains integrated. Use an accepted
  in-repo Rememorio figure as a brand-card reference when helpful, but keep the
  article's existing nodes, labels, and arrows as the mechanism source.
- Do not use a generic badge, alternate logo, image-model imitation of an
  unrelated mark, generic reading mascot, clover/leaf icon, route-line logo, or
  large avatar badge. A cute character is not sufficient; it must still read as
  a faithful Eyjafjalla-like redraw in the current figure style.
- Do not accept a brand card merely because it says `Rememorio`. The small
  character should preserve the canonical cues from `Eyjafjalla.png` in the new
  hand-drawn style: warm brown hair volume, curled light horns, red eyes, and
  white/red outfit hints. If those cues are missing, treat the figure as a
  failed brand treatment even when the mechanism diagram is otherwise usable.
- Deterministic compositing with the exact `Eyjafjalla.png` asset is an
  exception, not the default. Use it only when the user explicitly requires
  pixel-level brand fidelity or asks to use the exact source image. Before
  choosing that path, note that it can look less integrated than a whole-image
  generation.

Established Eyjafjalla mechanism-scene style:

- When a series has accepted figures where the Eyjafjalla-like character
  participates in the mechanism scene, treat those figures as the local style
  authority for that series. In this repository, examples include
  `claude-code/zh/assets/agent-fork-cover.png`,
  `claude-code/zh/assets/worker-tool-scope.png`,
  `claude-code/zh/assets/agent-selection-gates.png`, and other accepted
  sibling figures when present.
- The desired look is a warm hand-drawn technical notebook: parchment or
  off-white paper, slightly rough edges, ink outlines, soft watercolor fills,
  compact cards, ribbons, shelves, tools, ledgers, gates, arrows, and small
  source-note scraps that feel drawn into one scene.
- The character should be useful to the figure: pointing at a gate, standing by
  a ledger, holding a tool, or visually anchoring the runtime boundary. Add
  props such as magnifiers, stop signs, scales, boxes, or gears only when the
  mechanism calls for them. Do not cargo-cult a prop from a previous figure.
- Chinese-edition figures may use Chinese-majority labels with short English
  source identifiers. English-edition figures should keep the same composition
  and replace only the language-bearing labels.
- Do not drift into a generic mascot, dark tech poster, glossy UI mock,
  isolated avatar logo, or decorative character sheet. The figure is still a
  source-code mechanism diagram.

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
   If the figure includes or revises a Rememorio brand treatment, separately
   name the Eyjafjalla character reference and any accepted in-repo
   source-notes card used only for brand-card integration.

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
- Prefer concise English labels and source identifiers when they are
  language-neutral. For Chinese articles whose accepted series style uses
  Chinese-majority figure text, use short Chinese headings and callouts with
  English code identifiers rather than over-Englishing the diagram.
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

New standalone figure prompt pattern:

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
hand-drawn style, with a tiny Eyjafjalla-like character based on
`Eyjafjalla.png` (warm brown hair, curled horns, red eyes, white/red outfit
cues) and the exact words `Rememorio` and `source notes`.
No process notes, no prompt references, no meta text.
Use only these short English labels: [label list].
Avoid generic mascots, reading mascots, generic logos, clovers, leaves, pasted
stickers, competing badges, and any brand mark that looks detached from the
illustration.
Keep arrows clean and non-crossing; leave generous margins.
Professional editorial diagram, not cartoonish, not glossy 3D.
```

For a series that already has an accepted in-scene Eyjafjalla treatment, replace
the bottom-right-card line with this direction:

```text
Preserve the established series treatment: an Eyjafjalla-like chibi guide
drawn into the mechanism scene, faithful to the project character reference
(chestnut-brown hair, segmented curled horns, red eyes, red/white outfit cues)
and integrated with the same paper, ink, watercolor, shadows, and line weight
as the diagram. Do not add a separate bottom-right source-notes card unless the
source figure already has one.
Use Chinese-majority labels for a Chinese edition when that is the established
series style, with English retained for source identifiers and product names.
Use the same layout for English editions and localize only the language-bearing
labels.
```

Language-edition edit prompt pattern:

```text
Use case: text-localization.
Edit/reference <source-image-path> as the exact visual target.
Create the <target-language> version of this same figure.
Preserve the original composition, canvas size, paper texture, character,
boxes, icons, arrows, palette, shadows, and hand-drawn style.
Only replace visible <source-language> labels with concise <target-language>
labels from this list: [label map].
Keep code identifiers, product names, arrows, numbering, and all non-language
diagram structure unchanged.
Do not redesign, crop, simplify, add a new Rememorio card, remove the
Eyjafjalla-like guide, or change the figure's teaching role.
```

If exact labels matter, simplify the label list first and iterate with a
targeted whole-image generation or built-in edit. Use deterministic overlays
only for small exactness fixes where the result still looks like a single
finished raster image.

## Figure Localization And Language Editions

Treat a language-only variant of an accepted figure as an edit target, not as a
fresh illustration brief. The source-language PNG is the authority for layout,
character placement, colors, arrows, and visual density. The target-language
figure should read as the same image in another language.

When in-image text differs by language:

- keep mirrored language asset directories according to the article package,
  for example `<article>/<lang>/assets/` or the repository's established
  equivalent;
- keep the same filename set across language directories unless a figure is
  intentionally language-specific;
- generate or edit the target-language image from the accepted source-language
  raster, replacing text only;
- keep concise English source identifiers unchanged when they are already
  language-neutral;
- avoid programmatically rebuilding the diagram from scratch just to translate
  labels, because it usually drifts in character, layout, and visual style.

After localizing a batch, create a source/target contact sheet and inspect it
as one family. Then open high-risk figures full size, especially covers and any
figure with dense CJK or English labels. Fix a localized image when the
Eyjafjalla-like guide changes identity, a hand gains extra fingers or limbs,
the layout shifts, or the result feels like a sibling concept instead of the
same figure.

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

1. Include the Rememorio treatment in the figure brief before generation,
   naming the project character reference and, when useful, one accepted
   in-repo figure as the integration reference.
2. Choose the treatment deliberately: default bottom-right source-notes card
   for standalone figures, or the accepted in-scene Eyjafjalla guide for a
   series that already uses one. Do not mix both unless the source figure
   already does.
3. Derive the character from the reference: chestnut-brown hair, curled
   segmented horns, red eyes, and white/red outfit hints. Redraw those traits
   in the figure's sketch style instead of compositing the original asset.
4. Keep the character, wordmark or note if present, paper, line weight,
   shadows, and palette consistent with the rest of the figure.
5. If the character or brand area looks like a pasted sticker, generic icon,
   clover/leaf badge, unrelated mascot, generic reading mascot, detached logo,
   or a different character with only superficial horns, regenerate or use a
   built-in edit for the whole figure area rather than layering a patch on top.
6. Save the final integrated image as the only PNG referenced by the article.

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
  or into the accepted in-scene series treatment, and looks drawn with the same
  paper, line weight, palette, and shadows. The character visibly follows the
  project reference's broad traits: chestnut-brown hair, curled segmented
  horns, red eyes, and white/red outfit hints. No pasted sticker, alternate
  badge, generic icon, generic reading mascot, clover/leaf mark, distorted
  crop, recolor, or competing generated logo is visible.
- Compare the brand card or in-scene guide at full size against the project
  character reference and at least one accepted in-repo figure with the desired
  treatment. The comparison should confirm both identity fidelity and visual
  integration; if it only passes one of those two, regenerate or edit the whole
  brand area.
- For localized figures, compare source and target language versions side by
  side. They should match in layout, character identity, visual density, and
  teaching role; only the language-bearing text should materially change.
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
