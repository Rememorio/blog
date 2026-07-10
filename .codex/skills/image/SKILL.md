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

The Rememorio Blog repository's accepted article figures are the visual style
and quality reference. Before producing a new figure, inspect the repository
root `Eyjafjalla.png` character reference and a small contact sheet of accepted
local figures that match the target series and teaching role.

For a series that has already established an in-scene Eyjafjalla guide, inspect
one accepted sibling figure from that series before writing the prompt and
treat the guide's presence, scale, and role as part of the figure contract.

Default accepted mechanism-scene references in this repository:

- `claude-code/zh/assets/agent-fork-cover.png`
- `claude-code/zh/assets/worker-tool-scope.png`
- `claude-code/zh/assets/agent-selection-gates.png`
- `claude-code/zh/assets/mcp-tool-pool.png`
- `claude-code/zh/assets/commands-skills-mcp-cover.png`

Unless the user names another accepted visual family, the named Claude Code
figures are the default composition and style authority. Do not blend several
series' visual grammars in one prompt. Other repository figures may help verify
a mechanism, but they are not default visual calibration merely because they
exist nearby.

Use the Claude Code references for a clean, light technical-notebook page:
warm off-white paper, restrained grain and roughness, crisp ink, soft watercolor
fills, generous breathing room, scene-specific composition, short labels,
visible ownership boundaries, and an Eyjafjalla-like guide integrated into the
mechanism. Avoid dark aged parchment, tea-stained or treasure-map texture,
heavy brown shadows, repeated narrow dashboard/card columns, and oversized
decorative ledgers or shelves. The references are calibration for paper, line
weight, density, character integration, and mechanism clarity, not clip-art to
copy.

Do not use a work-in-progress article's current images as the style authority
for revising that same article. Existing local images are candidates to keep,
fix, or replace; accepted local references from this skill are the comparative
bar. If a figure generation tool accepts reference images, use the most
relevant `1` to `3` accepted local references by role. If the tool does not
accept reference images, explicitly translate the selected references' style
and layout roles into the prompt and later quality checks.

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
replace the character reference or the accepted local figures as the
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
- Treat the built-in imagegen output as the visual source of truth. When a
  generated figure has a visual defect such as bad anatomy, distorted brand
  treatment, broken layout, or drifting labels, fix it with built-in
  regeneration or built-in editing. Do not repair published article figures by
  stitching screenshots, painting over limbs, compositing character patches, or
  rebuilding the scene with local drawing code. Local scripts may decode,
  copy, resize, format-convert, or make contact sheets for inspection; they
  should not create the visual content of the final figure.
- Every final standalone figure needs one integrated `Rememorio` source-notes
  treatment generated as part of the same visual language as the figure. Use
  the default bottom-right card for new one-off figures. For an established
  series, preserve the accepted series treatment, including an Eyjafjalla-like
  guide character drawn into the mechanism scene. Do not use a pasted-looking
  logo, generic badge, clover/leaf icon, unrelated mascot, or competing
  generated mark.
- In an established in-scene guide series, a mechanism diagram that omits the
  Eyjafjalla-like guide is not acceptable even if the mechanism itself is
  otherwise correct. Regenerate or edit the whole figure so the guide appears
  as a recognizable participant in the workspace, not as an optional decoration
  to drop under layout pressure.
- Character anatomy is a binary gate for every in-scene guide and every visible
  brand-card pose. Prompt for exactly two arms and exactly two hands, one hand
  connected to each wrist/sleeve; use occlusion only when the pose still
  accounts unambiguously for both arms. Before acceptance, trace every visible
  hand back through one wrist and one arm. Reject duplicate, floating, hidden
  third, prop-emerging, or extra-fingered hands. Prefer one small held prop in
  total and never assign more than one prop to a hand. If the topology fails,
  simplify the pose and regenerate or edit the whole image; never patch a limb
  locally.
