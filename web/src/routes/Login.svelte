<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';
  import { auth } from '../lib/auth.js';

  // If already authenticated, redirect to dashboard
  onMount(() => {
    const unsub = auth.subscribe(state => {
      if (!state.loading && state.isAuthenticated) {
        push('/dashboard');
      }
    });
    return unsub;
  });
</script>

<div class="login-page">
  <div class="login-card">
    <div class="logo-wrap">
      <img src="/assets/crewport-logo.jpg" alt="CrewPort" class="logo" />
    </div>

    <div class="branding">
      <h1 class="title">CrewPort</h1>
      <p class="subtitle">Contract enforcement as a service for AI agent crews.</p>
    </div>

    <div class="auth-section">
      <a href="/auth/github" class="btn-github">
        <svg class="github-icon" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">
          <path d="M12 0C5.374 0 0 5.373 0 12c0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23A11.509 11.509 0 0 1 12 5.803c1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576C20.566 21.797 24 17.3 24 12c0-6.627-5.373-12-12-12z"/>
        </svg>
        Sign in with GitHub
      </a>
    </div>

    <p class="disclaimer">
      Access is restricted to authorized operators.
    </p>
  </div>
</div>

<style>
  .login-page {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: var(--ctp-crust);
    padding: 24px;
  }

  .login-card {
    width: 100%;
    max-width: 420px;
    background: var(--ctp-mantle);
    border: var(--border);
    border-radius: var(--radius);
    box-shadow: var(--shadow-lg);
    padding: 48px 40px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 32px;
  }

  .logo-wrap {
    width: 180px;
    height: 100px;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
    border-radius: var(--radius-sm);
  }

  .logo {
    width: 100%;
    height: 100%;
    object-fit: contain;
    filter: brightness(0.95);
  }

  .branding {
    text-align: center;
  }

  .title {
    font-size: 28px;
    font-weight: 700;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    color: var(--ctp-text);
    margin-bottom: 8px;
  }

  .subtitle {
    font-size: 13px;
    color: var(--ctp-subtext0);
    line-height: 1.5;
    max-width: 280px;
    margin: 0 auto;
  }

  .auth-section {
    width: 100%;
  }

  .btn-github {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
    width: 100%;
    padding: 12px 20px;
    background: var(--ctp-blue);
    color: var(--ctp-crust);
    border-radius: var(--radius-sm);
    font-size: 14px;
    font-weight: 600;
    letter-spacing: 0.02em;
    transition: all var(--transition);
    box-shadow: 0 2px 8px rgba(137, 180, 250, 0.2);
    text-decoration: none;
  }

  .btn-github:hover {
    opacity: 1;
    background: #a3c9fb;
    box-shadow: 0 4px 16px rgba(137, 180, 250, 0.35);
    transform: translateY(-1px);
  }

  .btn-github:active {
    transform: translateY(0);
  }

  .github-icon {
    width: 18px;
    height: 18px;
    flex-shrink: 0;
  }

  .disclaimer {
    font-size: 11px;
    color: var(--ctp-surface2);
    text-align: center;
  }

  @media (max-width: 480px) {
    .login-card {
      padding: 36px 24px;
    }

    .title {
      font-size: 24px;
    }
  }
</style>
