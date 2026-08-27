---
name: image
description: "Generate and maintain Rememorio hand-drawn technical article images with imagegen."
---

# Rememorio Imagegen Hand-Drawn Figures

Use this skill for every imagegen asset created for this repository.
The goal is a consistent Rememorio visual language for technical notes and
articles: precise enough for source-code explanation, warm enough for long-form
reading.

## Mandatory Reference-And-Pilot Loop

Style fidelity does not come from mentioning reference filenames in a prompt.
For every new figure, repair, or series refresh, execute this loop in order.
Do not start a batch by generating all figures in parallel.

### 1. Inspect The Pixels And Assign Authority

Open the repository-root `Eyjafjalla.png` at full size for character identity.
Then open the relevant accepted figures themselves, preferably in a contact
sheet and again at full size for the highest-risk character poses. Reading this
file, listing filenames, or relying on memory does not count as inspection.

Use the accepted references as a role hierarchy, not as interchangeable style
samples:

- Paper temperature, palette restraint, and negative space:
  `claude-code/zh/assets/commands-skills-mcp-cover.png`.
- Mechanism flow and scene rhythm:
  `claude-code/zh/assets/agent-fork-cover.png` or
  `claude-code/zh/assets/agent-selection-gates.png`.
- Guide-and-prop integration:
  `claude-code/zh/assets/worker-tool-scope.png`.
- Unusually dense source maps only:
  `claude-code/zh/assets/mcp-tool-pool.png`.
- Source-reading workshop covers and lifecycle routes:
  `deepseek-harness/zh/assets/cordis-self-referential-route.png`.
- Host/client or ownership-boundary workshop scenes:
  `deepseek-harness/zh/assets/cordis-host-client-boundary.png`.

The two Cordis figures are accepted examples of the same Rememorio family, not
a replacement palette: use them when a source-reading mechanism benefits from
physical workbenches, machines, shelves, gates, or cables. The cleaner
`commands-skills-mcp-cover.png` paper still wins any temperature conflict.

Before writing the prompt, record a short private calibration note containing:

- the chosen visual family;
- one paper/color authority;
- one composition authority;
- one character-integration authority;
- three concrete observations visible in those pixels, such as character
  scale, amount of unpainted paper, and how props connect to the mechanism.

Before the first imagegen call, give the user a concise commentary checkpoint
naming the chosen family and the exact references that were opened. This makes
reference inspection observable without asking for approval. If the agent
cannot name the files and concrete pixel observations, it is not ready to
generate.

Use at most `1` to `3` style references, assigning each one role. Do not ask the
model to average several references or blend unrelated series grammars. Do not
promote an existing work-in-progress image into an authority merely because it
is nearby.

### 2. Choose The Visual Family Before The Layout

If an article or series already has an accepted in-scene Eyjafjalla guide, the
new figure belongs to that family. The guide's presence, readable scale, and
useful role are mandatory; a detached chart plus a bottom-right mascot card is
the wrong family.

For a substantial technical or source-reading figure without an established
sibling, prefer the integrated mechanism-scene family: diagram, props, labels,
arrows, and guide form one illustrated workspace. Use the small bottom-right
source-notes card only for a genuinely simple standalone figure where an
in-scene guide would obscure the mechanism. Do not choose the card merely
because the article is new.

The composition must arise from the claim. Naming the right palette while
requesting generic panels, a dashboard grid, or repeated narrow cards still
produces the wrong style.

### 3. Write The Contract, Then Generate One Pilot

Write the compact figure brief from `Technical Diagram Method`, including the
language asset policy and character action. For a batch, also write one shared
style lock. Then generate only one representative source-language pilot,
normally the cover or overview. Do not generate the rest of the batch yet.

Compare the pilot beside its assigned authorities at article width and at full
size. It must pass all three gestalt gates before it can become a sibling
reference:

1. **Scene gate:** the mechanism reads as one workspace, not a flat flowchart,
   dashboard, or row of interchangeable cards with a pasted character.
2. **Character gate:** the guide is recognizably Eyjafjalla, correctly scaled,
   anatomically valid, and participating through a local prop.
3. **Material gate:** paper temperature, ink weight, watercolor restraint,
   density, edges, and shadows belong beside the accepted references.