- Before generating or accepting a bilingual figure, choose its language asset
  policy. A figure is `shared-neutral` only when every visible label is an exact
  code identifier, product or protocol name, numeral, or symbol that should be
  identical in both editions. If it contains a reader-facing title, invariant
  ribbon, explanation, generic system label, action, or state that changes with
  the page language, it requires `localized-siblings`. Short or readable English
  is not automatically language-neutral.
- Every final generated figure must be saved under the relevant article asset
  directory before publication: a shared `assets/` directory for a genuinely
  `shared-neutral` figure, or the matching `<lang>/assets/` directory for
  `localized-siblings`. Published article markup, home cards, and social
  metadata must use the correct language-specific page-relative path.
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
- Texture: subtle paper grain; no noisy parchment, tea-stained antique scroll
  texture, stains, or heavy shadows. Keep the paper closer to the accepted
  local figures' cleaner off-white notebook pages than to an aged treasure map.

Palette:

- Ink navy: `#102a43`.
- Rememorio teal: `#0f766e`.
- Rememorio blue: `#2d5f8b`.
- Muted forest green: `#1f7a4d` or `#15803d`.
- Slate gray: `#52606d` / `#64748b`.
- Warm amber: `#b7791f`.
- Soft paper: `#fffaf0` / `#fffffb`.
- Use bright red only for explicit danger or deletion semantics. A restrained
  red-brown may appear in title or invariant ribbons.

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
  maps. For runtime source-reading series, prefer a single paper-desk scene
  where the diagram, props, shelves, ledgers, ribbons, and guide character all
  belong to one illustrated workspace instead of a detached chart plus a logo.

Brand mark treatment:

- For new standalone figures, generate the Rememorio treatment as part of the
  whole figure, usually as a small bottom-right paper note or source-notes card
  that includes a tiny Eyjafjalla-like character and the word `Rememorio`.
  For a runtime source-reading series with accepted in-scene character
  treatment, keep that treatment instead of adding a second card or moving the
  character to the corner. In these series the character is the Rememorio
  treatment: a small Eyjafjalla-like guide participating in the mechanism scene,
  not an extra brand logo.
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
  The intended read is the familiar Rememorio little-sheep guide: horned,
  warm-haired, red-eyed, notebook-sized, and source-aware. It should not become
  a goat logo, generic horned mascot, robot helper, sheep animal, or unrelated
  anime character.
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
  `claude-code/zh/assets/agent-selection-gates.png`,
  `claude-code/zh/assets/mcp-tool-pool.png`,
  `claude-code/zh/assets/commands-skills-mcp-cover.png`, and other figures the
  user explicitly accepts as sibling references. Use them as composition and
  integration references, while `Eyjafjalla.png` remains the identity
  reference. Do not regress this treatment into a generic chart plus a
  bottom-right Rememorio card.
- Use only the named accepted figures in this section as default calibration
  unless the user explicitly identifies newer accepted examples. Do not promote
  experimental, disliked, or merely existing images into style authority just
  because they are nearby in the repository.
- The desired look is a clean, warm hand-drawn technical notebook: light
  off-white paper, restrained grain, slightly rough edges, crisp ink outlines,
  soft watercolor fills, and generous negative space. Cards, ribbons, shelves,
  tools, ledgers, gates, arrows, and desk props should form one
  mechanism-specific scene. Do not default to an aged parchment dashboard or a
  repeated row of narrow cards; the composition must arise from the mechanism.
- The recurring scene vocabulary may include a red-brown title ribbon, one
  clear mechanism lane, a bottom invariant strip, an Eyjafjalla-like guide near
  a local prop, and mechanism-specific shelves, ledgers, tool crates,
  envelopes, compasses, gates, funnels, or workbenches. This is a visual
  vocabulary, not a fixed template: vary the spatial composition to fit the
  claim and do not repeat the same card row across a batch.
- The guide is usually a visible full-body or upper-body chibi participant in
  the workspace, not a tiny corner avatar. Keep it secondary to the mechanism,
  but large enough that the horns, chestnut hair mass, red eyes, and white/red
  outfit cues remain readable at article width.
