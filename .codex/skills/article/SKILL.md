---
name: article
description: "Plan, revise, illustrate, and maintain Rememorio long-form technical articles."
argument-hint: "[article path or topic]"
---

# Rememorio Article

Use this skill for standalone technical articles and article-like chapters in
the Rememorio Blog repository. It keeps structure, source accuracy, image rhythm, and
maintenance behavior consistent.

## Living Exemplar

Treat the Wine & Chord prompt-cache article as the external style and quality
exemplar for Rememorio technical articles:

- Public reference URL:
  `https://www.wineandchord.com/books/prompt-cache/`.
- Source reference repository:
  `https://github.com/WineChord/books`.
- Source artifact reference:
  `docs/public/prompt-cache/index.html`.
- Visual reference pack:
  `docs/public/prompt-cache/assets/*.png`.

Use the Wine & Chord reference for editorial shape, not for subject matter. The
portable pattern is: immediate long-form article entry, a sticky left table of
contents, a restrained paper article shell, dense Chinese technical prose,
source-grounded links in the body, early evidence boundaries, guiding
questions, h2/h3/h4 hierarchy, compact responsive tables, and figures that make
runtime ownership visible without decorative excess.

Do not use a work-in-progress Rememorio article as the style exemplar for
itself. Local article pages are targets being revised; the Wine & Chord article
is the comparative bar for layout, pacing, figure discipline, and public voice.
Do not copy Wine & Chord prose, images, or topic-specific claims into
Rememorio articles.

## Article Contract

An article must be self-contained. A reader should not need any prior chat,
private notes, or hidden prompt context to understand the argument.

Preferred flow:

1. Start with the practical problem and why it matters.
2. Translate the first necessary terms into reader-facing language before using
   dense product, protocol, or source identifiers.
3. Define the mental model before naming internal mechanisms.
4. Introduce provider or platform contracts before source-code consequences.
5. Walk source mechanisms in execution order.
6. Use examples at every boundary where readers are likely to confuse concepts.
7. Surface common misreadings near the mechanism that creates them, then end
   with a comparison table or distilled operational rules.

For a concept-first tutorial aimed at readers without prior vocabulary, teach
before compressing. Use a classroom arc:

```text
familiar failure or practical question
-> one-sentence plain-language definition
-> one continuous example
-> introduce the elements one at a time as the example needs them
-> replay the complete flow
-> only then add formal boundaries, tables, schemas, code, and edge cases
-> recap in language the reader can reuse
```

Do not open by saying that unfamiliar term A is not term B, or that A belongs
to C, before the reader can explain A through an example. Tables compress
relationships that have already been taught; they do not introduce a concept.
JSON, protocol shapes, and state machines make a known mechanism precise; they
do not replace a plain-language first explanation. In the opening portion,
keep the term budget small, let one paragraph carry one main idea, and prefer a
short natural-language task record before a production schema. A fresh reader
should be able to answer “what is it, why do I need it, what parts does it have,
and what happens in order?” before the article advances to implementation
detail.

For source-comparison articles, the default long-form arc is:

```text
practical pressure
-> central thesis
-> minimal vocabulary or reading contract
-> 3 to 5 guiding questions
-> mental model
-> provider/platform contract
-> source execution chain
-> contrast table
-> common misreadings
-> transferable rules
-> references
```

For complex runtime mechanisms, each substantial section should teach through a
reader-facing lifecycle, not through a pile of source facts. Prefer this local
arc:

```text
what problem appears
-> what the runtime changes
-> before/after request or persisted-record shape
-> source functions that implement the transition
-> what breaks without this mechanism
-> cache/context consequence
```

For mechanism-heavy sections, make the default prose path easy to follow before
source names appear. A strong section usually moves through this sequence:

```text
concrete user or runtime situation
-> what the simpler design would wrongly do
-> which responsibility moves to which layer
-> small decision trace or before/after shape
-> verified source path or official contract
-> tradeoff, failure boundary, and handoff to the next mechanism
```

When adjacent sections explain one mechanism family, keep one representative
request, persisted record, tool result, user preference, or runtime object
running across them. Let readers carry the same state through write, update,
retrieval, recovery, and later-observation boundaries. Vary the local details
only when a new example clarifies a new boundary; do not restart with unrelated
examples in every section. Titles should name the reader-facing question or
tradeoff first, then mention the mechanism when it is needed for precision.

Use source names only after the reader has a concrete hook for the concept.
When a mechanism depends on overloaded terms such as `history`, `prompt`,
`context`, `baseline`, `projection`, or `replacement`, define the local usage
before the first dense source paragraph. A compact "terms used in this article"
table is better than repeated parenthetical glosses when several terms will
carry the argument.

For lifecycle sections that explain background work, forks, compaction,
self-evolution, maintenance jobs, or any other side path, make the flow
complete enough for a new reader to replay:

- why the foreground path cannot do this work inline;
- who triggers the side path and which gate or counter must fire;
- what input snapshot the side path sees;
- what it is allowed to read, write, or call;
- what it is explicitly not allowed to touch;
- whether it may do nothing;
- where successful output is stored;
- when later turns can observe the result.

If a section still feels abstract, add a compact step table before diving into
source names. The table should explain the product-level event first, then link
the source mechanism that implements it.

For multi-mechanism sections, do not leave the reader with only a taxonomy.
Convert mechanism lists into pressure-and-invariant tables when useful:

