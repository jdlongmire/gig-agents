<script>
  import Sidebar from './Sidebar.svelte';
  import Header from './Header.svelte';

  let sidebarCollapsed = false;
  let isMobile = false;
  let sidebarOpen = true;

  function handleToggleSidebar() {
    if (isMobile) {
      sidebarOpen = !sidebarOpen;
    } else {
      sidebarCollapsed = !sidebarCollapsed;
    }
  }

  // Responsive: detect mobile on mount and window resize
  import { onMount } from 'svelte';

  onMount(() => {
    checkMobile();
    const handler = () => checkMobile();
    window.addEventListener('resize', handler);
    return () => window.removeEventListener('resize', handler);
  });

  function checkMobile() {
    isMobile = window.innerWidth < 768;
    if (isMobile) {
      sidebarOpen = false;
    } else {
      sidebarOpen = true;
    }
  }
</script>

<div class="layout" class:sidebar-collapsed={sidebarCollapsed}>
  <!-- Mobile overlay -->
  {#if isMobile && sidebarOpen}
    <div class="overlay" on:click={() => (sidebarOpen = false)} role="presentation"></div>
  {/if}

  <!-- Sidebar -->
  <div class="sidebar-wrap" class:mobile-open={isMobile && sidebarOpen} class:mobile-hidden={isMobile && !sidebarOpen}>
    <Sidebar collapsed={sidebarCollapsed && !isMobile} />
  </div>

  <!-- Main area -->
  <div class="main">
    <Header onToggleSidebar={handleToggleSidebar} />
    <main class="content">
      <slot />
    </main>
  </div>
</div>

<style>
  .layout {
    display: flex;
    height: 100vh;
    overflow: hidden;
    background: var(--ctp-base);
  }

  .sidebar-wrap {
    flex-shrink: 0;
    height: 100%;
    z-index: 100;
    transition: transform var(--transition);
  }

  .main {
    flex: 1;
    display: flex;
    flex-direction: column;
    min-width: 0;
    height: 100%;
    overflow: hidden;
  }

  .content {
    flex: 1;
    overflow-y: auto;
    background: var(--ctp-base);
  }

  .overlay {
    display: none;
  }

  /* Mobile styles */
  @media (max-width: 767px) {
    .sidebar-wrap {
      position: fixed;
      top: 0;
      left: 0;
      height: 100%;
      transform: translateX(-100%);
    }

    .sidebar-wrap.mobile-open {
      transform: translateX(0);
    }

    .sidebar-wrap.mobile-hidden {
      transform: translateX(-100%);
    }

    .overlay {
      display: block;
      position: fixed;
      inset: 0;
      background: rgba(0, 0, 0, 0.5);
      z-index: 99;
    }
  }
</style>
