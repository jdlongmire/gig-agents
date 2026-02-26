<script>
  import { push } from 'svelte-spa-router';
</script>

<div class="not-found">
  <div class="content">
    <div class="code">404</div>
    <h1 class="title">Page not found</h1>
    <p class="desc">This route doesn't exist. The crew must have docked elsewhere.</p>
    <button class="btn-home" on:click={() => push('/dashboard')}>
      ← Back to dashboard
    </button>
  </div>
</div>

<style>
  .not-found {
    min-height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 40px;
  }

  .content {
    text-align: center;
    max-width: 400px;
  }

  .code {
    font-size: 80px;
    font-weight: 800;
    font-family: var(--font-mono);
    color: var(--bg-tertiary);
    line-height: 1;
    margin-bottom: 16px;
    letter-spacing: -0.04em;
  }

  .title {
    font-size: 22px;
    font-weight: 600;
    color: var(--text-primary);
    margin-bottom: 12px;
  }

  .desc {
    font-size: 14px;
    color: var(--text-tertiary);
    line-height: 1.5;
    margin-bottom: 32px;
  }

  .btn-home {
    padding: 10px 20px;
    background: var(--bg-tertiary);
    color: var(--text-primary);
    border: var(--border);
    border-radius: var(--radius-sm);
    font-size: 13px;
    font-weight: 600;
    transition: all var(--transition);
  }

  .btn-home:hover {
    background: var(--bg-secondary);
    border-color: var(--border-color);
  }
</style>