- The character should be useful to the figure: standing beside a local ledger,
  holding a notebook, carrying envelopes, sitting by a tool crate, holding a
  small stop sign, compass, magnifier, wrench, catalog, tag, or instruction
  card, or visually anchoring the runtime boundary. Add props only when the
  mechanism calls for them. Do not cargo-cult a prop from a previous figure,
  and do not repeat the same pose across an entire batch when a mechanism has a
  more natural local prop.
- Design every pose to pass the character anatomy gate. Avoid poses that
  require a chibi arm to reach across the canvas, a hand to touch distant
  cards, or a long pointer to be held at full extension; those often produce
  oversized hands, stretched arms, detached wrists, or missing limbs. Prefer a
  short prop held near the body, a notebook in both hands, a nearby tool, or a red
  highlight/attention mark on the target node. If the mechanism needs emphasis
  far from the guide character, mark the node in the diagram instead of
  stretching the character toward it.
- In every in-scene prompt, repeat the exact topology: exactly two arms and two
  hands, each hand connected to one visible wrist/sleeve and one arm, elbows
  near the torso, small proportional hands, and no duplicate or floating hand.
  If imagegen keeps failing, simplify the pose and regenerate the whole figure.
- For stable results, name the character and pose concretely in every prompt:
  `Eyjafjalla-like chibi guide`, `chestnut-brown hair`, `pale segmented curled
  horns`, `red eyes`, `white/red outfit cues`, and a local prop held close to
  the body. Avoid vague phrases such as `cute mascot`, `brand avatar`,
  `anime helper`, `little sheep logo`, or `Rememorio logo`; they invite the
  model to freely redesign the character.
- Chinese-edition figures must use Chinese for reader-facing titles, callouts,
  actions, states, and invariant ribbons. Keep short English source identifiers,
  product names, API fields, and protocol literals only when translation would
  make them less exact. English-edition figures should keep the same composition
  and replace only those language-bearing labels.
- Do not drift into a generic mascot, dark tech poster, glossy UI mock,
  isolated avatar logo, or decorative character sheet. The figure is still a
  source-code mechanism diagram.

## Reference Image Roles

Use accepted local Rememorio figures as a role library. Pick references by
teaching role, not by filename similarity or article title:

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

For a series overview or final map, the visual inventory must match the public
series inventory. If a chapter is added, removed, renamed, or reordered, update
the route-map figure, labels, alt text, and social image together, or keep the
old figure only if it remains explicitly framed as a partial snapshot.

Reference selection workflow:

1. Name the figure's teaching role: cover, overview, contract comparison,
   pressure ladder, lifecycle, persistence path, projection, replacement
   ledger, or final map.
2. Select the closest `1` to `3` accepted local references for that role.
3. State what each reference contributes: layout, palette, line treatment,
   density, or brand mark.
4. Do not reuse the reference's subject-specific nodes unless the new article
   is explaining the same mechanism.

## Technical Diagram Method

Before generating, write a compact figure brief:

1. Claim: the one source-level idea this figure must teach.
2. Scope: declare `full-map` or `focus-map`. A `full-map` inventory must match
   the public article or series. A `focus-map` may omit modules only when its
   title, alt text, and opening prose clearly frame the partial view.
3. Nodes: the exact entities that may appear.
4. Edges: the exact order or relationship between nodes.
5. Uncertainties: anything that must be described as public API contract rather
   than inferred provider internals.
6. Caption role: what the surrounding prose will explain so the image can stay
   visually clean.
7. References: which accepted local image roles should guide style and layout.
   If the figure includes or revises a Rememorio brand treatment, separately
   name the Eyjafjalla character reference and any accepted in-repo
   source-notes card used only for brand-card integration.
8. Character contract: say whether this figure belongs to an established
   in-scene guide series. If yes, the guide is mandatory and must remain large
   enough to identify the horns, chestnut hair mass, red eyes, and white/red
   outfit cues at article width.