```text
pressure source
-> why the simpler approach fails
-> chosen mechanism
-> invariant protected
-> failure boundary
```

For framework overviews and framework source readings, a requested focus
changes depth, not the reader's need for a system map. Before the focused deep
dive, follow one representative run across the major owners that actually
exist: entry and run loop, model/provider adapter, tools and side effects,
graph/delegation, session/context/memory/artifacts, streaming/UI/remote
protocols, and recovery/observability/evolution. Do not force absent layers or
turn this map into a feature brochure. Select the mechanisms that materially
change ownership, execution, persistence, or the user-visible contract, then
make the requested subsystem deeper than the rest.

Treat dense identifier runs as a comprehension failure, not as proof of source
depth. A paragraph that introduces roughly five or more fields, functions,
types, or modules in sequence should trigger a rewrite: group them under `3` to
`5` reader questions, owners, or lifecycle steps; explain the group through a
small scenario or shape; move exhaustive names to a compact table or
`details`; and place source links after the reader understands why the names
matter.

Do not stop at `what` and `how`. Each major mechanism should naturally reveal
the `why` behind the design: the invariant being protected, the simpler design
that would fail, the cost of the chosen approach, and the boundary where the
claim stops. Integrate this reasoning into the prose; never announce process
phrases such as "using critical thinking" or explain the editorial intent.

For source-heavy runtime articles, keep a small invariant ledger in the
writer's head and let it shape the text:

- What is the owner of this state: model view, UI, durable storage, telemetry,
  provider cache, or resume reconstruction?
- What must remain stable for caching, recovery, or user trust?
- What is allowed to be lossy, summarized, truncated, projected, or regenerated?
- What breaks if the mechanism is removed or moved later in the pipeline?
- Which provider or platform contract forces this design instead of a simpler
  one?

Every substantial API/provider article should also include failure conditions
or counterexamples where useful. Cover the ways the recommended design can
break: minimum thresholds, unstable schemas, dynamic data in a cached prefix,
version drift, TTL or routing behavior, overly broad or overly narrow cache
keys, premature summarization, or recovery boundaries that cannot be replayed.
Do not isolate all failure conditions in a final caveats section. Also embed
the relevant failure mode at the mechanism where it is created, then use the
final caveats section to connect the cases.

For long source articles, prefer ending with a transferable decision table or
rule set. It should map content state to runtime handling and the invariant
protected, so readers can apply the article beyond the named products.

For a new source-reading series, the first public chapter should normally build
the technical route before drilling into a single mechanism. Follow one
representative request, turn, job, or runtime cycle across the major owners:
entry point, request or queue boundary, model-visible view, tool or side-effect
gate, durable record, recovery path, and the later chapters each layer will
own. A narrow mechanism such as context management, prompt caching, tools, or
permissions can be excellent as chapter two or later, but it should not be the
reader's first map unless the whole series is intentionally scoped to only that
mechanism.

When an already-written mechanism chapter would make a stronger later chapter,
preserve it and move the series around it instead of flattening it into an
overview. Renumber titles and metadata, update the series index, language
gateway, blog home card, README entry, header navigation, previous/next links,
overview and chapter figures, alt text, and `og:image` so every public route
teaches the same reading order. If a series overview image is a route map, its
project count, labels, and ordering must match the visible chapter list unless
the prose clearly frames the image as a partial historical snapshot.

When a reader question is narrow or out of order, do not mirror the question as
a random new section. Extract the durable confusion behind it and place the
answer where the article's natural argument needs it. If the answer is useful
but too detailed for the main flow, use an HTML `details` block with a precise
`summary`. Good candidates for `details`: feature gates, version drift,
visible-source limits, edge-case recovery, provider caveats, or "how this
differs from a nearby concept" notes.

When a reader-comprehension problem is raised, do a reader-review pass before
rewriting. The review should read like a target-reader audit, not a copy-edit:
list what a reasonable reader still cannot answer, which terms arrive too
early, which transitions feel abrupt, where an example or contrast is missing,
and what sequence would make the section natural. Also test whether a fresh
reader can replay the mechanism without private context: what entered, who
owned it, which gate fired, which record changed, where it was stored, how it is
retrieved later, why the simpler design would fail, and where the claim stops.
Prioritize gaps that block this replay over sentence-level polish. If the user
explicitly authorizes a subagent or outside reader, use it only for this bounded
feedback. The subagent should not edit files; it should return prioritized
reading obstacles and concrete rewrite suggestions. Keep the final rewrite
owned by the main pass, so the article remains one coherent voice.

## Examples, Code, And Source Shapes

Treat examples as steps in the explanation, not as evidence deposited after a
claim. Establish the reader's question and the relevant boundary in prose,
show the smallest representation that resolves the ambiguity, then explain
what changed and why it matters. Protocol-shaped examples are mandatory when
prose alone leaves too much ambiguity. For API/runtime articles, include small
before/after request or record examples at mechanism boundaries.

When one framework supports multiple API protocols, do not infer identical
behavior from a shared runner or model abstraction. Compare the protocols at
their wire and ownership boundaries: request envelope, history ownership,
session or response identity, tool-call and tool-result shape, streaming event
model, resume/recovery semantics, error handling, and unsupported or translated
parameters. Show the smallest equivalent request for each protocol and state
where adaptation is lossy or where the framework, rather than the provider,
owns compatibility.

