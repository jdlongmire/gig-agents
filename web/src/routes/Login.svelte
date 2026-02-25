<script>
  import { onMount, onDestroy } from 'svelte';
  import { push } from 'svelte-spa-router';
  import { auth } from '../lib/auth.js';

  // Redirect if already authenticated
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
      desc: 'Set expectations for every agent task with structured, enforceable agreements.',
    },
    {
      title: 'Enforce SLAs',
      desc: 'Automated compliance monitoring — no delivery passes without meeting terms.',
    },
    {
      title: 'Audit Everything',
      desc: 'Full trail of agent actions, deliveries, and contract lifecycle events.',
    },
    {
      title: 'Operator Control',
      desc: 'You stay in command, always. No agent delivers without your sign-off.',
    }
  ];

  // ─── Agent Launch Animation ───────────────────────────────────────────────
  // The wheel has 6 triangles (crews). Periodically one "activates" and deploys
  // 2-3 agent circles (the hollow circles inside the logo triangles) rightward.
  // Wheel SVG coords: viewBox "0 0 800 800", center (400,400), radius ~320
  // The wheel is rendered at 150vh tall, positioned left:-75vh so center is at x=0
  // In viewport space, wheel center is at left edge. Right half visible.

  let agents = [];
  let agentId = 0;
  let launchInterval;
  // Index of the currently "active" (glowing) triangle on the wheel (0-5, or -1 for none)
  let activeTriangle = -1;

  // The 6 triangle positions on the wheel (angles in degrees, 0=right, CCW)
  // We pick from the visible right half: triangles 0 (0°), 1 (60°), 5 (300°) are most visible
  const triangleAngles = [0, 60, 120, 180, 240, 300];

  function spawnDeployment() {
    // Count existing agents — cap at 12 (up to 3 per event × 4 concurrent events)
    if (agents.length >= 12) return;

    // Pick a triangle to "activate" — bias toward the visible right half
    const visibleTriangles = [0, 1, 5]; // 0°, 60°, 300° — on the right side
    const triIdx = visibleTriangles[Math.floor(Math.random() * visibleTriangles.length)];
    const triAngleDeg = triangleAngles[triIdx];

    // Flash this triangle
    activeTriangle = triIdx;
    setTimeout(() => { activeTriangle = -1; }, 700);

    // Wheel radius in vh: wheel is 150vh so radius = 75vh, use 72vh to stay on visible rim
    const wheelRadiusVh = 72;
    const triAngleRad = triIdx === 0 ? 0
      : triIdx === 1 ? -60 * (Math.PI / 180)
      : triIdx === 5 ? 60 * (Math.PI / 180)
      : triAngleDeg * (Math.PI / 180);

    // Spawn point: right edge of the selected triangle on the wheel
    const baseSpawnX = wheelRadiusVh * Math.cos(triAngleRad);
    const baseSpawnY = 50 + wheelRadiusVh * Math.sin(triAngleRad);

    // Spawn 2-3 circles per deployment, slightly spread out
    const count = 2 + Math.floor(Math.random() * 2); // 2 or 3
    for (let i = 0; i < count; i++) {
      // Slight spread around spawn point so circles don't stack perfectly
      const spreadX = (Math.random() - 0.5) * 3; // ±1.5vw spread
      const spreadY = (Math.random() - 0.5) * 4; // ±2vh spread

      const spawnX = baseSpawnX + spreadX;
      const spawnY = baseSpawnY + spreadY;

      // Travel angle: mostly rightward, slight random vertical variation
      const travelAngle = (Math.random() * 40 - 20) * (Math.PI / 180);
      const duration = 8000 + Math.random() * 4000; // 8-12 seconds
      const size = 10 + Math.random() * 6; // 10-16px diameter

      const id = ++agentId;
      agents = [...agents, {
        id,
        spawnX,
        spawnY,
        travelAngle,
        duration,
        size,
        born: Date.now(),
      }];

      // Cleanup after animation completes
      setTimeout(() => {
        agents = agents.filter(a => a.id !== id);
      }, duration + 200);
    }
  }

  onMount(() => {
    // Stagger first two deployment events
    setTimeout(() => spawnDeployment(), 1500);
    setTimeout(() => spawnDeployment(), 4000);
    launchInterval = setInterval(() => spawnDeployment(), 4000 + Math.random() * 2000);
  });

  onDestroy(() => {
    clearInterval(launchInterval);
  });

  function agentStyle(agent) {
    const dx = Math.cos(agent.travelAngle) * 100; // vw travel distance
    const dy = Math.sin(agent.travelAngle) * 100; // vh travel
    return `
      left: ${agent.spawnX}vw;
      top: ${agent.spawnY}vh;
      width: ${agent.size}px;
      height: ${agent.size}px;
      --dx: ${dx}vw;
      --dy: ${dy}vh;
      animation-duration: ${agent.duration}ms;
    `;
  }