A pilot that is factually correct but misses any gestalt gate is rejected. Do
not use it as the style source for later figures. Rewrite the spatial
composition and character action, then regenerate; merely adding more style
adjectives to the same dashboard layout is not a repair.

### 4. Lock, Extend, Localize, Publish

After the pilot passes, make it the batch sibling reference while retaining the
designated accepted authorities. Generate or edit later figures against that
lock and stop at the first visible drift rather than allowing drift to
accumulate. Accept every source-language figure before deriving a localized
sibling from its raster. Only after the final family contact sheet and the
full-size quality gate pass may the figures be copied into publication paths.

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
- Every final figure needs one integrated `Rememorio` treatment generated as
  part of the same visual language. Use the in-scene guide for the
  mechanism-scene family and the bottom-right source-notes card only when the
  mandatory loop explicitly classifies the figure as simple-standalone. Do not
  use a pasted-looking logo, generic badge, clover/leaf icon, unrelated mascot,
  or competing generated mark.
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
- For a multi-figure refresh, select one paper target before the first
  generation and keep it for the whole batch. Default to `#fffaf0`, perceived
  as clean warm white rather than cream or yellow. Do not independently sample
  or reinterpret the background color for each figure.
- Use bright red only for explicit danger or deletion semantics. A restrained
  red-brown may appear in title or invariant ribbons.

Batch consistency lock:

- A batch is a set of sibling figures that will be viewed as one article or
  series. Write one shared style-lock paragraph before writing the individual
  figure briefs, and repeat it verbatim in every generation or edit prompt.
- The lock must name one paper/color authority, one line/composition authority,
  the exact paper target, ink color, accent palette, grain strength, edge
  treatment, shadow strength, and character rendering scale. Individual figure
  prompts may vary only the teaching role, mechanism, labels, composition, and
  local prop.
- Default batch lock for this repository: clean warm-white paper near
  `#fffaf0`; no overall amber/yellow wash; deep navy `#102a43` ink; restrained
  teal `#0f766e`, blue `#2d5f8b`, forest `#1f7a4d`, slate `#52606d`, and small
  amber `#b7791f` accents; red-brown only for title/invariant emphasis; subtle
  fine paper grain; light edge roughness; crisp medium-fine outlines; minimal
  soft shadows; generous unpainted breathing room.
- Do not ask imagegen to average several references. Assign each reference one
  explicit role and state that the paper/color authority wins any conflict.
- After the first accepted figure, use it as the sibling batch reference for
  later figures while retaining the designated accepted authorities. If a
  later figure drifts warmer, darker, more saturated, more textured, or more
  shadowed, repair it before continuing the batch; do not let drift accumulate.
- For a bilingual batch, accept the source-language figure first, then derive
  its localized sibling from that accepted raster. The localized sibling must
  inherit the same batch lock and must not become a second independent style
  sample.

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

Brand and character treatment:

- Follow the visual-family decision from the mandatory loop. In a
  mechanism-scene figure, the in-scene Eyjafjalla-like guide is the Rememorio
  treatment; do not add a second logo card or move the guide into a corner. In
  the explicitly chosen simple-standalone family, generate a small bottom-right
  source-notes card as part of the whole illustration, including the guide and
  the word `Rememorio`.
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

Mechanism-scene composition contract:

- Start with a physical metaphor for the claim, then place the exact mechanism
  into it. Prefer one or two dominant stations such as a workbench, machine,
  gate, ledger, shelf, map, or tool crate over one decorative card per node.
  Labels and arrows must connect those stations into a believable workspace.
- A title ribbon, one main mechanism lane, and a bottom invariant ribbon are a
  useful vocabulary, not a mandatory template. Vary the dominant silhouette
  and spatial rhythm across a batch; do not repeat the same grid with different
  nouns.
- Keep the guide secondary to the mechanism but large enough at article width
  to read the curled horns, chestnut hair mass, red eyes, and white/red outfit.
  Place the guide beside a local mechanism prop, not in unused corner space.
- Give the guide one simple action that explains the local mechanism: holding a
  nearby ledger, cable, envelope tray, tool, stop sign, compass, catalog, or
  instruction card. Do not copy a previous prop when it does not fit the claim.
- Keep the action anatomically easy: exactly two arms and two hands, one hand
  connected to each wrist/sleeve, elbows near the torso, and at most one small
  held prop. Highlight distant nodes in the diagram instead of stretching the
  guide across the canvas. If topology fails, simplify the pose and regenerate
  the whole image.
