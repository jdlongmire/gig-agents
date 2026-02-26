<script>
  import { onMount } from 'svelte';

  let isDark = true;

  onMount(() => {
    // Read the theme already set by main.js (no flash)
    isDark = document.documentElement.getAttribute('data-theme') !== 'light';
  });

  function toggle() {
    isDark = !isDark;
    const theme = isDark ? 'dark' : 'light';
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('crewport-theme', theme);
  }
</script>

<button
  class="theme-toggle"
  on:click={toggle}
  aria-label={isDark ? 'Switch to light theme' : 'Switch to dark theme'}
  title={isDark ? 'Light mode' : 'Dark mode'}
>
  {#if isDark}
    <!-- Sun icon — shown in dark mode to switch to light -->
    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
      <circle cx="12" cy="12" r="5"/>
      <line x1="12" y1="1" x2="12" y2="3"/>
      <line x1="12" y1="21" x2="12" y2="23"/>
      <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"/>
      <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"/>
      <line x1="1" y1="12" x2="3" y2="12"/>
      <line x1="21" y1="12" x2="23" y2="12"/>
      <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"/>
      <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"/>
    </svg>
  {:else}
    <!-- Moon icon — shown in light mode to switch to dark -->
    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
      <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"/>
    </svg>
  {/if}
</button>

<style>
  .theme-toggle {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 36px;
    height: 36px;
    border-radius: var(--radius-sm);
    background: transparent;
    color: var(--text-secondary);
    border: 1px solid transparent;
    cursor: pointer;
    transition: color var(--transition), background var(--transition), border-color var(--transition);
  }

  .theme-toggle:hover {
    color: var(--text-primary);
    background: var(--bg-tertiary);
    border-color: var(--border-color);
  }

  .theme-toggle:active {
    transform: scale(0.92);
  }

  .theme-toggle svg {
    display: block;
    flex-shrink: 0;
  }
</style>
