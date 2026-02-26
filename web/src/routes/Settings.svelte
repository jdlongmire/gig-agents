<script>
  import { auth } from '../lib/auth.js';

  $: user = $auth.user;
</script>

<div class="page">
  <div class="page-header">
    <h2 class="page-title">Settings</h2>
    <p class="page-subtitle">Account and platform configuration</p>
  </div>

  <div class="settings-sections">
    <!-- Account -->
    <section class="settings-section">
      <h3 class="section-title">Account</h3>
      <div class="settings-card">
        {#if user}
          <div class="profile-row">
            {#if user.avatar_url}
              <img src={user.avatar_url} alt={user.display_name || user.github_username} class="avatar" />
            {:else}
              <div class="avatar-fallback">{(user.display_name || user.github_username || '?')[0].toUpperCase()}</div>
            {/if}
            <div class="profile-info">
              <div class="profile-name">{user.display_name || user.github_username}</div>
              <div class="profile-meta">
                <span class="text-muted">@{user.github_username}</span>
                <span class="badge {user.role === 'operator' ? 'badge-blue' : 'badge-teal'}">{user.role || 'client'}</span>
              </div>
            </div>
          </div>
        {:else}
          <div class="text-muted">Loading account info...</div>
        {/if}
      </div>
    </section>

    <!-- Danger Zone -->
    <section class="settings-section">
      <h3 class="section-title danger">Session</h3>
      <div class="settings-card">
        <div class="danger-row">
          <div>
            <div class="danger-label">Sign out</div>
            <div class="danger-desc">End your current session. You'll need to sign in with GitHub again.</div>
          </div>
          <button class="btn-danger" on:click={() => $auth.logout()}>Sign out</button>
        </div>
      </div>
    </section>

    <!-- Platform Info -->
    <section class="settings-section">
      <h3 class="section-title">Platform</h3>
      <div class="settings-card info-grid">
        <div class="info-item">
          <span class="info-label">Version</span>
          <span class="info-value text-mono">phase1-skeleton</span>
        </div>
        <div class="info-item">
          <span class="info-label">Stack</span>
          <span class="info-value">Go · SQLite · Svelte</span>
        </div>
        <div class="info-item">
          <span class="info-label">Status</span>
          <span class="badge badge-yellow">Pre-alpha</span>
        </div>
      </div>
    </section>
  </div>
</div>

<style>
  .page {
    padding: 32px;
    max-width: 720px;
  }

  .page-header {
    margin-bottom: 32px;
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

  .settings-sections {
    display: flex;
    flex-direction: column;
    gap: 28px;
  }

  .section-title {
    font-size: 12px;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.08em;
    color: var(--text-tertiary);
    margin-bottom: 10px;
  }

  .section-title.danger {
    color: var(--text-tertiary);
  }

  .settings-card {
    background: var(--bg-secondary);
    border: var(--border);
    border-radius: var(--radius);
    box-shadow: var(--shadow);
    padding: 20px;
  }

  .profile-row {
    display: flex;
    align-items: center;
    gap: 16px;
  }

  .avatar {
    width: 48px;
    height: 48px;
    border-radius: 50%;
    border: 2px solid var(--border-color);
    object-fit: cover;
    flex-shrink: 0;
  }

  .avatar-fallback {
    width: 48px;
    height: 48px;
    border-radius: 50%;
    background: var(--bg-tertiary);
    border: 2px solid var(--border-color);
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 700;
    font-size: 18px;
    color: var(--text-primary);
    flex-shrink: 0;
  }

  .profile-name {
    font-size: 15px;
    font-weight: 600;
    color: var(--text-primary);
    margin-bottom: 4px;
  }

  .profile-meta {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 13px;
  }

  .danger-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 16px;
  }

  .danger-label {
    font-size: 14px;
    font-weight: 600;
    color: var(--text-primary);
    margin-bottom: 4px;
  }

  .danger-desc {
    font-size: 12px;
    color: var(--text-tertiary);
    line-height: 1.4;
  }

  .btn-danger {
    padding: 8px 16px;
    background: var(--bg-tertiary);
    color: var(--text-primary);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-sm);
    font-size: 13px;
    font-weight: 600;
    white-space: nowrap;
    flex-shrink: 0;
  }

  .btn-danger:hover {
    background: var(--bg-secondary);
    border-color: var(--text-tertiary);
  }

  .info-grid {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .info-item {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .info-label {
    font-size: 13px;
    color: var(--text-tertiary);
  }

  .info-value {
    font-size: 13px;
    color: var(--text-primary);
  }

  @media (max-width: 600px) {
    .page {
      padding: 20px 16px;
    }

    .danger-row {
      flex-direction: column;
      align-items: flex-start;
    }
  }
</style>