- Name the guide concretely in every prompt: `Eyjafjalla-like chibi guide`,
  `chestnut-brown hair`, `pale segmented curled horns`, `red eyes`, and
  `white/red outfit cues`. Vague requests such as `cute mascot`, `brand avatar`,
  or `little sheep logo` authorize an unwanted redesign.
- Chinese-edition figures use Chinese for reader-facing titles, callouts,
  actions, states, and invariant ribbons. English source identifiers, product
  names, API fields, and protocol literals remain English when exactness
  requires it. English editions preserve the accepted scene and replace only
  language-bearing labels.
- Reject generic dashboards, component-card galleries, dark tech posters,
  isolated character stickers, character sheets, and decorative workspaces
  whose props do not encode the mechanism.

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

Choose the role before the visual metaphor, then follow the reference hierarchy
in the mandatory loop. Borrow composition grammar, density, palette, or guide
integration—not the reference's subject-specific nodes.

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

Simple-standalone prompt skeleton:

Use this only after the mandatory family decision concludes that an in-scene
guide would obscure a genuinely simple figure.

```text
High-resolution 16:9 hand-drawn technical article infographic for Rememorio.
Reference authority: [exact accepted path] controls paper and palette;
[exact accepted path] controls composition. Observed cues to preserve:
[three visible observations from the reference pass].
Clean light warm off-white notebook paper, restrained grain and roughness,
precise ink sketch lines, subtle watercolor fills. No aged parchment.
Subject: [one precise source-level idea].
Composition: [nodes and arrows in exact order].
Integrate a small bottom-right Rememorio source-notes paper card into the same
hand-drawn style, with a tiny Eyjafjalla-like character based on
`Eyjafjalla.png` (warm brown hair, curled horns, red eyes, white/red outfit
cues) and the exact words `Rememorio` and `source notes`.
If the character's hands are visible, show exactly two arms and exactly two
hands, each hand connected to one wrist/sleeve; no extra or floating hand.
Visible labels only: [edition-appropriate label list].
Avoid generic mascots, reading mascots, generic logos, clovers, leaves, pasted
stickers, competing badges, and any brand mark that looks detached from the
illustration.
Keep arrows clean and non-crossing; leave generous margins.
No process notes, prompt references, meta text, watermark, glossy 3D, or neon.
```

Mechanism-scene prompt skeleton:

Use this for substantial technical figures and every established in-scene
series. Fill every bracket with one concrete choice. Do not leave a menu of
possible props in the prompt; that encourages a decorative asset collage.

