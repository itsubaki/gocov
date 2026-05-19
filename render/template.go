package render

const reportTemplate = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Coverage report</title>
  <style>
    :root {
      color-scheme: light;
      --bg: #f7f7f3;
      --panel: #ffffff;
      --ink: #24272e;
      --muted: #69707d;
      --border: #dcded6;
      --green: #2f9e44;
      --green-soft: #e9f7ed;
      --red: #d94848;
      --red-soft: #fff0f1;
      --amber: #c7861b;
      --amber-soft: #fff7df;
      --teal: #178a8a;
      --code: #272b33;
      --line: #eef0ea;
      --shadow: 0 18px 45px rgba(35, 39, 46, 0.10);
    }

    * { box-sizing: border-box; }

    body {
      margin: 0;
      background: var(--bg);
      color: var(--ink);
      font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      letter-spacing: 0;
    }

    a { color: inherit; text-decoration: none; }

    .layout {
      display: grid;
      grid-template-columns: minmax(260px, 320px) minmax(0, 1fr);
      min-height: 100vh;
    }

    .sidebar {
      position: sticky;
      top: 0;
      height: 100vh;
      overflow: auto;
      border-right: 1px solid var(--border);
      background: #fbfbf8;
      padding: 24px 18px;
    }

    .search {
      width: 100%;
      height: 40px;
      border: 1px solid var(--border);
      border-radius: 8px;
      background: #fff;
      color: var(--ink);
      padding: 0 12px;
      font: inherit;
      outline: none;
    }

    .search:focus {
      border-color: var(--teal);
      box-shadow: 0 0 0 3px rgba(23, 138, 138, 0.14);
    }

    .side-meta {
      margin: 14px 0 18px;
      color: var(--muted);
      font-size: 12px;
      line-height: 1.45;
    }

    .file-list {
      display: grid;
      gap: 8px;
    }

    .file-link {
      display: grid;
      gap: 6px;
      padding: 10px;
      border: 1px solid var(--border);
      border-radius: 8px;
      background: #fff;
    }

    .file-link:hover {
      border-color: #bfc4b8;
      box-shadow: 0 8px 22px rgba(35, 39, 46, 0.08);
    }

    .file-link-top {
      display: flex;
      justify-content: space-between;
      gap: 10px;
      align-items: baseline;
    }

    .file-name {
      min-width: 0;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", monospace;
      font-size: 12px;
      color: var(--code);
    }

    .file-percent {
      flex: 0 0 auto;
      font-size: 12px;
      font-weight: 800;
    }

    .bar {
      display: block;
      height: 7px;
      border-radius: 999px;
      background: #edf0e8;
      overflow: hidden;
    }

    .bar-fill {
      display: block;
      height: 100%;
      width: var(--coverage);
      border-radius: inherit;
      background: var(--coverage-color, var(--green));
    }

    main {
      min-width: 0;
      padding: 34px clamp(18px, 4vw, 48px) 56px;
    }

    .hero {
      margin-bottom: 24px;
    }

    h1 {
      margin: 0 0 10px;
      font-size: clamp(34px, 4vw, 58px);
      line-height: 0.95;
      letter-spacing: 0;
    }

    .subtitle {
      margin: 0;
      color: var(--muted);
      line-height: 1.55;
      max-width: 840px;
      overflow-wrap: anywhere;
    }

    .stats {
      display: grid;
      grid-template-columns: repeat(4, minmax(0, 1fr));
      gap: 12px;
      margin-bottom: 22px;
    }

    .stat {
      background: var(--panel);
      border: 1px solid var(--border);
      border-radius: 8px;
      padding: 16px;
      box-shadow: 0 10px 24px rgba(35, 39, 46, 0.05);
    }

    .stat-label {
      color: var(--muted);
      font-size: 12px;
      text-transform: uppercase;
      font-weight: 800;
    }

    .stat-value {
      margin-top: 8px;
      font-size: 28px;
      font-weight: 900;
    }

    .toolbar {
      position: sticky;
      top: 0;
      z-index: 20;
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
      align-items: center;
      padding: 12px;
      margin: 0 0 18px;
      border: 1px solid var(--border);
      border-radius: 8px;
      background: rgba(255, 255, 255, 0.92);
      backdrop-filter: blur(10px);
    }

    .tool-button {
      border: 1px solid var(--border);
      background: #fff;
      color: var(--ink);
      border-radius: 8px;
      height: 34px;
      padding: 0 12px;
      font: inherit;
      font-size: 13px;
      font-weight: 800;
      cursor: pointer;
    }

    .tool-button.active {
      background: var(--ink);
      border-color: var(--ink);
      color: #fff;
    }

    .toolbar-divider {
      width: 1px;
      height: 24px;
      background: var(--border);
      margin: 0 2px;
    }

    .section-title {
      display: flex;
      align-items: baseline;
      justify-content: space-between;
      gap: 14px;
      margin: 26px 0 12px;
    }

    h2 {
      margin: 0;
      font-size: 22px;
      letter-spacing: 0;
    }

    .directory-overview {
      display: grid;
      grid-template-columns: minmax(220px, 340px) minmax(0, 1fr);
      gap: 22px;
      align-items: center;
      margin-bottom: 16px;
      padding: 18px;
      border: 1px solid var(--border);
      border-radius: 8px;
      background: var(--panel);
      box-shadow: 0 12px 26px rgba(35, 39, 46, 0.06);
    }

    .directory-pie-wrap {
      display: grid;
      place-items: center;
      gap: 12px;
    }

    .directory-pie {
      position: relative;
      width: min(100%, 300px);
      aspect-ratio: 1;
      border-radius: 50%;
      display: grid;
      place-items: center;
      box-shadow: 0 16px 34px rgba(35, 39, 46, 0.10);
    }

    .directory-pie-svg {
      position: absolute;
      inset: 0;
      width: 100%;
      height: 100%;
      overflow: visible;
      border-radius: 50%;
      filter: drop-shadow(0 0 0 rgba(0, 0, 0, 0));
    }

    .directory-slice {
      cursor: pointer;
      stroke: var(--panel);
      stroke-width: 4px;
      transition: filter 150ms ease;
      fill-rule: evenodd;
      vector-effect: non-scaling-stroke;
    }

    .directory-slice:hover,
    .directory-slice:focus {
      filter: brightness(1.08) saturate(1.06);
      outline: none;
    }

    .directory-pie-center {
      display: grid;
      place-items: center;
      position: relative;
      z-index: 2;
      width: 34%;
      aspect-ratio: 1;
      border-radius: 50%;
      background: var(--panel);
      box-shadow: 0 0 0 1px var(--border);
      text-align: center;
      pointer-events: none;
    }

    .directory-pie-value {
      font-size: 28px;
      line-height: 1;
      font-weight: 900;
    }

    .directory-pie-label {
      margin-top: 4px;
      color: var(--muted);
      font-size: 11px;
      font-weight: 800;
      max-width: 104px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .directory-pie-meta {
      color: var(--muted);
      font-size: 12px;
      font-weight: 700;
      text-align: center;
    }

    .directory-tooltip {
      position: absolute;
      z-index: 4;
      display: none;
      width: 240px;
      padding: 8px 10px;
      border-radius: 8px;
      background: var(--ink);
      color: #fff;
      box-shadow: 0 12px 24px rgba(35, 39, 46, 0.18);
      font-size: 12px;
      font-weight: 800;
      line-height: 1.35;
      pointer-events: none;
      transform: translate(-50%, -118%);
      overflow-wrap: anywhere;
    }

    .directory-tooltip.visible {
      display: block;
    }

    .directory-legend {
      display: grid;
      gap: 8px;
      max-height: 380px;
      overflow: auto;
      padding-right: 4px;
    }

    .directory-row {
      display: grid;
      grid-template-columns: 14px minmax(0, 1fr) auto;
      gap: 8px 10px;
      align-items: center;
      padding: 10px 0;
      border-bottom: 1px solid #eceee6;
    }

    .directory-row:last-child {
      border-bottom: 0;
    }

    .directory-swatch {
      width: 14px;
      height: 14px;
      border-radius: 50%;
      background: var(--coverage-color);
      box-shadow: inset 0 0 0 1px rgba(35, 39, 46, 0.18);
    }

    .directory-row-main {
      min-width: 0;
    }

    .directory-row-name {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", monospace;
      font-size: 13px;
      font-weight: 800;
    }

    .directory-row-meta {
      margin-top: 3px;
      color: var(--muted);
      font-size: 12px;
    }

    .directory-row-percent {
      font-size: 18px;
      font-weight: 900;
      text-align: right;
    }

    .directory-coverage {
      grid-column: 2 / -1;
      height: 6px;
      border-radius: 999px;
      background: #edf0e8;
      overflow: hidden;
    }

    .directory-coverage-fill {
      width: var(--coverage);
      height: 100%;
      border-radius: inherit;
      background: var(--coverage-color);
    }

    .directory-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
      gap: 10px;
      margin-bottom: 18px;
    }

    .directory {
      background: var(--panel);
      border: 1px solid var(--border);
      border-radius: 8px;
      padding: 12px;
    }

    .directory-top {
      display: flex;
      justify-content: space-between;
      gap: 12px;
      margin-bottom: 8px;
      font-size: 13px;
      font-weight: 800;
    }

    .directory-name {
      min-width: 0;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", monospace;
    }

    .missing {
      border: 1px solid #efb1b1;
      border-radius: 8px;
      background: #fff7f7;
      color: #842029;
      padding: 14px 16px;
      margin: 18px 0;
      overflow-wrap: anywhere;
    }

    .file-card {
      margin: 0 0 14px;
      background: var(--panel);
      border: 1px solid var(--border);
      border-radius: 8px;
      box-shadow: 0 12px 26px rgba(35, 39, 46, 0.06);
      overflow: clip;
    }

    details { overflow: hidden; }
    details[open] Stats { border-bottom: 1px solid var(--border); }

    Stats {
      list-style: none;
      cursor: pointer;
      padding: 14px 16px;
    }

    Stats::-webkit-details-marker { display: none; }

    .file-header {
      display: grid;
      grid-template-columns: minmax(0, 1fr) auto;
      gap: 16px;
      align-items: center;
    }

    .file-title {
      margin: 0 0 8px;
      font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", monospace;
      font-size: 15px;
      overflow-wrap: anywhere;
    }

    .file-meta {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
      color: var(--muted);
      font-size: 12px;
    }

    .pill {
      display: inline-flex;
      align-items: center;
      min-height: 24px;
      padding: 0 8px;
      border-radius: 999px;
      border: 1px solid var(--border);
      background: #fafaf7;
      font-weight: 700;
    }

    .coverage-number {
      font-size: 24px;
      font-weight: 900;
      text-align: right;
    }

    .source-wrap {
      overflow: auto;
      max-height: 780px;
      background: #fff;
    }

    table.source {
      width: max-content;
      min-width: 100%;
      border-collapse: collapse;
      table-layout: auto;
      font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", monospace;
      font-size: 13px;
      line-height: 1.55;
    }

    .line-no {
      width: 72px;
      color: #8a9099;
      background: #fafaf7;
      text-align: right;
      user-select: none;
      border-right: 1px solid var(--line);
    }

    .line-hits {
      width: 74px;
      color: #69707d;
      text-align: right;
      user-select: none;
      border-right: 1px solid var(--line);
      font-size: 12px;
      font-weight: 800;
    }

    .line-code {
      color: var(--code);
      white-space: pre;
      overflow: visible;
    }

    .source td {
      padding: 0 12px;
      vertical-align: top;
      height: 24px;
      border-bottom: 1px solid #f1f2ed;
    }

    tr.line-covered td { background: var(--green-soft); }
    tr.line-covered .line-no { border-left: 4px solid var(--green); }

    tr.line-missed td { background: var(--red-soft); }
    tr.line-missed .line-no { border-left: 4px solid var(--red); }

    tr.line-partial td { background: var(--amber-soft); }
    tr.line-partial .line-no { border-left: 4px solid var(--amber); }

    tr.line-neutral .line-no { border-left: 4px solid transparent; }

    body[data-line-filter="missed"] tr.line-covered,
    body[data-line-filter="missed"] tr.line-neutral {
      display: none;
    }

    body[data-line-filter="covered"] tr.line-missed,
    body[data-line-filter="covered"] tr.line-partial,
    body[data-line-filter="covered"] tr.line-neutral {
      display: none;
    }

    body[data-line-filter="changed"] tr.line-neutral {
      display: none;
    }

    .hidden { display: none !important; }

    @media (max-width: 960px) {
      .layout {
        grid-template-columns: 1fr;
      }

      .sidebar {
        position: relative;
        height: auto;
      }

      .stats {
        grid-template-columns: repeat(2, minmax(0, 1fr));
      }

      .directory-overview {
        grid-template-columns: 1fr;
      }
    }

    @media (max-width: 620px) {
      main {
        padding-inline: 12px;
      }

      h1 {
        font-size: 36px;
      }

      .stats {
        grid-template-columns: 1fr;
      }

      .file-header {
        grid-template-columns: 1fr;
      }

      .coverage-number {
        text-align: left;
      }

      .line-no {
        width: 54px;
      }

      .line-hits {
        width: 58px;
      }

      table.source {
        font-size: 12px;
      }
    }
  </style>