</script>

<!-- ─── Page Root ─────────────────────────────────────────────────────────── -->
<div class="landing">

  <!-- ─── Background: Animated Wheel Layer ─────────────────────────────── -->
  <div class="wheel-layer" aria-hidden="true">
    <!--
      SVG viewBox: 800×800, center at 400,400
      Hexagonal ring: 6 equilateral triangles arranged in a ring
      Each triangle apex points outward from center
      Ring radius (center-to-triangle-center): ~300
      Triangle size: ~110px side length

      Angle math for 6 positions (0°=right, going counterclockwise):
        0:   0°   → right
        1:  60°   → upper-right
        2: 120°   → upper-left
        3: 180°   → left
        4: 240°   → lower-left
        5: 300°   → lower-right
    -->
    <svg
      class="wheel-svg"
      viewBox="0 0 800 800"
      xmlns="http://www.w3.org/2000/svg"
      preserveAspectRatio="xMidYMid meet"
    >
      <!-- The entire wheel rotates counterclockwise -->
      <g class="wheel-ring" transform="translate(400,400)">

        <!--
          Each triangle group:
          - Centered at ring radius ~300 from origin
          - Rotated so apex points outward
          - Size: equilateral triangle with circumradius ~90 (side ~156)
          - Inner circles: 3 small hollow circles in triangle arrangement
          - Connection lines to adjacent triangles

          Triangle circumradius R=90: vertices at
            apex (outward):  (0, -90)
            base-right:      (78, 45)
            base-left:       (-78, 45)
          (rotated so apex points away from center)
        -->

        <!-- ── Triangle 0: right (0°) — index 0 ── -->
        <g transform="rotate(0)" class={activeTriangle === 0 ? 'tri-active' : ''}>
          <g transform="translate(300, 0) rotate(90)">
            <polygon
              points="0,-90 78,45 -78,45"
              fill="none"
              stroke="white"
              stroke-width="1.5"
              stroke-linejoin="round"
            />
            <circle cx="0" cy="-60" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="-52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
          </g>
        </g>

        <!-- ── Triangle 1: upper-right (60°) — index 1 ── -->
        <g transform="rotate(-60)" class={activeTriangle === 1 ? 'tri-active' : ''}>
          <g transform="translate(300, 0) rotate(90)">
            <polygon
              points="0,-90 78,45 -78,45"
              fill="none"
              stroke="white"
              stroke-width="1.5"
              stroke-linejoin="round"
            />
            <circle cx="0" cy="-60" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="-52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
          </g>
        </g>

        <!-- ── Triangle 2: upper-left (120°) — index 2 ── -->
        <g transform="rotate(-120)" class={activeTriangle === 2 ? 'tri-active' : ''}>
          <g transform="translate(300, 0) rotate(90)">
            <polygon
              points="0,-90 78,45 -78,45"
              fill="none"
              stroke="white"
              stroke-width="1.5"
              stroke-linejoin="round"
            />
            <circle cx="0" cy="-60" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="-52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
          </g>
        </g>

        <!-- ── Triangle 3: left (180°) — index 3 ── -->
        <g transform="rotate(-180)" class={activeTriangle === 3 ? 'tri-active' : ''}>
          <g transform="translate(300, 0) rotate(90)">
            <polygon
              points="0,-90 78,45 -78,45"
              fill="none"
              stroke="white"
              stroke-width="1.5"
              stroke-linejoin="round"
            />
            <circle cx="0" cy="-60" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="-52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
          </g>
        </g>

        <!-- ── Triangle 4: lower-left (240°) — index 4 ── -->
        <g transform="rotate(-240)" class={activeTriangle === 4 ? 'tri-active' : ''}>
          <g transform="translate(300, 0) rotate(90)">
            <polygon
              points="0,-90 78,45 -78,45"
              fill="none"
              stroke="white"
              stroke-width="1.5"
              stroke-linejoin="round"
            />
            <circle cx="0" cy="-60" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="-52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
          </g>
        </g>

        <!-- ── Triangle 5: lower-right (300°) — index 5 ── -->
        <g transform="rotate(-300)" class={activeTriangle === 5 ? 'tri-active' : ''}>
          <g transform="translate(300, 0) rotate(90)">
            <polygon
              points="0,-90 78,45 -78,45"
              fill="none"
              stroke="white"
              stroke-width="1.5"
              stroke-linejoin="round"
            />
            <circle cx="0" cy="-60" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
            <circle cx="-52" cy="30" r="7" fill="none" stroke="white" stroke-width="1.2"/>
          </g>
        </g>

        <!-- ── Ring connector lines between adjacent triangles ── -->
        <!--
          Each triangle base vertices closest to adjacent triangle.
          Base-right vertex of triangle N → Base-left vertex of triangle N+1
          For triangle at angle θ with circumradius R=90, translate(300,0) rotate(90):
            base-right in wheel coords: rotate(θ) * (translate(300,0) * (52, 30 rotated 90°))
            But let's just draw arcs/lines between approximate points.

          Approximate connection points (outer vertices of adjacent base corners):
          Connecting vertex positions in SVG (before rotation) at radius 300+45=345 from center.
          We'll draw these as a series of line segments connecting adjacent triangle bases.

          For 6 triangles equally spaced, base edges face each other at ~60° apart.
          The inner ring edge is at approximately r=300-90=210 from center.

          Let's draw simple straight connector lines between adjacent triangle base-corners.
          Each triangle T at angle α (in degrees, CCW from right):
            Center at (300*cos(α), -300*sin(α))   [SVG y-inverted]
            Local base-right after rotate(90): originally (52, 30) → rotate 90° → (30, -52)
              But then rotated by α... complex. Let's use a simpler approach:
              Just connect the ring with a circle outline for connectors.
        -->

        <!-- Thin inner connecting ring (represents the hub connections) -->
        <circle
          cx="0" cy="0" r="215"
          fill="none"
          stroke="white"
          stroke-width="0.6"
          stroke-dasharray="8 16"
          opacity="0.5"
        />

        <!-- Outer ring guide (very subtle) -->
        <circle
          cx="0" cy="0" r="390"
          fill="none"
          stroke="white"
          stroke-width="0.4"
          stroke-dasharray="4 20"
          opacity="0.3"
        />

        <!-- Spoke lines from center to each triangle (subtle) -->
        <line x1="0" y1="0" x2="210" y2="0" stroke="white" stroke-width="0.5" opacity="0.25"/>
        <line x1="0" y1="0" x2="105" y2="-182" stroke="white" stroke-width="0.5" opacity="0.25"/>
        <line x1="0" y1="0" x2="-105" y2="-182" stroke="white" stroke-width="0.5" opacity="0.25"/>
        <line x1="0" y1="0" x2="-210" y2="0" stroke="white" stroke-width="0.5" opacity="0.25"/>
        <line x1="0" y1="0" x2="-105" y2="182" stroke="white" stroke-width="0.5" opacity="0.25"/>
        <line x1="0" y1="0" x2="105" y2="182" stroke="white" stroke-width="0.5" opacity="0.25"/>

        <!-- Center hub dot -->
        <circle cx="0" cy="0" r="8" fill="none" stroke="white" stroke-width="1.2" opacity="0.6"/>
        <circle cx="0" cy="0" r="3" fill="white" opacity="0.5"/>

      </g><!-- end .wheel-ring -->
    </svg>
  </div><!-- end .wheel-layer -->

  <!-- ─── Agent Particles (hollow circles = deployed agents/tasks) ─────── -->
  {#each agents as agent (agent.id)}
    <div
      class="agent-particle"
      style={agentStyle(agent)}
      aria-hidden="true"
    >
      <!-- Motion trail lines (behind the circle, to its left) -->
      <div class="agent-trails">
        <div class="trail trail-1"></div>
        <div class="trail trail-2"></div>
        <div class="trail trail-3"></div>
      </div>
      <!-- Agent circle: hollow white outline matching logo's inner circles -->
      <svg
        class="agent-svg"
        viewBox="0 0 1 1"
        xmlns="http://www.w3.org/2000/svg"
        width="100%"
        height="100%"
      >
        <circle
          cx="0.5" cy="0.5" r="0.42"
          fill="none"
          stroke="white"
          stroke-width="0.1"
        />
      </svg>
    </div>
  {/each}

  <!-- ─── Content Layer ─────────────────────────────────────────────────── -->
  <div class="content-layer">

    <!-- Hero -->
    <section class="hero">
      <div class="hero-inner">

        <!-- Logo -->
        <div class="logo-wrap">
          <img src="/assets/crewport-logo.jpg" alt="CrewPort" class="logo" />
        </div>

        <!-- Wordmark + Tagline -->
        <div class="hero-text">
          <h1 class="wordmark">CrewPort</h1>
          <p class="tagline-primary">Deploy agent crews that deliver.</p>
          <p class="tagline-secondary">
            Contract enforcement, SLA monitoring, and audit trails for AI agent operations.
          </p>
        </div>

        <!-- CTA -->
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

    <!-- Feature Highlights -->
    <section class="features">
      <div class="features-inner">
        {#each features as f}
          <div class="feature-card">
            <div class="feature-icon" aria-hidden="true">
              <svg viewBox="0 0 24 24" width="14" height="14" xmlns="http://www.w3.org/2000/svg">
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

  </div><!-- end .content-layer -->

</div><!-- end .landing -->

<style>
  /* ─── Reset / Root ──────────────────────────────────────────────────── */
  .landing {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    background: #0a0a0a;
    color: #ffffff;
    position: relative;
    overflow: hidden;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', system-ui, sans-serif;
  }

  /* ─── Wheel Layer ───────────────────────────────────────────────────── */
  .wheel-layer {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    z-index: 0;
    pointer-events: none;
    /* Clip left overflow so only right half of wheel is visible */
    overflow: hidden;
  }

  .wheel-svg {
    /*
      Wheel is 150vh × 150vh.
      Center the wheel at viewport left edge (x=0), vertically centered.
      So position: left = -75vh (half the width, off-screen left)
                   top  = 50vh - 75vh = -25vh (centered vertically)
    */
    position: absolute;
    width: 150vh;
    height: 150vh;
    left: -75vh;
    top: calc(50vh - 75vh);
    opacity: 0.17;
    animation: wheel-rotate 60s linear infinite;
  }

  @keyframes wheel-rotate {
    from { transform: rotate(0deg); }
    to   { transform: rotate(-360deg); }
  }

  /* ─── Triangle activation glow (brief pulse when a crew deploys) ────── */
  /* filter: brightness() works on SVG <g> elements in modern browsers.    */
  /* Boosts the triangle to appear bright against the faded wheel.         */
  .wheel-svg .tri-active {
    filter: brightness(8) drop-shadow(0 0 6px white);
    transition: filter 0.1s ease;
  }

  /* ─── Agent Particles (hollow circles) ──────────────────────────────── */
  .agent-particle {
    position: fixed;
    z-index: 1;
    pointer-events: none;
    opacity: 0;
    animation: agent-fly linear forwards;
    transform-origin: center center;
  }

  @keyframes agent-fly {
    0%   { opacity: 0;    transform: translate(0, 0); }
    8%   { opacity: 0.18; }
    70%  { opacity: 0.14; }
    90%  { opacity: 0;    }
    100% { opacity: 0;    transform: translate(var(--dx), var(--dy)); }
  }

  .agent-svg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    filter: drop-shadow(0 0 2px rgba(255,255,255,0.25));
  }

  /* Motion trail lines — positioned to the left of the circle */
  .agent-trails {
    position: absolute;
    top: 50%;
    right: 100%;
    transform: translateY(-50%);
    display: flex;
    flex-direction: column;
    gap: 3px;
    padding-right: 4px;
  }

  .trail {
    height: 1px;
    background: linear-gradient(to left, rgba(255,255,255,0.25), transparent);
    border-radius: 1px;
  }
  .trail-1 { width: 22px; margin-left: 0; }
  .trail-2 { width: 15px; margin-left: 3px; opacity: 0.65; }
  .trail-3 { width: 9px;  margin-left: 6px; opacity: 0.4; }

  /* ─── Content Layer ─────────────────────────────────────────────────── */
  .content-layer {
    position: relative;
    z-index: 10;
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }

  /* ─── Hero ──────────────────────────────────────────────────────────── */
  .hero {
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

  /* ─── Logo ──────────────────────────────────────────────────────────── */
  .logo-wrap {
    width: 180px;
    height: 100px;
    display: flex;
    align-items: center;
    justify-content: center;
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 2px;
    overflow: hidden;
  }

  .logo {
    width: 100%;
    height: 100%;
    object-fit: contain;
  }

  /* ─── Hero Text ─────────────────────────────────────────────────────── */
  .hero-text {
    text-align: center;
  }

  .wordmark {
    font-size: 36px;
    font-weight: 700;
    letter-spacing: 0.24em;
    text-transform: uppercase;
    color: #ffffff;
    margin: 0 0 20px;
    text-shadow: 0 0 40px rgba(255,255,255,0.08);
  }

  .tagline-primary {
    font-size: 22px;
    font-weight: 300;
    color: #ffffff;
    letter-spacing: 0.02em;
    margin: 0 0 12px;
    line-height: 1.4;
  }

  .tagline-secondary {
    font-size: 13px;
    color: #888888;
    line-height: 1.7;
    letter-spacing: 0.02em;
    max-width: 380px;
    margin: 0 auto;
  }

  /* ─── CTA ───────────────────────────────────────────────────────────── */
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
    padding: 13px 24px;
    background: transparent;
    color: #ffffff;
    border: 1px solid rgba(255, 255, 255, 0.35);
    border-radius: 2px;
    font-size: 12px;
    font-weight: 600;
    letter-spacing: 0.10em;
    text-transform: uppercase;
    text-decoration: none;
    transition: border-color 150ms ease, background 150ms ease, box-shadow 150ms ease;
  }

  .btn-github:hover {
    background: rgba(255, 255, 255, 0.05);
    border-color: rgba(255, 255, 255, 0.7);
    box-shadow: 0 0 20px rgba(255, 255, 255, 0.08);
    color: #ffffff;
  }

  .btn-github:active {
    background: rgba(255, 255, 255, 0.08);
  }

  .github-icon {
    width: 16px;
    height: 16px;
    flex-shrink: 0;
    opacity: 0.8;
  }

  .disclaimer {
    font-size: 11px;
    color: #555555;
    text-align: center;
    margin: 0;
    letter-spacing: 0.03em;
  }

  /* ─── Features ──────────────────────────────────────────────────────── */
  .features {
    border-top: 1px solid rgba(255, 255, 255, 0.07);
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(2px);
    padding: 48px 24px;
  }

  .features-inner {
    max-width: 900px;
    margin: 0 auto;
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    border: 1px solid rgba(255, 255, 255, 0.07);
  }

  .feature-card {
    padding: 28px 22px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    border-right: 1px solid rgba(255, 255, 255, 0.07);
    transition: background 150ms ease;
  }

  .feature-card:last-child {
    border-right: none;
  }

  .feature-card:hover {
    background: rgba(255, 255, 255, 0.025);
  }

  .feature-icon {
    color: #555555;
    line-height: 0;
    margin-bottom: 2px;
  }

  .feature-title {
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    color: #ffffff;
    margin: 0;
  }

  .feature-desc {
    font-size: 12px;
    color: #666666;
    line-height: 1.65;
    margin: 0;
  }

  /* ─── Footer ────────────────────────────────────────────────────────── */
  .landing-footer {
    border-top: 1px solid rgba(255, 255, 255, 0.05);
    padding: 18px 24px;
    text-align: center;
    background: #000000;
  }

  .footer-text {
    font-size: 11px;
    color: #444444;
    letter-spacing: 0.07em;
    text-transform: uppercase;
  }

  /* ─── Responsive ────────────────────────────────────────────────────── */
  @media (max-width: 768px) {
    /* Shrink wheel on smaller screens */
    .wheel-svg {
      width: 100vh;
      height: 100vh;
      left: -50vh;
      top: calc(50vh - 50vh);
      opacity: 0.12;
    }

    .features-inner {
      grid-template-columns: repeat(2, 1fr);
    }

    .feature-card:nth-child(2) {
      border-right: none;
    }

    .feature-card:nth-child(1),
    .feature-card:nth-child(2) {
      border-bottom: 1px solid rgba(255, 255, 255, 0.07);
    }

    .wordmark {
      font-size: 28px;
      letter-spacing: 0.20em;
    }

    .tagline-primary {
      font-size: 18px;
    }

    .hero {
      padding: 60px 24px 48px;
    }
  }

  @media (max-width: 480px) {
    /* Hide wheel entirely on very small screens */
    .wheel-layer {
      display: none;
    }

    .features-inner {
      grid-template-columns: 1fr;
    }

    .feature-card {
      border-right: none;
      border-bottom: 1px solid rgba(255, 255, 255, 0.07);
    }

    .feature-card:last-child {
      border-bottom: none;
    }

    .wordmark {
      font-size: 24px;
      letter-spacing: 0.18em;
    }

    .tagline-primary {
      font-size: 16px;
    }

    .logo-wrap {
      width: 150px;
      height: 84px;
    }

    .hero {
      padding: 48px 20px 40px;
    }

    .hero-inner {
      gap: 30px;
    }
  }
</style>
