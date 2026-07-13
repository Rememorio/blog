(() => {
  const article = document.querySelector(".article-body");
  const toc = document.querySelector("#toc");
  if (!article || !toc) return;

  document.querySelectorAll("table").forEach((table) => {
    const headers = Array.from(table.querySelectorAll("thead th")).map((th) =>
      th.textContent.trim()
    );
    table.querySelectorAll("tbody tr").forEach((row) => {
      Array.from(row.children).forEach((cell, index) => {
        if (headers[index]) cell.dataset.label = headers[index];
      });
    });
  });

  const headings = Array.from(article.querySelectorAll("h2, h3, h4"));
  if (!headings.length) return;

  const counts = new Map();
  const slugify = (text) => {
    const base = text
      .toLowerCase()
      .trim()
      .replace(/[^\p{Letter}\p{Number}]+/gu, "-")
      .replace(/^-+|-+$/g, "");
    return base || "section";
  };

  const list = document.createElement("ol");
  list.className = "toc-list";
  const links = [];
  const lastItems = new Map();

  const childListFor = (item) => {
    let child = item.querySelector(":scope > ol");
    if (!child) {
      child = document.createElement("ol");
      child.className = "toc-sublist";
      item.appendChild(child);
    }
    return child;
  };

  const parentListFor = (level) => {
    if (level === 2) return list;
    if (level === 3 && lastItems.get(2)) return childListFor(lastItems.get(2));
    if (level === 4 && lastItems.get(3)) return childListFor(lastItems.get(3));
    if (level === 4 && lastItems.get(2)) return childListFor(lastItems.get(2));
    return list;
  };

  headings.forEach((heading, index) => {
    if (!heading.id) {
      const base = slugify(heading.textContent);
      const seen = counts.get(base) || 0;
      counts.set(base, seen + 1);
      heading.id = seen === 0 ? base : `${base}-${seen + 1}`;
    }

    const level = Number(heading.tagName.slice(1));
    const item = document.createElement("li");
    const link = document.createElement("a");
    link.href = `#${heading.id}`;
    link.textContent = heading.textContent.trim();
    link.className = `toc-link toc-level-${level}`;
    if (index === 0) link.classList.add("active");
    item.appendChild(link);
    parentListFor(level).appendChild(item);
    lastItems.set(level, item);
    if (level <= 2) lastItems.delete(3);
    if (level <= 3) lastItems.delete(4);
    links.push(link);
  });

  toc.appendChild(list);

  const activate = (id) => {
    links.forEach((link) => {
      link.classList.toggle("active", link.getAttribute("href") === `#${id}`);
    });
  };

  const observer = new IntersectionObserver(
    (entries) => {
      const visible = entries
        .filter((entry) => entry.isIntersecting)
        .sort((a, b) => a.boundingClientRect.top - b.boundingClientRect.top);
      if (visible[0]) activate(visible[0].target.id);
    },
    { rootMargin: "-20% 0px -72% 0px", threshold: 0 }
  );

  headings.forEach((heading) => observer.observe(heading));
})();