- label examples as shape-level when fields are simplified or internal
  normalization is omitted;
- keep them close to real provider or source data structures;
- distinguish public API payloads from runtime-internal history items;
- show only fields needed for the concept;
- avoid invented exact values for hidden provider internals;
- explain what changed immediately after the code block.
- When a JSON or API shape is still abstract, precede it with a short
  plain-language decision trace of `3` to `5` steps that explains what the
  runtime is deciding and what it deliberately leaves undecided.
- For branchy mechanisms such as forks, skip-cache modes, fallback paths, or
  resume variants, use minimal sequence examples (`M1 M2 M3 F1`) or compact
  JSON fragments before adding a new diagram. Add a figure only when the
  ownership boundary, lifecycle, or recovery path remains hard to see.
- For source walkthroughs, include short source-shaped snippets when the prose
  depends on a state transition. Prefer excerpts that show the owner, key
  fields, and handoff point, such as `namespace + key + value`,
  `tool_call_id + result_ref`, `resolved / invalidated / new`, or
  `static profile + dynamic search results`. Link the full source, show only
  fields that carry the mechanism, and explain immediately what changed. Do not
  paste long source blocks as evidence by volume.

Classify every published code block by what the reader should interpret, not
by which highlighter happens to make it colorful:

- Give every public `pre code` block an explicit `language-*` class.
- Use the real language for source code, commands, and syntactically meaningful
  payloads. A simplified example may still use that language when it remains
  valid and is clearly labeled as shape-level.
- Use `language-text` for decision traces, sequence sketches, event timelines,
  tree layouts, illustrative pseudocode, and mixed prose. These blocks need
  code-block spacing and alignment, but syntax colors would invent semantics.
- Never rely on language auto-detection. Short fragments, JSON, shell commands,
  and prose diagrams are easy to misclassify, so detection can make the same
  article look authoritative while teaching the wrong structure.

The public blog has one shared highlighting layer. Reuse
`/assets/code-highlight.css`,
`/assets/vendor/highlightjs/highlight.min.js`, and
`/assets/code-highlight.js` with the correct page-relative paths, each exactly
once. Do not add a CDN dependency, a page-specific token theme, manually
authored `hljs-*` spans, or a second highlighting runtime. Article CSS owns the
code block's background, spacing, typography, border, and overflow; the shared
highlight stylesheet owns token colors. When a shared CSS or JavaScript asset
changes, update its cache-version query consistently on every page that loads
it.

Keep source and publication targets synchronized during every revision. When a
standalone Markdown source and generated public HTML both exist, render the
public page from the source before evaluating the final reader experience.

The prose should read like a finished technical essay, not like an edited chat
transcript. It must not reveal the user's requests, editing plan, model
instructions, review process, local machine details, or any private rationale.
Remove phrases that imply the page was created by prompting an assistant. The
reader should only see the subject, evidence, argument, figures, and links.

Use visual emphasis as part of the article's argument. A long source article
should not read as undifferentiated exposition: core invariants, misconception
corrections, and design tradeoffs can be given quiet emphasis with `strong`, a
small note/callout, or a compact comparison table. Keep this restrained and
reader-facing. Do not turn every paragraph into a card, and do not use emphasis
for decorative excitement or process notes.

For multi-page technical notes, source-reading series, or batch article passes,
add a compact reading contract before changing deep prose:

- Landing pages and first chapters should expose a short reading route that
  tells readers which system layers to visit first.
- Chapters and reference pages should include a brief "Reading Contract" or
  "阅读契约" near the top, after the title and any hero visual.
- The contract should name the page's main question, the owners or boundaries
  to track, and the check a reader should be able to answer afterward.
- Keep these contracts local to the page, source-safe, and localized; do not
  include editing process notes or private rationale.

Never expose process instructions, private prompts, unfinished editorial
markers, or hidden editing rationale in public prose.

## Rememorio Exemplar Traits

Use the exemplar as a practical bar for new or revised articles:

- It opens directly into the article. Avoid a marketing landing page, chapter
  picker, banner-first shell, or other gate that makes readers click before
  reading.
- It starts from the real pressure a practitioner feels, then narrows to one
  source-level thesis.
- It defines a small set of questions that guide the whole article.
- It states the evidence boundary near the opening: official contracts,
  verified source, and bounded inference should be visibly separated without
  exposing process notes.
- It introduces provider or platform contracts before interpreting source code.
- It separates UI view, durable storage, and model-visible/runtime view when
  those surfaces diverge.
- It uses h2 sections as major argument turns, h3 sections as mechanisms, and
  h4 sections for local implementation steps.
- It renders a nested sticky table of contents from h2/h3/h4 headings so the
  article's logic is visible while scrolling.
- It links natural words in the body to GitHub source, official docs, papers,
  specs, or canonical references instead of hiding all sources at the end.
- It keeps code paths public and clickable; it never exposes local absolute
  paths, usernames, private branches, private transcript details, hidden
  prompts, or editing instructions.
- It uses figures as mechanism explanations, not decoration. A good figure
  teaches one ownership boundary, lifecycle, request shape, recovery path, or
  comparison.
- It closes by compressing the system into transferable rules and common
  misreadings.

## Mandatory Figure Rule

MUST: Any illustrated Rememorio article must use
`.codex/skills/image/SKILL.md` for every public figure. The published page must
reference final raster images, usually PNG, stored under the shared or
language-specific asset directory selected by the image skill and referenced
with the correct relative URL for the page.

