from __future__ import annotations

import html
import re
from pathlib import Path


ROOT = Path(__file__).resolve().parent
CHAPTERS = ROOT / "chapters"

ORDER = [
    "00-series-map",
    "01-learning-loop",
    "02-memory",
    "03-skills",
    "04-session-search",
    "05-runtime-triggers",
    "06-summary-and-questions",
]


def local_href(href: str) -> str:
    if href.endswith(".md"):
        return href[:-3] + ".html"
    return href


def inline(text: str) -> str:
    code_parts: list[str] = []

    def stash_code(match: re.Match[str]) -> str:
        code = html.escape(match.group(1), quote=False)
        code_parts.append(f"<code>{code}</code>")
        return f"\ue000{len(code_parts) - 1}\ue000"

    text = re.sub(r"`([^`]+)`", stash_code, text)
    escaped = html.escape(text, quote=False)

    def link(match: re.Match[str]) -> str:
        label = match.group(1)
        href = local_href(match.group(2).strip())
        return f'<a href="{html.escape(href, quote=True)}">{label}</a>'

    escaped = re.sub(r"\[([^\]]+)\]\(([^)]+)\)", link, escaped)
    escaped = re.sub(r"\*\*([^*]+)\*\*", r"<strong>\1</strong>", escaped)

    for idx, code in enumerate(code_parts):
        escaped = escaped.replace(f"\ue000{idx}\ue000", code)
    return escaped


def slugify(text: str, seen: dict[str, int]) -> str:
    base = re.sub(r"<[^>]+>", "", text)
    base = re.sub(r"[`*_~\[\]().,，。？！：:；;、/\\]+", " ", base)
    words = re.findall(r"[A-Za-z0-9]+|[\u4e00-\u9fff]+", base)
    slug = "-".join(words).lower() or "section"
    count = seen.get(slug, 0)
    seen[slug] = count + 1
    return slug if count == 0 else f"{slug}-{count + 1}"


def split_table_row(line: str) -> list[str]:
    line = line.strip()
    if line.startswith("|"):
        line = line[1:]
    if line.endswith("|"):
        line = line[:-1]
    return [cell.strip() for cell in line.split("|")]


def is_table_separator(line: str) -> bool:
    cells = split_table_row(line)
    return bool(cells) and all(re.fullmatch(r":?-{3,}:?", cell.strip()) for cell in cells)


def render_table(rows: list[str]) -> str:
    header = split_table_row(rows[0])
    body = [split_table_row(row) for row in rows[2:]]
    parts = ["<table>", "<thead><tr>"]
    for cell in header:
        parts.append(f"<th>{inline(cell)}</th>")
    parts.append("</tr></thead>")
    if body:
        parts.append("<tbody>")
        for row in body:
            parts.append("<tr>")
            for cell in row:
                parts.append(f"<td>{inline(cell)}</td>")
            parts.append("</tr>")
        parts.append("</tbody>")
    parts.append("</table>")
    return "\n".join(parts)


def render_markdown(markdown: str) -> tuple[str, list[tuple[int, str, str]], bool]:
    lines = markdown.splitlines()
    parts: list[str] = []
    toc: list[tuple[int, str, str]] = []
    paragraph: list[str] = []
    list_type: str | None = None
    in_code = False
    code_lang = ""
    code_lines: list[str] = []
    seen_slugs: dict[str, int] = {}
    has_mermaid = False
    idx = 0

    def flush_paragraph() -> None:
        nonlocal paragraph
        if paragraph:
            parts.append(f"<p>{inline(' '.join(paragraph))}</p>")
            paragraph = []

    def close_list() -> None:
        nonlocal list_type
        if list_type:
            parts.append(f"</{list_type}>")
            list_type = None

    while idx < len(lines):
        line = lines[idx]

        if in_code:
            if line.startswith("```"):
                raw = "\n".join(code_lines)
                if code_lang == "mermaid":
                    has_mermaid = True
                    parts.append(f'<div class="mermaid">{html.escape(raw, quote=False)}</div>')
                else:
                    lang_class = f" language-{html.escape(code_lang, quote=True)}" if code_lang else ""
                    parts.append(
                        f'<pre><code class="{lang_class.strip()}">{html.escape(raw, quote=False)}</code></pre>'
                    )
                in_code = False
                code_lang = ""
                code_lines = []
            else:
                code_lines.append(line)
            idx += 1
            continue

        if line.startswith("```"):
            flush_paragraph()
            close_list()
            in_code = True
            code_lang = line[3:].strip().lower()
            code_lines = []
            idx += 1
            continue

        if not line.strip():
            flush_paragraph()
            close_list()
            idx += 1
            continue

        if (
            line.lstrip().startswith("|")
            and idx + 1 < len(lines)
            and is_table_separator(lines[idx + 1])
        ):
            flush_paragraph()
            close_list()
            rows = [line, lines[idx + 1]]
            idx += 2
            while idx < len(lines) and lines[idx].lstrip().startswith("|"):
                rows.append(lines[idx])
                idx += 1
            parts.append(render_table(rows))
            continue

        heading = re.match(r"^(#{1,6})\s+(.+?)\s*$", line)
        if heading:
            flush_paragraph()
            close_list()
            level = len(heading.group(1))
            text = heading.group(2)
            anchor = slugify(text, seen_slugs)
            toc.append((level, anchor, re.sub(r"`", "", text)))
            parts.append(f'<h{level} id="{anchor}">{inline(text)}</h{level}>')
            idx += 1
            continue

        quote = re.match(r"^>\s*(.*)$", line)
        if quote:
            flush_paragraph()
            close_list()
            parts.append(f"<blockquote><p>{inline(quote.group(1))}</p></blockquote>")
            idx += 1
            continue

        unordered = re.match(r"^\s*[-*]\s+(.+)$", line)
        ordered = re.match(r"^\s*\d+\.\s+(.+)$", line)
        if unordered or ordered:
            flush_paragraph()
            wanted = "ul" if unordered else "ol"
            if list_type != wanted:
                close_list()
                parts.append(f"<{wanted}>")
                list_type = wanted
            item = (unordered or ordered).group(1)
            parts.append(f"<li>{inline(item)}</li>")
            idx += 1
            continue

        paragraph.append(line.strip())
        idx += 1

    flush_paragraph()
    close_list()
    return "\n".join(parts), toc, has_mermaid