9. Character action: when a guide character appears, describe a pose that can
   be drawn with normal proportions, and describe how distant nodes are
   emphasized without forcing the character to reach across the diagram.
   Include the local prop and body constraints, for example `holding a small
   ledger close to the body`, `exactly two arms and two connected hands`,
   `small proportional hands`, and `no long pointing arm`.

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
- Prefer concise English only for exact source identifiers, product names, API
  fields, and protocol literals that remain unchanged across editions. A short
  English heading, generic noun, action, state, explanation, or invariant is
  still reader-facing language and must match the publication language. For
  Chinese articles, use short Chinese headings and callouts with English code
  identifiers rather than over-Englishing the diagram.
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
- When a diagram is likely to appear in multiple language editions, first reduce
  its label set, then classify what remains. Share the raster only when all
  labels are exact identifiers, product/protocol names, numerals, or symbols.
  If any title, ribbon, callout, action, state, or generic system label needs
  translation, create localized sibling rasters from the accepted source image;
  localized alt text and prose do not substitute for localizing visible
  reader-facing image text.

New standalone figure prompt pattern:

```text
High-resolution 16:9 hand-drawn technical article infographic for Rememorio.
Use the selected Claude Code figures, or a user-approved series family, as
style references: [reference filenames and what each contributes].
Clean light warm off-white notebook paper, restrained grain and roughness,
precise ink sketch lines, subtle watercolor fills. No aged parchment.
Palette: deep navy, Rememorio blue, muted forest green, slate gray, small
amber highlights.
Subject: [one precise source-level idea].
Composition: [nodes and arrows in exact order].
Integrate a small bottom-right Rememorio source-notes paper card into the same
hand-drawn style, with a tiny Eyjafjalla-like character based on
`Eyjafjalla.png` (warm brown hair, curled horns, red eyes, white/red outfit
cues) and the exact words `Rememorio` and `source notes`.
If the character's hands are visible, show exactly two arms and exactly two
hands, each hand connected to one wrist/sleeve; no extra or floating hand.
No process notes, no prompt references, no meta text.
Use only these short English labels: [label list].
Avoid generic mascots, reading mascots, generic logos, clovers, leaves, pasted
stickers, competing badges, and any brand mark that looks detached from the
illustration.
Keep arrows clean and non-crossing; leave generous margins.
Professional editorial diagram, not cartoonish, not glossy 3D.
```

Eyjafjalla mechanism-scene prompt pattern:

Use this when the article series has adopted the in-scene little-sheep guide
treatment. Fill every bracket before generation; omitting the character action
or negative constraints is a common cause of style drift.

```text
Use case: infographic-diagram
Asset type: Rememorio technical article figure, 16:9 PNG.
Create a high-resolution hand-drawn technical notebook infographic in the
established Rememorio Eyjafjalla mechanism-scene style.

Style references: use the selected Claude Code mechanism-scene figures, or a
user-approved series family, as composition references and the project
`Eyjafjalla.png` as the identity reference. Clean light warm off-white notebook
paper, restrained grain and edge roughness, crisp ink sketch lines, soft
watercolor fills, clean non-crossing arrows, and mechanism-specific props
integrated into one illustrated workspace. No aged parchment or repeated
dashboard/card-column layout.

Subject: [one source-level mechanism or invariant].
Composition: top title ribbon: "[short title]"; center mechanism lane:
[exact nodes and arrows]; side or rejection lane: [optional]; bottom paper
ribbon: "[one short invariant]".

Eyjafjalla-like guide: place a recognizable full-body or upper-body chibi guide
near [local prop] in [lower-left/lower-center/lower-right/side margin],
integrated into the same paper scene. Faithful stylized redraw of the canonical
character: chestnut-brown rounded hair with loose side locks, pale segmented
curled horns, red eyes, white/red outfit cues, dark glove or boot accents.
The guide holds [one small local prop] close to the body; exactly two arms and
exactly two hands, each hand connected one-to-one to a visible wrist/sleeve and
arm; elbows near the torso; small proportional hands; no long pointing arm,
duplicate hand, floating hand, hidden third hand, detached wrist, missing limb,
or extra finger.

Visible labels only: [short labels].
Avoid: flat flowchart with pasted character, generic Rememorio logo/card,
sheep or goat animal mascot, robot helper, unrelated anime mascot, detached
badge, missing curled horns, missing red eyes, missing chestnut hair mass,
guide not participating in the mechanism, glossy 3D, neon gradients, decorative
blobs, process notes, prompt text, watermark.
```