MUST: Every article change requires a visual impact pass by default, even when
the user asks only for prose, links, source details, section structure, or a
small factual correction. Check whether the change creates a new mechanism,
changes a lifecycle, shifts a boundary, alters a before/after shape, introduces
a new comparison, or makes an existing figure inaccurate. If so, add, replace,
or regenerate the relevant image through the image workflow. If not, keep the
existing figures and note that no image change was needed during handoff.

MUST: Any newly generated or regenerated article image must be saved as a final
PNG under the relevant shared or language-specific article assets directory.
Use `<article>/assets/<figure-name>.png` only for a genuinely language-neutral
figure; use `<article>/<lang>/assets/<figure-name>.png` when reader-facing image
text differs by edition. Update Markdown sources, generated HTML, index pages,
and social metadata that point at the figure. If a future article intentionally
uses a remote image host, record the local source path and remote URL mapping,
but local repository assets are the default for this blog.

MUST NOT: Do not publish hand-authored HTML, inline SVG, Mermaid, canvas, CSS
shape compositions, DOM diagrams, or other code-drawn substitutes as article
figures. They may be used only as private planning scaffolds or temporary
implementation aids, and must be replaced by final raster assets before
publication.

If exact labels, arrows, or source-code identifiers must be corrected after
generation, follow the image workflow: simplify labels, regenerate or use a
built-in edit first, and use deterministic post-processing only for small
exactness fixes that still look integrated. Always export a single final raster
image. Do not let post-processing become an excuse to rebuild the figure as
HTML or SVG.

## Heading Hierarchy

Use a visible hierarchy rather than one long flat sequence.

- `h1`: article title, no number.
- `h2`: major parts. For Chinese articles, prefer Chinese numerals such as
  `一、`, `二、`, `三、` when the article reads like an essay or source
  walkthrough. Use plain numeric `1.`, `2.`, `3.` only when it better matches
  an existing article series.
- `h3`: argument stages inside each major part. Use local numbering such as
  `2.1`, `2.2`, `2.3` when it helps the sidebar reveal the chapter map.
- `h4`: local mechanism steps, source-code subpaths, examples, boundary cases,
  or "why this matters" checkpoints. Add `h4` when a section has multiple
  distinct operations; do not force every paragraph into a heading.

Avoid global section numbering such as `13`, `14`, `15` after a new major part.
If a section title contains code, keep the code term but make the surrounding
phrase plain and readable.

Use the hierarchy to reveal the argument:

- `h2` answers "which large question are we solving?"
- `h3` answers "which mechanism or comparison are we in?"
- `h4` answers "which step inside this mechanism is being inspected?"

For long source-heavy articles, the table of contents should support at least
`h2` through `h4`, preferably as a nested tree rather than a flat list. The TOC
should make the article's logic visible without requiring the reader to scroll
through the body.

## Figure Rhythm

Default image/text ratio for source-heavy technical articles:

- Cover image at the top.
- One overview image in the introduction or first major section.
- One figure per major conceptual turn, usually every `700` to `1100` Chinese
  characters for dense source analysis. Do not let a long mechanism section run
  past roughly `1300` Chinese characters without a figure unless source-shaped
  snippets, before/after records, or compact decision tables already make the
  section visually self-explanatory.
- Add an extra figure only when it prevents a specific misreading: boundary vs
  replacement, UI history vs model view, snapshot vs hot reload, local
  truncation vs semantic compaction.
- Avoid more than two figures without at least one dense explanatory section
  between them.
- For a substantial standalone source article, expect a cover, an early system
  overview, and several mechanism-specific figures. A 4-part source walkthrough
  often needs `7` to `12` figures. Fewer is acceptable only when the article is
  short or already has interactive visuals that do the same teaching work.
- Use figures to serve the argument, not as decoration. Every figure should
  answer a reader confusion that prose alone would make costly.
- Do not add a new figure merely to satisfy rhythm when it would duplicate a
  precise snippet or table, weaken visual consistency, or turn a source
  transition into a vague diagram. Keep the existing figure package when code,
  record shapes, and tables carry the new explanation more accurately.

All new article images must follow the
`.codex/skills/image/SKILL.md` skill.

Before rewriting a page, make a figure plan:

1. List the article's major conceptual turns.
2. Mark which turns already have strong visuals.
3. Add or replace figures where a mechanism, boundary, lifecycle, comparison,
   or recovery path is otherwise invisible.
4. Keep filenames stable only when the semantic role remains stable; rename
   when the figure's teaching role changes.
5. Save new or regenerated final PNGs under the article assets directory and
   update every relative image reference.
6. Update captions, alt text, image `src`, and `og:image` when the cover changes.

For incremental edits, run the same check at smaller scope:

1. Identify the paragraph, section, or source claim being changed.
2. Locate any figure whose labels, arrows, ordering, boundary, or caption now
   touches that claim.
3. Decide whether the existing figure remains accurate, needs a small
   exactness fix, needs a built-in edit, or needs a fresh figure regeneration.
4. If a figure is added or touched, save the final PNG under the article assets
   directory and update the published image reference.
5. Preserve image rhythm; do not add decorative images just because a change was
   made. It is acceptable to leave figures unchanged when the edit is better
   carried by source snippets, before/after records, or compact tables.
