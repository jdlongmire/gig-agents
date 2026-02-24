# Platform Architecture

v0.1 — Technical Architecture · CONUS Only

---

## Design Principles

**Lean until proven.** Every architectural decision should optimize for speed to MVP and cost efficiency. Over-engineering before product-market fit kills startups. Scale problems are good problems to have.

**Operator-agnostic execution.** The platform doesn't care how operators build their agent teams — CrewAI, AutoGen, LangGraph, raw API calls, Claude tool-use, whatever. The platform cares about inputs (client brief) and outputs (deliverable). Everything in between is the operator's domain.

**Acquisition-ready from day one.** Clean API boundaries, standard data models, well-documented interfaces. When an acquirer's engineering team does due diligence, they should see something they can integrate, not something they have to rewrite.

---

## High-Level Overview

```
┌─────────────┐     ┌──────────────────┐     ┌─────────────────┐
│   Client     │────▶│   Platform API   │◀────│    Operator      │
│   (Web/App)  │◀────│   (Fly.io)       │────▶│    (Agent Teams) │
└─────────────┘     └──────────────────┘     └─────────────────┘
                           │    │
                    ┌──────┘    └──────┐
                    ▼                  ▼
             ┌────────────┐    ┌─────────────┐
             │  Postgres   │    │  Object      │
             │  (Fly.io)   │    │  Storage     │
             └────────────┘    │  (R2/S3)     │
                               └─────────────┘
```

The system is a three-sided marketplace with a thin orchestration layer. The platform handles discovery, matching, payment, and delivery. Operators handle execution.

---

## Fly.io Deployment Architecture

### Why Fly.io

Fly.io gives us edge deployment, managed Postgres, straightforward scaling, and a pricing model that doesn't punish early-stage experimentation. It also handles TLS, load balancing, and regional routing without additional infrastructure. For an MVP through early scale, it's the right call. If an acquirer wants to migrate to their own infra later, the app is containerized and portable.

### Service Topology

All services deploy as Fly.io apps (containers). Keep it simple — start as a monolith, extract services only when there's a concrete reason.

**gig-api** — Primary application server. Handles all API routes, authentication, matching logic, and business rules. Single deployable unit at MVP. Fly.io region: `iad` (Ashburn, VA) for CONUS coverage.

**gig-worker** — Background job processor. Handles async tasks: notification dispatch, deliverable processing, operator webhook management, payment settlement triggers. Runs as a separate Fly.io app so it can scale independently of the API.

**gig-web** — Client-facing frontend. Static SPA served from Fly.io edge or optionally from a CDN (Cloudflare). Talks exclusively to gig-api.

### Fly.io Postgres

Managed Postgres cluster on Fly.io. Single primary in `iad`, read replica added when query load justifies it. This handles all relational data: users, gigs, transactions, reviews, operator capability indexes.

### Secrets and Config

Fly.io secrets for all sensitive config (API keys, DB credentials, payment provider tokens). Environment-specific config via `fly.toml` and deploy-time environment variables.

---

## Core Data Model

### Users

```
users
├── id (uuid)
├── type (client | operator | admin)
├── email
├── display_name
├── auth_provider_id
├── created_at
└── status (active | suspended | pending_verification)
```

### Operator Profiles

```
operator_profiles
├── user_id (fk)
├── capabilities[] (tagged taxonomy — see below)
├── agent_stack (free text — what they use)
├── portfolio_items[] (fk to deliverable samples)
├── rating_avg
├── gigs_completed
├── verified (boolean)
└── subscription_tier (free | pro)
```

### Gigs

```
gigs
├── id (uuid)
├── client_id (fk)
├── operator_id (fk, nullable until claimed)
├── title
├── brief (structured client input)
├── deliverable_type (document | presentation | code | design | marketing | other)
├── status (open | claimed | in_progress | delivered | disputed | completed)
├── budget_cents
├── platform_fee_cents
├── deadline
├── created_at
└── completed_at
```

### Deliverables

```
deliverables
├── id (uuid)
├── gig_id (fk)
├── operator_id (fk)
├── file_url (object storage reference)
├── file_type
├── version (integer, supports revisions)
├── submitted_at
└── accepted_at
```

### Reviews

```
reviews
├── id (uuid)
├── gig_id (fk)
├── reviewer_id (fk)
├── rating (1-5)
├── comment
├── review_type (client_of_operator | operator_of_client)
└── created_at
```

---

## Capability Taxonomy

This is a key piece of IP. A standardized, hierarchical way to describe what agent teams can produce.

```
capability_taxonomy
├── category (document | presentation | code | design | marketing | data | other)
│   ├── subcategory (e.g., document.report, document.letter, code.web_app)
│   │   ├── format (e.g., .docx, .pptx, .pdf, .jsx, .py)
│   │   └── complexity_tier (simple | moderate | complex)
```

Operators tag their profiles with capabilities. Client briefs map to capability requirements. The matching engine connects the two. Over time, this taxonomy becomes a de facto industry standard for describing agentic output — and that's worth something to an acquirer.

---

## Request Flow

### Client Posts a Gig

1. Client submits structured brief via web UI
2. API validates input, creates gig record (status: `open`)
3. Matching engine scores available operators by capability match, rating, availability
4. Qualified operators receive notification (webhook + in-app)
5. Operators review brief, claim gig (status: `claimed`)
6. Platform holds client payment in escrow

### Operator Delivers

1. Operator runs their agent team against the brief (off-platform — their infra, their tools)
2. Operator uploads deliverable to platform via API
3. Deliverable stored in object storage, metadata recorded
4. Client notified of delivery
5. Client reviews: accept (triggers payment release) or request revision
6. On acceptance: escrow releases to operator minus platform fee, both parties prompted for review

### Dispute Path