</head>
<body data-line-filter="all">
  <div class="layout">
    <aside class="sidebar">
      <input class="search" id="fileSearch" type="search" placeholder="Filter files" aria-label="Filter files">
      <div class="side-meta">
        <div>{{.Stats.TotalFiles}} files</div>
        <div>{{.Stats.CoveredStatements}} / {{.Stats.TotalStatements}} statements covered</div>
      </div>
      <nav class="file-list" id="fileList">
        {{range .Files}}
        <a class="file-link" href="#{{.ID}}" data-file-link data-name="{{.DisplayPath}}" data-coverage="{{.Stats.Percent}}">
          <span class="file-link-top">
            <span class="file-name" title="{{.DisplayPath}}">{{.DisplayPath}}</span>
            <span class="file-percent">{{pct .Stats.Percent}}</span>
          </span>
          <span class="bar" aria-hidden="true"><span class="bar-fill" style="--coverage: {{stylePct .Stats.Percent}}; --coverage-color: {{coverageColor .Stats.Percent}}"></span></span>
        </a>
        {{end}}
      </nav>
    </aside>

    <main>
      <header class="hero">
        <div>
          <h1>Coverage report</h1>
          <p class="subtitle">
            {{if .ModulePath}}Module: {{.ModulePath}} {{end}}
          </p>
          <p class="subtitle">
            Generated {{generatedAt .GeneratedAt}} from {{.ProfilePath}}
          </p>
        </div>
      </header>

      <section class="stats" aria-label="Coverage totals">
        <div class="stat">
          <div class="stat-label">Tracked</div>
          <div class="stat-value">{{.Stats.TotalLines}}</div>
        </div>
        <div class="stat">
          <div class="stat-label">Covered</div>
          <div class="stat-value">{{.Stats.CoveredLines}}</div>
        </div>
        <div class="stat">
          <div class="stat-label">Partial</div>
          <div class="stat-value">{{.Stats.PartialLines}}</div>
        </div>
        <div class="stat">
          <div class="stat-label">Missed</div>
          <div class="stat-value">{{.Stats.MissedLines}}</div>
        </div>
      </section>

      <div class="toolbar">
        <button class="tool-button active" type="button" data-sort="a-z" aria-pressed="true">A-Z</button>
        <button class="tool-button" type="button" data-sort="coverage-asc" aria-pressed="false">Coverage asc</button>
        <button class="tool-button" type="button" data-sort="coverage-desc" aria-pressed="false">Coverage desc</button>
      </div>

      {{if .Directories}}
      <div class="section-title">
        <h2>Directories</h2>
      </div>

      <section class="directory-overview" aria-label="Directory coverage by line share">
        <div class="directory-pie-wrap">
          <div class="directory-pie" aria-label="Directory coverage pie chart">
            <svg class="directory-pie-svg" viewBox="-1 -1 2 2" role="img" aria-label="Directory coverage by coverable lines">
              {{range directorySlices .}}
              <path class="directory-slice" d="{{.Path}}" fill="{{.Color}}" tabindex="0" data-directory-slice data-name="{{.Name}}" data-coverage="{{.Coverage}}" data-lines="{{.Lines}}" data-share="{{.Share}}" data-depth="{{.Depth}}" />
              {{end}}
            </svg>
            <div class="directory-pie-center" data-pie-center data-default-value="{{pct .Stats.Percent}}" data-default-label="total">
              <div>
                <div class="directory-pie-value" data-pie-value>{{pct .Stats.Percent}}</div>
                <div class="directory-pie-label" data-pie-label>total</div>
              </div>
            </div>
            <div class="directory-tooltip" data-directory-tooltip role="status"></div>
          </div>
          <div class="directory-pie-meta">{{.Stats.TotalLines}} coverable lines across {{len .Directories}} directories</div>
        </div>
        <section class="directory-grid" data-directory-grid>
          {{range .Directories}}
          <div class="directory" data-directory-card data-name="{{.Name}}" data-coverage="{{.Stats.Percent}}">
            <div class="directory-top">
              <span class="directory-name" title="{{.Name}}">{{.Name}}</span>
              <span>{{pct .Stats.Percent}}</span>
            </div>
            <div class="bar" aria-hidden="true"><span class="bar-fill" style="--coverage: {{stylePct .Stats.Percent}}; --coverage-color: {{coverageColor .Stats.Percent}}"></span></div>
            <div class="directory-row-meta" style="margin-top:8px; color:var(--muted); font-size:12px;">
              {{.Stats.TotalLines}} lines, {{sharePct .Stats.TotalLines $.Stats.TotalLines}} of total
            </div>
          </div>
          {{end}}
        </section>
      </section>
      {{end}}

      {{if .MissingFiles}}
      <div class="missing">
        Source files not found under the repository root:
        {{range .MissingFiles}}<div>{{.}}</div>{{end}}
      </div>
      {{end}}

      <div class="section-title">
        <h2>Files</h2>
      </div>

      <div class="toolbar">
        <button class="tool-button active" type="button" data-filter="all">All lines</button>
        <button class="tool-button" type="button" data-filter="changed">Covered + missed</button>
        <button class="tool-button" type="button" data-filter="missed">Missed only</button>
        <button class="tool-button" type="button" data-filter="covered">Covered only</button>
        <span class="toolbar-divider" aria-hidden="true"></span>
        <button class="tool-button" type="button" id="expandAll">Expand all</button>
        <button class="tool-button" type="button" id="collapseAll">Collapse all</button>
      </div>

      <div data-file-cards>
      {{range .Files}}
      <section class="file-card" id="{{.ID}}" data-file-card data-name="{{.DisplayPath}}" data-coverage="{{.Stats.Percent}}">
        <details open>
          <Stats>
            <div class="file-header">
              <div>
                <h3 class="file-title">{{.DisplayPath}}</h3>
                <div class="file-meta">
                  <span class="pill">{{.Stats.CoveredStatements}} / {{.Stats.TotalStatements}} statements</span>
                  <span class="pill">{{.Blocks}} blocks</span>
                  <span class="pill">{{.Stats.MissedLines}} missed lines</span>
                  {{if not .Found}}<span class="pill">source missing</span>{{end}}
                </div>
              </div>
              <div class="coverage-number">{{pct .Stats.Percent}}</div>
            </div>
          </Stats>
          {{if .Found}}
          <div class="source-wrap">
            <table class="source" aria-label="Source coverage for {{.DisplayPath}}">
              <tbody>
                {{range .Lines}}
                <tr class="{{lineClass .State}}">
                  <td class="line-no">{{.Number}}</td>
                  <td class="line-hits">{{.Hits}}</td>
                  <td class="line-code">{{.Code}}</td>
                </tr>
                {{end}}
              </tbody>
            </table>
          </div>
          {{end}}
        </details>
      </section>
      {{end}}
      </div>
    </main>
  </div>

  <script>
    const search = document.getElementById('fileSearch');
    const links = [...document.querySelectorAll('[data-file-link]')];
    const cards = [...document.querySelectorAll('[data-file-card]')];
    const buttons = [...document.querySelectorAll('[data-filter]')];
    const sortButtons = [...document.querySelectorAll('[data-sort]')];
    const fileList = document.getElementById('fileList');
    const fileCards = document.querySelector('[data-file-cards]');
    const directoryLegend = document.querySelector('[data-directory-legend]');
    const directoryGrid = document.querySelector('[data-directory-grid]');
    const directoryTooltip = document.querySelector('[data-directory-tooltip]');
    const pieCenter = document.querySelector('[data-pie-center]');
    const pieValue = document.querySelector('[data-pie-value]');
    const pieLabel = document.querySelector('[data-pie-label]');
    const nameCollator = new Intl.Collator(undefined, { numeric: true, sensitivity: 'base' });

    search.addEventListener('input', () => {
      const query = search.value.trim().toLowerCase();
      for (const link of links) {
        const visible = link.dataset.name.toLowerCase().includes(query);
        link.classList.toggle('hidden', !visible);
      }
      for (const card of cards) {
        const visible = card.dataset.name.toLowerCase().includes(query);
        card.classList.toggle('hidden', !visible);
      }
    });

    for (const button of buttons) {
      button.addEventListener('click', () => {
        document.body.dataset.lineFilter = button.dataset.filter;
        for (const other of buttons) {
          other.classList.toggle('active', other === button);
        }
      });
    }

    const compareName = (left, right) => {
      return nameCollator.compare(left.dataset.name || '', right.dataset.name || '');
    };

    const compareCoverage = (left, right, direction) => {
      const leftCoverage = Number.parseFloat(left.dataset.coverage || '0');
      const rightCoverage = Number.parseFloat(right.dataset.coverage || '0');
      const coverageDiff = (leftCoverage - rightCoverage) * direction;
      if (coverageDiff !== 0) return coverageDiff;
      return compareName(left, right);
    };

    const compareItems = (mode) => {
      if (mode === 'coverage-asc') return (left, right) => compareCoverage(left, right, 1);
      if (mode === 'coverage-desc') return (left, right) => compareCoverage(left, right, -1);
      return compareName;
    };

    const sortChildren = (container, selector, mode) => {
      if (!container) return;
      const items = [...container.querySelectorAll(selector)].sort(compareItems(mode));
      for (const item of items) container.appendChild(item);
    };

    const sortReport = (mode) => {
      sortChildren(fileList, '[data-file-link]', mode);
      sortChildren(fileCards, '[data-file-card]', mode);
      sortChildren(directoryLegend, '[data-directory-row]', mode);
      sortChildren(directoryGrid, '[data-directory-card]', mode);

      for (const button of sortButtons) {
        const active = button.dataset.sort === mode;
        button.classList.toggle('active', active);
        button.setAttribute('aria-pressed', active ? 'true' : 'false');
      }
    };

    for (const button of sortButtons) {
      button.addEventListener('click', () => sortReport(button.dataset.sort));
    }

    const showDirectory = (slice, event) => {
      if (!slice || !pieCenter || !pieValue || !pieLabel) return;
      pieValue.textContent = slice.dataset.coverage;
      pieLabel.textContent = slice.dataset.name;

      if (!directoryTooltip) return;
      directoryTooltip.textContent = slice.dataset.name + ': ' + slice.dataset.coverage + '; ' + slice.dataset.lines + ' lines, ' + slice.dataset.share + ' of total';
      directoryTooltip.classList.add('visible');

      const pie = slice.closest('.directory-pie');
      if (!pie) return;
      const rect = pie.getBoundingClientRect();
      if (event) {
        directoryTooltip.style.left = (event.clientX - rect.left) + 'px';
        directoryTooltip.style.top = (event.clientY - rect.top) + 'px';
      } else {
        directoryTooltip.style.left = '50%';
        directoryTooltip.style.top = '18%';
      }
    };

    const resetDirectory = () => {
      if (pieCenter && pieValue && pieLabel) {
        pieValue.textContent = pieCenter.dataset.defaultValue;
        pieLabel.textContent = pieCenter.dataset.defaultLabel;
      }
      if (directoryTooltip) {
        directoryTooltip.classList.remove('visible');
      }
    };

    for (const slice of document.querySelectorAll('[data-directory-slice]')) {
      slice.addEventListener('mouseenter', event => showDirectory(slice, event));
      slice.addEventListener('mousemove', event => showDirectory(slice, event));
      slice.addEventListener('mouseleave', resetDirectory);
      slice.addEventListener('focus', () => showDirectory(slice));
      slice.addEventListener('blur', resetDirectory);
    }

    document.getElementById('expandAll').addEventListener('click', () => {
      for (const item of document.querySelectorAll('details')) item.open = true;
    });

    document.getElementById('collapseAll').addEventListener('click', () => {
      for (const item of document.querySelectorAll('details')) item.open = false;
    });
  </script>
</body>
</html>
`
