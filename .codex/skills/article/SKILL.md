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
2. Define the mental model before naming internal mechanisms.
3. Introduce provider or platform contracts before source-code consequences.
4. Walk source mechanisms in execution order.
5. Use examples at every boundary where readers are likely to confuse concepts.
6. End with a comparison table or distilled operational rules.

For source-comparison articles, the default long-form arc is:

```text
practical pressure
-> central thesis
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

For multi-mechanism sections, do not leave the reader with only a taxonomy.
Convert mechanism lists into pressure-and-invariant tables when useful:

```text
pressure source
-> why the simpler approach fails
-> chosen mechanism
-> invariant protected
-> failure boundary
```

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

When a reader question is narrow or out of order, do not mirror the question as
a random new section. Extract the durable confusion behind it and place the
answer where the article's natural argument needs it. If the answer is useful
but too detailed for the main flow, use an HTML `details` block with a precise
`summary`. Good candidates for `details`: feature gates, version drift,
visible-source limits, edge-case recovery, provider caveats, or "how this
differs from a nearby concept" notes.

Protocol-shaped examples are mandatory when prose alone leaves too much
ambiguity. For API/runtime articles, include small before/after request or
record examples at mechanism boundaries:

- label examples as shape-level when fields are simplified or internal
  normalization is omitted;
- keep them close to real provider or source data structures;
- distinguish public API payloads from runtime-internal history items;
- show only fields needed for the concept;
- avoid invented exact values for hidden provider internals;
- explain what changed immediately after the code block.
- For branchy mechanisms such as forks, skip-cache modes, fallback paths, or
  resume variants, use minimal sequence examples (`M1 M2 M3 F1`) or compact
  JSON fragments before adding a new diagram. Add a figure only when the
  ownership boundary, lifecycle, or recovery path remains hard to see.

Keep source and publication targets synchronized during every revision. When a
standalone Markdown source and generated public HTML both exist, render the
public page from the source before evaluating the final reader experience.

The prose should read like a finished technical essay, not like an edited chat
transcript. It must not reveal the user's requests, editing plan, model
instructions, review process, local machine details, or any private rationale.
Remove phrases that imply the page was created by prompting an assistant. The
reader should only see the subject, evidence, argument, figures, and links.

For multi-page technical notes or batch article passes, add a compact reading
contract before changing deep prose:

- Landing pages should expose a short reading route that tells readers which
  system layers to visit first.
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
reference final raster images, usually PNG, stored under the article's local
`assets/` directory and referenced with the correct relative URL for the page.

MUST: Every article change requires a visual impact pass by default, even when
the user asks only for prose, links, source details, section structure, or a
small factual correction. Check whether the change creates a new mechanism,
changes a lifecycle, shifts a boundary, alters a before/after shape, introduces
a new comparison, or makes an existing figure inaccurate. If so, add, replace,
or regenerate the relevant image through the image workflow. If not, keep the
existing figures and note that no image change was needed during handoff.

MUST: Any newly generated or regenerated article image must be saved as a final
PNG under the relevant article assets directory, for example
`<article>/assets/<figure-name>.png`. Update Markdown sources, generated HTML,
index pages, and social metadata that point at the figure. If a future article
intentionally uses a remote image host, record the local source path and remote
URL mapping, but local repository assets are the default for this blog.

MUST NOT: Do not publish hand-authored HTML, inline SVG, Mermaid, canvas, CSS
shape compositions, DOM diagrams, or other code-drawn substitutes as article
figures. They may be used only as private planning scaffolds or temporary
implementation aids, and must be replaced by final raster assets before
publication.

If exact labels, arrows, or source-code identifiers must be corrected after
generation, do it only as part of the image workflow's deterministic
post-processing step, and still export a single final raster image. Do not let
post-processing become an excuse to rebuild the figure as HTML or SVG.

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
  past roughly `1300` Chinese characters without a figure unless the section is
  intentionally code/table-heavy and visually self-explanatory.
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
3. Decide whether the existing figure remains accurate, needs a deterministic
   label/overlay fix, or needs a fresh figure regeneration.
4. If a figure is added or touched, save the final PNG under the article assets
   directory and update the published image reference.
5. Preserve image rhythm; do not add decorative images just because a change was
   made.
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

Figures may use concise English labels, source identifiers, and product terms
even inside a Chinese article. This keeps diagrams legible, reduces generated
text errors, and makes later English localization cheaper. Explain the meaning
in surrounding Chinese prose when an English label might be ambiguous.

Do not force Chinese labels into generated figures when the same idea can be
carried by short English terms such as `model view`, `snapshot`, `fork`, or
`ledger`. If a future English version is added, keep heading slugs, asset
filenames, source links, and figure concepts stable so the translation can share
the same evidence and visual package.

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
  the link target or record the snapshot internally, but keep the reader-facing
  text human: file paths, function names, and mechanisms matter more than raw
  commit hashes.
- Do not make `based on owner/repo@sha`, raw commit IDs, local branch names, or
  revision metadata a headline, hero line, caption, navigation label, or other
  prominent reader-facing motif.
- Before publishing, check that representative source URLs return 200 and that
  visible file paths, function names, and line anchors still point to the
  intended code.
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
- For long-lived source articles, prefer immutable source URLs or record the
  source snapshot being explained. Moving `main` links are acceptable only when
  the article intentionally tracks live source and the link target is checked
  during the current edit.
- Audit representative GitHub line anchors before publishing: verify that they
  still point at the named symbol, field, or file section.

## Images, Alt Text, And Captions

Images must be useful for readers who see them and readers who rely on text:

- Every public `img` needs informative `alt` text that names the concept, not
  generic text such as "diagram" or "image".
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

## Tables And Responsive Reading

Use tables for compact comparison only when the rows are easy to scan. For
wide comparison tables in standalone HTML articles:

- Keep the desktop table dense and aligned.
- On mobile, avoid forcing page-level horizontal scrolling.
- Prefer a card-style row layout where each body cell carries the corresponding
  header via `data-label`.
- Generate `data-label` from the table headers in page script rather than
  duplicating labels by hand in prose.
- Validate the mobile viewport for body scroll width, heading overflow, image
  overflow, and table/card overflow.

## Update Protocol

For any article change:

1. Update the source article first when a source package exists.
2. Render generated HTML from Markdown sources when the article package uses a
   renderer.
3. Keep filenames stable unless the semantic role of a figure changes.
4. Update `og:image` if the cover changes.
5. Check the generated table of contents after heading changes.
6. Verify desktop and mobile widths for overflow, cramped figures, and broken
   code/table scrolling.
7. Keep the README public index current for new public entry points.

Before committing public article changes, run the smallest relevant validation
that covers the edit:

- Check asset references exist and every public image has non-empty alt text.
- Search for private leakage: local absolute paths, local file URLs, temporary
  paths, unfinished editorial markers, hidden process language, private request
  text, or private rationale.
- Check generated TOC, table `data-label` behavior, and mobile body
  `scrollWidth`.
- If links were added or changed, sample-check official docs and representative
  GitHub line anchors.
- Run repository content/build checks when the touched files participate in a
  renderer or build step.

## Evolution

When future revisions reveal a durable writing or maintenance rule, update this
skill in the same change. The rule must be reusable, verified, and phrased as a
general article practice. Do not record private conversations or transient
editing instructions.
