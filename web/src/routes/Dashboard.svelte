<script>
  import { auth } from '../lib/auth.js';
</script>

<div class="page">
  <div class="page-header">
    <h2 class="page-title">Dashboard</h2>
    <p class="page-subtitle">Mission control overview</p>
  </div>

  <div class="stats-grid">
    <div class="stat-card">
      <div class="stat-label">Active Crews</div>
      <div class="stat-value">—</div>
    </div>
    <div class="stat-card">
      <div class="stat-label">Open Contracts</div>
      <div class="stat-value">—</div>
    </div>
    <div class="stat-card">
      <div class="stat-label">In Progress</div>
      <div class="stat-value">—</div>
    </div>
    <div class="stat-card">
      <div class="stat-label">Completed</div>
      <div class="stat-value">—</div>
    </div>
  </div>

  <div class="panel-grid">
    <div class="panel">
      <div class="panel-header">
        <span class="panel-title">Recent Contracts</span>
        <a href="#/contracts" class="panel-link">View all →</a>
      </div>
      <div class="panel-empty">
        <span class="empty-icon">📋</span>
        <p>No contracts yet.</p>
      </div>
    </div>

    <div class="panel">
      <div class="panel-header">
        <span class="panel-title">Active Shells</span>
        <a href="#/shells" class="panel-link">View all →</a>
      </div>
      <div class="panel-empty">
        <span class="empty-icon">⬡</span>
        <p>No shells registered.</p>
      </div>
    </div>
  </div>
</div>

<style>
  .page {
    padding: 32px;
    max-width: 1200px;
  }

  .page-header {
    margin-bottom: 32px;
  }

  .page-title {
    font-size: 22px;
    font-weight: 600;
    color: var(--ctp-text);
    letter-spacing: 0.02em;
    margin-bottom: 4px;
  }

  .page-subtitle {
    font-size: 13px;
    color: var(--ctp-subtext0);
  }

  .stats-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 16px;
    margin-bottom: 32px;
  }

  .stat-card {
    background: var(--ctp-mantle);
    border: var(--border);
    border-radius: var(--radius);
    padding: 20px;
    box-shadow: var(--shadow);
  }

  .stat-label {
    font-size: 12px;
    color: var(--ctp-subtext0);
    text-transform: uppercase;
    letter-spacing: 0.06em;
    margin-bottom: 8px;
  }

  .stat-value {
    font-size: 28px;
    font-weight: 700;
    color: var(--ctp-text);
    font-family: var(--font-mono);
  }

  .panel-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 16px;
  }

  .panel {
    background: var(--ctp-mantle);
    border: var(--border);
    border-radius: var(--radius);
    box-shadow: var(--shadow);
    overflow: hidden;
  }

  .panel-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px 20px;
    border-bottom: var(--border);
  }

  .panel-title {
    font-size: 13px;
    font-weight: 600;
    color: var(--ctp-subtext1);
    text-transform: uppercase;
    letter-spacing: 0.06em;
  }

  .panel-link {
    font-size: 12px;
    color: var(--ctp-blue);
  }

  .panel-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 48px 24px;
    color: var(--ctp-subtext0);
    gap: 12px;
  }

  .empty-icon {
    font-size: 28px;
    opacity: 0.4;
  }

  .panel-empty p {
    font-size: 13px;
  }

  @media (max-width: 900px) {
    .stats-grid {
      grid-template-columns: repeat(2, 1fr);
    }

    .panel-grid {
      grid-template-columns: 1fr;
    }
  }

  @media (max-width: 600px) {
    .page {
      padding: 20px 16px;
    }

    .stats-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }
</style>
