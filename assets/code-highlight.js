(() => {
  const SKIPPED_LANGUAGES = new Set(["text", "txt", "plaintext"]);

  function highlightCodeBlocks() {
    if (!window.hljs) return;

    document.querySelectorAll("pre code").forEach((block) => {
      const languageClass = Array.from(block.classList).find((className) =>
        className.startsWith("language-"),
      );
      const language = languageClass?.slice("language-".length);

      if (!language || SKIPPED_LANGUAGES.has(language)) {
        block.classList.add("nohighlight");
        return;
      }

      if (window.hljs.getLanguage(language)) {
        window.hljs.highlightElement(block);
      }
    });
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", highlightCodeBlocks, {
      once: true,
    });
  } else {
    highlightCodeBlocks();
  }
})();
