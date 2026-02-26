<script>
  import { location } from 'svelte-spa-router';

  export let collapsed = false;

  const navItems = [
    { href: '/dashboard', label: 'Dashboard', icon: '⊞' },
    { href: '/contracts', label: 'Contracts', icon: '📋' },
    { href: '/shells',    label: 'Shells',     icon: '⬡' },
    { href: '/settings',  label: 'Settings',   icon: '⚙' },
  ];

  function isActive(href) {
    return $location === href || ($location === '/' && href === '/dashboard');
  }
</script>

<aside class="sidebar" class:collapsed>
  <!-- Logo / Brand -->
  <div class="sidebar-brand">
    <img src="/assets/crewport-logo.jpg" alt="CrewPort" class="brand-logo" />
    {#if !collapsed}
      <span class="brand-name">CrewPort</span>
    {/if}
  </div>

  <!-- Nav -->
  <nav class="sidebar-nav">
    {#each navItems as item}
      <a
        href={`#${item.href}`}
        class="nav-item"
        class:active={isActive(item.href)}
        title={collapsed ? item.label : ''}
      >
        <span class="nav-icon" aria-hidden="true">{item.icon}</span>
        {#if !collapsed}
          <span class="nav-label">{item.label}</span>
        {/if}
      </a>
    {/each}
  </nav>

  <!-- Bottom area -->
  <div class="sidebar-footer">
    <div class="status-dot" title="Connected"></div>
    {#if !collapsed}
      <span class="status-label">Online</span>
    {/if}
  </div>
</aside>

<style>
  .sidebar {
    width: var(--sidebar-width);
    min-width: var(--sidebar-width);
    height: 100%;
    background: var(--bg-secondary);
    border-right: var(--border);
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
    transition: width var(--transition), min-width var(--transition);
    overflow: hidden;
  }

  .sidebar.collapsed {
    width: 56px;
    min-width: 56px;
  }

  /* Brand */
  .sidebar-brand {
    height: var(--header-height);
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 0 14px;
    border-bottom: var(--border);
    flex-shrink: 0;
    overflow: hidden;
  }

  .brand-logo {
    width: 28px;
    height: 28px;
    object-fit: contain;
    border-radius: 4px;
    flex-shrink: 0;
    filter: brightness(0.9);
  }

  .brand-name {
    font-size: 14px;
    font-weight: 700;
    letter-spacing: 0.1em;
    text-transform: uppercase;
    color: var(--text-primary);
    white-space: nowrap;
  }

  /* Nav */
  .sidebar-nav {
    flex: 1;
    padding: 12px 8px;
    display: flex;
    flex-direction: column;
    gap: 2px;
    overflow-y: auto;
  }

  .nav-item {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 9px 10px;
    border-radius: var(--radius-sm);
    color: var(--text-tertiary);
    text-decoration: none;
    font-size: 13px;
    font-weight: 500;
    transition: all var(--transition);
    border-left: 2px solid transparent;
    white-space: nowrap;
    overflow: hidden;
  }

  .nav-item:hover {
    color: var(--text-primary);
    background: var(--bg-tertiary);
    opacity: 1;
  }

  .nav-item.active {
    color: var(--text-primary);
    background: var(--bg-tertiary);
    border-left-color: var(--text-primary);
  }

  .nav-icon {
    font-size: 15px;
    flex-shrink: 0;
    width: 18px;
    text-align: center;
  }

  .nav-label {
    flex: 1;
  }

  /* Footer */
  .sidebar-footer {
    padding: 12px 14px;
    border-top: var(--border);
    display: flex;
    align-items: center;
    gap: 8px;
    flex-shrink: 0;
  }

  .status-dot {
    width: 7px;
    height: 7px;
    border-radius: 50%;
    background: var(--text-secondary);
    box-shadow: 0 0 6px var(--border-color);
    flex-shrink: 0;
  }

  .status-label {
    font-size: 11px;
    color: var(--text-tertiary);
    text-transform: uppercase;
    letter-spacing: 0.06em;
  }
</style>