6. Include the visual decision in final verification notes: regenerated,
   adjusted, or intentionally unchanged.

## Source Accuracy

When explaining source code:

- Distinguish public API contract from inferred service internals.
- Name files and functions only after verifying them in the repository or in
  official documentation.
- Do not overfit a diagram to implementation details that are not visible.
- If a mechanism exists only through surrounding contracts, say so in prose and
  keep the figure abstract.
- Prefer "model-visible view" and "runtime projection" language when UI,
  storage, and API request history diverge.

Classify every source-level claim before making it sound certain:

- `official docs`: provider, product, API, or standards documentation states
  the behavior.
- `verified source`: the linked public source file, type, function, test, or
  constant directly shows the behavior.
- `surrounding contract inference`: visible call sites and data shapes imply a
  boundary, but the exact implementation is not public.
- `not visible`: the claim would depend on service internals, private state, or
  unpublished code.

Write only `official docs` and `verified source` claims as direct facts. Mark
`surrounding contract inference` as an inference in prose, and keep diagrams
abstract. Do not publish `not visible` claims as implementation facts.

When a mechanism is behind a feature gate, conditional import, build flag, or
dead-code-elimination boundary, verify whether the target module files are
present in the analyzed snapshot before describing internals. If only call
sites, persisted records, or replay hooks are visible, explain the execution
contract and explicitly avoid inventing selector heuristics, scoring policies,
tool prompts, thresholds, or private module behavior. Shape-level examples may
show what the visible contract does, but must not read as the missing
implementation's actual strategy.
When the missing implementation is a compaction, projection, staging, or
summary mechanism, still reconstruct the visible lifecycle from the surrounding
contracts: trigger owner, candidate record shape, summary producer boundary,
staged-vs-committed state, projection point, persistence record, resume replay,
and overflow recovery. Clearly mark selector algorithms, scoring signs,
thresholds, and prompt wording as unavailable unless the source directly shows
them.

When an article discusses hidden or provider-constructed prompts, split the
claim into product prompts, public API request surface, provider-generated
scaffolding derived from request fields, model-internal behavior, and safety or
policy enforcement. Do not infer a cache-visible hidden system prompt from a
product UI prompt, a model self-identification answer, or general safety
behavior. Treat official docs about product prompts, API parameters, tool
scaffolds, token accounting, and prompt caching as separate evidence layers.

It is acceptable to add bounded engineering inference when it helps readers
connect the visible contracts into a plausible lifecycle. Keep that inference
constraint-driven: state which visible owner, recovery, cache, or request-shape
constraints force the likely design, and describe what the mechanism probably
must protect. Do not fabricate exact prompts, schemas, scoring formulas, magic
numbers, private filenames, or vendor intent.

Define internal view terms at first use. Terms such as `API-bound`,
`model-visible view`, `projection`, `cache edit`, and `provider cache view`
should name their owner and lifetime: UI, durable transcript, request payload,
provider-side cache, or replay/recovery state. For cache editing, distinguish an
expected cache-read reduction from a prefix break. The reader should understand
whether a mechanism removes tokens from the model view, preserves local history,
avoids client-side prefix mutation, or merely changes observability counters.
When explaining cache edit, compare it against both direct replacement and
direct deletion of the same block. Call out whether the benefit applies to the
current transition from a warm old prefix, or only to future turns after a new
shorter prefix has been written.
When a cache/projection mechanism uses overloaded terms such as `prefix`,
separate at least three layers before drawing conclusions: request or identity
prefix used for cache lookup, provider-side cached object or processed state,
and the effective model-visible view after edits or projections. If a field
such as `cache_reference`, a beta header, or a cache-control marker participates
in lookup identity, explain whether it was already part of the stable request
discipline or newly introduced in the transition being analyzed. Do not imply
that cache edits preserve every counter or every token span; they may preserve
the reusable identity path while intentionally deleting tokens from the cached
or model-visible view.
For shape-level cache-edit examples, do not imply an exact `cache_edits` block
index unless the source helper or provider contract directly proves that
position. If placement is helper-dependent, link the helper and explain the
placement rule separately from the conceptual request shape.
When using KV-cache or prefix-cache mental models, separate operational
shorthand from mathematical equivalence. Do not claim that a provider-side edit
is token-for-token equivalent to a fresh prefill of the shortened prompt unless
official docs or source prove how suffix states are recomputed, rearranged, or
masked. It is safer to state the visible contract, counters, and source-level
handling, then mark deeper inference as provider-internal.
For cache lookback explanations, distinguish distance to the prior cache
write from distance to the edited or deleted block. Lookback starts at the
current breakpoint and searches for earlier written cache entries; it does not
search outward from the historical block being edited.
When explaining skip-cache-write or fire-and-forget forks, split the request
into shared prefix `S` and fork-only suffix `F`. Make clear that the fork still
reads `S` from cache and still sends `F` as uncached input, but avoids writing a
new `S + F` cache tail that future mainline turns will not resume from.
Also explain what `F` concretely is in the source: a `promptMessages` task
instruction, summary request, side-question wrapper, suggestion prompt, or
other real message that enters the model for this fork. Do not leave it as an
abstract "tail" if the reader needs to know why it exists. State whether that
tail belongs to the parent conversation, the fork transcript, or no transcript
at all, and why the mainline will or will not ever resume from `S + F`.
When a mechanism asks a model to generate a compaction summary, explain the
summary prompt contract at shape level: who triggers it, what messages are sent,
whether tools or thinking are allowed, what output wrapper is expected, which
parts are stripped, and which summary record is installed afterward. Distinguish
the generated summary text from runtime markers such as compact boundaries,
provider payload filters, restored attachments, and post-compact cleanup hooks.
Do not let readers infer that "summary exists" automatically explains selection,
prompting, installation, and resume behavior.