For a series that already has an accepted in-scene Eyjafjalla treatment, replace
the bottom-right-card line with this direction:

```text
Use case: infographic-diagram
Asset type: Rememorio source-reading article figure, 16:9 PNG.
High-resolution hand-drawn technical notebook infographic in the established
Rememorio mechanism-scene style.
Style: clean light warm off-white notebook paper with restrained grain and edge
roughness, crisp ink sketch lines, soft watercolor fills, clean arrows, and
small desk props, ledgers, shelves, gates, tools, envelopes, funnels,
workbenches, or maps chosen to match the mechanism. Use cards and ribbons only
when the mechanism benefits from them. No aged parchment or repeated
dashboard/card-column layout.
Character: preserve the established in-scene Rememorio treatment: an
Eyjafjalla-like chibi guide drawn into the mechanism scene, faithful to the
project character reference: chestnut-brown hair with rounded volume and loose
side locks, pale segmented curled horns, red eyes, white/red outfit cues, dark
glove or boot accents. The guide should read as the familiar Rememorio
little-sheep guide, not a generic mascot, robot, sheep animal, goat logo,
detached badge, or unrelated anime character.
Character action: [place the guide near a local prop that belongs to this
mechanism, for example a ledger, notebook, envelope tray, tool crate, compass,
small stop sign, magnifier, wrench, catalog, tag, instruction card, or map].
Show exactly two arms and exactly two hands, each hand connected one-to-one to
a visible wrist/sleeve and arm; keep elbows near the torso and hands small and
proportional. No long pointing arm, duplicate or floating hand, hidden third
hand, detached wrist, missing limb, or extra finger. If the target node is far
away, highlight it in the diagram instead of making the character reach across
the canvas.
Subject: [one precise source-level idea].
Composition: [title ribbon; exact nodes and arrows; side lanes; bottom
invariant ribbon; where the guide and local prop sit].
Visible labels only: [short label list].
No process notes, no prompt references, no meta text, no watermark.
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
Preserve the exact limb count, hand-to-wrist connections, pose, and held props.
Only replace visible <source-language> labels with concise <target-language>
labels from this list: [label map].
Keep code identifiers, product names, arrows, numbering, and all non-language
diagram structure unchanged.
Translate every reader-facing title, invariant ribbon, callout, action, state,
and generic system label in the map. Do not leave narrative text untranslated
merely because it is short or familiar to technical readers.
Do not redesign, crop, simplify, add a new Rememorio card, remove the
Eyjafjalla-like guide, or change the figure's teaching role.
```

If exact labels matter, simplify the label list first and iterate with a
targeted whole-image generation or built-in edit. Use deterministic overlays
only for small exactness fixes where the result still looks like a single
finished raster image.

## Figure Localization And Language Editions

Before deciding whether an accepted figure can be reused, classify every
visible label as either `identifier/product/protocol` or `reader-facing
narrative`. The latter includes titles, invariant ribbons, explanations,
generic component names, actions, states, and audience-facing callouts. Reuse a
single raster only when the second class is empty. If any reader-facing label
changes with the publication language, treat the language variant as an edit
target, not as a fresh illustration brief. The source-language PNG is the
authority for layout, character placement, colors, arrows, and visual density.
The target-language figure should read as the same image in another language.

When in-image text differs by language:

- keep mirrored language asset directories according to the article package,
  for example `<article>/<lang>/assets/` or the repository's established
  equivalent;
- keep the same filename set across language directories unless a figure is
  intentionally language-specific;
- generate or edit the target-language image from the accepted source-language
  raster, replacing text only;
- keep exact English source identifiers, product names, API fields, and protocol
  literals unchanged when they are semantically identical across editions;
