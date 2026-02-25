<script>
  import { location } from 'svelte-spa-router';
  import { auth } from '../lib/auth.js';

  export let onToggleSidebar = () => {};

  const pageTitles = {
    '/':          'Dashboard',
    '/dashboard': 'Dashboard',
    '/contracts': 'Contracts',
    '/shells':    'Shells',
    '/settings':  'Settings',
  };

  $: pageTitle = pageTitles[$location] || 'CrewPort';
  $: user = $auth.user;
</script>

<header class="header">
  <div class="header-left">
    <button class="toggle-btn" on:click={onToggleSidebar} title="Toggle sidebar" aria-label="Toggle sidebar">
      <svg viewBox="0 0 20 20" fill="currentColor" class="menu-icon">
        <rect x="2" y="4"  width="16" height="2" rx="1"/>
        <rect x="2" y="9"  width="16" height="2" rx="1"/>
        <rect x="2" y="14" width="16" height="2" rx="1"/>
      </svg>
    </button>

    <h1 class="page-title">{pageTitle}</h1>
  </div>

  <div class="header-right">
    {#if user}
      <div class="user-info">
        <span class="username">@{user.github_username}</span>
        {#if user.role}
          <span class="badge {user.role === 'operator' ? 'badge-blue' : 'badge-teal'}">{user.role}</span>
        {/if}
        {#if user.avatar_url}
          <img src={user.avatar_url} alt={user.display_name || user.github_username} class="avatar" />
        {:else}
          <div class="avatar-fallback">
            {(user.display_name || user.github_username || '?')[0].toUpperCase()}
          </div>
        {/if}
      </div>

      <button class="logout-btn" on:click={() => $auth.logout()} title="Sign out">
        <svg viewBox="0 0 20 20" fill="currentColor" class="logout-icon">
          <path fill-rule="evenodd" d="M3 3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h7a1 1 0 1 0 0-2H4V5h6a1 1 0 1 0 0-2H3zm11.293 4.293a1 1 0 0 1 1.414 0l3 3a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414-1.414L15.586 12H9a1 1 0 1 1 0-2h6.586l-1.293-1.293a1 1 0 0 1 0-1.414z" clip-rule="evenodd"/>
        </svg>
      </button>
    {/if}
  </div>
</header>

<style>
  .header {
    height: var(--header-height);
    background: var(--ctp-mantle);
    border-bottom: var(--border);
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 20px;
    flex-shrink: 0;
    gap: 16px;
  }

  .header-left {
    display: flex;
    align-items: center;
    gap: 12px;
    min-width: 0;
  }

  .toggle-btn {
    width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: transparent;
    color: var(--ctp-subtext0);
    border-radius: var(--radius-sm);
    flex-shrink: 0;
    transition: all var(--transition);
  }

  .toggle-btn:hover {
    background: var(--ctp-surface0);
    color: var(--ctp-text);
  }

  .menu-icon {
    width: 16px;
    height: 16px;
  }

  .page-title {
    font-size: 15px;
    font-weight: 600;
    color: var(--ctp-text);
    letter-spacing: 0.02em;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .header-right {
    display: flex;
    align-items: center;
    gap: 12px;
    flex-shrink: 0;
  }

  .user-info {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .username {
    font-size: 13px;
    color: var(--ctp-subtext1);
    font-family: var(--font-mono);
  }

  .avatar {
    width: 28px;
    height: 28px;
    border-radius: 50%;
    border: 1px solid var(--ctp-surface1);
    object-fit: cover;
  }

  .avatar-fallback {
    width: 28px;
    height: 28px;
    border-radius: 50%;
    background: var(--ctp-surface0);
    border: 1px solid var(--ctp-surface1);
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    font-weight: 700;
    color: var(--ctp-blue);
  }

  .logout-btn {
    width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: transparent;
    color: var(--ctp-subtext0);
    border-radius: var(--radius-sm);
    transition: all var(--transition);
  }

  .logout-btn:hover {
    background: rgba(243, 139, 168, 0.1);
    color: var(--ctp-red);
  }

  .logout-icon {
    width: 16px;
    height: 16px;
  }

  @media (max-width: 480px) {
    .username {
      display: none;
    }

    .header {
      padding: 0 12px;
    }
  }
</style>
