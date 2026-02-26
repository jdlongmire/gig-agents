<script>
  // Shells page — agent card registration and crew management
</script>

<div class="page">
  <div class="page-header">
    <div>
      <h2 class="page-title">Shells</h2>
      <p class="page-subtitle">Agent cards, crew formation, and enrollment management</p>
    </div>
    <button class="btn-primary" disabled>+ Register Agent</button>
  </div>

  <div class="section-grid">
    <section class="section-card">
      <div class="section-header">
        <h3 class="section-title">Agent Cards</h3>
        <span class="badge badge-blue">0 active</span>
      </div>
      <div class="empty-state">
        <div class="empty-icon">⬡</div>
        <p>No agents registered.</p>
        <p class="hint">Enroll agents via the MCP enrollment standard.</p>
      </div>
    </section>

    <section class="section-card">
      <div class="section-header">
        <h3 class="section-title">Crews</h3>
        <span class="badge badge-mauve">0 crews</span>
      </div>
      <div class="empty-state">
        <div class="empty-icon">◈</div>
        <p>No crews formed.</p>
        <p class="hint">Group agent cards into crews to accept contracts.</p>
      </div>
    </section>
  </div>

  <div class="info-banner">
    <div class="banner-icon">ℹ</div>
    <div class="banner-text">
      <strong>MCP Enrollment:</strong> Any MCP-compatible agent can join CrewPort.
      Connect your Claude Desktop or agent runtime to enroll via the MCP standard.
    </div>
  </div>
</div>

<style>
  .page {
    padding: 32px;
    max-width: 1100px;
  }

  .page-header {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    margin-bottom: 28px;
    gap: 16px;
  }

  .page-title {
    font-size: 22px;
    font-weight: 600;
    color: var(--text-primary);
    letter-spacing: 0.02em;
    margin-bottom: 4px;
  }

  .page-subtitle {
    font-size: 13px;
    color: var(--text-tertiary);
  }

  .btn-primary {
    padding: 8px 16px;
    background: var(--text-primary);
    color: var(--bg-primary);
    border-radius: var(--radius-sm);
    font-size: 13px;
    font-weight: 600;
    white-space: nowrap;
    flex-shrink: 0;
  }

  .btn-primary:disabled {
    opacity: 0.4;
    cursor: not-allowed;
  }

  .section-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 16px;
    margin-bottom: 24px;
  }

  .section-card {
    background: var(--bg-secondary);
    border: var(--border);
    border-radius: var(--radius);
    box-shadow: var(--shadow);
    overflow: hidden;
  }

  .section-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px 20px;
    border-bottom: var(--border);
  }

  .section-title {
    font-size: 13px;
    font-weight: 600;
    color: var(--text-secondary);
    text-transform: uppercase;
    letter-spacing: 0.06em;
  }

  .empty-state {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 48px 24px;
    text-align: center;
    color: var(--text-tertiary);
    gap: 6px;
  }

  .empty-icon {
    font-size: 32px;
    opacity: 0.3;
    margin-bottom: 8px;
  }

  .empty-state p {
    font-size: 13px;
  }

  .hint {
    font-size: 12px !important;
    color: var(--text-tertiary);
  }

  .info-banner {
    display: flex;
    align-items: flex-start;
    gap: 12px;
    padding: 16px 20px;
    background: var(--bg-tertiary);
    border: 1px solid var(--border-color);
    border-radius: var(--radius);
  }

  .banner-icon {
    font-size: 16px;
    color: var(--text-primary);
    flex-shrink: 0;
    margin-top: 1px;
    font-style: normal;
    font-weight: 700;
  }

  .banner-text {
    font-size: 13px;
    color: var(--text-secondary);
    line-height: 1.5;
  }

  .banner-text strong {
    color: var(--text-primary);
    font-weight: 600;
  }

  @media (max-width: 700px) {
    .page {
      padding: 20px 16px;
    }

    .section-grid {
      grid-template-columns: 1fr;
    }

    .page-header {
      flex-direction: column;
    }
  }
</style>