- update each language page's figure paths, `og:image`, gateway or home-card
  cover, and any other public image reference so it resolves to that language's
  raster. Do not keep a Chinese page pointed at an English sibling merely
  because the filename is the same;
- avoid programmatically rebuilding the diagram from scratch just to translate
  labels, because it usually drifts in character, layout, and visual style.
- preserve accepted character anatomy during localization. If the localized
  edit changes a hand, arm, face, horns, or pose while translating labels, treat
  the localized image as failed and rerun the built-in text-localization edit
  with stricter preservation constraints. Do not paste a character region from
  one language image into another to repair the drift.

After localizing a batch, create a source/target contact sheet and inspect it
as one family. Then open high-risk figures full size, especially covers and any
figure with dense CJK or English labels. Fix a localized image when the
Eyjafjalla-like guide changes identity, a hand gains extra fingers or limbs,
the layout shifts, or the result feels like a sibling concept instead of the
same figure.

For established little-sheep mechanism-scene series, prefer generating the
source language first, accepting it only after a full-size anatomy and style
check, and then creating the target-language version from that accepted visual
target. The target-language image may differ in labels and bottom invariant
wording, but should preserve the same title-ribbon placement, diagram density,
guide pose, local prop, shelves, ledgers, arrows, and paper texture. If the
target language produces a different guide, a new pose, or a redesigned
workspace, treat it as a failed localization, not as an acceptable variant.

## Capturing Built-In Imagegen Outputs

The built-in imagegen tool may show the image in the Codex UI before the final
workspace path is obvious. For project-bound assets, always persist the actual
imagegen PNG into the article assets directory; do not screenshot the preview
or substitute a locally rebuilt diagram.

Preferred capture sequence:

1. Add a unique internal asset marker to the prompt, for example
   `Internal asset id: article-figure-language-date-attempt. Do not render this
   id.` Keep it portable and free of local absolute paths.
2. After generation, first look under `$CODEX_HOME/generated_images/` for the
   matching recent imagegen output.
3. If the path is not obvious, locate the latest built-in generation event by
   parsing `$CODEX_HOME/sessions/**/*.jsonl` as JSON and matching the internal
   asset marker in `payload.revised_prompt` for records whose
   `payload.type` is `image_generation_call`.
4. Decode that event's `payload.result` base64 field directly to the target
   PNG. Do not print the base64 line to the terminal, and do not use broad
   text searches that dump megabytes of image data.
5. Open the decoded image at full size before copying or overwriting the
   article asset. Confirm it is the intended imagegen figure, not a screenshot,
   contact sheet, browser capture, or unrelated cached image.

Avoid scavenging arbitrary `data:image` blobs across all sessions. If several
candidates exist, select by the unique internal asset marker and timestamp, then
visually verify. This keeps the workflow portable across machines while
preventing accidental replacement with stale or unrelated images.

## Post-Processing

Use post-processing sparingly. The preferred final figure is one coherent image
whose text, arrows, and Rememorio treatment were generated or edited together.
Local processing is appropriate for format conversion, downscaling, light
cropping, or a small exactness correction that does not make the figure look
assembled from layers.

- Keep the generated figure as the primary visual source.
- Prefer regenerating or built-in editing when labels, arrows, or the brand
  treatment drift.
- For character, brand, and layout defects, do not use local compositing,
  inpainting, masking, or patch assembly as the repair path. Whole-image
  generation or a built-in edit keeps line weight, paper texture, shadows,
  anatomy, and text treatment coherent. If a repeated edit makes the character
  worse, change the composition or pose and regenerate the figure.
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
shared-neutral:
  <article>/assets/<figure-name>.png

localized-siblings:
  <article>/zh/assets/<figure-name>.png
  <article>/en/assets/<figure-name>.png