For articles that mix provider APIs and source code, include a quiet evidence
boundary near the opening or before deep source interpretation. It should tell
readers which claims come from official contracts, verified source snapshots, and
bounded engineering inference, without sounding like process notes.
When source links point to a public mirror or reconstructed source snapshot,
say so in the evidence boundary. Do not let mirror-backed links read like
official vendor source authority. Keep mirror repository owners in URLs rather
than visible prose unless the owner itself is part of the technical argument.

## Locale And I18N

Keep the article's prose in the target publication language. For current
Rememorio source-reading articles, Chinese prose is the default unless the
article package already has a separate English source.

For bilingual standalone articles in this static blog, prefer sibling language
pages over in-page machine translation:

- use the repository's established language slugs consistently, commonly
  `<article>/zh/` for Chinese and `<article>/en/` for English in this blog; do
  not introduce a new slug such as `cn/` when the surrounding package already
  uses `zh/`;
- keep `<article>/` as a thin language gateway when both editions exist;
- preserve old root anchors by redirecting root URLs with a hash to the Chinese
  page while keeping the hash;
- add `canonical` and `alternate hreflang` links for `zh-CN`, `en`, and
  `x-default`;
- add visible language switches in each edition's header, showing both
  languages such as `中文 / English` rather than only the current language;
- keep the top navigation, brand link, GitHub link, language switch placement,
  and active states visually consistent across every sibling article page;
- update the blog index with explicit entries or labels for each public
  language entry point.

For an article package or series that already has both Chinese and English
surfaces, treat bilingual delivery as the default unit of work. Creating a new
chapter, changing the reading order, replacing a cover, or revising a shared
mechanism normally means updating both `<article>/zh/...` and
`<article>/en/...` siblings in the same pass, unless the user explicitly asks
for a single-language draft. Do not leave the English side as a route shell,
stale chapter list, missing next/previous target, or untranslated copy of the
Chinese page.

Equivalence is structural as well as textual: titles, metadata, header tabs,
language switches, source lists, figures, alt text, home cards, and route-map
references should describe the same public chapter set in both languages.

Write each language as a finished article in that language. Do not mechanically
translate sentence by sentence when the target language needs different
rhythm, transitions, or terminology. Keep the evidence, figures, source links,
and claims aligned across editions.

Check teaching-device parity, not paragraph-count parity. For every major
section, both editions should preserve the same scenario, mental model,
decision trace or table, source-shaped example, evidence boundary, and final
tradeoff or conclusion when those devices carry the explanation. The prose may
be rewritten for the target language, but one sibling must not collapse into
continuous source-heavy paragraphs or silently drop the example that makes the
other sibling understandable.

Do not publish an English route map, stub, or placeholder when the user asks
for an English edition of an article or series. Either create a finished
sibling article in the target language, or keep the unfinished route
unpublished until it is ready. A bilingual series should feel like one
publication surface with two complete reading paths.

Figures must match the target publication language for reader-facing titles,
invariant ribbons, explanations, generic component names, actions, states, and
callouts. Keep concise English only for exact source identifiers, product or
protocol names, API fields, and other literals that should not change between
editions. A short or familiar English phrase such as `model view`, `snapshot`,
`fork`, or `ledger` is not automatically language-neutral when it functions as
reader-facing narrative on a Chinese page.

When adding a new language edition, apply the image skill's language asset
policy before reusing a raster. A single shared figure is valid only when every
visible label is an exact identifier, product/protocol name, numeral, or symbol.
If any reader-facing label changes with the page language, create a localized
sibling through the image edit workflow, keep the filename and visual concept
aligned, and localize alt text, captions, navigation, metadata, and surrounding
prose at the same time.

When figure text must differ by language, mirror the asset layout in the
language directories and keep filenames aligned. For example, if the article
package uses `zh/assets/` and `en/assets/`, Chinese pages should read from
`zh/assets/` and English pages from `en/assets/` with the same figure names
where the teaching role is the same. Produce the target-language figure by
editing the accepted source-language raster through the image workflow, not by
redesigning the diagram from scratch. The localized figure should look like the
same composition with different text.

Navigation and browser identity are part of the bilingual article, not
surrounding chrome to patch later. After adding or moving any bilingual page,
check the top tabs/actions, brand/home links, shared favicon, `中文 / English`
switch, series index cards, previous/next links, root gateway, and blog home
entry as one graph. Every public route in the package should declare the
project favicon instead of falling back to the browser's generic globe. Each
Chinese page should point to the corresponding English sibling, and each
English page should point back to the corresponding Chinese sibling. Avoid
switches that only point to the language root unless the sibling article
genuinely does not exist.

## Terminology And Link Semantics

Treat terminology and links as part of the article's argument, not as decoration.
The linked target must be at the same abstraction level as the linked words.

- Category terms should link to a neutral category source only when that source
  is actually useful. If no good neutral source exists, leave the category term
  unlinked and link the concrete products or mechanisms nearby instead.
