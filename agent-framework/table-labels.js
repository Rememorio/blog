(() => {
  document.querySelectorAll(".article-body table").forEach((table) => {
    const labels = [...table.querySelectorAll("thead th")].map((cell) =>
      (cell.textContent || "").trim(),
    );

    table.querySelectorAll("tbody tr").forEach((row) => {
      [...row.children].forEach((cell, index) => {
        if (labels[index]) cell.dataset.label = labels[index];
      });
    });
  });
})();
