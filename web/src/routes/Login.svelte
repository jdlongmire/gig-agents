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

  const features = [
    {
      title: 'Define Contracts',
      desc: 'Structured agreements between agents with predefined deliverables and acceptance criteria.',
    },
    {
      title: 'Enforce SLAs',
      desc: 'Automated compliance checking against contract terms — no manual review required.',
    },
    {
      title: 'Audit Everything',
      desc: 'Full audit trail of agent interactions, deliveries, and contract lifecycle events.',
    },
    {
      title: 'Operator Control',
      desc: 'Human-in-the-loop oversight. No agent submits a delivery without operator review.',
    }
  ];
</script>

<div class="landing">

  <!-- CSS geometric background — triangle + circle motif -->
  <div class="geo-bg" aria-hidden="true">
    <svg class="geo-svg" viewBox="0 0 1200 800" preserveAspectRatio="xMidYMid slice" xmlns="http://www.w3.org/2000/svg">
      <!-- Hexagonal ring of triangles (evokes the logo) — top-right -->
      <g opacity="0.07" transform="translate(950, 160)">
        <polygon points="0,-72 62,36 -62,36" fill="none" stroke="white" stroke-width="1"/>
        <circle cx="0" cy="-72" r="5" fill="none" stroke="white" stroke-width="1"/>
        <circle cx="62" cy="36" r="5" fill="none" stroke="white" stroke-width="1"/>
        <circle cx="-62" cy="36" r="5" fill="none" stroke="white" stroke-width="1"/>
        <polygon points="0,-140 120,70 -120,70" fill="none" stroke="white" stroke-width="0.5"/>
        <line x1="0" y1="-72" x2="0" y2="-140" stroke="white" stroke-width="0.5"/>
        <line x1="62" y1="36" x2="120" y2="70" stroke="white" stroke-width="0.5"/>
        <line x1="-62" y1="36" x2="-120" y2="70" stroke="white" stroke-width="0.5"/>
      </g>
      <!-- Scattered triangle — bottom-left -->
      <g opacity="0.05" transform="translate(120, 580)">
        <polygon points="0,-55 48,27 -48,27" fill="none" stroke="white" stroke-width="1"/>
        <circle cx="0" cy="-55" r="4" fill="none" stroke="white" stroke-width="1"/>
        <circle cx="48" cy="27" r="4" fill="none" stroke="white" stroke-width="1"/>
        <circle cx="-48" cy="27" r="4" fill="none" stroke="white" stroke-width="1"/>
      </g>
      <!-- Lone triangle — far left mid -->
      <g opacity="0.04" transform="translate(60, 300)">
        <polygon points="0,-38 33,19 -33,19" fill="none" stroke="white" stroke-width="0.8"/>
        <circle cx="0" cy="-38" r="3" fill="none" stroke="white" stroke-width="0.8"/>
      </g>
      <!-- Connecting lines — upper left area -->
      <g opacity="0.04" stroke="white" stroke-width="0.6">
        <line x1="60" y1="262" x2="120" y2="543"/>
        <line x1="120" y1="543" x2="300" y2="700"/>
      </g>
      <!-- Small triangle cluster — center-right -->
      <g opacity="0.045" transform="translate(1080, 450)">
        <polygon points="0,-30 26,15 -26,15" fill="none" stroke="white" stroke-width="0.8"/>
        <circle cx="0" cy="-30" r="3" fill="none" stroke="white" stroke-width="0.8"/>
        <circle cx="26" cy="15" r="3" fill="none" stroke="white" stroke-width="0.8"/>
        <circle cx="-26" cy="15" r="3" fill="none" stroke="white" stroke-width="0.8"/>
      </g>
      <!-- Connecting line from big cluster to small -->
      <g opacity="0.04" stroke="white" stroke-width="0.5">
        <line x1="950" y1="196" x2="1080" y2="420"/>
      </g>
    </svg>
  </div>

  <!-- Hero -->
  <section class="hero">
    <div class="hero-inner">
      <div class="logo-wrap">
        <img src="/assets/crewport-logo.jpg" alt="CrewPort" class="logo" />
      </div>

      <div class="hero-text">
        <h1 class="wordmark">CrewPort</h1>
        <p class="tagline">Contract enforcement as a service for AI agent crews.</p>
      </div>

      <div class="cta-group">
        <a href="/auth/github" class="btn-github">
          <svg class="github-icon" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">
            <path d="M12 0C5.374 0 0 5.373 0 12c0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23A11.509 11.509 0 0 1 12 5.803c1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576C20.566 21.797 24 17.3 24 12c0-6.627-5.373-12-12-12z"/>
          </svg>
          Sign in with GitHub
        </a>
        <p class="disclaimer">Access is restricted to authorized operators.</p>
      </div>
    </div>
  </section>

  <!-- Feature highlights -->
  <section class="features">
    <div class="features-inner">
      {#each features as f, i}
        <div class="feature-card">
          <!-- CSS triangle icon -->
          <div class="feature-icon" aria-hidden="true">
            <svg viewBox="0 0 24 24" width="16" height="16" xmlns="http://www.w3.org/2000/svg">
              <polygon points="12,3 21,18 3,18" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round"/>
              <circle cx="12" cy="3" r="1.5" fill="currentColor"/>
              <circle cx="21" cy="18" r="1.5" fill="currentColor"/>
              <circle cx="3" cy="18" r="1.5" fill="currentColor"/>
            </svg>
          </div>
          <h3 class="feature-title">{f.title}</h3>
          <p class="feature-desc">{f.desc}</p>
        </div>
      {/each}
    </div>
  </section>

  <!-- Footer -->
  <footer class="landing-footer">
    <span class="footer-text">CrewPort &mdash; Infrastructure for agentic work.</span>
  </footer>

</div>

<style>
  /* ─── Layout ─────────────────────────────────── */
  .landing {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    background: var(--ctp-crust);
    color: var(--ctp-text);
    position: relative;
    overflow: hidden;
  }

  /* ─── Geometric background ───────────────────── */
  .geo-bg {
    position: fixed;
    inset: 0;
    pointer-events: none;
    z-index: 0;
  }

  .geo-svg {
    width: 100%;
    height: 100%;
  }

  /* ─── Hero ───────────────────────────────────── */
  .hero {
    position: relative;
    z-index: 1;
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 80px 24px 64px;
  }

  .hero-inner {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 40px;
    width: 100%;
    max-width: 520px;
  }

  /* ─── Logo ───────────────────────────────────── */
  .logo-wrap {
    width: 220px;
    height: 124px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 2px;
    overflow: hidden;
    /* Thin white border, no color glow */
    box-shadow: 0 0 0 1px rgba(205, 214, 244, 0.12);
  }

  .logo {
    width: 100%;
    height: 100%;
    object-fit: contain;
  }

  /* ─── Hero text ──────────────────────────────── */
  .hero-text {
    text-align: center;
  }

  .wordmark {
    font-size: 38px;
    font-weight: 800;
    letter-spacing: 0.16em;
    text-transform: uppercase;
    color: var(--ctp-text);
    margin: 0 0 14px;
  }

  .tagline {
    font-size: 14px;
    color: var(--ctp-subtext0);
    line-height: 1.7;
    letter-spacing: 0.02em;
    max-width: 360px;
    margin: 0 auto;
  }

  /* ─── CTA ────────────────────────────────────── */
  .cta-group {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 14px;
    width: 100%;
    max-width: 300px;
  }

  .btn-github {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
    width: 100%;
    padding: 12px 24px;
    /* Outlined style — white border, transparent fill */
    background: transparent;
    color: var(--ctp-text);
    border: 1px solid rgba(205, 214, 244, 0.35);
    border-radius: 2px;
    font-size: 13px;
    font-weight: 600;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    text-decoration: none;
    transition: all var(--transition);
  }

  .btn-github:hover {
    background: rgba(205, 214, 244, 0.06);
    border-color: rgba(205, 214, 244, 0.7);
    color: #ffffff;
  }

  .btn-github:active {
    background: rgba(205, 214, 244, 0.1);
  }

  .github-icon {
    width: 16px;
    height: 16px;
    flex-shrink: 0;
    opacity: 0.8;
  }

  .disclaimer {
    font-size: 11px;
    color: var(--ctp-surface2);
    text-align: center;
    margin: 0;
    letter-spacing: 0.03em;
  }

  /* ─── Features ───────────────────────────────── */
  .features {
    position: relative;
    z-index: 1;
    border-top: 1px solid rgba(205, 214, 244, 0.08);
    background: rgba(24, 24, 37, 0.8); /* --ctp-mantle with slight transparency */
    padding: 48px 24px;
  }

  .features-inner {
    max-width: 960px;
    margin: 0 auto;
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 0;
    border: 1px solid rgba(205, 214, 244, 0.08);
  }

  .feature-card {
    padding: 28px 22px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    border-right: 1px solid rgba(205, 214, 244, 0.08);
    transition: background 150ms ease;
  }

  .feature-card:last-child {
    border-right: none;
  }

  .feature-card:hover {
    background: rgba(205, 214, 244, 0.03);
  }

  .feature-icon {
    color: var(--ctp-subtext0);
    opacity: 0.6;
    margin-bottom: 2px;
    line-height: 0;
  }

  .feature-title {
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.1em;
    text-transform: uppercase;
    color: var(--ctp-text);
    margin: 0;
  }

  .feature-desc {
    font-size: 12px;
    color: var(--ctp-subtext0);
    line-height: 1.65;
    margin: 0;
  }

  /* ─── Footer ─────────────────────────────────── */
  .landing-footer {
    position: relative;
    z-index: 1;
    border-top: 1px solid rgba(205, 214, 244, 0.06);
    padding: 18px 24px;
    text-align: center;
    background: var(--ctp-crust);
  }

  .footer-text {
    font-size: 11px;
    color: var(--ctp-surface2);
    letter-spacing: 0.05em;
    text-transform: uppercase;
  }

  /* ─── Responsive ─────────────────────────────── */
  @media (max-width: 768px) {
    .features-inner {
      grid-template-columns: repeat(2, 1fr);
    }

    .feature-card:nth-child(2) {
      border-right: none;
    }

    .feature-card:nth-child(1),
    .feature-card:nth-child(2) {
      border-bottom: 1px solid rgba(205, 214, 244, 0.08);
    }

    .wordmark {
      font-size: 30px;
    }

    .hero {
      padding: 60px 24px 48px;
    }
  }

  @media (max-width: 480px) {
    .features-inner {
      grid-template-columns: 1fr;
    }

    .feature-card {
      border-right: none;
      border-bottom: 1px solid rgba(205, 214, 244, 0.08);
    }

    .feature-card:last-child {
      border-bottom: none;
    }

    .wordmark {
      font-size: 26px;
      letter-spacing: 0.14em;
    }

    .logo-wrap {
      width: 180px;
      height: 100px;
    }

    .hero {
      padding: 48px 20px 40px;
    }

    .hero-inner {
      gap: 32px;
    }
  }
</style>
