package render

const reportTemplate = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Coverage report</title>
  <script>
    (function () {
      try {
        const theme = localStorage.getItem('gocov-theme');
        if (theme === 'dark' || theme === 'light') {
          document.documentElement.dataset.theme = theme;
        }
      } catch (_) {}
    })();
  </script>
  <style>
    :root {
      color-scheme: light;
      --bg: #f7f7f3;
      --panel: #ffffff;
      --surface: #ffffff;
      --surface-soft: #fafaf7;
      --sidebar: #fbfbf8;
      --ink: #24272e;
      --muted: #69707d;
      --border: #dcded6;
      --border-strong: #bfc4b8;
      --green: #2f9e44;
      --green-soft: #e9f7ed;
      --red: #d94848;
      --red-soft: #fff0f1;
      --amber: #c7861b;
      --amber-soft: #fff7df;
      --teal: #178a8a;
      --code: #272b33;
      --line: #eef0ea;
      --bar-track: #edf0e8;
      --toolbar-bg: rgba(255, 255, 255, 0.92);
      --button-bg: #ffffff;
      --button-active-bg: var(--ink);
      --button-active-ink: #ffffff;
      --focus-ring: rgba(23, 138, 138, 0.14);
      --row-border: #eceee6;
      --source-bg: #ffffff;
      --gutter-bg: #fafaf7;
      --line-no: #8a9099;
      --line-hits: #69707d;
      --source-row: #f1f2ed;
      --missing-border: #efb1b1;
      --missing-bg: #fff7f7;
      --missing-ink: #842029;
      --tooltip-bg: var(--ink);
      --tooltip-ink: #ffffff;
      --swatch-ring: rgba(35, 39, 46, 0.18);
      --shadow: 0 18px 45px rgba(35, 39, 46, 0.10);
      --shadow-soft: 0 10px 24px rgba(35, 39, 46, 0.05);
      --shadow-card: 0 12px 26px rgba(35, 39, 46, 0.06);
      --shadow-hover: 0 8px 22px rgba(35, 39, 46, 0.08);
      --shadow-pie: 0 16px 34px rgba(35, 39, 46, 0.10);
      --shadow-tooltip: 0 12px 24px rgba(35, 39, 46, 0.18);
    }

    @media (prefers-color-scheme: dark) {
      :root:not([data-theme="light"]) {
        color-scheme: dark;
        --bg: #11130f;
        --panel: #181b16;
        --surface: #181b16;
        --surface-soft: #20241d;
        --sidebar: #141610;
        --ink: #eef2e8;
        --muted: #a6ad9f;
        --border: #32372d;
        --border-strong: #596151;
        --green: #6ed082;
        --green-soft: rgba(54, 150, 77, 0.22);
        --red: #ff7a7a;
        --red-soft: rgba(217, 72, 72, 0.18);
        --amber: #f4c15c;
        --amber-soft: rgba(199, 134, 27, 0.20);
        --teal: #50c7c7;
        --code: #edf1e8;
        --line: #2a2f26;
        --bar-track: #2a3026;
        --toolbar-bg: rgba(24, 27, 22, 0.92);
        --button-bg: #12150f;
        --button-active-bg: #eef2e8;
        --button-active-ink: #11130f;
        --focus-ring: rgba(80, 199, 199, 0.20);
        --row-border: #2a2f26;
        --source-bg: #11140f;
        --gutter-bg: #181b16;
        --line-no: #a0a89a;
        --line-hits: #a6ad9f;
        --source-row: #242920;
        --missing-border: #7c3737;
        --missing-bg: #2a1717;
        --missing-ink: #ffb5b5;
        --tooltip-bg: #f0f5e8;
        --tooltip-ink: #11130f;
        --swatch-ring: rgba(238, 242, 232, 0.22);
        --shadow: 0 18px 45px rgba(0, 0, 0, 0.38);
        --shadow-soft: 0 10px 24px rgba(0, 0, 0, 0.26);
        --shadow-card: 0 12px 26px rgba(0, 0, 0, 0.28);
        --shadow-hover: 0 8px 22px rgba(0, 0, 0, 0.30);
        --shadow-pie: 0 16px 34px rgba(0, 0, 0, 0.34);
        --shadow-tooltip: 0 12px 24px rgba(0, 0, 0, 0.34);
      }
    }

    :root[data-theme="dark"] {
      color-scheme: dark;
      --bg: #11130f;
      --panel: #181b16;
      --surface: #181b16;
      --surface-soft: #20241d;
      --sidebar: #141610;
      --ink: #eef2e8;
      --muted: #a6ad9f;
      --border: #32372d;
      --border-strong: #596151;
      --green: #6ed082;
      --green-soft: rgba(54, 150, 77, 0.22);
      --red: #ff7a7a;
      --red-soft: rgba(217, 72, 72, 0.18);
      --amber: #f4c15c;
      --amber-soft: rgba(199, 134, 27, 0.20);
      --teal: #50c7c7;
      --code: #edf1e8;
      --line: #2a2f26;
      --bar-track: #2a3026;
      --toolbar-bg: rgba(24, 27, 22, 0.92);
      --button-bg: #12150f;
      --button-active-bg: #eef2e8;
      --button-active-ink: #11130f;
      --focus-ring: rgba(80, 199, 199, 0.20);
      --row-border: #2a2f26;
      --source-bg: #11140f;
      --gutter-bg: #181b16;
      --line-no: #a0a89a;
      --line-hits: #a6ad9f;
      --source-row: #242920;
      --missing-border: #7c3737;
      --missing-bg: #2a1717;
      --missing-ink: #ffb5b5;
      --tooltip-bg: #f0f5e8;
      --tooltip-ink: #11130f;
      --swatch-ring: rgba(238, 242, 232, 0.22);
      --shadow: 0 18px 45px rgba(0, 0, 0, 0.38);
      --shadow-soft: 0 10px 24px rgba(0, 0, 0, 0.26);
      --shadow-card: 0 12px 26px rgba(0, 0, 0, 0.28);
      --shadow-hover: 0 8px 22px rgba(0, 0, 0, 0.30);
      --shadow-pie: 0 16px 34px rgba(0, 0, 0, 0.34);
      --shadow-tooltip: 0 12px 24px rgba(0, 0, 0, 0.34);
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
      background: var(--sidebar);
      padding: 24px 18px;
    }

    .sidebar-controls {
      display: grid;
      grid-template-columns: minmax(0, 1fr);
      gap: 8px;
      align-items: center;
    }

    .search {
      width: 100%;
      height: 40px;
      border: 1px solid var(--border);
      border-radius: 8px;
      background: var(--surface);
      color: var(--ink);
      padding: 0 12px;
      font: inherit;
      outline: none;
    }

    .search:focus {
      border-color: var(--teal);
      box-shadow: 0 0 0 3px var(--focus-ring);
    }

    .theme-toggle {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      gap: 7px;
      height: 40px;
      min-width: 88px;
      border: 1px solid var(--border);
      border-radius: 8px;
      background: var(--button-bg);
      color: var(--ink);
      cursor: pointer;
      font: inherit;
      font-size: 12px;
      font-weight: 800;
    }

    .theme-toggle-mobile {
      display: none;
    }

    .theme-toggle:hover {
      border-color: var(--border-strong);
    }

    .theme-toggle:focus-visible {
      outline: none;
      border-color: var(--teal);
      box-shadow: 0 0 0 3px var(--focus-ring);
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
      background: var(--surface);
    }

    .file-link:hover {
      border-color: var(--border-strong);
      box-shadow: var(--shadow-hover);
    }

    .file-link-top {
      display: flex;
      justify-content: space-between;
      gap: 10px;
      align-items: baseline;
    }

    .file-name {
      min-width: 0;
      max-width: 200px;
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
      background: var(--bar-track);
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
      display: grid;
      grid-template-columns: minmax(0, 1fr) auto;
      gap: 16px;
      align-items: start;
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
      box-shadow: var(--shadow-soft);
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

    .lines-label {
      font-size: 14px;
      color: var(--muted);
      margin-left: 8px;
      font-weight: 600;
      letter-spacing: 0;
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
      background: var(--toolbar-bg);
      backdrop-filter: blur(10px);
    }

    .tool-button {
      border: 1px solid var(--border);
      background: var(--button-bg);
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
      background: var(--button-active-bg);
      border-color: var(--button-active-bg);
      color: var(--button-active-ink);
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
      box-shadow: var(--shadow-card);
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
      box-shadow: var(--shadow-pie);
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
      width: max-content;
      padding: 8px 10px;
      border-radius: 8px;
      background: var(--tooltip-bg);
      color: var(--tooltip-ink);
      box-shadow: var(--shadow-tooltip);
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
      border-bottom: 1px solid var(--row-border);
    }

    .directory-row:last-child {
      border-bottom: 0;
    }

    .directory-swatch {
      width: 14px;
      height: 14px;
      border-radius: 50%;
      background: var(--coverage-color);
      box-shadow: inset 0 0 0 1px var(--swatch-ring);
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
      background: var(--bar-track);
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
      border: 1px solid var(--missing-border);
      border-radius: 8px;
      background: var(--missing-bg);
      color: var(--missing-ink);
      padding: 14px 16px;
      margin: 18px 0;
      overflow-wrap: anywhere;
    }

    .file-card {
      margin: 0 0 14px;
      background: var(--panel);
      border: 1px solid var(--border);
      border-radius: 8px;
      box-shadow: var(--shadow-card);
      overflow: clip;
    }

    details { overflow: hidden; }
    details[open] summary { border-bottom: 1px solid var(--border); }

    summary {
      list-style: none;
      cursor: pointer;
      padding: 14px 16px;
    }

    summary::-webkit-details-marker { display: none; }

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
      background: var(--surface-soft);
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
      background: var(--source-bg);
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
      color: var(--line-no);
      background: var(--gutter-bg);
      text-align: right;
      user-select: none;
      border-right: 1px solid var(--line);
    }

    .line-hits {
      width: 74px;
      color: var(--line-hits);
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
      border-bottom: 1px solid var(--source-row);
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

      .sidebar-controls {
        grid-template-columns: minmax(0, 1fr) auto;
      }

      .theme-toggle-desktop {
        display: none;
      }

      .theme-toggle-mobile {
        display: inline-flex;
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

      .hero {
        grid-template-columns: 1fr;
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
      <div class="sidebar-controls">
        <input class="search" id="fileSearch" type="search" placeholder="Filter files" aria-label="Filter files">
        <button class="theme-toggle theme-toggle-mobile" type="button" data-theme-toggle aria-label="Toggle color theme" aria-pressed="false">
          <span data-theme-label>Light</span>
        </button>
      </div>
      <div class="side-meta">
        <div>{{.Stats.TotalFiles}} files</div>
        <div>{{.Stats.CoveredStatements}} / {{.Stats.TotalStatements}} ({{pct .Stats.Percent}}) statements covered.</div>
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
            Generated {{generatedAt .GeneratedAt}}
          </p>
        </div>
        <button class="theme-toggle theme-toggle-desktop" type="button" data-theme-toggle aria-label="Toggle color theme" aria-pressed="false">
          <span data-theme-label>Light</span>
        </button>
      </header>

      <section class="stats" aria-label="Coverage totals">
        <div class="stat">
          <div class="stat-label">Tracked</div>
          <div class="stat-value">{{.Stats.TotalLines}}<span class="lines-label">lines</span></div>
        </div>
        <div class="stat">
          <div class="stat-label">Covered</div>
          <div class="stat-value">{{.Stats.CoveredLines}}<span class="lines-label">lines</span></div>
        </div>
        <div class="stat">
          <div class="stat-label">Partial</div>
          <div class="stat-value">{{.Stats.PartialLines}}<span class="lines-label">lines</span></div>
        </div>
        <div class="stat">
          <div class="stat-label">Missed</div>
          <div class="stat-value">{{.Stats.MissedLines}}<span class="lines-label">lines</span></div>
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
            <svg class="directory-pie-svg" viewBox="-1 -1 2 2" role="img" aria-label="Directory coverage by coverable statements">
              {{range directories .}}
              <path class="directory-slice" d="{{.Path}}" fill="{{.Color}}" tabindex="0" data-directory-slice data-name="{{.Name}}" data-coverage="{{.Coverage}}" data-statements="{{.Statements}}" data-share="{{.Share}}" data-depth="{{.Depth}}" />
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
          <div class="directory-pie-meta">{{.Stats.TotalStatements}} coverable statements across {{len .Directories}} directories</div>
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
             {{.Stats.CoveredStatements}} / {{.Stats.TotalStatements}} statements covered.
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
        <details>
          <summary>
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
          </summary>
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
    const themeToggles = [...document.querySelectorAll('[data-theme-toggle]')];
    const themeLabels = [...document.querySelectorAll('[data-theme-label]')];
    const themeQuery = window.matchMedia('(prefers-color-scheme: dark)');
    const nameCollator = new Intl.Collator(undefined, { numeric: true, sensitivity: 'base' });

    const activeTheme = () => {
      const theme = document.documentElement.dataset.theme;
      if (theme === 'dark' || theme === 'light') return theme;
      return themeQuery.matches ? 'dark' : 'light';
    };

    const updateThemeToggle = () => {
      if (themeToggles.length === 0) return;
      const theme = activeTheme();
      for (const toggle of themeToggles) {
        toggle.dataset.themeMode = theme;
        toggle.setAttribute('aria-pressed', theme === 'dark' ? 'true' : 'false');
      }
      for (const label of themeLabels) {
        label.textContent = theme === 'dark' ? 'Dark' : 'Light';
      }
    };

    if (themeToggles.length > 0) {
      for (const toggle of themeToggles) {
        toggle.addEventListener('click', () => {
          const next = activeTheme() === 'dark' ? 'light' : 'dark';
          document.documentElement.dataset.theme = next;
          try {
            localStorage.setItem('gocov-theme', next);
          } catch (_) {}
          updateThemeToggle();
        });
      }

      const handleThemeQueryChange = () => {
        if (!document.documentElement.dataset.theme) updateThemeToggle();
      };

      if (themeQuery.addEventListener) {
        themeQuery.addEventListener('change', handleThemeQueryChange);
      } else if (themeQuery.addListener) {
        themeQuery.addListener(handleThemeQueryChange);
      }

      updateThemeToggle();
    }

    function openDetailsForHash() {
      if (!location.hash) return;
      const id = location.hash.slice(1);
      const card = document.getElementById(id);
      if (card) {
        const details = card.querySelector('details');
        if (details) details.open = true;

        const toolbar = document.querySelector('.toolbar');
        let offset = 0;
        if (toolbar) {
          const rect = toolbar.getBoundingClientRect();
          offset = rect.height + 10; // 10px余裕
        }

        const cardRect = card.getBoundingClientRect();
        const scrollTop = window.pageYOffset + cardRect.top - offset;
        window.scrollTo({ top: scrollTop, behavior: 'smooth' });
      }
    }

    window.addEventListener('hashchange', openDetailsForHash);
    window.addEventListener('DOMContentLoaded', openDetailsForHash);

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
      directoryTooltip.textContent = slice.dataset.name;
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

    // Collapse all details by default on page load
    for (const item of document.querySelectorAll('details')) item.open = false;

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
