<script>
  import { onMount } from 'svelte';
  import Router, { push } from 'svelte-spa-router';
  import { wrap } from 'svelte-spa-router/wrap';

  import { auth } from './lib/auth.js';
  import Layout from './components/Layout.svelte';
  import Login from './routes/Login.svelte';
  import Dashboard from './routes/Dashboard.svelte';
  import Contracts from './routes/Contracts.svelte';
  import Shells from './routes/Shells.svelte';
  import Settings from './routes/Settings.svelte';
  import NotFound from './routes/NotFound.svelte';

  // Auth guard — redirects to login if not authenticated
  function authGuard(detail) {
    if (!$auth.isAuthenticated && !$auth.loading) {
      push('/');
      return false;
    }
    return true;
  }

  // Route config
  const routes = {
    // Root: redirect to login or dashboard based on auth state
    '/': wrap({
      component: Login,
      conditions: [
        () => {
          // If authenticated, redirect to dashboard
          if ($auth.isAuthenticated && !$auth.loading) {
            push('/dashboard');
            return false;
          }
          return true;
        }
      ]
    }),

    '/dashboard': wrap({
      component: Dashboard,
      userData: { layout: true },
      conditions: [authGuard],
    }),

    '/contracts': wrap({
      component: Contracts,
      userData: { layout: true },
      conditions: [authGuard],
    }),

    '/shells': wrap({
      component: Shells,
      userData: { layout: true },
      conditions: [authGuard],
    }),

    '/settings': wrap({
      component: Settings,
      userData: { layout: true },
      conditions: [authGuard],
    }),

    '*': NotFound,
  };

  // Initialize auth on mount
  onMount(async () => {
    await auth.init();
  });

  $: isAuthenticated = $auth.isAuthenticated;
  $: loading = $auth.loading;

  // Current route needs layout wrapper
  let currentRouteNeedsLayout = false;

  function handleRouteLoaded(event) {
    currentRouteNeedsLayout = !!event.detail?.userData?.layout;
  }

  function handleConditionsFailed(event) {
    // Redirected by guard — already handled in conditions
  }
</script>

{#if loading}
  <!-- Boot splash -->
  <div class="boot-screen">
    <img src="/assets/crewport-logo.jpg" alt="CrewPort" class="boot-logo" />
    <div class="boot-spinner"></div>
  </div>
{:else if isAuthenticated}
  <Layout>
    <Router {routes}
      on:routeLoaded={handleRouteLoaded}
      on:conditionsFailed={handleConditionsFailed}
    />
  </Layout>
{:else}
  <Router {routes}
    on:routeLoaded={handleRouteLoaded}
    on:conditionsFailed={handleConditionsFailed}
  />
{/if}

<style>
  .boot-screen {
    height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 32px;
    background: var(--bg-primary);
  }

  .boot-logo {
    width: 160px;
    height: auto;
    object-fit: contain;
    opacity: 0.85;
    animation: fadeIn 0.4s ease;
  }

  .boot-spinner {
    width: 24px;
    height: 24px;
    border: 2px solid var(--border-color);
    border-top-color: var(--text-primary);
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }

  @keyframes fadeIn {
    from { opacity: 0; transform: scale(0.95); }
    to   { opacity: 0.85; transform: scale(1); }
  }
</style>