```text
Use case: infographic-diagram
Asset type: Rememorio source-reading article figure, 16:9 PNG.

Reference authority:
- [exact accepted path] controls paper/color.
- [exact accepted path] controls scene geometry and density.
- project `Eyjafjalla.png` controls character identity.
Observed reference cues: [three concrete pixel observations].
Batch style lock: [paste the shared lock verbatim].

Claim: [one precise source-level mechanism or invariant].
Physical metaphor: [one coherent workspace metaphor].
Scene geometry: [exact left/center/right or foreground/background placement].
Mechanism truth: [exact nodes, ownership boundaries, arrows, and order].
Allowed scene props only: [small closed list chosen for this mechanism].
Title or invariant ribbons: [exact short edition-language text, or none].

Eyjafjalla-like chibi guide: [full-body or upper-body] beside [one local prop]
at [specific location], performing [one simple action] that explains the local
mechanism. Faithful stylized redraw: chestnut-brown rounded hair and loose side
locks, pale segmented curled horns, red eyes, white/red outfit cues, and dark
glove or boot accents. Exactly two arms and exactly two hands, one connected to
each visible wrist/sleeve and arm; elbows near the torso; small proportional
hands; at most one held prop. Emphasize distant nodes in the diagram instead of
making the guide reach across the canvas.

Visible labels only: [short label list].
Clean light warm-off-white paper, restrained grain, crisp medium-fine navy ink,
soft restrained watercolor, minimal shadows, generous unpainted breathing
room, and clean non-crossing arrows. The diagram, props, labels, and guide must
read as one hand-drawn workspace.

Reject: generic dashboard or component-card gallery; flat flowchart with a
pasted character; unrelated decorative shelves or ledgers; repeated narrow
columns; aged yellow parchment; heavy brown shadows; generic mascot, sheep or
goat animal, robot helper, detached badge, missing identity anchors; duplicate,
floating, prop-emerging, or extra-fingered hands; glossy 3D; neon; decorative
blobs; process notes; prompt text; watermark.
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
2. Use the treatment selected by the mandatory visual-family decision: the
   in-scene Eyjafjalla guide for mechanism scenes or the bottom-right
   source-notes card for simple standalone figures. Do not mix both unless the
   accepted source figure already does.
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

Apply these gates in order and stop at the first failure. A later strength does
not compensate for an earlier failure: correct source labels do not rescue the
wrong visual family, and beautiful styling does not rescue incorrect arrows.
Inspect every image at full size and article width; use contact sheets for
family comparison, never as the only anatomy or typography check.

### Gate 0: Process Evidence

- The references were actually opened, and the private calibration note names
  their roles and three visible observations.
- The visual family, language asset policy, figure brief, and batch lock were
  chosen before generation.
- A representative pilot passed beside the accepted authorities before the
  rest of the batch was generated. If no pilot was accepted, stop here.

### Gate 1: Mechanism Truth

- The scope, node inventory, ownership boundaries, order, arrows, and
  failure/recovery paths match verified article claims.
- Arrows do not imply immediate execution where the relationship is periodic,
  optional, delayed, or storage-mediated. Foreground, background, and
  maintenance paths remain distinct when their triggers differ.
- Props expose the mechanism's reason or protected invariant rather than adding
  decoration. The figure remains understandable without invented source facts.

### Gate 2: Visual Family And Composition

- Beside the assigned authorities, the figure matches paper temperature, ink
  weight, watercolor restraint, negative space, edge treatment, shadow
  strength, and overall density.
- A mechanism-scene figure reads as one workspace with a dominant physical
  metaphor. The guide, props, arrows, and labels belong to that workspace.
- Reject a flat flowchart, generic dashboard, repeated component-card grid,
  aged parchment, decorative asset collage, or pasted character even if its
  palette and individual nodes are correct.
- Sibling figures share the batch lock but vary their dominant silhouette and
  spatial composition according to the claim.

### Gate 3: Character Identity, Integration, And Anatomy

- Compare the guide full size with `Eyjafjalla.png` and the selected integration
  authority. It preserves chestnut hair volume, pale segmented curled horns,
  red eyes, and white/red outfit cues while sharing the scene's line, paper,
  palette, and shadows.
- In an in-scene family the guide is present, large enough to identify at
  article width, beside a useful local prop, and performing a mechanism-related
  action. It is not a card, corner logo, sticker, sheep/goat animal, robot,
  generic horned mascot, or unrelated anime character.
- Anatomy passes the binary topology trace: exactly two arms and no more than
  two hands; every visible hand connects through one wrist/sleeve to one arm;
  no duplicate, floating, hidden third, prop-emerging, extra-fingered, detached,
  or impossibly reaching limb.

### Gate 4: Text And Localization

- All text is readable and correctly spelled, with no missing glyphs, clipped
  words, escaped labels, bad code-identifier wraps, overlaps, or ambiguous
  hand-drawn lettering. Shorten or relayout before shrinking text.
- Reader-facing titles, ribbons, actions, states, and generic labels match the
  edition language; exact identifiers, products, API fields, and protocol
  literals remain source-accurate.
- Source and localized siblings match side by side in canvas, scene geometry,
  character identity and pose, arrows, props, palette, density, and teaching
  role. Only language-bearing text may materially change.

### Gate 5: Publication

- The final raster exists under the selected shared or language-specific asset
  directory, and page markup, social metadata, gateways, and home cards resolve
  to the correct edition's file.
- No prompt, private instruction, model name, process note, watermark, or
  unfinished editorial marker is visible.
- Local desktop and mobile rendering show no broken path, crop, overflow, or
  unreadable figure at article width.

If a gate fails, regenerate or use a built-in whole-image edit so the figure
stays visually unified. Change the composition or simplify the character pose
when those are the cause; do not attempt to fix a structural failure with more
style adjectives or local patch assembly.

## Evolution

Update this skill when a repeated image correction becomes a reusable rule.
Keep changes small, general, and independent of any one conversation. Do not
record private prompts, one-off preferences, or unpublished source details.