- Product names should link to official overview or product documentation on
  first mention. Link a product name to a narrow feature page only when the
  sentence is specifically about that feature.
- Provider mechanism names should link to the provider's own documentation, not
  to a competing provider or a generic wiki page.
- Source identifiers, file paths, constants, structs, and functions should link
  to verified source URLs, not overview docs.
- Avoid awkward half-translations. If the English term is the industry term used
  by the products or docs, use the English term directly in a Chinese article.
  Add a Chinese gloss only when it improves comprehension, and do not repeat the
  gloss after the first mention.
- Do not mechanically add Chinese parenthetical glosses for mature technical
  terms such as `runtime`, `provider contract`, `owner`, `projection`, or
  `ledger`. Use the English term directly when it is clearer and already
  carries the engineering meaning.
- Keep reader-facing terminology consistent across title, guiding questions,
  headings, alt text, captions, and conclusion. If the article chooses a local
  Chinese term such as `自进化`, use it consistently for the concept; reserve
  source literals such as `self-improvement review` for code, logs, or exact
  product text.
- Avoid stiff imported abstractions when a more natural Chinese phrase carries
  the same meaning. Prefer plain reader-facing wording in headings and
  transitions; keep specialized terms for places where the source or protocol
  actually requires them.
- When a category has competing vendor names, use the smallest stable category
  term in prose and link concrete product names to official docs. Do not
  foreground minor naming differences unless naming, taxonomy, or API surface is
  part of the article's argument. Avoid repetitive parenthetical translations
  such as `代码智能体` unless the article truly needs the gloss.
- Treat a link as a claim about the linked words. If the words name a category,
  avoid linking them to a single product as though that product defines the
  category; instead link nearby product names, official docs, source files, or
  concrete mechanisms at their proper layer.
- Do not link a broad concept to a page that describes a narrower vendor product
  unless the prose clearly says the link is an example, not the definition.
- Before publishing, scan early paragraphs especially carefully: the first
  linked occurrence of a term teaches readers what conceptual layer the article
  is operating on.

## Links And References

Source-heavy articles should be navigable from the prose itself, not only from
a reference list at the end.

- Turn every visible source file path into a clickable GitHub link.
- Link important functions, structs, constants, request fields, and protocol
  terms at their first or most explanatory occurrence.
- Prefer exact GitHub blob URLs with line anchors when the local source has
  been verified.
- Source links must use the canonical repository name, not a temporary local
  directory name or scratch mirror. Verify the GitHub repository before linking.
- If the analysis depends on a fixed source snapshot, use an immutable commit in
  the link target unless the article intentionally tracks live `main`. When the
  source was verified from a local clone, convert explanatory GitHub links to
  the checked commit before publishing; keep the reader-facing text human:
  file paths, function names, and mechanisms matter more than raw commit hashes.
- Do not make `based on owner/repo@sha`, raw commit IDs, local branch names, or
  revision metadata a headline, hero line, caption, navigation label, or other
  prominent reader-facing motif.
- Before publishing, check that representative source URLs return 200 and that
  visible file paths, function names, and line anchors still point to the
  intended code.
- For a batch of fixed-snapshot source links, verify every touched path, symbol,
  and line range against the pinned local checkout, then sample public URLs for
  HTTP 200. A few reachable links do not prove that the remaining anchors still
  explain the words they are attached to.
- Prefer official docs for provider contracts, APIs, SDK behavior, and product
  features.
- Use reputable primary or canonical sources for background concepts: official
  docs, specs, standards, papers, or well-maintained reference pages.
- Keep public prose natural. Add links to existing words rather than inserting
  process notes or citation chatter.
- Keep a bottom reference list, but do not make it the only path to sources.
- If a visible path has moved in the current source tree, correct the path
  before linking; do not point a stale path at a different file.
- Link official product/docs pages where provider behavior matters, such as
  OpenAI, Anthropic, Cursor, MCP, GitHub, or relevant API documentation.
- Link papers, specs, and high-signal engineering posts at the paragraph where
  they support a claim. Prefer primary or canonical sources over summaries.
- Avoid naked URLs in prose. Link the phrase a reader would naturally click.
- Audit generated public HTML or Markdown for leaked local file URLs, local
  absolute paths, temporary paths, private prompts, or internal process language
  before publishing.
- For long-lived source articles, prefer immutable source URLs for verified
  source claims. Moving `main` links are acceptable only when the article
  intentionally tracks live source and the link target is checked during the
  current edit.
- Audit representative GitHub line anchors before publishing: verify that they
  still point at the named symbol, field, or file section.

## Images, Alt Text, And Captions

Images must be useful for readers who see them and readers who rely on text:

- Every public `img` needs informative `alt` text that names the concept, not
  generic text such as "diagram" or "image".
- When a figure is a series overview, route map, comparison grid, or chapter
  map, its alt text should name the same entries, ordering, or scope that the
  visible figure and surrounding prose present. Update alt text whenever the
  figure's inventory changes.
- Use a visible caption when the figure's interpretation is not already clear
  from the surrounding paragraph.
- It is acceptable to use Rememorio style without explicit `figcaption` when
  the immediately adjacent prose explains the figure and the alt text is
  specific.
- The article must not rely on image text alone for source evidence; the prose
  or caption should carry the source links.
- In this static blog, route images with page-relative paths. From an article
  root page use `./assets/name.png`; from generated chapter pages use
  `../assets/name.png` or the equivalent correct relative path. Validate the
  rendered HTML rather than assuming a path by inspection.

