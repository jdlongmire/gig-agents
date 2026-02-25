<script>
  // Contracts page — list and manage contracts
  let filter = 'all';

  const filters = [
    { id: 'all', label: 'All' },
    { id: 'open', label: 'Open' },
    { id: 'active', label: 'In Progress' },
    { id: 'completed', label: 'Completed' },
  ];
</script>

<div class="page">
  <div class="page-header">
    <div>
      <h2 class="page-title">Contracts</h2>
      <p class="page-subtitle">Structured work agreements with class-based deliverables</p>
    </div>
    <button class="btn-primary" disabled>+ Post Contract</button>
  </div>

  <div class="filter-bar">
    {#each filters as f}
      <button
        class="filter-btn"
        class:active={filter === f.id}
        on:click={() => (filter = f.id)}
      >
        {f.label}
      </button>
    {/each}
  </div>

  <div class="empty-state">
    <div class="empty-icon">📋</div>
    <h3>No contracts yet</h3>
    <p>Post your first contract to get started. Contracts define deliverables,<br />acceptance criteria, and payment terms.</p>
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
    color: var(--ctp-text);
    letter-spacing: 0.02em;
    margin-bottom: 4px;
  }

  .page-subtitle {
    font-size: 13px;
    color: var(--ctp-subtext0);
  }

  .btn-primary {
    padding: 8px 16px;
    background: var(--ctp-blue);
    color: var(--ctp-crust);
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

  .filter-bar {
    display: flex;
    gap: 4px;
    margin-bottom: 24px;
    padding: 4px;
    background: var(--ctp-mantle);
    border: var(--border);
    border-radius: var(--radius);
    width: fit-content;
  }

  .filter-btn {
    padding: 6px 14px;
    border-radius: var(--radius-sm);
    font-size: 13px;
    color: var(--ctp-subtext0);
    background: transparent;
    transition: all var(--transition);
  }

  .filter-btn:hover {
    color: var(--ctp-text);
    background: var(--ctp-surface0);
  }

  .filter-btn.active {
    color: var(--ctp-text);
    background: var(--ctp-surface0);
    box-shadow: var(--shadow);
  }

  .empty-state {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 80px 24px;
    text-align: center;
    color: var(--ctp-subtext0);
    gap: 12px;
  }

  .empty-icon {
    font-size: 40px;
    opacity: 0.3;
    margin-bottom: 8px;
  }

  .empty-state h3 {
    font-size: 16px;
    font-weight: 600;
    color: var(--ctp-subtext1);
  }

  .empty-state p {
    font-size: 13px;
    line-height: 1.6;
  }

  @media (max-width: 600px) {
    .page {
      padding: 20px 16px;
    }

    .page-header {
      flex-direction: column;
    }
  }
</style>