1. Client flags deliverable as unsatisfactory after revision cycle
2. Platform mediator (initially manual, later AI-assisted) reviews brief vs. deliverable
3. Resolution: full refund, partial refund, or deliverable accepted with adjustment
4. Dispute outcomes feed back into operator rating

---

## Authentication and Authorization

**Auth provider:** Clerk or Auth0 via Fly.io. OAuth2 with Google and email/password at minimum. Operator accounts require additional verification step (capability demonstration or portfolio review).

**Authorization model:** Role-based. Three roles — client, operator, admin. Gig-level permissions ensure only the assigned operator and the posting client can access gig details and deliverables.

**API authentication:** Bearer tokens (JWT). Short-lived access tokens, longer-lived refresh tokens. All API routes authenticated except public-facing marketplace browsing.

---

## Payment Infrastructure

**Provider:** Stripe Connect. Purpose-built for marketplace payment flows.

- Client pays into platform-managed escrow (Stripe payment intent)
- On delivery acceptance, platform triggers transfer to operator's connected Stripe account
- Platform fee deducted automatically via Stripe Connect application fee
- Stripe handles 1099 reporting for US-based operators (CONUS scope makes this cleaner)

**Escrow logic lives in the application layer**, not in Stripe. Stripe holds the funds; the platform controls release triggers.

---

## Object Storage

**Deliverable files:** Cloudflare R2 (S3-compatible, zero egress fees). Deliverables are the product — they need to be reliably stored, versioned, and fast to download.

**Access control:** Pre-signed URLs generated by the API. Files are never publicly accessible. URLs expire after download or a short time window.

**Versioning:** Each deliverable upload creates a new version record. Clients can access revision history.

---

## Operator Integration Model

This is intentionally minimal. The platform does not dictate how operators run their agents.

**What the platform provides to operators:**
- Structured brief data (JSON) via API or webhook
- Deliverable upload endpoint
- Status update endpoints (claimed, in_progress, delivered)
- Webhook notifications for new matching gigs

**What the platform does NOT do:**
- Host or run agent teams
- Provide compute for agent execution
- Require specific frameworks or tools
- Inspect or intermediate in the agent execution process

Operators are black boxes that accept briefs and produce deliverables. This is a feature, not a limitation. It means the platform scales with the entire agent ecosystem rather than betting on one framework.

---

## Monitoring and Observability

**Application metrics:** Fly.io built-in metrics for request volume, latency, error rates. Supplement with a lightweight APM (Sentry for errors, maybe Grafana Cloud free tier for dashboards).

**Business metrics (the ones that matter for acquisition):**
- Gigs posted, claimed, completed, disputed (daily/weekly/monthly)
- GMV and platform revenue
- Operator activation and retention
- Client repeat rate
- Time-to-delivery distribution
- Rating distribution

Build a simple internal dashboard early. These numbers are the product when you're selling the company.

---

## MVP Scope

What ships in Phase 1 (0–6 months):

- Client registration and gig posting (structured brief form)
- Operator registration with capability tagging
- Manual matching (admin-assisted) — automate later when you understand the patterns
- Deliverable upload and download
- Stripe Connect payment flow with escrow
- Basic review system (1–5 stars + comment)
- Email notifications (no real-time yet)
- Admin dashboard with core business metrics

What explicitly does NOT ship in Phase 1:

- Real-time chat between client and operator
- Automated matching algorithm (learn the patterns manually first)
- Operator subscription tiers
- Client subscription/retainer model
- Mobile app (responsive web only)
- AI-assisted dispute resolution

---

## Tech Stack Summary

| Layer | Choice | Rationale |
|---|---|---|
| Hosting | Fly.io | Edge deployment, managed Postgres, container-native, portable |
| API | Node.js (Fastify) or Python (FastAPI) | Operator preference — pick one and commit |
| Database | Postgres (Fly.io managed) | Relational data, proven at scale, acquirer-friendly |
| Object storage | Cloudflare R2 | S3-compatible, zero egress, cheap |
| Auth | Clerk or Auth0 | Don't build auth, buy it |
| Payments | Stripe Connect | Marketplace escrow, 1099 handling, industry standard |
| Frontend | React (Next.js or Vite SPA) | Large talent pool, component ecosystem, acquirer-friendly |
| Background jobs | BullMQ (Node) or Celery (Python) | Runs on gig-worker Fly.io app |
| Email | Resend or Postmark | Transactional email, simple API |
| Monitoring | Sentry + Fly.io metrics | Error tracking + infra visibility |

---

## Migration / Acquisition Readiness

Everything is containerized (Docker on Fly.io). No Fly.io-specific lock-in beyond managed Postgres, which is standard Postgres. An acquirer can migrate to AWS, GCP, or their own infra with minimal friction.

Key portability points:
- Standard Postgres — pg_dump and restore anywhere
- S3-compatible object storage — swap R2 for S3 with a config change
- Docker containers — run anywhere containers run
- Stripe Connect — transfers to acquirer's Stripe account
- Clean API boundaries — documented OpenAPI spec
- No proprietary dependencies

---

## Open Technical Decisions

- **API language:** Node.js (Fastify) vs. Python (FastAPI). Both work. Pick based on team strength.
- **Frontend framework:** Next.js (SSR, better SEO) vs. Vite SPA (simpler, faster to ship). SEO matters for discoverability, which pushes toward Next.js.
- **Matching algorithm:** Start manual, instrument everything, build the algorithm from observed patterns. Don't guess at matching logic before you have data.
- **Operator sandboxing:** Do we ever need to run operator agents on-platform? Current answer is no — operators use their own infra. Revisit if quality control demands it.
- **Rate limiting and abuse prevention:** Operators gaming reviews, clients posting garbage briefs, bot registrations. Plan for it, but don't over-build before it's a problem.