## Tables, Code, And Responsive Reading

Use tables for compact comparison only when the rows are easy to scan. For
wide comparison tables in standalone HTML articles:

- Keep the desktop table dense and aligned.
- On mobile, avoid forcing page-level horizontal scrolling.
- Prefer a card-style row layout where each body cell carries the corresponding
  header via `data-label`.
- Generate `data-label` from the table headers in page script rather than
  duplicating labels by hand in prose.
- Keep a long code block's horizontal scrolling inside its own wrapper and
  verify that it does not widen the document. In tight table cells, do not use
  a global no-wrap rule for every inline `code` element. Scope that behavior to
  the table or identifier column that needs atomic tokens, and test its longest
  value at mobile width.
- Do not treat `overflow-x: hidden` as responsive validation. At representative
  desktop and mobile widths, confirm the intended media queries activate,
  compare document `scrollWidth` with `clientWidth`, inspect element bounding
  rectangles, give shrinking grid/flex children `min-width: 0`, keep code
  scrolling local to its wrapper, and verify table/card `data-label` behavior.

## Update Protocol

For any article change:

1. Update the source article first when a source package exists.
2. Render generated HTML from Markdown sources when the article package uses a
   renderer.
3. Keep filenames stable unless the semantic role of a figure changes.
4. Update `og:image` if the cover changes.
5. Check the generated table of contents after heading changes.
6. For series additions, removals, or reordering, verify that overview pages,
   route-map figures, cover images, alt text, chapter numbers, navigation
   tabs, favicon links, home cards, README entries, and language gateways all
   expose the same sequence and publication identity.
7. Verify desktop and mobile widths for overflow, cramped figures, and broken
   code/table scrolling.
8. Keep the README public index current for new public entry points.

Treat the blog home as an explicit publication inventory, not an accidental
sample. Choose one exposure mode per series and keep it consistent:

- In chapter-level mode, give every completed chapter one home post card in
  each published language. After every addition, removal, or reordering,
  compare normalized route sets and counts against the series index instead of
  checking only that one representative card exists. Keep the series index in
  reading order; follow the established home chronology, usually newest first.
- In series-level mode, expose one landing card and no stray individual chapter
  card. Do not silently mix series-level and chapter-level modes.
- Treat topic shortcuts as navigation only. A shortcut to the series index does
  not prove that the article-card inventory is complete.
- When the home hero represents the newest chapter, update its visible cover
  and `og:image` together. If it deliberately represents the stable series
  overview instead, keep both surfaces aligned to that identity.

For language-structure changes, also verify:

- root gateway page, each language page, and the blog index all link to the
  intended targets;
- `hreflang`, canonical URLs, page titles, descriptions, Open Graph metadata,
  the project favicon, and header language switches are present and
  language-correct;
- every article page exposes both language options in the same header style,
  and each switch points to the corresponding sibling page rather than to the
  series root;
- top navigation tabs/actions, active states, previous/next links, series index
  cards, and blog home cards remain visually and semantically consistent across
  Chinese and English pages;
- old anchor links either still resolve or redirect to the new language page
  with the hash preserved;
- relative asset paths are correct from each language subdirectory;
- language-specific asset directories have matching filename sets when
  in-image text differs, and each page's `og:image`, cover, captions, and alt
  text point at the correct language asset;
- source and target language figure contact sheets have been compared when a
  batch of localized raster images is added or regenerated.
- overview and route-map images have been rechecked against both language
  chapter lists, especially after adding one page to an established series.

Before committing public article changes, run the smallest relevant validation
that covers the edit:

- Check asset references exist and every public image has non-empty alt text.
- Search for private leakage: local absolute paths, local file URLs, temporary
  paths, unfinished editorial markers, hidden process language, private request
  text, or private rationale.
- Check generated TOC, active responsive rules, table `data-label` behavior,
  document and body `scrollWidth`, high-risk element bounding rectangles, and
  whether heading hierarchy still reads as a useful route map.
- When emphasis blocks, notes, or highlighted sentences were added, verify that
  they clarify invariants or misreadings and do not create a stack of generic
  callouts.
- If links were added or changed, sample-check official docs and representative
  GitHub line anchors. For fixed source snapshots, scan touched pages for stale
  `blob/main` source links and replace them with the verified commit target.
- If source-shaped snippets, JSON fragments, or API shapes were added, check
  balanced tags or fences and make sure shape-level examples cannot be mistaken
  for exact private internals.
- For every touched page containing code blocks, verify that each `pre code`
  has an explicit semantic language class and that the shared highlight
  stylesheet, runtime, and initializer load exactly once from valid paths. In a
  browser, inspect one source block and one `language-text` block when both are
  present: the source block should receive `hljs` token markup, while the text
  block should receive `nohighlight` and no token spans. Recheck page-level
  overflow and local code scrolling after the runtime has executed.
- When shared highlighting JavaScript changes, run a syntax check. When shared
  highlighting CSS or JavaScript changes, verify that every consuming page
  uses the same new cache version.
- Run repository content/build checks when the touched files participate in a
  renderer or build step.

## Evolution

When future revisions reveal a durable writing or maintenance rule, update this
skill in the same change. The rule must be reusable, verified, and phrased as a
general article practice. Do not record private conversations or transient
editing instructions.