def nav_link(slug: str | None, label: str, rel: str) -> str:
    if not slug:
        return f'<span class="chapter-disabled">{label}</span>'
    return f'<a href="./{slug}.html" rel="{rel}">{label}</a>'


def render_page(slug: str, markdown: str, index: int) -> str:
    body, toc, has_mermaid = render_markdown(markdown)
    title = toc[0][2] if toc else slug
    prev_slug = ORDER[index - 1] if index > 0 else None
    next_slug = ORDER[index + 1] if index + 1 < len(ORDER) else None
    toc_links = "\n".join(
        f'<a class="toc-level-{level}" href="#{anchor}">{html.escape(text)}</a>'
        for level, anchor, text in toc
        if level in {2, 3}
    )
    if not toc_links:
        toc_links = '<a href="../index.html#chapters">返回章节目录</a>'
    mermaid_script = ""
    if has_mermaid:
        mermaid_script = """
    <script type="module">
      import mermaid from "https://cdn.jsdelivr.net/npm/mermaid@11/dist/mermaid.esm.min.mjs";
      mermaid.initialize({ startOnLoad: true, theme: "neutral" });
    </script>"""

    return f"""<!doctype html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>{html.escape(title)} | Hermes Agent 源码阅读</title>
    <meta name="theme-color" content="#f6f4ef" />
    <link rel="stylesheet" href="../styles.css" />
  </head>
  <body>
    <header class="site-header">
      <a class="brand" href="../" aria-label="Hermes Agent 源码阅读首页">
        <span class="brand-mark">H</span>
        <span>Hermes Agent Notes</span>
      </a>
      <nav aria-label="页面导航">
        <a href="../#chapters">章节</a>
        <a href="../#sources">源码</a>
        <a href="https://github.com/NousResearch/hermes-agent/tree/3edd09a46">源码版本</a>
      </nav>
    </header>

    <main class="page-shell chapter-shell">
      <aside class="toc chapter-toc" aria-label="章节目录">
        <strong>本章目录</strong>
        {toc_links}
      </aside>
      <article class="article markdown-body">
        <nav class="chapter-links" aria-label="章节导航">
          {nav_link(prev_slug, "上一章", "prev")}
          <a href="../#chapters">全部章节</a>
          {nav_link(next_slug, "下一章", "next")}
        </nav>
        {body}
        <nav class="chapter-links chapter-links-bottom" aria-label="章节导航">
          {nav_link(prev_slug, "上一章", "prev")}
          <a href="../#chapters">全部章节</a>
          {nav_link(next_slug, "下一章", "next")}
        </nav>
      </article>
    </main>

    <footer class="site-footer">
      <p>Hermes Agent 源码阅读笔记。</p>
      <a href="../">返回首页</a>
    </footer>{mermaid_script}
  </body>
</html>
"""


def main() -> None:
    for index, slug in enumerate(ORDER):
        source = CHAPTERS / f"{slug}.md"
        target = CHAPTERS / f"{slug}.html"
        target.write_text(render_page(slug, source.read_text(encoding="utf-8"), index), encoding="utf-8")


if __name__ == "__main__":
    main()