```

When mirroring a standalone source article into generated HTML, keep the same
filename in the selected shared or language-specific `assets/` directory and
update both Markdown and HTML references.

Static blog publication protocol:

1. Finish the built-in image generation or edit pass, label fixes, and
   integrated `Rememorio` treatment first.
2. Save the final PNG locally under the asset directory selected by the
   language policy so the source package remains reproducible. Use a shared
   article `assets/` directory only for `shared-neutral` figures; otherwise save
   matching filenames under each language's `assets/` directory.
3. Reference the final PNG with the correct page-relative path for that package.
   Verify the resolved path from the published page instead of assuming that a
   shared `../assets/` path is suitable for both languages.
4. If the article source is Markdown and generated HTML is checked in, update
   the source first, then render or mirror the generated page.
5. Update captions, alt text, cover references, and social metadata when the
   image's role changes.
6. For series overview images, verify that project counts, labels, ordering,
   alt text, and surrounding route lists describe the same public series.
7. Verify the rendered page locally for broken image paths, overflow, and
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

Inspect every generated figure at full size and at its intended article width
before publishing. A contact sheet alone is not an anatomy or readability pass:

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
- For the established little-sheep mechanism-scene style, the figure passes
  only if the guide is part of the mechanism workspace: placed near a local
  prop, visually secondary to the diagram but clearly present, and consistent
  with the title ribbon, bottom invariant ribbon, desk props, shelves, ledgers,
  arrows, and paper texture. Reject images where the guide becomes a corner
  logo, a standalone character sticker, a generic mascot, a sheep animal, a
  robot helper, or an unrelated horned character.
- In an established in-scene guide series, reject any generated figure where
  the guide is missing entirely, cropped away, too small to identify, or
  replaced by a generic `Rememorio` card. Presence of the character is a style
  invariant, not an optional enhancement.
- Reject and regenerate any established little-sheep mechanism-scene figure
  that reads as a flat flowchart with a pasted character, generic
  Rememorio/logo card instead of an in-scene guide, sheep or goat animal
  mascot, unrelated anime mascot, detached badge, missing curled horns, missing
  red eyes, missing chestnut hair mass, or a guide that does not participate in
  the mechanism.
- Character anatomy passes the binary topology gate: exactly two arms and no
  more than two hands, with every visible hand traceable through one wrist or
  sleeve to one arm. There is no duplicate, floating, hidden third, or
  prop-emerging hand, no extra finger, and no prop that forces impossible
  reach. If the pose fails, redesign it so the character stays local and the
  diagram node carries the visual emphasis.
- The guide's action matches the mechanism rather than merely copying an old
  prop: ledgers for context/history, envelopes for events, tool crates or
  wrenches for tool runtime, stop signs or gates for permissions, funnels for
  projection, compasses or maps for recovery, notebooks or instruction cards
  for context assembly. The prop should stay near the body; distant emphasis
  belongs in the diagram.
- Compare the brand card or in-scene guide at full size against the project
  character reference and at least one accepted in-repo figure with the desired
  treatment. The comparison should confirm both identity fidelity and visual
  integration; if it only passes one of those two, regenerate or edit the whole
  brand area.
- For localized figures, compare source and target language versions side by
  side. They should match in layout, character identity, visual density, and
  teaching role; only the language-bearing text should materially change. Also
  verify that every reader-facing title, ribbon, callout, action, state, and
  generic label matches the page language, while exact identifiers and product
  names remain source-accurate.
- For `localized-siblings`, verify that the Chinese and English pages, social
  metadata, gateways, and home cards resolve to their own language asset paths.
  A visually correct localized PNG is not published if the page still points
  at the other edition's file.
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
accepted local reference images and inspect the set as one family. The batch
should look systematic, not like unrelated one-off diagrams. Also open every
figure with a visible character at full size; contact sheets can hide extra
hands, disconnected wrists, and small identity drift.

For any figure with deterministic text overlays, also inspect at least the
highest-risk images at full size. Contact sheets are useful for family style,
but they can hide missing glyphs, cramped labels, and bad wrapping.

## Evolution

Update this skill when a repeated image correction becomes a reusable rule.
Keep changes small, general, and independent of any one conversation. Do not
record private prompts, one-off preferences, or unpublished source details.
