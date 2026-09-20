(() => {
  const form = document.querySelector('.post-search');
  const input = document.querySelector('#post-query');
  const list = document.querySelector('.post-list');
  if (!form || !input || !list) return;

  const cards = [...list.querySelectorAll('.post-card')];
  const buttons = [...document.querySelectorAll('[data-post-view]')];
  const clear = form.querySelector('[data-search-clear]');
  const status = document.querySelector('.post-view-status');
  const empty = document.querySelector('.post-empty');
  const chinese = document.documentElement.lang.startsWith('zh');
  const normalize = value => value.normalize('NFKC').toLowerCase().replace(/[^\p{L}\p{N}]/gu, '');
  const index = cards.map(card => normalize(card.textContent));
  let view = 'latest';
  let composing = false;

  function render(updateUrl = true) {
    const query = input.value.trim();
    const terms = query.split(/\s+/u).map(normalize).filter(Boolean);
    const searching = terms.length > 0;
    let shown = 0;
    cards.forEach((card, i) => {
      const matches = terms.every(term => index[i].includes(term));
      card.hidden = !matches || (!searching && view !== 'all' && i >= 12);
      if (!card.hidden) shown++;
    });
    list.dataset.view = searching ? 'search' : view;
    buttons.forEach(button => {
      button.disabled = searching && button.dataset.postView === 'latest';
      button.setAttribute('aria-pressed', String(button.dataset.postView === (searching ? 'all' : view)));
    });
    clear.hidden = !input.value;
    empty.hidden = shown !== 0;
    status.textContent = chinese
      ? (searching ? `找到 ${shown} 篇匹配文章，共 ${cards.length} 篇。` : `显示 ${shown} / ${cards.length} 篇；搜索覆盖全部文章。`)
      : (searching ? `${shown} ${shown === 1 ? 'match' : 'matches'} across ${cards.length} articles.` : `Showing ${shown} of ${cards.length}; search covers every article.`);

    const url = new URL(location.href);
    query ? url.searchParams.set('q', query) : url.searchParams.delete('q');
    view === 'all' ? url.searchParams.set('view', 'all') : url.searchParams.delete('view');
    if (updateUrl) history.replaceState(null, '', url);
    document.querySelectorAll('.language-switch a').forEach(link => {
      const sibling = new URL(link.href);
      sibling.search = url.search;
      link.href = sibling.href;
    });
  }

  function restore() {
    const params = new URL(location.href).searchParams;
    view = params.get('view') === 'all' ? 'all' : 'latest';
    input.value = params.get('q') || '';
    render(false);
  }

  buttons.forEach(button => button.addEventListener('click', () => {
    view = button.dataset.postView;
    render();
  }));
  input.addEventListener('compositionstart', () => { composing = true; });
  input.addEventListener('compositionend', () => { composing = false; render(); });
  input.addEventListener('input', event => {
    if (!composing && !event.isComposing) render();
  });
  form.addEventListener('submit', event => { event.preventDefault(); if (!composing) render(); });
  clear.addEventListener('click', () => { input.value = ''; render(); input.focus(); });
  window.addEventListener('popstate', restore);
  restore();
  document.documentElement.classList.add('home-enhanced');
})();
