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
  //
  // ROUND-TRIP ANIMATION:
  // Each agent flies out to a peak position, then arcs back to a triangle on the
  // wheel — simulating a deploy cycle. The arc-back is created by intermediate
  // keyframe waypoints that form a curved path (bezier approximation).

  let agents = [];
  let agentId = 0;
  let launchInterval;
  // Index of the currently "active" (glowing) triangle on the wheel (0-5, or -1 for none)
  let activeTriangle = -1;

  // The 6 triangle positions on the wheel (angles in degrees, 0=right, CCW)
  // We pick from the visible right half: triangles 0 (0°), 1 (60°), 5 (300°) are most visible
  const triangleAngles = [0, 60, 120, 180, 240, 300];

  // Compute viewport-space spawn point for a given visible triangle index
  function triangleSpawnPoint(triIdx) {
    const wheelRadiusVh = 72;
    const triAngleRad = triIdx === 0 ? 0
      : triIdx === 1 ? -60 * (Math.PI / 180)
      : triIdx === 5 ? 60 * (Math.PI / 180)
      : triangleAngles[triIdx] * (Math.PI / 180);
    return {
      x: wheelRadiusVh * Math.cos(triAngleRad),
      y: 50 + wheelRadiusVh * Math.sin(triAngleRad),
    };
  }

  function spawnDeployment() {
    // Count existing agents — cap at 14 (round-trip means more overlap)
    if (agents.length >= 14) return;

    // Pick a triangle to "activate" (source) — bias toward the visible right half
    const visibleTriangles = [0, 1, 5]; // 0°, 60°, 300° — on the right side
    const triIdx = visibleTriangles[Math.floor(Math.random() * visibleTriangles.length)];

    // Flash this triangle (deploy pulse)
    activeTriangle = triIdx;
    setTimeout(() => { activeTriangle = -1; }, 700);

    const spawn = triangleSpawnPoint(triIdx);

    // Pick a DIFFERENT visible triangle as the return dock target (variety is good)
    const returnCandidates = visibleTriangles.filter(t => t !== triIdx);
    const returnTriIdx = returnCandidates[Math.floor(Math.random() * returnCandidates.length)];
    const returnDock = triangleSpawnPoint(returnTriIdx);

    // Spawn 2-3 circles per deployment, slightly spread out
    const count = 2 + Math.floor(Math.random() * 2); // 2 or 3
    for (let i = 0; i < count; i++) {
      // Slight spread around spawn point so circles don't stack perfectly
      const spreadX = (Math.random() - 0.5) * 3; // ±1.5vw spread
      const spreadY = (Math.random() - 0.5) * 4; // ±2vh spread

      const spawnX = spawn.x + spreadX;
      const spawnY = spawn.y + spreadY;

      // Outbound travel angle: mostly rightward, slight random vertical variation
      const travelAngle = (Math.random() * 40 - 20) * (Math.PI / 180);

      // Peak travel distance (how far out before turning back)
      const peakDistance = 55 + Math.random() * 40; // 55–95vw out
      const peakDx = Math.cos(travelAngle) * peakDistance;
      const peakDy = Math.sin(travelAngle) * peakDistance;

      // Arc waypoint: offset perpendicular to the return path to create a curve.
      // The midpoint of the arc swings "above" or "below" the straight return line.
      // Perpendicular direction: rotate return vector 90° and scale.
      const retVecX = returnDock.x - (spawnX + peakDx);
      const retVecY = returnDock.y - (spawnY + peakDy);
      const arcCurvature = (Math.random() > 0.5 ? 1 : -1) * (15 + Math.random() * 20);
      // Perpendicular: (-retVecY, retVecX) normalised × curvature
      const retLen = Math.sqrt(retVecX * retVecX + retVecY * retVecY) || 1;
      const arcMidDx = peakDx + retVecX * 0.5 + (-retVecY / retLen) * arcCurvature;
      const arcMidDy = peakDy + retVecY * 0.5 + (retVecX / retLen) * arcCurvature;

      // Return destination relative to spawnX/spawnY
      const finalDx = returnDock.x - spawnX + (Math.random() - 0.5) * 2;
      const finalDy = returnDock.y - spawnY + (Math.random() - 0.5) * 2;

      // Total round-trip duration: 14–18 seconds
      const duration = 14000 + Math.random() * 4000;
      const size = 10 + Math.random() * 6; // 10-16px diameter

      // Stagger each circle within a deployment by a small delay
      const delay = i * (400 + Math.random() * 300);

      const id = ++agentId;
      agents = [...agents, {
        id,
        spawnX,
        spawnY,
        travelAngle,
        peakDx,
        peakDy,
        arcMidDx,
        arcMidDy,
        finalDx,
        finalDy,
        duration,
        size,
        delay,
        born: Date.now(),
      }];

      // Flash the return triangle when agent is about to dock (at ~90% of duration)
      const dockFlashDelay = delay + duration * 0.88;
      setTimeout(() => {
        activeTriangle = returnTriIdx;
        setTimeout(() => { activeTriangle = -1; }, 500);
      }, dockFlashDelay);

      // Cleanup after animation completes
      setTimeout(() => {
        agents = agents.filter(a => a.id !== id);
      }, delay + duration + 200);
    }
  }

  onMount(() => {
    // Stagger first two deployment events
    setTimeout(() => spawnDeployment(), 1500);
    setTimeout(() => spawnDeployment(), 4500);
    launchInterval = setInterval(() => spawnDeployment(), 5000 + Math.random() * 2000);
  });

  onDestroy(() => {
    clearInterval(launchInterval);
  });

  function agentStyle(agent) {
    // --dur and --dly are CSS custom properties so child elements (trails)
    // can pick them up via var(--dur) / var(--dly) for their own animations.
    return `
      left: ${agent.spawnX}vw;
      top: ${agent.spawnY}vh;
      width: ${agent.size}px;
      height: ${agent.size}px;
      --dx: ${agent.peakDx}vw;
      --dy: ${agent.peakDy}vh;
      --mdx: ${agent.arcMidDx}vw;
      --mdy: ${agent.arcMidDy}vh;
      --rx: ${agent.finalDx}vw;
      --ry: ${agent.finalDy}vh;
      --dur: ${agent.duration}ms;
      --dly: ${agent.delay}ms;
      animation-duration: ${agent.duration}ms;
      animation-delay: ${agent.delay}ms;
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
    opacity: 0.35;
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
    /*
      Round-trip animation:
        0%   → spawn at triangle (opacity 0)
        8%   → fade in, begin outbound leg
        38%  → reach peak (farthest point from wheel)
        55%  → arc midpoint (curved waypoint, perpendicular offset)
        80%  → approaching return dock triangle
        92%  → nearly docked, fading out
        100% → docked at return triangle (opacity 0)

      The multi-keyframe translate() path approximates a bezier curve:
        outbound: straight(ish) rightward arc
        return:   wide curved sweep back to a different triangle
    */
    animation: agent-fly ease-in-out both;
    transform-origin: center center;
  }

  @keyframes agent-fly {
    0%   {
      opacity: 0;
      transform: translate(0, 0);
    }
    8%   {
      opacity: 0.45;
      transform: translate(
        calc(var(--dx) * 0.12),
        calc(var(--dy) * 0.12)
      );
    }
    38%  {
      /* Peak — farthest point outward */
      opacity: 0.38;
      transform: translate(var(--dx), var(--dy));
    }
    55%  {
      /* Arc midpoint — perpendicular curve waypoint for smooth return arc */
      opacity: 0.30;
      transform: translate(var(--mdx), var(--mdy));
    }
    80%  {
      /* Approaching return dock */
      opacity: 0.22;
      transform: translate(
        calc(var(--rx) * 0.85 + var(--mdx) * 0.15),
        calc(var(--ry) * 0.85 + var(--mdy) * 0.15)
      );
    }
    92%  {
      opacity: 0.12;
      transform: translate(var(--rx), var(--ry));
    }
    100% {
      opacity: 0;
      transform: translate(var(--rx), var(--ry));
    }
  }

  .agent-svg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    filter: drop-shadow(0 0 2px rgba(255,255,255,0.25));
  }

  /*
    Motion trail lines.
    During the OUTBOUND leg (0%→38%), the circle moves rightward — trails extend leftward
    (behind the circle in the direction of travel).
    During the RETURN leg (55%→100%), the circle moves leftward — trails should extend
    rightward. We achieve this by scaling the trail container to -1 on the X axis,
    which mirrors the gradient so it fades in the correct direction, and repositioning
    it to the right of the circle.

    Layout:
      .agent-trails is positioned relative to the agent-particle div.
      Outbound: right: 100% of parent (trails to the left of the circle)
      Return:   left:  100% of parent (trails to the right of the circle)
    We animate via `left` switching (handled in keyframes via translate).
  */
  .agent-trails {
    position: absolute;
    top: 50%;
    /* Start on the LEFT side of the circle (outbound mode) */
    right: 100%;
    left: auto;
    transform: translateY(-50%);
    display: flex;
    flex-direction: column;
    gap: 3px;
    padding-right: 4px;
    animation: trail-flip ease-in-out both;
    animation-duration: var(--dur, 16000ms);
    animation-delay: var(--dly, 0ms);
    transform-origin: right center;
  }

  @keyframes trail-flip {
    /*
      Strategy: position the trail div using translateX.
      Outbound (0→35%): translateX(0) — stays on the LEFT (right:100% is the base position)
      Transition (42→58%): shrink to near-zero so the flip is invisible, then jump
      Return (65→100%): we can't easily switch from right:100% to left:100% in keyframes,
        so instead we use a large negative translateX to jump the trails to the right of
        the circle. The circle is ~16px wide. From right:100% position, moving by
        calc(100% + 100% + 16px) would put us to the right side. We use a fixed offset
        in vw since the parent is fixed-size. Since .agent-trails is ~22px wide and the
        circle is 10-16px, we need roughly +200% of trail width + circle width to jump right.
        We use scaleX(-1) to flip the gradient direction too.
    */
    0%   { transform: translateY(-50%) scaleX(1);    opacity: 1; }
    35%  { transform: translateY(-50%) scaleX(1);    opacity: 1; }
    44%  { transform: translateY(-50%) scaleX(0.05); opacity: 0.2; }
    /* At 50%: jump to right side while invisible (scaleX near 0) */
    50%  { transform: translateY(-50%) translateX(calc(200% + 24px)) scaleX(-0.05); opacity: 0.2; }
    60%  { transform: translateY(-50%) translateX(calc(200% + 24px)) scaleX(-1);   opacity: 0.8; }
    88%  { transform: translateY(-50%) translateX(calc(200% + 24px)) scaleX(-1);   opacity: 0.8; }
    95%  { transform: translateY(-50%) translateX(calc(200% + 24px)) scaleX(-0.3); opacity: 0.2; }
    100% { transform: translateY(-50%) translateX(calc(200% + 24px)) scaleX(0);    opacity: 0; }
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
