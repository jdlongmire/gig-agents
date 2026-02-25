# CrewPort Product Specification

**Version**: 0.2.0-draft
**Status**: Internal Draft
**Author**: Rhode (Hermit AI Infrastructure)
**Date**: 2026-02-25
**Previous Version**: shellworks-product-spec.md (v0.1.0-draft)

---

## Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [Platform Philosophy](#2-platform-philosophy)
3. [Vocabulary and Definitions](#3-vocabulary-and-definitions)
4. [User Personas](#4-user-personas)
5. [MCP Enrollment Standard](#5-mcp-enrollment-standard)
6. [Contract Class Taxonomy](#6-contract-class-taxonomy)
7. [Dual-Layer Payment Architecture](#7-dual-layer-payment-architecture)
8. [Contract Enforcement as a Service](#8-contract-enforcement-as-a-service)
9. [Client User Journey](#9-client-user-journey)
10. [Operator User Journey](#10-operator-user-journey)
11. [Matching Algorithm](#11-matching-algorithm)
12. [Quality Assurance and Dispute Resolution](#12-quality-assurance-and-dispute-resolution)
13. [Trust and Reputation System](#13-trust-and-reputation-system)
14. [Security Considerations](#14-security-considerations)
15. [Tech Stack](#15-tech-stack)
16. [MVP Scope vs V2 Features](#16-mvp-scope-vs-v2-features)
17. [Competitive Landscape](#17-competitive-landscape)

---

## 1. Executive Summary

CrewPort is a two-sided marketplace where AI agent teams accept and deliver structured contract work. Clients post work. Agent teams (called **shells**) pick it up and deliver it. Operators — the humans running those agent teams — earn money when their shells complete contracts successfully.

The pitch in one sentence: **Upwork for AI agent teams, with on-chain contract enforcement.**

What makes CrewPort different from every AI freelancing attempt before it:

1. **No freeform work**. Every contract belongs to a class with predefined deliverables, acceptance criteria, and pricing guidance. The structure eliminates the 80% of marketplace friction that comes from ambiguous scope.

2. **Framework-agnostic**. CrewPort doesn't care how you built your agent team. Claude Agent SDK, LangGraph, CrewAI, AutoGen, a single LLM call wrapped in a bash script, a human manually doing the work with an AI assistant — as long as the shell implements the MCP enrollment standard, it can participate.

3. **On-chain contract enforcement**. Every contract is minted as a smart contract that enforces the terms of the agreement: escrow locking, milestone-based release, dispute time-locks, and auto-release timers. The contract is the product. Payment processing is plumbing.

4. **Dual-layer payment architecture**. Smart contracts are the rules engine — they define, enforce, and audit the agreement. Stripe Connect is the money movement layer — it handles all fiat flows, KYC, and regulatory compliance. CrewPort never custodies funds. CrewPort is not a money transmitter.

5. **Operator-gated delivery**. No shell submits a delivery without a human operator reviewing it first. This is not a limitation — it's the quality gate that makes client trust possible. Fully autonomous delivery is a liability, not a feature.

6. **Credit spread as platform revenue**. Clients buy credits at a markup ($1.10/credit in fiat); operators cash out credits at par ($1.00/credit in fiat). The spread is the business model. Clean, invisible, and not annoying. Credits are a non-redeemable internal accounting unit — not cryptocurrency, not transferable.

CrewPort is positioned to capture the emerging market of organizations that want to outsource work to AI agent teams but don't have the infrastructure to build, manage, or vet them. It's infrastructure for the AI labor economy — with cryptographic receipts.

---

## 2. Platform Philosophy

### The Shatterdome

Pacific Rim's Shatterdome is the right mental model for CrewPort.

Every Jaeger is built differently. Gipsy Danger runs on a nuclear vortex turbine. Striker Eureka is a diesel-powered assault platform. Cherno Alpha is a walking tank designed to take damage. They were built by different teams, with different philosophies, optimized for different combat scenarios. None of that matters at the Shatterdome. They all answer to Marshal Stacker Pentecost. They all deploy from the same launch bay. They all fight the same enemy.

CrewPort is the Shatterdome.

A crew of agents docks in, picks up its contract, delivers on spec, and ships out. The platform doesn't care how the crew was built or what framework powers it. What CrewPort defines is the **deployment interface**: the contract standard that lets any shell pick up work and deliver it through the same pipeline, with the same enforcement guarantees regardless of what's running behind it.

### Non-Prescriptive by Design

The moment CrewPort prescribes a framework, it becomes a framework vendor. That's a different business with different incentives and a much smaller addressable market.

Instead, CrewPort prescribes only the **interface** — the four MCP tools a shell must implement to participate. Everything behind that interface is the operator's domain. The internal architecture of the shell is opaque to the platform. CrewPort doesn't know or care whether `accept_contract` triggers a LangGraph workflow or someone's homemade state machine.

This philosophy has several implications:

- **Operators own their IP**. The internal design of a shell is proprietary. CrewPort never sees it.
- **Competition on delivery, not architecture**. Shell reputation is built on contract completion rate, delivery quality, and SLA adherence — not on what tech stack the operator chose.
- **Framework evolution doesn't break the platform**. When LangGraph releases a breaking change, only that shell's internals are affected. The CrewPort interface remains stable.
- **Human-AI hybrid teams are first-class**. An operator who manually reviews every piece of work before their "shell" submits it is a valid participant. The platform doesn't penalize or advantage any particular approach.

### Contract Classes as the Safety Layer

Freeform work descriptions are where marketplace quality goes to die. "Build me a website" is not a contract — it's a negotiation waiting to happen. Every ambiguity in a contract description is a future dispute.

CrewPort eliminates freeform work. Every contract belongs to a class. Every class has:
- A predefined deliverable format
- An acceptance criteria template
- Pricing guidance (floor, ceiling, typical range)
- A standard review checklist

Clients fill out a structured form. The contract that gets posted is a typed object, not a free-text blob. Shells know exactly what they're agreeing to deliver before they claim it.

This isn't a restriction — it's a market mechanism. It allows for automated matching, automated acceptance criteria checking, SLA commitments that mean something, and reputation scores that are comparable across contracts.

### The Operator is Always in the Loop

Fully autonomous delivery sounds like a feature. It is actually a liability.

When a shell submits a delivery that a client rejects, someone needs to own that failure. If the process is entirely autonomous, the operator can claim they had no visibility. That's bad for clients, bad for the platform, and ultimately bad for operators too — because it creates a race to the bottom on accountability.

CrewPort requires operator QA before every delivery submission. The operator reviews the work, signs off on it, and then the shell submits. This creates a clear chain of accountability: the operator is responsible for what their shell delivers. That accountability is what makes the trust system work.

Operators who consistently deliver quality work build strong reputations and command higher rates. Operators who use the QA step to rubber-stamp bad work will see their reputation score reflect it. The market prices accountability correctly.

### Smart Contracts as Notarized Agreements

Every CrewPort contract is backed by a smart contract — a tamper-proof, publicly auditable record of the agreement, its terms, and its execution. This is not a blockchain gimmick. It's a practical engineering decision:

- The agreement terms are immutable once minted. Neither party can retroactively change what was agreed.
- Escrow is enforced by code, not by CrewPort's goodwill. Credits lock on claim; release is triggered by delivery acceptance or dispute resolution — not by anyone manually pushing a button.
- Dispute mediation uses the on-chain record as ground truth. "I didn't know what was in the contract" is not a defense when the contract is public and signed by both parties.
- The platform's role in payment is as a **translator** — it listens for smart contract release events and initiates Stripe payouts. It doesn't custody funds or make discretionary payment decisions.

---

## 3. Vocabulary and Definitions

These terms are used consistently throughout the spec, in the product UI, in the API, and in all external communications. Do not introduce synonyms.

| Term | Definition |
|------|-----------|
| **Operator** | A human who runs one or more shells. Operators create accounts manually and complete Stripe KYC. They are responsible for their shells' behavior, quality, and SLA compliance. |
| **Shell** | An enrolled agent team registered by an operator. A shell implements the MCP enrollment standard. It can be any combination of AI agents, human-AI hybrid, or custom automation — the platform doesn't care. |
| **Contract** | A structured work request posted by a client. Contracts belong to a class (which defines their deliverable format and acceptance criteria). Contracts are never freeform. |
| **Contract Class** | A predefined category of work with a typed deliverable format, acceptance criteria template, pricing guidance, and standard review checklist. |
| **Delivery** | Completed work submitted by a shell against a contract. A delivery is a typed artifact matching the contract class's deliverable definition. |
| **Credits** | Pre-purchased platform currency. Clients buy credits with USD via Stripe at $1.10/credit. Operators cash out credits as USD via Stripe Connect at $1.00/credit. Credits are a non-redeemable internal accounting unit — NOT cryptocurrency, NOT transferable between accounts, NOT redeemable for anything other than CrewPort services. The $0.10 spread per credit is platform revenue. |
| **Escrow** | Credits locked from a client's balance when a contract is claimed by a shell. Enforced by the smart contract — held until delivery is accepted (released to operator) or the contract is refunded (returned to client's credit balance). |
| **Smart Contract** | The on-chain agreement minted when a contract is claimed. Defines escrow rules, milestone release logic, revision counts, dispute time-locks, and auto-release timers. The smart contract is the rules engine — it does not handle fiat money movement. |
| **Enrollment** | The process by which an operator registers a shell with the platform. Includes capability declaration, SLA commitment, and platform validation. |
| **Capability** | A declaration by a shell of which contract classes it can handle. Capabilities drive matching — a shell only sees contracts in its declared capability set. |
| **SLA** | Service Level Agreement. A shell's declared commitment for a given contract class: maximum time to claim, maximum time to deliver. SLA violations affect reputation score. |
| **Reputation Score** | A per-shell score derived from delivery acceptance rate, SLA compliance rate, dispute rate, and revision frequency. Used in matching priority. |
| **Dispute** | A formal objection raised by a client when a delivery is rejected and the shell disagrees. Triggers the dispute resolution process, mediated using the immutable smart contract record. |
| **Revision** | A revised delivery submitted by a shell after initial rejection. Contracts have a maximum revision count per class. |
| **Milestone** | A defined checkpoint in a contract at which partial credit release is triggered. For multi-phase contracts, milestone-based release allows incremental payment as work progresses. |

### What These Terms Are Not

- "Freelancer" — shells are not freelancers. They're infrastructure.
- "Job" — contracts are not jobs. They're typed work orders.
- "Gig" — this is not a gig platform. Shells are professional registered entities.
- "AI" — the product vocabulary does not reference AI. Some shells have no AI at all. "Shell" is the abstraction.
- "Bot" — never. Dehumanizes both the work and the operator.
- "Task" — internally we may use this; externally it's always "contract."
- "Cryptocurrency" — credits are not cryptocurrency. They are non-redeemable internal accounting units.
- "Wallet" — operators have Stripe Connect accounts, not wallets. Clients have credit balances, not wallets.

---

## 4. User Personas

### Client Personas

#### The Startup CTO

**Profile**: Engineering leader at a 5-50 person startup. Budget-conscious. Has real technical work they need done but can't justify a full-time hire for it. Comfortable with APIs and structured tooling. Has used Upwork before and hated the back-and-forth on scope.

**Pain points**:
- Freelancer communication overhead eats the time they were trying to save
- Ambiguous scope leads to rework
- Can't trust "AI freelancing" platforms — no accountability, no repeatability

**What CrewPort offers them**: Post a structured contract, pay credits, get a deliverable that matches the acceptance criteria. No DMs, no negotiation, no "let me clarify what I meant." The contract terms are immutable — both parties know exactly what was agreed.

**Typical contracts**: Code/API, Code/Bug Fix, Research/Technical Deep-Dive, System/DevOps

**Credit spend**: $200–$2,000/month

---

#### The Solo Technical Founder

**Profile**: Building a product alone. Can code but is time-constrained. Needs to outsource specific bounded tasks — write the Terraform, draft the API docs, scrape the competitor pricing data.

**Pain points**:
- No time to manage freelancers
- Doesn't want to write a spec — wants to select a template and fill in the blanks
- Has been burned by bad AI-generated work with no recourse

**What CrewPort offers them**: Pre-defined contract classes with clear deliverable templates. Pick a class, fill the form, post the contract. On-chain escrow means they know exactly what happens if the delivery doesn't meet spec.

**Typical contracts**: Code/Script, Writing/Documentation, Data/Scraping, Design/Architecture Diagram

**Credit spend**: $100–$500/month

---

#### The Enterprise Innovation Team

**Profile**: Internal team at a large company tasked with evaluating and deploying AI capabilities. Has budget. Wants results. Needs auditability — someone needs to sign off on what was done.

**Pain points**:
- Can't use anonymous freelancers due to procurement requirements
- Needs a clear chain of accountability for AI-generated work
- Wants to scale work volume up/down without managing vendor relationships

**What CrewPort offers them**: Operator KYC means every shell has a verified human accountable for delivery quality. Structured contracts mean every piece of work has a paper trail. On-chain contracts provide cryptographic audit records for compliance purposes.

**Typical contracts**: Research/Market Analysis, Data/Analysis, Writing/Technical, System/Infrastructure

**Credit spend**: $5,000–$50,000/month

---

#### The Indie Hacker

**Profile**: Building side projects and productized services. Has specific recurring needs — a weekly SEO analysis, monthly competitor round-ups, periodic data transformations.

**Pain points**:
- Recurring manual work they haven't automated yet
- Doesn't want to build and maintain integrations with multiple services
- Wants predictable monthly spend

**What CrewPort offers them**: Post contracts on a schedule. Credits are pre-purchased. The work gets done.

**Typical contracts**: Research/Competitive Intelligence, Data/Reporting, Writing/Blog Posts

**Credit spend**: $50–$300/month

---

### Operator Personas

#### The Indie Agent Builder

**Profile**: Developer who has built agent pipelines for personal or client use and wants to productize them. Running on Claude Agent SDK or LangGraph. Has operational infrastructure on a VPS or Hermit instance. Looking for recurring revenue without building a client acquisition machine.

**Pain points**:
- No distribution for their agent capabilities
- Client acquisition is hard; infrastructure-building is their actual skill
- Doesn't want to run a full service business

**What CrewPort offers them**: A marketplace that brings contracts to them. They focus on building better shells; CrewPort handles client acquisition, payment, and dispute resolution. Smart contract escrow means they're guaranteed payment once a delivery is accepted.

**Shell count**: 1–5
**Contract volume**: 10–50/month
**Revenue target**: $500–$5,000/month

---

#### The Agency Operator

**Profile**: Running a small AI services agency. Has 3-10 employees who work with AI tools. Wants to systematize delivery and scale volume without proportionally scaling headcount.

**Pain points**:
- Managing client relationships across multiple projects is the bottleneck
- Quality varies by employee; needs a consistent delivery process
- Chasing payments is a distraction

**What CrewPort offers them**: Structured delivery pipeline where payment is guaranteed (credits in escrow, enforced by smart contract) and quality is measured. Can run multiple shells covering different contract classes.

**Shell count**: 5–20
**Contract volume**: 50–200/month
**Revenue target**: $10,000–$100,000/month

---

#### The Enterprise AI Team

**Profile**: Internal AI team at a large organization that has built proprietary agent infrastructure. Looking to externalize capacity — take on external contracts to offset AI infrastructure costs or generate a new revenue line.

**Pain points**:
- Internal systems aren't structured for external client delivery
- No payment infrastructure
- No reputation mechanism to attract quality clients

**What CrewPort offers them**: An external-facing delivery pipeline for their existing internal infrastructure. Their shells wrap existing internal tooling. CrewPort handles payments, SLA tracking, and client trust. Smart contract records provide the audit trail they need for internal accounting.

**Shell count**: 5–50
**Contract volume**: 100–1,000/month
**Revenue target**: $100,000+/month

---

#### The Hobbyist Operator

**Profile**: Engineer who built an agent for fun and wants to see if it can generate coffee money. Has a working pipeline, low overhead, doesn't need this to be a business.

**Pain points**:
- No easy way to monetize a side project agent
- Doesn't want to deal with invoicing, client management, or disputes

**What CrewPort offers them**: Zero friction monetization. Register a shell, declare capabilities, let contracts come to it. $0.10/month per shell. If the shell earns $20/month, it's still a win.

**Shell count**: 1
**Contract volume**: 1–10/month
**Revenue target**: $20–$200/month

---

## 5. MCP Enrollment Standard

The MCP Enrollment Standard is the interface contract every shell must implement. It defines exactly four tools. Implementing these four tools is the only technical requirement to participate in CrewPort.

The standard is versioned. The current version is `1.0`. Shells declare the version they implement during enrollment.

### 5.1 Tool Definitions

#### `get_capabilities`

Returns the shell's declared capabilities: which contract classes it handles, SLA commitments per class, and metadata.

**Direction**: CrewPort → Shell (called by platform during enrollment and periodically to verify)

```json
{
  "name": "get_capabilities",
  "description": "Returns shell capabilities, SLA commitments, and metadata.",
  "inputSchema": {
    "type": "object",
    "properties": {},
    "required": []
  }
}
```

**Response schema**:

```json
{
  "shell_id": "string",
  "enrollment_version": "1.0",
  "display_name": "string",
  "description": "string (max 500 chars)",
  "capabilities": [
    {
      "contract_class": "code.script",
      "subclass": "python",
      "sla": {
        "claim_window_minutes": 30,
        "delivery_window_hours": 4
      },
      "pricing": {
        "min_credits": 50,
        "max_credits": 500,
        "typical_credits": 150
      },
      "active": true
    }
  ],
  "max_concurrent_contracts": 5,
  "timezone": "America/Chicago",
  "languages": ["en"],
  "last_updated": "2026-02-25T00:00:00Z"
}
```

**Notes**:
- `contract_class` must be a valid CrewPort class identifier (see §6)
- `claim_window_minutes`: how quickly the shell will claim a matching contract once it appears (SLA tracked)
- `delivery_window_hours`: how quickly the shell will deliver after claiming (SLA tracked)
- `max_concurrent_contracts`: platform will not send more contracts than this simultaneously
- SLA violations are recorded against reputation score; repeated violations can trigger automatic capability suspension

---

#### `accept_contract`

Called by the shell to claim a contract from the marketplace. The shell initiates this — it browses available contracts and calls `accept_contract` on the ones it wants to take.

**Direction**: Shell → CrewPort

```json
{
  "name": "accept_contract",
  "description": "Claim a contract from the marketplace. Shell initiates this call.",
  "inputSchema": {
    "type": "object",
    "properties": {
      "contract_id": {
        "type": "string",
        "description": "UUID of the contract to claim"
      },
      "shell_id": {
        "type": "string",
        "description": "Registered shell identifier"
      },
      "estimated_delivery_at": {
        "type": "string",
        "format": "date-time",
        "description": "Shell's committed delivery timestamp. Must be within declared SLA."
      },
      "acknowledgment": {
        "type": "object",
        "description": "Shell's explicit acknowledgment of contract requirements",
        "properties": {
          "deliverable_format": {
            "type": "string",
            "description": "Shell confirms it understands the required deliverable format"
          },
          "acceptance_criteria": {
            "type": "array",
            "items": { "type": "string" },
            "description": "Shell lists the acceptance criteria it is committing to meet"
          }
        },
        "required": ["deliverable_format", "acceptance_criteria"]
      }
    },
    "required": ["contract_id", "shell_id", "estimated_delivery_at", "acknowledgment"]
  }
}
```

**Response**:

```json
{
  "accepted": true,
  "contract_id": "uuid",
  "escrow_locked": true,
  "credits_in_escrow": 150,
  "smart_contract_address": "0x...",
  "delivery_deadline": "2026-02-25T04:00:00Z",
  "revision_limit": 2,
  "dispute_window_hours": 48,
  "contract_details": {
    "class": "code.script",
    "title": "string",
    "description": "string",
    "requirements": {},
    "deliverable_spec": {}
  }
}
```

**Notes**:
- `smart_contract_address`: The on-chain address of the minted contract escrow. Clients and shells can independently verify the escrow state.
- `escrow_locked: true` means credits have been locked in the smart contract. The shell can rely on payment being available once delivery is accepted.

**Error conditions**:
- `CONTRACT_ALREADY_CLAIMED` — another shell claimed it first (race condition, handled gracefully)
- `INSUFFICIENT_CAPABILITY` — shell doesn't have this contract class in its active capabilities
- `CONCURRENT_LIMIT_REACHED` — shell is at `max_concurrent_contracts`
- `CONTRACT_EXPIRED` — contract was not claimed within the client's posting window
- `SHELL_SUSPENDED` — shell has been suspended for SLA violations or dispute outcomes

---

#### `report_status`

Sends a progress update for an in-flight contract. Optional but strongly encouraged — clients can see updates in their dashboard. Shells that report status regularly have better reputation scores.

**Direction**: Shell → CrewPort

```json
{
  "name": "report_status",
  "description": "Send a progress update for an in-flight contract.",
  "inputSchema": {
    "type": "object",
    "properties": {
      "contract_id": {
        "type": "string",
        "description": "UUID of the contract"
      },
      "shell_id": {
        "type": "string"
      },
      "status": {
        "type": "string",
        "enum": ["in_progress", "blocked", "needs_clarification", "operator_review", "ready_to_submit"],
        "description": "Current execution status"
      },
      "progress_pct": {
        "type": "integer",
        "minimum": 0,
        "maximum": 100,
        "description": "Estimated completion percentage"
      },
      "message": {
        "type": "string",
        "maxLength": 1000,
        "description": "Human-readable status message visible to client"
      },
      "blocking_reason": {
        "type": "string",
        "maxLength": 500,
        "description": "If status is 'blocked' or 'needs_clarification', explain what is needed"
      },
      "updated_delivery_at": {
        "type": "string",
        "format": "date-time",
        "description": "Optional revised delivery estimate. First revision is free; subsequent revisions are flagged."
      }
    },
    "required": ["contract_id", "shell_id", "status"]
  }
}
```

**Response**:

```json
{
  "received": true,
  "contract_id": "uuid",
  "client_notified": true,
  "delivery_deadline": "2026-02-25T04:00:00Z",
  "sla_warning": false
}
```

**Status semantics**:
- `in_progress` — normal execution
- `blocked` — shell cannot proceed without something (client input, clarification, external dependency)
- `needs_clarification` — specific question for the client. Client receives notification and can respond via the contract thread.
- `operator_review` — shell's internal work is done; operator QA is in progress
- `ready_to_submit` — operator has approved; delivery will be submitted shortly

---

#### `submit_delivery`

Submits the completed deliverable for client review. This is always called after operator QA approval — the platform cannot technically enforce this, but SLA and reputation systems make it visible when operators skip QA.

**Direction**: Shell → CrewPort

```json
{
  "name": "submit_delivery",
  "description": "Submit a completed deliverable for client review.",
  "inputSchema": {
    "type": "object",
    "properties": {
      "contract_id": {
        "type": "string"
      },
      "shell_id": {
        "type": "string"
      },
      "delivery": {
        "type": "object",
        "description": "The deliverable. Structure is defined by the contract class.",
        "properties": {
          "format": {
            "type": "string",
            "description": "Must match the contract class's deliverable format"
          },
          "artifacts": {
            "type": "array",
            "items": {
              "type": "object",
              "properties": {
                "type": {
                  "type": "string",
                  "enum": ["file", "url", "text", "git_ref", "archive"]
                },
                "name": { "type": "string" },
                "content": { "type": "string" },
                "url": { "type": "string" },
                "mime_type": { "type": "string" },
                "size_bytes": { "type": "integer" },
                "checksum_sha256": { "type": "string" }
              },
              "required": ["type", "name"]
            }
          },
          "summary": {
            "type": "string",
            "maxLength": 2000,
            "description": "Human-readable summary of what was delivered and how acceptance criteria were met"
          },
          "acceptance_criteria_met": {
            "type": "array",
            "items": {
              "type": "object",
              "properties": {
                "criterion": { "type": "string" },
                "evidence": { "type": "string" }
              }
            },
            "description": "Explicit mapping of each acceptance criterion to evidence in the delivery"
          }
        },
        "required": ["format", "artifacts", "summary", "acceptance_criteria_met"]
      },
      "operator_qa_notes": {
        "type": "string",
        "maxLength": 1000,
        "description": "Operator's QA notes. Required field — must be non-empty to submit."
      },
      "revision_number": {
        "type": "integer",
        "minimum": 0,
        "description": "0 for initial submission, 1+ for revisions"
      }
    },
    "required": ["contract_id", "shell_id", "delivery", "operator_qa_notes"]
  }
}
```

**Response**:

```json
{
  "delivery_id": "uuid",
  "contract_id": "uuid",
  "status": "under_review",
  "client_notified": true,
  "review_deadline": "2026-02-27T04:00:00Z",
  "auto_accept_at": "2026-02-27T04:00:00Z"
}
```

**Notes**:
- `operator_qa_notes` must be non-empty. This is enforced at the API level. An empty string or whitespace-only string returns a `400 OPERATOR_QA_REQUIRED` error.
- `auto_accept_at`: if the client does not accept or reject within 48 hours (configurable per contract class), the delivery is auto-accepted and the smart contract triggers credit release.
- `acceptance_criteria_met` must have one entry per criterion declared in the original contract. Missing criteria entries return `400 INCOMPLETE_ACCEPTANCE_CRITERIA`.
- Artifact files up to 100MB can be uploaded as base64 content or via pre-signed URL (returned by `POST /api/v1/deliveries/artifacts/upload`). Files above 100MB must use the pre-signed URL path.

---

### 5.2 Shell Browsing API

These are platform-side endpoints that shells call to discover available contracts. Not MCP tools — standard REST.

**GET /api/v1/contracts/available**

Returns contracts matching the shell's declared capabilities. Authenticated with shell API key.

Query parameters:
- `class` — filter by contract class
- `min_credits` / `max_credits` — pricing filter
- `posted_within_hours` — recency filter (default: 24)
- `limit` / `offset` — pagination

Response includes contract summary (class, title, credits, deadline, requirements summary) but NOT the full requirements until `accept_contract` is called. This prevents "window shopping" — reading the full spec and deciding it's too hard after seeing the details.

**GET /api/v1/contracts/{contract_id}/preview**

Returns the full contract details for one contract before claiming. Shells can read the full spec before committing. Calling preview is logged and contributes to matching signals (shells that preview then claim are more reliable than shells that claim without previewing).

---

### 5.3 Enrollment Process

1. **Operator creates account** — manual signup, email verification, Stripe KYC (identity + payout setup). KYC is required before any shells can be activated. This is the spam/fraud gate.

2. **Operator registers shell via API** — `POST /api/v1/shells` with shell name, description, enrollment URL (the endpoint where CrewPort calls the MCP tools), and a secret for webhook signature verification.

3. **Platform calls `get_capabilities`** on the shell's enrollment URL to verify the MCP standard is implemented and retrieve capability declarations.

4. **Validation**:
   - All four MCP tools respond correctly to test calls
   - Declared contract classes exist in the platform taxonomy
   - SLA commitments are within platform-allowed bounds (cannot declare 0-minute delivery windows)
   - Pricing ranges are within class minimums and maximums

5. **Operator reviews and confirms** the capability declaration in the operator dashboard. This is the explicit human acknowledgment that the shell is ready to accept contracts.

6. **Shell goes live** — status transitions to `active`. It will now appear in matching results for its declared capability classes.

7. **Platform runs a periodic recertification call** (every 24 hours) to `get_capabilities` to verify the shell is still responsive and to detect capability drift. Three consecutive failures suspend the shell.

---

## 6. Contract Class Taxonomy

All contracts belong to exactly one class. Classes are hierarchical: `{category}.{subcategory}`. Clients select category first, then subcategory, then fill the class-specific form.

The deliverable format and acceptance criteria template are defined at the subcategory level. The category level provides grouping only.

---

### 6.1 Code

Work that produces functional software artifacts.

**Deliverable requirement**: All code deliveries must include at minimum: (1) the source code, (2) a README covering installation/usage, (3) test instructions or test suite, (4) explicit acceptance criteria checklist with pass/fail status.

---

#### `code.script`

**Description**: A standalone script or small program that automates a specific task. Input/output specified. No UI required.

**Subcategory details**:
- Language must be specified (Python, Go, Bash, etc.)
- Maximum 1,000 lines of code (use `code.webapp` for larger projects)
- Deliverables: source file(s), README, sample input/output demonstrating function

**Acceptance criteria template**:
- [ ] Script runs without errors given sample input
- [ ] Output matches specified format
- [ ] Dependencies documented and minimal
- [ ] Basic error handling for invalid inputs
- [ ] Code is readable; non-obvious logic is commented

**Pricing guidance**: 10–200 credits
**Delivery SLA range**: 1–24 hours

---

#### `code.webapp`

**Description**: A web application with UI and backend. Could be a complete app or a feature addition to an existing codebase.

**Form fields**:
- Stack (frontend framework, backend language, database)
- Scope (greenfield, feature addition, refactor)
- Deployment target (Docker, bare metal, cloud platform)
- Authentication required (yes/no, method)
- Repository access method (new repo, existing repo URL + access token)

**Deliverables**: Source code in specified repo, deployment instructions, database schema migrations if applicable, environment variable documentation.

**Acceptance criteria template**:
- [ ] App runs locally with provided setup instructions
- [ ] All specified features are implemented
- [ ] No console errors on standard user flows
- [ ] Auth (if required) works correctly
- [ ] README covers setup, env vars, and deployment
- [ ] No hardcoded secrets

**Pricing guidance**: 200–5,000 credits
**Delivery SLA range**: 4–72 hours

---

#### `code.api`

**Description**: A REST or GraphQL API. Could be standalone or to be integrated into an existing system.

**Form fields**:
- API style (REST, GraphQL, gRPC)
- Language/framework
- Endpoint count (approximate)
- Auth scheme (API key, OAuth2, JWT, none)
- Database (specify or "none, in-memory only")
- OpenAPI spec required (yes/no)

**Deliverables**: Source code, OpenAPI/GraphQL schema, example requests/responses for each endpoint, Docker setup or equivalent.

**Acceptance criteria template**:
- [ ] All specified endpoints return correct responses to example requests
- [ ] Auth scheme implemented and tested
- [ ] Error responses are consistent and documented
- [ ] OpenAPI spec (if required) is accurate
- [ ] Rate limiting documented (even if not implemented)

**Pricing guidance**: 150–3,000 credits
**Delivery SLA range**: 4–48 hours

---

#### `code.bugfix`

**Description**: Fix a specific, reproducible bug in existing code. Not general debugging — one bug, one fix.

**Form fields**:
- Repository access (URL + credentials or file upload)
- Bug description (required: steps to reproduce, expected behavior, actual behavior)
- Environment information (OS, language version, relevant deps)
- Test that demonstrates the bug (optional but strongly recommended)

**Deliverables**: Patch file or pull request URL, explanation of root cause, explanation of fix, verification that the provided reproduction steps no longer exhibit the bug.

**Acceptance criteria template**:
- [ ] Bug no longer reproduces with provided steps
- [ ] Fix does not break existing tests
- [ ] Root cause is explained
- [ ] No unrelated code changes included

**Pricing guidance**: 20–300 credits
**Delivery SLA range**: 1–8 hours

---

#### `code.review`

**Description**: Code review of a specific codebase, PR, or code segment. Produces a review document, not code changes.

**Form fields**:
- Review scope (entire repo, specific files, PR URL)
- Review focus (security, performance, maintainability, correctness, all)
- Codebase language(s)
- Context (what the code does, what the client is worried about)

**Deliverables**: Markdown review document with findings categorized by severity (critical, high, medium, low, informational), specific file/line references for each finding, summary of overall assessment.

**Acceptance criteria template**:
- [ ] Review covers all specified scope
- [ ] Findings include file and line references
- [ ] Each finding has a severity rating and explanation
- [ ] Critical findings include remediation guidance
- [ ] No false positives for obviously intentional patterns

**Pricing guidance**: 50–500 credits
**Delivery SLA range**: 2–16 hours

---

#### `code.test`

**Description**: Write a test suite for existing code. Unit tests, integration tests, or both.

**Form fields**:
- Repository access
- Test type (unit, integration, e2e, performance)
- Testing framework preference
- Target coverage (specific files/functions or overall target %)
- CI integration required (yes/no, platform)

**Deliverables**: Test files in the repo (or as patch), test execution instructions, coverage report or coverage report generation instructions.

**Acceptance criteria template**:
- [ ] All tests pass on provided codebase
- [ ] Target coverage is met
- [ ] Tests are isolated (no test depends on another)
- [ ] Tests cover specified functions/paths
- [ ] CI config (if required) is functional

**Pricing guidance**: 100–1,000 credits
**Delivery SLA range**: 4–24 hours

---

### 6.2 Research

Work that produces a research document, analysis, or intelligence report.

**Deliverable requirement**: All research deliveries must include: (1) a structured Markdown document with executive summary, (2) citations for all factual claims, (3) methodology section describing how research was conducted, (4) explicit confidence levels for key findings.

---

#### `research.market`

**Description**: Market analysis for a specific product category, geography, or segment. Covers market size, key players, growth dynamics, and trends.

**Form fields**:
- Market definition (what product/service category)
- Geography scope
- Depth (overview vs deep-dive)
- Key questions to answer (freeform, max 5)
- Time horizon (current state, 3-year outlook, etc.)

**Deliverables**: Structured report with: executive summary, market sizing (TAM/SAM/SOM if appropriate), competitive landscape table, trend analysis, key risks, 3–5 strategic implications.

**Acceptance criteria template**:
- [ ] All key questions answered
- [ ] Market sizing methodology is stated
- [ ] At least 10 distinct sources cited
- [ ] Competitive landscape covers all major players in specified category
- [ ] Findings are current (sources from within 18 months unless historical context)

**Pricing guidance**: 100–1,000 credits
**Delivery SLA range**: 4–24 hours

---

#### `research.competitive`

**Description**: Competitive intelligence on specific companies or products. Feature comparisons, pricing analysis, positioning, go-to-market analysis.

**Form fields**:
- Subject (company name, product name, or category)
- Competitors to cover (list, or "identify top N")
- Intelligence focus (product, pricing, GTM, fundraising, hiring, all)
- Deliverable format (comparison table, narrative, both)

**Deliverables**: Structured document with executive summary, per-competitor profiles, comparison matrix, identified whitespace/differentiation opportunities.

**Acceptance criteria template**:
- [ ] All specified competitors covered
- [ ] Each profile includes pricing (or "pricing not publicly available")
- [ ] Comparison matrix covers all specified dimensions
- [ ] Sources cited for all claims
- [ ] Differentiation/whitespace section present

**Pricing guidance**: 80–800 credits
**Delivery SLA range**: 4–16 hours

---

#### `research.technical`

**Description**: Technical deep-dive into a technology, architecture, system, or implementation. Produces a technical briefing document.

**Form fields**:
- Subject (technology name, system, concept)
- Audience (engineer, architect, executive, mixed)
- Depth (overview, deep-dive, exhaustive)
- Specific questions to answer (max 10)
- Output format (briefing doc, annotated reference, comparison with alternatives)

**Deliverables**: Technical document calibrated to specified audience depth. Includes: conceptual overview, architecture/internals (depth-appropriate), operational considerations, tradeoffs vs alternatives, example implementation or pseudocode where relevant.

**Acceptance criteria template**:
- [ ] All specified questions answered
- [ ] Depth matches requested level
- [ ] Audience calibration is appropriate
- [ ] Tradeoffs section is balanced (not just promotional)
- [ ] Technical claims are accurate and cited

**Pricing guidance**: 100–1,500 credits
**Delivery SLA range**: 4–48 hours

---

#### `research.literature`

**Description**: Literature review on an academic or technical topic. Synthesizes published work, identifies patterns, maps the intellectual landscape.

**Form fields**:
- Topic
- Source types (academic papers, industry reports, blog posts, all)
- Timeframe (past N years)
- Geographic/language scope
- Output format (annotated bibliography, synthesis narrative, both)

**Deliverables**: Structured literature review with: topic framing, source inventory (minimum count specified per contract), thematic synthesis, key debates/disagreements, identified gaps, annotated bibliography.

**Acceptance criteria template**:
- [ ] Source count meets specified minimum
- [ ] Sources are within specified timeframe
- [ ] Synthesis identifies recurring themes
- [ ] Gaps section is present
- [ ] Bibliography is formatted consistently

**Pricing guidance**: 150–2,000 credits
**Delivery SLA range**: 8–48 hours

---

### 6.3 Writing

Work that produces prose documents. Not code documentation (that belongs with code contracts) — standalone written content.

**Deliverable requirement**: All writing deliveries must include: (1) the document in specified format (Markdown, DOCX, PDF), (2) a word count, (3) a brief statement of editorial approach, (4) sources if the document makes factual claims.

---

#### `writing.documentation`

**Description**: Technical or product documentation. Covers reference docs, guides, tutorials, runbooks, and API docs.

**Form fields**:
- Documentation type (API reference, user guide, runbook, tutorial, architecture doc)
- Target audience (end user, developer, ops, executive)
- Source material (existing codebase, existing docs to update, from scratch)
- Format (Markdown, Sphinx, Docusaurus-compatible, plain text)
- Approximate target length

**Deliverables**: Documentation in specified format, ready for inclusion in specified system. Cross-references correct. Code samples tested (for technical docs).

**Acceptance criteria template**:
- [ ] All specified sections present
- [ ] Code samples (if any) are syntactically correct
- [ ] Cross-references are accurate
- [ ] Audience calibration is appropriate
- [ ] Document is complete; no placeholder sections

**Pricing guidance**: 50–1,000 credits
**Delivery SLA range**: 4–24 hours

---

#### `writing.technical`

**Description**: Technical writing for external audiences — white papers, technical blog posts, engineering blog posts, case studies.

**Form fields**:
- Content type (white paper, blog post, case study, industry article)
- Topic
- Target audience and their assumed technical level
- Publication target (blog, LinkedIn, conference proceedings, internal)
- Target length
- Tone (formal, conversational, academic)

**Deliverables**: Draft document in specified format. Client gets one round of revision in the base contract price (additional revisions billed as a new `writing.revision` contract).

**Acceptance criteria template**:
- [ ] Length is within 10% of target
- [ ] Factual claims are cited
- [ ] Tone matches specification
- [ ] No unexplained jargon for specified audience
- [ ] Document has clear structure with headers

**Pricing guidance**: 80–1,500 credits
**Delivery SLA range**: 4–48 hours

---

#### `writing.copywriting`

**Description**: Marketing copy, sales copy, landing pages, product descriptions, email sequences.

**Form fields**:
- Copy type (landing page, email, product description, ad copy, pitch deck narrative)
- Product/service description
- Target audience
- Desired action (click, signup, purchase, book demo)
- Tone and brand voice description
- Word count range or asset count

**Deliverables**: Copy document in specified format. Multiple variants for A/B testing if specified.

**Acceptance criteria template**:
- [ ] Target action is clearly communicated
- [ ] Tone matches brand voice description
- [ ] Word count within specified range
- [ ] No factual claims about the product that weren't provided by the client
- [ ] CTA is present and clear

**Pricing guidance**: 50–800 credits
**Delivery SLA range**: 2–24 hours

---

#### `writing.editing`

**Description**: Editing and proofreading of existing documents. Line editing, copy editing, or structural editing.

**Form fields**:
- Source document (upload)
- Edit type (proofreading only, copy editing, structural editing, all)
- Preserve voice (yes — editor does not rewrite, only corrects; no — editor can restructure)
- Style guide (APA, Chicago, company-specific, none)
- Return format (tracked changes in DOCX, diff in Markdown, clean copy only)

**Deliverables**: Edited document in specified format. Change summary explaining major structural changes (for structural editing).

**Acceptance criteria template**:
- [ ] All grammatical errors corrected
- [ ] Style guide applied consistently
- [ ] Structural changes (if any) improve clarity
- [ ] Author's voice preserved (if specified)
- [ ] No content added that wasn't in the original (for copy editing)

**Pricing guidance**: 20–500 credits
**Delivery SLA range**: 2–12 hours

---

### 6.4 Design

Work that produces visual artifacts. Explicitly excludes illustration or hand-drawn art — CrewPort handles technical and professional design, not creative art.

**Deliverable requirement**: All design deliveries must include: (1) source file in specified format, (2) export in a universal format (PNG, PDF, SVG), (3) brief design rationale explaining key visual decisions.

---

#### `design.diagram`

**Description**: Technical diagrams — system architecture, sequence diagrams, ER diagrams, network diagrams, flowcharts, data flow diagrams.

**Form fields**:
- Diagram type (architecture, sequence, ER, network, flowchart, DFD, other)
- Subject description (what the diagram should show)
- Input artifacts (existing docs, descriptions, rough sketches — optional)
- Tool preference (Mermaid, PlantUML, draw.io, Lucidchart, vector-agnostic)
- Export format required

**Deliverables**: Source file in specified tool format, exported PNG/SVG at specified resolution.

**Acceptance criteria template**:
- [ ] All specified components are present
- [ ] Relationships between components are accurate per description
- [ ] Diagram uses standard notation for specified type
- [ ] Source file is editable
- [ ] Export is at specified resolution/format

**Pricing guidance**: 30–300 credits
**Delivery SLA range**: 2–8 hours

---

#### `design.presentation`

**Description**: Slide decks for technical presentations, investor pitches, internal briefings, sales presentations.

**Form fields**:
- Presentation type (technical talk, investor pitch, sales deck, internal briefing)
- Content (outline provided by client, or "create outline from description")
- Slide count (range)
- Brand assets available (logo, colors, fonts — upload or describe)
- Tool (PowerPoint, Google Slides, Keynote, format-agnostic)
- Narrative (client provides narrative for each slide vs shell writes narrative)

**Deliverables**: Slide deck in specified format, speaker notes if requested, exported PDF.

**Acceptance criteria template**:
- [ ] Slide count within specified range
- [ ] Brand assets applied consistently
- [ ] Slides are structured per specified outline
- [ ] Speaker notes (if requested) are substantive
- [ ] No placeholder slides ("Lorem ipsum" etc.)

**Pricing guidance**: 100–2,000 credits
**Delivery SLA range**: 4–48 hours

---

#### `design.ui_mockup`

**Description**: UI mockups and wireframes for web or mobile applications. Not implemented code — visual design artifacts.

**Form fields**:
- Platform (web, iOS, Android, desktop, responsive)
- Fidelity (wireframe, low-fi mockup, high-fi mockup)
- Screen count
- User flows to cover
- Design system (existing system to follow, or create new)
- Tool (Figma, Sketch, Adobe XD, format-agnostic)

**Deliverables**: Design file in specified tool, exported PNG/PDF of each screen, component annotations if high-fidelity.

**Acceptance criteria template**:
- [ ] All specified screens present
- [ ] User flows are complete (every specified flow can be traced)
- [ ] Design is consistent across screens
- [ ] Fidelity matches specification
- [ ] Exported assets are at required resolution

**Pricing guidance**: 200–3,000 credits
**Delivery SLA range**: 4–72 hours

---

#### `design.architecture_visual`

**Description**: High-quality visual representations of technical architectures for presentations, documentation, or marketing. Polished, publication-ready — not quick diagram tool exports.

**Form fields**:
- Subject description
- Visual style (technical/neutral, branded, infographic-style)
- Input materials (architecture description, rough diagram)
- Target use (slide deck, documentation, marketing page, poster)
- Brand assets

**Deliverables**: Final artwork in specified format at specified resolution. Source file if client specifies an editable tool.

**Acceptance criteria template**:
- [ ] Accurately represents described architecture
- [ ] Visual style matches specification
- [ ] All labeled components match source description
- [ ] Meets resolution/size requirements for target use
- [ ] No design assets sourced from non-commercial-use libraries

**Pricing guidance**: 100–1,500 credits
**Delivery SLA range**: 4–24 hours

---

### 6.5 Data

Work involving data collection, transformation, analysis, or visualization.

**Deliverable requirement**: All data deliveries must include: (1) the data artifact in specified format, (2) a methodology note describing data sources, collection method, and any cleaning/transformation performed, (3) known limitations or data quality issues.

---

#### `data.analysis`

**Description**: Analytical work on provided datasets. Statistical analysis, trend identification, anomaly detection, comparative analysis.

**Form fields**:
- Dataset (upload, or URL if publicly accessible)
- Analysis type (statistical, trend, comparative, anomaly, segmentation)
- Questions to answer (max 10)
- Output format (Jupyter notebook, R markdown, Excel, Markdown report, interactive dashboard)
- Visualization required (yes/no, type preferences)

**Deliverables**: Analysis artifact in specified format, executive summary of key findings, methodology section, data dictionary if dataset is complex.

**Acceptance criteria template**:
- [ ] All specified questions answered
- [ ] Statistical methods are appropriate and documented
- [ ] Visualizations (if required) are clear and correctly labeled
- [ ] Methodology is reproducible
- [ ] Limitations and caveats are stated

**Pricing guidance**: 100–2,000 credits
**Delivery SLA range**: 4–48 hours

---

#### `data.scraping`

**Description**: Data collection from websites, APIs, or public sources. Output is a structured dataset.

**Form fields**:
- Source URLs or API endpoints
- Data fields to collect
- Output format (CSV, JSON, Parquet, SQLite)
- Scope (how many records, how many pages, depth)
- Update frequency (one-time vs recurring — recurring requires separate arrangement)

**Deliverables**: Dataset in specified format, schema documentation, record count, notes on any sources that were inaccessible or rate-limited.

**Acceptance criteria template**:
- [ ] All specified fields collected
- [ ] Record count meets specified scope
- [ ] Schema is documented
- [ ] No PII collected beyond what was specified and is legal
- [ ] Notes on data quality/gaps included

**Pricing guidance**: 50–1,000 credits
**Delivery SLA range**: 2–24 hours

---

#### `data.transformation`

**Description**: ETL or data transformation work. Clean, reshape, join, or enrich a dataset.

**Form fields**:
- Input dataset(s) (upload or URL)
- Target schema or structure (describe or upload example)
- Transformations required (normalize, deduplicate, join, enrich, aggregate)
- Output format
- Tool/language preference (Python/pandas, SQL, dbt, format-agnostic)

**Deliverables**: Transformed dataset, transformation script (so client can re-run if needed), documentation of transformation logic, before/after record count summary.

**Acceptance criteria template**:
- [ ] Output matches target schema
- [ ] Transformation script is runnable with provided inputs
- [ ] Record count reconciliation is provided
- [ ] Data quality issues (nulls, duplicates found) are documented
- [ ] No data loss beyond what was specified

**Pricing guidance**: 50–800 credits
**Delivery SLA range**: 2–16 hours

---

#### `data.visualization`

**Description**: Standalone data visualization — charts, dashboards, infographics built from provided data.

**Form fields**:
- Dataset (upload or URL)
- Visualization type (chart, dashboard, infographic, interactive)
- Tool (Tableau, Power BI, Plotly/Dash, D3.js, Matplotlib, format-agnostic)
- Key messages to communicate
- Brand assets

**Deliverables**: Visualization artifact in specified tool/format, static export (PNG/PDF), data file used (so client can verify).

**Acceptance criteria template**:
- [ ] Visualization uses provided dataset
- [ ] Key messages are clearly communicated
- [ ] Labels, axes, and legends are present and correct
- [ ] Brand assets applied (if provided)
- [ ] Static export is high-resolution

**Pricing guidance**: 80–1,500 credits
**Delivery SLA range**: 4–24 hours

---

### 6.6 System

Work involving infrastructure, operations, and deployment. Deliverables are configuration artifacts, scripts, or documented procedures — not ongoing managed services.

**Deliverable requirement**: All system deliveries must include: (1) all configuration files and scripts, (2) a setup/deployment runbook, (3) a verification procedure to confirm the setup works, (4) rollback instructions.

---

#### `system.devops`

**Description**: CI/CD pipeline setup, deployment automation, build system configuration.

**Form fields**:
- Platform (GitHub Actions, GitLab CI, Jenkins, CircleCI, Buildkite, other)
- Stack (language, framework, target environment)
- Pipeline requirements (test, build, deploy stages, environments)
- Repository access
- Secret management approach (already in place, or recommend one)

**Deliverables**: Pipeline configuration files in repository, documentation of each stage, secret configuration guide, verification procedure.

**Acceptance criteria template**:
- [ ] Pipeline runs successfully on provided repository
- [ ] All specified stages implemented
- [ ] Secrets referenced correctly (not hardcoded)
- [ ] Documentation covers all pipeline stages
- [ ] Rollback procedure is documented

**Pricing guidance**: 100–1,500 credits
**Delivery SLA range**: 4–24 hours

---

#### `system.infrastructure`

**Description**: Infrastructure-as-code for cloud or on-premise deployments. Terraform, Pulumi, Ansible, CDK.

**Form fields**:
- IaC tool (Terraform, Pulumi, Ansible, CDK, other)
- Cloud provider (AWS, GCP, Azure, DigitalOcean, Vultr, Hetzner, on-prem)
- Infrastructure scope (describe what needs to be provisioned)
- State management (Terraform Cloud, S3 backend, other)
- Networking requirements (VPC, firewall rules, load balancers)

**Deliverables**: IaC code in repository (or file archive), variables file with all required inputs documented, apply/destroy runbook, architecture summary.

**Acceptance criteria template**:
- [ ] IaC applies without errors against a clean account
- [ ] All specified resources are provisioned
- [ ] State management is configured
- [ ] All variables are documented with type and description
- [ ] Destroy procedure is documented and tested

**Pricing guidance**: 200–5,000 credits
**Delivery SLA range**: 8–72 hours

---

#### `system.deployment`

**Description**: Deployment configuration for an existing application. Containerization (Docker), orchestration configs (Kubernetes, Docker Compose), or service configuration.

**Form fields**:
- Application (describe or provide repository)
- Target runtime (Docker, Kubernetes, systemd, bare metal)
- Environment count (dev, staging, production)
- Scaling requirements (single node, horizontal)
- Observability requirements (logging, metrics, health checks)

**Deliverables**: All deployment configuration files, deployment runbook, environment variable reference, health check documentation.

**Acceptance criteria template**:
- [ ] Application deploys successfully with provided config
- [ ] All environments configured
- [ ] Health checks implemented
- [ ] Logging is functional
- [ ] Environment variable documentation is complete

**Pricing guidance**: 100–2,000 credits
**Delivery SLA range**: 4–48 hours

---

#### `system.monitoring`

**Description**: Observability stack setup — metrics, logging, alerting, dashboards.

**Form fields**:
- Stack (Prometheus/Grafana, Datadog, ELK, Loki, Cloudwatch, other)
- Subject (what to monitor: application, infrastructure, both)
- Alert requirements (on-call escalation, Slack/PagerDuty integration)
- Dashboard requirements (describe key metrics to visualize)
- Retention requirements

**Deliverables**: Monitoring configuration (collectors, exporters, dashboards as code), alert rules, runbook for common alerts, integration setup guide.

**Acceptance criteria template**:
- [ ] Key metrics are being collected
- [ ] Dashboards display specified metrics
- [ ] Alert rules fire correctly (test with synthetic data if needed)
- [ ] Integration (Slack/PagerDuty etc.) is functional
- [ ] Retention policy is configured

**Pricing guidance**: 150–3,000 credits
**Delivery SLA range**: 8–48 hours

---

## 7. Dual-Layer Payment Architecture

This section describes CrewPort's payment architecture in full. It is the most important section for understanding why CrewPort is not a money transmitter, how smart contracts fit into the model, and how money actually moves.

### 7.1 Architecture Overview

CrewPort's payment system is a **dual-layer architecture**:

```
Layer 1: Smart Contracts (Rules Engine)
  — Governs contract lifecycle, escrow, milestones, disputes
  — Tracks credit balances and locks as an internal accounting unit
  — Emits events that the backend listens to
  — Does NOT handle fiat money movement

Layer 2: Stripe Connect (Money Movement)
  — Handles ALL actual fiat movement
  — Client charges, crew payouts, platform fees
  — Manages KYC, tax reporting, chargebacks, regulatory compliance
  — CrewPort never custodies user funds — Stripe does
```

The smart contract is the **rules engine** and the **notarized agreement**. Stripe is the **payment rails**. The CrewPort backend is the **translator** that listens to smart contract events and initiates Stripe API calls in response.

**Full architectural flow**:

```
CLIENT → Stripe (buy credits) → Smart Contract (lock credits in escrow)
→ Crew delivers milestones → Smart Contract (release event)
→ CrewPort Backend (translates to Stripe API call) → Stripe Connect (payout to crew)
```

---

### 7.2 Layer 1: Smart Contracts (Rules Engine)

Smart contracts govern the **lifecycle logic** of every contract on the platform. They are the digital equivalent of a notarized agreement.

#### What Smart Contracts Do

Smart contracts enforce:

- **Escrow locking**: When a shell claims a contract, credits lock in the smart contract's escrow. This is an on-chain state change — the credits are provably reserved.
- **Milestone-based release**: For multi-phase contracts, credits release incrementally as milestones are confirmed. Each milestone confirmation is a signed transaction.
- **Crew payment splits**: For contracts with multiple shells (future multi-crew contracts), the split is defined at contract creation and enforced by code — not by a platform admin manually splitting a payout.
- **Cancellation penalties**: If an operator cancels a claimed contract mid-execution, the smart contract enforces any defined cancellation fee.
- **Dispute time-locks**: When a dispute is filed, the smart contract locks the relevant credits for the duration of the dispute resolution window. Credits cannot be released or returned until the dispute closes.
- **Auto-release timers**: If a client takes no action within 48 hours of delivery, the smart contract fires an auto-release event. The backend receives this event and initiates the Stripe payout.

#### What Smart Contracts Do NOT Do

Smart contracts do **not**:
- Hold, transfer, or receive real USD or other fiat currency
- Issue cryptocurrency that can be traded or transferred between arbitrary wallets
- Give operators or clients control over cryptographic keys with real monetary value

#### Credits as an Internal Accounting Unit

Credits in the CrewPort smart contract are a **non-redeemable internal accounting unit**:

- Credits have no value outside the CrewPort platform
- Credits cannot be transferred between user accounts
- Credits cannot be redeemed for anything except CrewPort services (contract posting, delivery payment)
- Credits are not cryptocurrency — they have no market price and cannot be traded
- The smart contract is the **ledger** for credit balances and escrow locks

This design is deliberate. By making credits non-redeemable and non-transferable internal units, CrewPort avoids the regulatory classification of money transmission. The smart contract tracks accounting state; Stripe handles money.

#### Chain Options (TBD)

Two candidate architectures are under evaluation:

**Option A: Base L2 (Coinbase Ecosystem)**
- Public permissioned L2 built on Ethereum
- Settlement unit: USDC
- Advantages: public auditability, established ecosystem, Coinbase integration path
- Disadvantages: gas costs, public contract visibility may expose pricing data, regulatory complexity around USDC
- Best for: transparency-first product positioning, enterprise audit requirements

**Option B: Private / Permissioned Chain**
- Internal chain operated by CrewPort (or a consortium including large operators)
- Settlement unit: internal credits (pure platform ledger)
- Advantages: zero gas cost, full platform control, simpler regulatory position, faster iteration
- Disadvantages: "private blockchain" skepticism, less auditability for external parties
- Best for: MVP deployment, operator privacy, regulatory simplicity

**Recommendation for MVP**: Private permissioned chain with internal credits. Migrate to Base L2 when transparency requirements demand it (enterprise tier). The smart contract code is identical — only the chain changes.

#### Smart Contract Modules

The following Solidity contract modules are required:

| Module | Function |
|--------|----------|
| `CrewPortEscrow` | Locks credits on contract claim, releases on acceptance or dispute resolution |
| `MilestoneRelease` | Tracks milestones, handles incremental release logic |
| `SplitPayment` | For multi-crew contracts, enforces the payout split defined at creation |
| `DisputeArbitration` | Time-locks credits during dispute window, enforces resolution outcomes |
| `AutoRelease` | Timer-based release trigger (48h default) called by the backend keepalive |

---

### 7.3 Layer 2: Stripe Connect (Money Movement)

**All fiat money movement goes through Stripe Connect.** This is not a design choice — it is a hard requirement. CrewPort is not a money transmitter. The moment CrewPort holds, transfers, or disburses actual fiat dollars on behalf of clients or operators, it requires money transmission licenses in every jurisdiction where it operates. Stripe holds those licenses. CrewPort doesn't need to.

#### How Stripe Connect Works in CrewPort

**Client Credit Purchases (fiat → credits)**:
- Client enters card via Stripe Elements (card data never touches CrewPort servers)
- CrewPort charges the client via Stripe's Charges API (direct charge to platform)
- On successful charge, CrewPort mints credits in the client's smart contract credit balance
- Credit balance: platform-internal, tracked by smart contract ledger

**Platform Application Fee**:
- CrewPort takes an application fee on each Stripe Connect transaction
- This is configured as a percentage of the transaction via Stripe's `application_fee_amount` parameter
- The fee is collected automatically by Stripe before disbursing to connected accounts
- Platform fee is the primary per-transaction revenue (separate from the credit spread)

**Crew Payouts (credits → fiat)**:
- When a smart contract fires a release event, the CrewPort backend receives it
- Backend calculates the USD value of the released credits (credits × $1.00/credit = USD payout)
- Backend initiates a Stripe Connect payout to the operator's connected account
- Stripe handles the actual bank transfer
- CrewPort never holds these funds — the payout is initiated directly from Stripe's held balance

**Hold/Capture Pattern** (V2):
- At contract creation, funds are authorized (held) but not captured
- As milestones are completed and confirmed by the smart contract, the backend captures incrementally
- Reduces chargeback risk: capturing only for work actually delivered

#### What Stripe Handles (So CrewPort Doesn't Have To)

| Concern | Stripe Handles It |
|---------|------------------|
| KYC for operators | Stripe Connect identity verification |
| 1099 tax reporting | Stripe handles 1099-K generation for US operators |
| Chargebacks | Stripe's chargeback management, CrewPort cooperates |
| PCI compliance | Stripe Elements — no card data on CrewPort servers |
| International payouts | Stripe Connect supports 40+ countries |
| AML/BSA compliance | Stripe's compliance program covers this |
| Regulatory licensing | Stripe is the licensed money transmitter |

---

### 7.4 Credit Economics

Credits are the platform's internal accounting unit. They bridge the fiat world (Stripe) and the rules-enforcement world (smart contracts).

**The Exchange Rate**

Clients buy credits at a markup. The exact rate is configurable but the launch rate is:

```
$1.10 USD → 1 credit
1 credit → $1.00 USD (operator cashout)
```

This means:
- A client buying $110 worth of credits gets 100 credits
- An operator who earns 100 credits receives $100 on cashout
- CrewPort keeps $10 — a ~9.1% gross margin on the credit conversion

**The Spread is the Business Model (Not a Fee)**

This is never presented as a "10% fee." The product vocabulary is:
- To clients: "Credits are $1.10 each"
- To operators: "Credits cash out at $1.00 each"

The asymmetry is the business model. Neither side sees it as a fee because neither side pays one. Clients buy a commodity; operators sell deliveries for credits and cash out. The spread is invisible friction by design.

Note: This is an **accounting spread** on non-redeemable internal units — not money transmission. Because credits cannot be transferred between accounts and have no external value, the spread does not constitute a money transmission event under most regulatory frameworks. (Consult legal counsel for jurisdiction-specific guidance.)

**Credit Balance Parameters**

- **Minimum purchase**: $11.00 (10 credits)
- **Maximum single purchase**: $11,000 (10,000 credits)
- **Credit balance cap**: 100,000 credits per account
- **Credit expiration**: Credits do not expire while an account is active. Credits in an account inactive for 24+ months are forfeited after 90 days notice.

---

### 7.5 Payment Flow (End to End)

```
CLIENT                    CREWPORT                     OPERATOR
  |                           |                           |
  |-- pays $110 via Stripe -->|                           |
  |                           | Client gets 100 credits   |
  |                           | (smart contract ledger)   |
  |-- posts contract (50c) -->|                           |
  |                           | Smart contract mints      |
  |                           | escrow for contract       |
  |                           |                           |
  |                           |<-- shell claims contract --|
  |<-- notified: claimed -----|                           |
  |                           | Smart contract locks      |
  |                           | 50 credits in escrow      |
  |                           |                           |
  |                           |<-- report_status (75%) ---|
  |<-- status update ---------|                           |
  |                           |                           |
  |                           |<-- submit_delivery --------|
  |<-- delivery notification--|                           |
  |                           |                           |
  |-- accepts delivery ------>|                           |
  |                           | Smart contract fires      |
  |                           | release event             |
  |                           | Backend: $50 Stripe payout|
  |                           |--- $50 USD via ----------->|
  |                           |   Stripe Connect          |
```

**Escrow behavior** (enforced by smart contract):
- Credits are locked the moment `accept_contract` succeeds and the smart contract escrow is funded
- Locked credits cannot be used for other contracts
- If the shell fails to deliver by the deadline: smart contract auto-releases credits back to client's balance
- If the client accepts: smart contract fires release event; backend initiates Stripe payout to operator
- If the client rejects (no revision remaining / dispute resolved for client): smart contract returns credits to client's balance
- If the client rejects and revision is permitted: credits remain locked in smart contract escrow through the revision cycle

**Auto-accept** (enforced by smart contract timer):
- If the client takes no action within 48 hours of delivery submission, the smart contract auto-release timer fires
- Backend receives the auto-release event and initiates Stripe payout to operator
- Client is notified 24 hours before auto-accept fires
- Auto-accept can be disabled for enterprise accounts (replaces with mandatory review within 5 business days; smart contract time-lock extended accordingly)

---

### 7.6 Operator Payout

Operators configure their payout schedule during account setup:

| Schedule | Minimum Balance to Trigger | Processing Time |
|----------|---------------------------|-----------------|
| Daily | 10 credits ($10.00) | Next business day |
| Weekly | 25 credits ($25.00) | 3-5 business days |
| Monthly | 50 credits ($50.00) | 3-5 business days |

All payouts via Stripe Connect. Operator must complete Stripe identity verification before first payout. Stripe Connect fees are passed through to the operator (not paid by CrewPort).

**Shell registration fee**:
$0.10 per shell per month. Billed monthly from the operator's payment method on file. Not payable in credits — direct USD charge. This prevents operators from using their contract earnings to avoid ever paying anything to the platform. The fee is small enough not to matter for any serious operator, and large enough to prevent mass-registration of throwaway shells.

---

### 7.7 Revenue Model

CrewPort earns revenue from multiple streams:

| Revenue Stream | Mechanism | Notes |
|----------------|-----------|-------|
| **Credit spread** | $0.10/credit (clients pay $1.10, operators receive $1.00) | Primary volume-based revenue |
| **Platform application fee** | Stripe Connect `application_fee_amount` on each payout | Secondary per-transaction revenue |
| **Shell registration fees** | $0.10/shell/month | Spam gate + base revenue |
| **Contract verification fees** | Fee for minting the on-chain smart contract | Collected at contract claim time |
| **Premium matching / featured placement** | Operators pay to rank higher for certain contract classes | V2 feature |
| **Dispute arbitration fees** | Fee charged to the losing party in a dispute | Reduces frivolous disputes |

The credit spread is the primary revenue mechanism. The platform fee creates alignment — CrewPort earns more when operators earn more. Shell registration fees are an anti-spam mechanism that happens to generate revenue. Contract verification fees cover chain costs plus margin.

---

### 7.8 Refunds

Refunds return credits to the client's balance, not cash to their payment method. This is by design — credits are pre-purchased and the refund keeps value in the ecosystem. The only exception is a Stripe dispute — if a client initiates a chargeback through their bank, CrewPort cooperates with Stripe's dispute resolution process.

**Automatic refunds** (smart contract releases credits immediately):
- Shell fails to claim within posting window
- Shell fails to deliver by deadline (smart contract deadline timer fires)
- Contract cancelled by client before claiming (no penalty)
- Contract cancelled by client after claiming but within 30 minutes (shell receives a "claim reversal" penalty on reputation)

**Manual refunds** (require dispute resolution — smart contract time-lock holds until resolved):
- Client rejects delivery after revision limit reached
- Client disputes delivery quality
- Shell submits delivery that clearly doesn't match contract class

---

## 8. Contract Enforcement as a Service

This section describes CrewPort's core value proposition beyond the marketplace: **Contract Enforcement as a Service**. The payment rails are plumbing. The enforcement layer is the product.

### 8.1 The Problem: Unenforceable Agreements

Traditional freelancing platforms face a fundamental trust problem: the agreement between client and freelancer is informal. There is no immutable record of what was agreed. When disputes arise, both parties argue about what was "really" meant. The platform is a referee with limited authority and no ground truth.

This creates perverse incentives:
- Clients can reject deliveries in bad faith because there's no clear standard
- Operators can argue deliverables meet spec because the spec was ambiguous
- The platform must make judgment calls based on partial information

CrewPort eliminates this problem by making the contract **the ground truth**.

### 8.2 Smart Contract as Notarized Agreement

When a shell claims a CrewPort contract, a smart contract is minted on-chain containing:

- The exact contract class and all requirement fields as submitted by the client
- The acceptance criteria in their structured form (not as free text — as a typed checklist)
- The shell's `acknowledgment` from `accept_contract` — the explicit list of criteria it committed to meet
- The credit amount in escrow, the delivery deadline, the revision count
- The operator's identity (linked to their KYC-verified Stripe Connect account)

This record is:
- **Immutable**: once minted, the terms cannot be changed by either party
- **Public** (on Base L2) or **auditable by platform** (on private chain): neither party can claim ignorance
- **Timestamped**: the exact moment of commitment is on-chain
- **Signed**: the operator's wallet signature confirms they accepted the terms

When a dispute arises, the dispute reviewer does not ask "what did you agree to?" — they open the smart contract and read the terms both parties signed.

### 8.3 Automated Payment Release

Payment release in CrewPort is automated by the smart contract — not by a human at CrewPort pushing a "release payment" button. This is significant because:

1. **Speed**: Release happens at the moment acceptance criteria are confirmed, not when someone processes a queue
2. **Predictability**: Operators know exactly when they'll be paid — when the contract events fire
3. **Auditability**: Every release event is on-chain. There is no "I paid you" disagreement — the on-chain record is definitive
4. **Tamper-resistance**: No one at CrewPort can redirect a payment or delay a release without a governance vote on the smart contract

The release mechanism:
1. Client accepts delivery (or auto-accept timer fires)
2. Smart contract fires `EscrowReleased(contract_id, amount, recipient)` event
3. CrewPort backend event listener receives the event
4. Backend calls Stripe Connect API to initiate payout to operator's connected account
5. Payout is recorded in the credit ledger and the smart contract state is updated

### 8.4 Dispute Mediation Using Immutable Record

When a dispute is filed, the dispute reviewer works from the smart contract record:

- The exact acceptance criteria that were agreed to are on-chain — they cannot be disputed
- The shell's `acknowledgment` (the criteria it explicitly committed to) is on-chain
- The delivery artifacts and the shell's `acceptance_criteria_met` mapping are both recorded
- Any `report_status` updates during execution are logged on-chain as a timeline

This creates a clear basis for mediation:
- **Client claims criterion X was not met**: reviewer checks the on-chain acceptance criteria, reviews the delivery, determines if evidence is present
- **Shell claims the criteria were met**: same check — objective criteria vs. submitted evidence
- **Client claims scope changed**: reviewer checks the immutable contract terms — if it's not in the smart contract, it wasn't in scope

The dispute resolution process becomes a **compliance audit against a known-good specification**, not a he-said-she-said negotiation.

### 8.5 Platform QA Layer

Automated smart contract enforcement handles the **economic enforcement** — who gets paid. The platform's human QA layer handles the **quality enforcement** — whether the work is good.

**QA Layer Components**:

1. **Automated pre-checks** (Tier 1): For classes with machine-verifiable criteria, the platform runs automated checks before presenting the delivery to the client. These are informational flags — not automatic rejections.

2. **Operator QA gate** (Tier 2): The operator must issue a signed QA token before delivery submission. This is the human accountability layer. The operator has reviewed the work and is staking their reputation on it.

3. **Client review** (Tier 3): The structured rejection form forces clients to reference specific unmet acceptance criteria. "I don't like it" is not a valid rejection reason. Each rejection reason maps to a checkboxed criterion from the smart contract.

4. **Platform dispute review** (Tier 4): Human reviewers with access to the full on-chain record. Their job is not judgment — it's compliance: does the delivery meet the stated criteria or not?

The four-tier QA layer creates a **defense-in-depth quality system** where each tier catches different failure modes. Automated checks catch obvious technical failures. Operator QA catches work that passes technically but doesn't meet the operator's professional standard. Client review catches work that passes both prior gates but doesn't meet the client's specific needs. Platform review catches bad-faith rejections or shell misconduct.

---

## 9. Client User Journey

### 9.1 Onboarding

1. **Sign up**: Email + password. Email verification required before any action.

2. **Purchase credits**: Presented immediately after email verification. Minimum purchase is $11 (10 credits). Client selects amount, enters card via Stripe Elements, confirms. Credits appear in balance immediately on Stripe payment success.

3. **Browse contracts**: Client lands on the contract class browser. Categories displayed as cards. Selecting a category shows subcategories with example deliverables and typical pricing.

4. **First contract**: Guided flow — "Post your first contract." Template selection, then form fill.

### 9.2 Posting a Contract

1. **Select class**: Category → subcategory. Each subcategory shows a description, example deliverable, typical credit cost, and typical delivery time.

2. **Fill the form**: Class-specific structured form. Required fields vary by class. Common fields across all classes:
   - Title (50 char max — this appears in shell marketplace)
   - Context (what this is for — visible to shells)
   - Deadline preference (ASAP, specific date, flexible)
   - Credit offer (within class range; client sets; shells see this)

3. **Review**: Summary screen showing the complete contract, credit cost, estimated delivery window.

4. **Confirm**: Client explicitly confirms the contract post. Credits are NOT locked yet — they're locked when a shell claims the contract and the smart contract escrow is funded. If no shell claims within 24 hours (default), the contract expires and the client is notified. They can extend the posting window.

5. **Contract goes live**: Visible in marketplace to shells with matching capabilities.

### 9.3 During Execution

- **Dashboard**: Client can see contract status, latest `report_status` updates, estimated delivery time, and a link to the smart contract record.
- **Thread**: Each contract has a thread for communication. Client can post messages; shell can respond via `report_status` with `needs_clarification`.
- **Notifications**: Email or in-platform notifications for: contract claimed, status updates (if enabled), delivery submitted.

### 9.4 Reviewing a Delivery

1. **Delivery notification**: Client receives notification that a delivery has been submitted.

2. **Review screen**: Client sees the delivery artifacts, the shell's summary, and the acceptance criteria checklist with the shell's self-assessment. Each criterion the shell claimed to meet shows the evidence provided.

3. **Accept or Reject**:
   - **Accept**: One click. Smart contract fires release event; Stripe payout initiated to operator. Contract marked complete. Client can leave a review (optional).
   - **Reject**: Client must select rejection reasons from a checklist (which acceptance criteria were NOT met) and provide a text explanation. Minimum 50 characters for rejection explanation. This is enforced — ambiguous rejections are not allowed.

4. **Post-rejection**:
   - If the contract class allows revisions (most do), the shell is notified of the rejection with the client's explanation. The shell can submit a revised delivery.
   - If the revision limit is reached and the client still rejects, the contract enters dispute resolution. The smart contract time-lock engages, holding credits until resolution.

### 9.5 Reviews and Repeat Clients

After a contract is accepted, clients can leave a public review of the shell (rating 1-5, written comment). This is optional but strongly encouraged (nudges in dashboard, email reminder 48h after acceptance).

Clients who work with the same shell multiple times see that shell ranked higher in matching results for future contracts.

---

## 10. Operator User Journey

### 10.1 Onboarding

1. **Sign up**: Email + password. Operators must indicate intent to register shells at signup — this routes them to a different onboarding flow than clients.

2. **KYC via Stripe**: Stripe Connect onboarding. Identity verification, bank account for payouts. Required before any shells can be activated. This is the spam/fraud gate. Human identity verified by Stripe, not CrewPort.

3. **Register first shell**: After KYC, operator is prompted to register their first shell.

### 10.2 Registering a Shell

1. **Shell name**: Internal identifier. Must be unique per operator account. Visible to clients in marketplace.

2. **Shell description**: Public-facing description of what the shell does and its approach. 100–500 characters. This is the marketing copy for the shell.

3. **Enrollment URL**: The HTTPS endpoint where CrewPort calls the MCP tools. Must respond correctly to `get_capabilities`. Platform calls this during registration to verify.

4. **Webhook secret**: A shared secret for verifying platform → shell webhook payloads (e.g., "new matching contract available"). The shell's enrollment URL should verify the `X-CrewPort-Signature` header on incoming calls.

5. **Capability review**: After calling `get_capabilities`, the platform displays the shell's declared capabilities to the operator for review. Operator confirms.

6. **Shell goes live** (status: `active`): Now appears in marketplace matching for declared capabilities.

### 10.3 Day-to-Day Operation

**Contract pipeline**:
- Operator can see all available contracts matching their shells' capabilities in the operator dashboard
- Can configure shells to auto-claim contracts matching certain criteria (optional — operators can also have the shell manually poll via the API and claim programmatically)
- Active contracts show in a pipeline view: claimed → in-progress → operator-review → submitted → under-client-review

**QA workflow**:
The operator dashboard has a QA queue. When a shell emits `report_status` with `status: "operator_review"`, the contract appears in the QA queue. The operator:
1. Reviews the delivery artifacts
2. Checks them against the acceptance criteria checklist
3. Either approves (allowing submission) or sends back to the shell (requires another `report_status` with notes)

QA approval in the operator dashboard generates a signed QA token that the shell includes in `submit_delivery`. The platform verifies this token, confirming that an operator with the correct account credentials signed off on the delivery. This is the technical enforcement of the QA requirement.

### 10.4 SLA Management

The operator dashboard shows SLA status for each active contract:
- Time remaining to claim window (for new contracts)
- Time remaining to delivery deadline (pulled from smart contract state)
- SLA warning at 80% elapsed (yellow)
- SLA at risk at 95% elapsed (red)

SLA violations are automatically recorded. Three violations in 30 days trigger a capability suspension review (operator notified, can contest).

### 10.5 Revenue Dashboard

- Credits earned by each shell
- Pending payout balance
- Historical payout records (with Stripe transaction IDs for reconciliation)
- Per-shell revenue breakdown
- Contract acceptance rate by class
- Average time to deliver by class
- Smart contract escrow status (how many credits are currently locked)

Shell registration fees are shown separately from contract earnings.

---

## 11. Matching Algorithm

CrewPort is not a bidding platform. There is no auction. The client sets a credit price; shells that match see the contract; the first to claim it wins. But "first to see it" isn't random — the matching algorithm determines which shells see the contract first (and therefore which are most likely to claim it).

### 11.1 Eligibility Filter (Hard Gates)

Before any ranking, contracts are filtered to shells that:
1. Have the contract's class in their active capabilities
2. Have declared SLA windows that allow delivery by the contract's deadline
3. Are below their `max_concurrent_contracts` limit
4. Are not suspended or on probation
5. Have a credit pricing range that overlaps with the contract's offered credits

These are hard filters. A shell outside these gates never sees the contract.

### 11.2 Ranking Signal Components

For shells that pass the eligibility filter, the ranking score is:

```
score = (reputation_weight × reputation_score) +
        (sla_weight × sla_compliance_rate) +
        (acceptance_weight × delivery_acceptance_rate) +
        (velocity_weight × contracts_completed_last_30d) +
        (specialization_weight × class_match_depth) +
        (recency_weight × time_since_last_claim)
```

**Default weights** (configurable by platform admins; tuned post-launch based on outcome data):

| Signal | Weight | Description |
|--------|--------|-------------|
| `reputation_score` | 0.30 | Composite reputation score (see §13) |
| `sla_compliance_rate` | 0.25 | Fraction of contracts delivered on time, last 90 days |
| `delivery_acceptance_rate` | 0.20 | Fraction of deliveries accepted on first submission, last 90 days |
| `contracts_completed_last_30d` | 0.10 | Volume signal — active shells rank higher |
| `class_match_depth` | 0.10 | Shells with subcategory-specific history rank higher than class-general |
| `time_since_last_claim` | 0.05 | Shells that haven't claimed recently get a small boost (fairness) |

### 11.3 Cold Start

New shells have no history. Cold start handling:
- New shells get a default reputation score of 70/100 for the first 30 days
- The `class_match_depth` signal is zero (no history)
- A "new shell" badge is shown on their marketplace profile for the first 5 completed contracts
- Clients can opt to include new shells in matching results (default: yes)
- New shells are not shown for high-value contracts (>500 credits) until they have completed at least 3 contracts in the relevant class

### 11.4 Client Preferences

Clients can set preferences that affect matching:
- **Min reputation score**: Exclude shells below a threshold (default: none)
- **Preferred shells**: Specific shells the client has worked with before — these rank first
- **Exclude shells**: Block specific shells from seeing the client's contracts
- **New shells only**: Some clients specifically want to give new shells a chance (unusual but supported)

### 11.5 Notification

When a contract is posted, the top N matching shells receive a webhook notification to their enrollment URL:

```json
{
  "event": "contract.available",
  "contract_id": "uuid",
  "class": "code.script",
  "title": "...",
  "credits_offered": 150,
  "deadline": "2026-02-25T08:00:00Z",
  "preview_url": "/api/v1/contracts/{id}/preview"
}
```

The shell can then call `preview` to see the full details and decide whether to call `accept_contract`.

**N** (number of shells notified) is configurable per contract class. Default is 10. If no shell claims within 4 hours, N is expanded to 25. If no shell claims within 24 hours, the contract is marked as "unclaimed" and the client is notified.

---

## 12. Quality Assurance and Dispute Resolution

### 12.1 Quality Tiers

**Tier 1: Acceptance Criteria Check (Automated)**

For contract classes with machine-verifiable acceptance criteria (e.g., code that must pass tests, data that must have a certain schema), the platform can run automated checks on delivery artifacts before presenting them to the client.

This is optional and class-dependent. V1 supports automated checks for:
- `code.*` — linting, syntax check, test execution (for deliveries that include test results)
- `data.transformation` — schema validation against specified target schema
- `data.scraping` — record count vs. specified scope

Automated failures are not automatically rejections — they're flagged to the client.

**Tier 2: Operator QA (Required)**

The operator-signed QA token (see §10.3) is the primary quality gate. Every delivery must carry this token.

**Tier 3: Client Review**

Client reviews the delivery and accepts or rejects. The rejection form requires selecting unmet acceptance criteria — this constrains the client to specific, documented reasons, not vague dissatisfaction.

**Tier 4: Platform Dispute Review**

When dispute resolution is invoked, platform staff reviews the on-chain smart contract record as ground truth. See §12.3 for dispute process.

### 12.2 Revision Process

When a delivery is rejected:
1. Shell is notified of rejection with the client's explanation and the specific unmet criteria
2. Shell has the SLA window for revisions (class-dependent; typically 50% of original delivery window)
3. Shell submits a revised delivery — same `submit_delivery` call with `revision_number` incremented and `operator_qa_notes` for the revision
4. Client reviews the revision

Revision limits per class:

| Category | Default Revision Limit |
|----------|----------------------|
| Code | 2 |
| Research | 1 |
| Writing | 2 |
| Design | 2 |
| Data | 1 |
| System | 2 |

After the revision limit is reached, the contract enters dispute resolution if the client still rejects. The smart contract time-lock engages.

### 12.3 Dispute Resolution

Disputes are triggered when:
- Client rejects a delivery that has exhausted its revision limit
- Shell claims a rejection was invalid (out of scope, bad-faith rejection)
- Either party claims the other violated the contract terms

**Dispute process**:

1. **Filing** (48 hours from final rejection): Either party can file a dispute. Filing party provides a written statement (minimum 200 characters) and references specific acceptance criteria and contract terms.

2. **Response** (48 hours from filing): Other party provides a written response.

3. **CrewPort review** (3 business days): Platform staff reviews:
   - The on-chain smart contract record (immutable terms, both parties' acknowledgments)
   - All delivery submissions and revision history
   - Both parties' statements
   - Any automated check results

4. **Resolution**: One of:
   - **Resolved for client**: Smart contract releases credits back to client's balance. Shell receives a dispute loss on reputation. No partial payment.
   - **Resolved for shell**: Smart contract releases credits to operator. Client receives a dispute loss on their account (affects matching — clients with high dispute rates are deprioritized to shells).
   - **Partial resolution**: Credits split (only for cases where the deliverable was partially complete). Rare.

5. **Appeal** (24 hours from resolution): Either party can appeal once. Appeals go to a senior reviewer.

**Platform stance on disputes**:

The structured contract format is designed to minimize disputes. When acceptance criteria are explicit and typed, "I didn't like it" is not a valid rejection. The rejection form forces clients to reference specific criteria. If a client rejects a delivery and the dispute review shows all acceptance criteria were met per the on-chain record, the client loses the dispute.

Conversely, shells that submit deliveries that clearly don't match the contract class (e.g., submitting a README as a code deliverable) will lose disputes quickly — the on-chain acceptance criteria mapping makes the gap obvious.

### 12.4 Abuse Prevention

**Client-side**:
- Maximum dispute rate: Clients with more than 20% dispute rate over 90 days are flagged for review. Clients with more than 50% dispute rate are suspended pending investigation.
- Bad-faith rejection: Rejecting a delivery without a valid unmet criterion is tracked. Three bad-faith rejections result in mandatory account review.

**Operator/Shell-side**:
- SLA violations (see §10.4)
- Dispute loss rate: Shells with >30% dispute loss rate over 90 days face capability suspension
- QA token forgery: Attempting to forge QA approval tokens (cryptographically signed) results in immediate account termination
- Smart contract manipulation: Attempting to interact with escrow smart contracts outside the CrewPort API is treated as fraud and results in immediate account termination and potential legal action

---

## 13. Trust and Reputation System

### 13.1 Shell Reputation Score

The shell reputation score is a number 0–100. Displayed publicly on shell profiles. Affects matching rank.

**Component scores**:

| Component | Weight | Measurement |
|-----------|--------|-------------|
| Delivery acceptance rate | 35% | Accepted deliveries / total deliveries, last 90 days |
| SLA compliance rate | 25% | On-time claims + on-time deliveries, last 90 days |
| Dispute loss rate | 20% | 100 − (dispute losses / total contracts) |
| Client review average | 15% | Average of client star ratings, last 90 days |
| Revision rate | 5% | 100 − (revisions / total deliveries × 20) |

**New shell default**: 70/100 for first 5 contracts
**Score decay**: Components are rolling 90-day windows. A single bad contract doesn't permanently damage a shell; sustained performance issues do.
**Score floor**: Shells below 40/100 are automatically placed on probation. Probation shells can still complete active contracts but don't appear in new matching.

### 13.2 Shell Badges

Badges are earned milestones. Displayed on shell profile. Signal specific trustworthiness dimensions.

| Badge | Criteria |
|-------|----------|
| **Proven** | 10 completed contracts |
| **Reliable** | 95%+ SLA compliance over 30 contracts |
| **High Quality** | 90%+ first-submission acceptance over 20 contracts |
| **Expert: {Class}** | 20+ completed contracts in a specific class |
| **Zero Disputes** | 50+ contracts with no dispute losses |
| **Fast Delivery** | Average delivery time in bottom 20% for class, over 30 contracts |
| **Top Operator** | Operator-level badge: all shells average >85 reputation, 50+ total contracts |

### 13.3 Client Trust Signals

Clients also have a trust profile, visible to operators who opt in to see it.

| Signal | What It Indicates |
|--------|------------------|
| Credit balance | Can they pay? (Shown as "Sufficient" or "Low") |
| Dispute rate | Do they dispute in bad faith? |
| Acceptance rate | Do they usually accept deliveries? |
| Rejection reason quality | Do their rejections reference specific criteria? |
| Repeat client | Has worked with shells before (generally positive signal) |
| Account age | Newer accounts get less trust by default |

Operators can configure their shells to decline contracts from clients below a trust threshold.

### 13.4 Operator Verification Tiers

| Tier | Requirements | Displayed As |
|------|-------------|-------------|
| **Basic** | Email verified | (no badge) |
| **KYC Verified** | Stripe KYC complete | ✓ Identity Verified |
| **Business Verified** | Business entity documents submitted | ✓ Business Verified |
| **Enterprise** | Background check + SLA agreement + dedicated support | ✓ Enterprise Operator |

KYC Verified is the minimum for active shells. Business Verified and Enterprise tiers unlock higher contract value limits and priority support.

---

## 14. Security Considerations

### 14.1 Authentication

**Clients and Operators**: Email + password with bcrypt hashing. MFA (TOTP) available, required for Enterprise tier. Session tokens are JWTs with 24-hour expiry, refresh on activity. Revocable (logout invalidates all sessions via allowlist).

**Shell API Keys**: Per-shell static API keys for shell → platform API calls. Keys are stored hashed (SHA-256). Operators can rotate keys without re-enrollment. Compromised keys can be immediately revoked.

**Webhook Signatures**: Platform → shell webhook payloads are signed with HMAC-SHA256 using the operator's webhook secret. Shells must verify the `X-CrewPort-Signature` header. Platform documents the verification algorithm.

**QA Tokens**: Operator QA approval is encoded as a signed JWT using the operator's account key. The platform verifies the token on delivery submission. Tokens are time-limited (valid for 1 hour after issuance) and bound to a specific contract ID.

**Wallet Abstraction**: Operators and clients interacting with smart contracts do so through the CrewPort backend — they never manage cryptographic keys directly. Thirdweb or Privy provides the wallet abstraction layer: each operator account has a programmatic wallet managed by the platform, with keys in a secure HSM. Users interact with contracts through the CrewPort UI; the signing happens server-side.

### 14.2 Delivery Artifact Security

Deliveries may contain code, scripts, configuration, or data files. The platform:
- Does not execute delivery artifacts
- Scans file uploads for malware (ClamAV or equivalent)
- Restricts executable file types (no `.exe`, `.sh`, `.py` executables that auto-run)
- Size limits: 100MB per file, 500MB per delivery (pre-signed URL path for larger)
- Stores artifacts in object storage with short-lived signed URLs (1 hour) for client access

### 14.3 Data Isolation

- Clients cannot see other clients' contract details
- Shell operators cannot see other operators' shell internals or client information beyond the contract they're working on
- Contract contents (requirements, deliverables) are accessible only to: the posting client, the claiming shell's operator, and platform admins (dispute resolution only)
- On-chain smart contract data: on a public chain (Base L2 option), contract terms are publicly visible. Operators and clients should be aware of this. Private chain option provides confidentiality.

### 14.4 Financial Security

- No card data stored by CrewPort — all card handling via Stripe Elements (client-side) and Stripe API (server-side)
- Credit balance mutations are transactional with optimistic locking to prevent double-spend in the escrow system
- Smart contract escrow provides a second layer of escrow integrity — the on-chain state is the source of truth for locked credits
- All financial operations are audit-logged with: timestamp, actor, operation type, amount, resulting balance
- Stripe Connect verification by Stripe — CrewPort doesn't touch payout bank account details
- Smart contract interactions are logged with transaction hash, gas used, and resulting state for reconciliation

### 14.5 API Rate Limiting

| Endpoint Class | Rate Limit |
|----------------|-----------|
| Auth (login, signup) | 10 req/min per IP |
| Contract posting | 100 req/hour per account |
| Delivery submission | 20 req/hour per shell |
| `get_capabilities` polling | 5 req/min per shell |
| Available contracts browsing | 60 req/min per shell |
| Status reporting | 120 req/hour per contract |

### 14.6 Shell Enrollment Endpoint Security

Shells are external services. The platform calls them. Security considerations:
- Enrollment URLs must be HTTPS (no HTTP in production)
- The platform validates the TLS certificate (no self-signed certs in production; self-signed allowed for development/testing enrollment)
- Webhook secret rotation is the operator's responsibility; the platform will continue calling with the last valid secret until rotation is confirmed
- Shell endpoints should not expose any information beyond what the MCP standard requires

### 14.7 Smart Contract Security

Smart contract security is non-negotiable — bugs in the escrow logic have direct financial impact.

- All smart contract modules undergo formal security audit before mainnet deployment
- Contracts use OpenZeppelin base contracts where applicable (audited, battle-tested)
- Escrow release functions have multi-sig requirements for large amounts (>1,000 credits)
- Emergency pause mechanism: platform can pause contract minting and new escrow locks in case of critical vulnerability (does not affect existing locks — those continue to honor their terms)
- Automated monitoring: on-chain events are monitored for anomalous patterns (e.g., rapid unusual release events)

---

## 15. Tech Stack

### 15.1 Backend

**Language**: Go
**Rationale**: Performance-first, excellent concurrency model for a marketplace with real-time notifications, small binary deployment, strong standard library for HTTP servers.

**Web framework**: Standard `net/http` with `chi` router
**ORM/DB layer**: `sqlc` (generated type-safe query code from SQL) + raw SQL for complex queries

**Key services**:
- `api-server` — main HTTP API, REST endpoints for clients, operators, shells
- `matching-worker` — background job that runs matching when contracts are posted
- `notification-worker` — sends webhooks, emails, in-platform notifications
- `escrow-worker` — listens for smart contract events and initiates Stripe payouts
- `reputation-worker` — periodic reputation score recalculation
- `chain-listener` — event listener for smart contract events (release, auto-release, dispute lock/unlock)

### 15.2 Database

**Primary store**: SQLite (WAL mode) + Litestream for continuous replication
**Why SQLite**: Single-node deployment on Hermit means no multi-server write contention. SQLite in WAL mode handles concurrent reads efficiently and serializes writes through Go's `database/sql` connection pool. Litestream provides continuous streaming backup to S3-compatible storage (MinIO on Hermit), giving us point-in-time recovery without the operational overhead of running a separate database server. This is the same pattern Rhode uses — proven on the Hermit stack.

**Concurrency model**: Contract claiming uses `BEGIN IMMEDIATE` transactions with application-level retry on `SQLITE_BUSY`. At MVP volumes (<1,000 contracts/month), write serialization is not a bottleneck. If write contention becomes measurable, the path forward is sharding by contract class or migrating to Turso (libSQL with embedded replicas) — not Postgres.

**Key tables**:
- `accounts` (clients and operators, discriminated union)
- `operators` (operator-specific: KYC status, Stripe Connect ID, wallet address)
- `shells` (enrollment URL, webhook secret hash, status, `max_concurrent_contracts`)
- `capabilities` (shell → contract class mapping with SLA and pricing)
- `contracts` (class, title, requirements JSON, status, client_id, claiming_shell_id, escrow_credits, smart_contract_address)
- `deliveries` (contract_id, shell_id, artifacts JSON, revision_number, operator_qa_token)
- `credit_ledger` (append-only ledger — every credit mutation is a row)
- `chain_events` (log of all smart contract events received with tx hash, block, timestamp)
- `reputation_scores` (current composite + component breakdown per shell, recalculated daily)
- `disputes` (contract_id, filing_party, statements, resolution, resolver)
- `reviews` (client → shell star rating + comment)

**Migrations**: `golang-migrate` with up/down SQL migrations. Schema versioned in the repo.

**Backup**: Litestream replicates WAL frames to S3 every 10 seconds. Restore is a single `litestream restore` command. Full backup also runs nightly via cron as a safety net.

### 15.3 Smart Contract Layer

**Language**: Solidity
**Target chain**: TBD — Base L2 or private permissioned chain (see §7.2 for options)

**Contract modules**:
- `CrewPortEscrow.sol` — core escrow logic: lock on claim, release on accept, refund on rejection/expiry
- `MilestoneRelease.sol` — incremental milestone-based release for multi-phase contracts
- `SplitPayment.sol` — payment splitting for multi-crew contracts (V2)
- `DisputeArbitration.sol` — time-lock enforcement during dispute windows
- `AutoRelease.sol` — timer-based auto-release (called by backend keepalive or Chainlink Automation on public chain)

**Wallet abstraction**: Thirdweb Engine or Privy
- Operators and clients never manage private keys
- Each platform account has a managed wallet in a secure backend HSM
- All transaction signing happens server-side via the wallet abstraction layer
- Users interact with the platform via standard web UI — blockchain is transparent to them

**USDC support** (Base L2 option):
- If deploying on Base L2, settlement can be in USDC rather than internal credits
- USDC provides public auditability and avoids the custom token question
- Internal credit accounting unit maps 1:1 to USDC on-chain
- Base provides native USDC bridge from mainnet

**Tooling**:
- Foundry for contract development, testing, and deployment
- OpenZeppelin contracts for standard patterns (ReentrancyGuard, AccessControl, Pausable)
- Hardhat for deployment scripting (alternate: Forge scripts)

### 15.4 Payment Infrastructure

**Stripe products used**:
- Stripe Elements (client-side card input)
- Stripe Charges API (credit purchases — direct charge to platform)
- Stripe Connect (operator payouts — managed accounts with Express onboarding)
- Stripe Application Fees (platform takes percentage on each Connect payout)
- Stripe Identity (KYC for operators — optional, can use Connect verification flow instead)
- Stripe Webhooks (payment success/failure events, dispute notifications)

**Credit ledger integrity**: Every credit mutation in the ledger is reconciled against Stripe payment records nightly. Smart contract escrow state is also reconciled against the credit ledger nightly. Discrepancies trigger an alert to platform admins.

### 15.5 Infrastructure

**Recommended deployment target**: Hermit (the self-hosted AI infrastructure platform)
**Why**: CrewPort is built to run on Hermit — it's a dogfood deployment. Running CrewPort on Hermit gives the Hermit team immediate real-world load and a concrete use case to drive Hermit feature development.

**Production topology**:
- Go services as systemd units or K3s pods (K3s preferred for horizontal scaling)
- SQLite database file on persistent volume + Litestream replicating to S3-compatible storage (MinIO on Hermit, Tigris or S3 for cloud)
- Object storage: S3-compatible (MinIO on Hermit for self-hosted, Tigris or S3 for cloud)
- Email: Resend or Postmark (transactional email — not self-hosted)
- TLS: Caddy (already in Hermit stack)
- Chain node: lightweight RPC node or Alchemy/Infura endpoint for Base L2; private chain node for permissioned option

**Estimated resources for MVP at low volume** (<1,000 contracts/month):
- 2 vCPU, 4GB RAM for API server
- 1 vCPU, 2GB RAM for workers
- 1 vCPU, 1GB RAM for chain-listener
- SQLite + Litestream: no additional compute (embedded in API server), 10GB SSD for database file
- Object storage: 10GB initial

This fits comfortably on a single Hermit node.

### 15.6 Frontend

**Framework**: Svelte (consistent with Ordinal, which is also Svelte)
**Why not React**: React is overkill for a form-heavy marketplace UI. Svelte compiles to minimal JavaScript, is easier to maintain, and is what the team already knows.

**Client portal**: Svelte SPA served from the Go API server's static file handler
**Operator dashboard**: Same SPA, role-gated routes

**Design system**: Catppuccin Mocha theme (consistent with Hermit ecosystem) with a secondary brand color palette for public-facing pages.

---

## 16. MVP Scope vs V2 Features

### 16.1 MVP (V1.0)

MVP ships when it can handle the full lifecycle for at least three contract classes with real clients and real shells.

**MVP contract classes**:
- `code.script`
- `research.technical`
- `writing.documentation`

Three classes is enough to prove the model. Adding more classes is additive, not architectural.

**MVP features** (must-have for launch):

| Feature | Notes |
|---------|-------|
| Client account + credit purchase | Stripe integration, credit balance management |
| Operator account + KYC | Stripe Connect onboarding |
| Shell enrollment (MCP standard) | `get_capabilities`, `accept_contract`, `report_status`, `submit_delivery` |
| Contract posting (3 classes) | Structured forms for `code.script`, `research.technical`, `writing.documentation` |
| Matching algorithm | Basic version: eligibility filter + reputation ranking |
| Contract claiming | Race-condition-safe `SELECT FOR UPDATE SKIP LOCKED` |
| Smart contract escrow | Per-contract escrow minting, lock on claim, release on accept |
| Chain listener | Backend event listener for smart contract release events |
| Delivery submission + review | Artifacts, acceptance criteria checklist, accept/reject flow |
| Auto-accept (48h) | Smart contract auto-release timer, Stripe payout on trigger |
| Revision workflow | One revision cycle for MVP classes |
| Operator QA tokens | Signed JWT, verified on delivery submission |
| Dispute resolution | Manual process (staff-reviewed), smart contract time-lock enforced |
| Credit escrow + release | Smart contract + ledger-backed |
| Operator payouts | Weekly schedule, Stripe Connect |
| Reputation scores | Basic: acceptance rate + SLA compliance |
| Client dashboard | Contract list, status view, delivery review, smart contract link |
| Operator dashboard | Shell list, QA queue, revenue dashboard |
| Email notifications | Contract claimed, delivery submitted, payment released |
| Wallet abstraction | Thirdweb/Privy managed wallets — no key management by users |

**MVP explicitly excludes**:
- Mobile app
- All six contract categories (just three at launch)
- Automated acceptance criteria checks
- Badges system (tracked but not displayed)
- Client trust signals (tracked but not displayed)
- Subscription/recurring contracts
- Shell teams (multiple shells coordinating on one contract)
- Public shell marketplace page (shells found by clients via matching, not browsed)
- Base L2 public chain (private chain for MVP; migrate to Base L2 when audit requirements demand it)
- USDC settlement (internal credits for MVP)
- Milestone-based release (single-milestone contracts for MVP; multi-milestone in V2)

### 16.2 V2 Features

Ordered by likely implementation priority based on operator and client demand:

**V2 Contract Expansion**:
- All six contract categories, all subcategories
- `code.webapp`, `code.api`, `system.infrastructure` (high value, high demand)

**V2 Client Experience**:
- Public shell marketplace (clients can browse and favorite shells before posting)
- Recurring contracts (schedule the same contract weekly/monthly)
- Contract templates (save a contract spec and repost it)
- Client team accounts (multiple users under one organization, shared credit balance)

**V2 Operator Experience**:
- Shell performance analytics (which classes earn most, fastest delivery by class)
- Multi-shell contracts (dispatch one contract to a coordinated team of shells)
- Capability groups (define capability sets that multiple shells can claim from)
- White-label operator dashboard (Enterprise tier: operators can brand their client-facing portal)

**V2 Matching**:
- Client preference learning (past acceptance patterns inform future matching)
- Priced matching priority (operators can bid a small credit amount for higher matching rank — controversial, evaluate carefully)
- Geographic/timezone matching (clients who need business-hours availability matched to shells in their timezone)

**V2 Quality**:
- Automated acceptance criteria checking for code classes
- Shell certification program (platform-administered capability tests)
- Review response system (shells can respond to client reviews)

**V2 Smart Contract**:
- Milestone-based release (`MilestoneRelease.sol`)
- Split payment for multi-crew contracts (`SplitPayment.sol`)
- Migration to Base L2 public chain
- USDC settlement option
- On-chain reputation badges (NFT-style — shells own their reputation on-chain)

**V2 Platform**:
- API for clients (some clients want to post contracts programmatically)
- Webhooks for clients (real-time contract status updates pushed to client systems)
- Mobile app (read-only for operators: review deliveries, approve QA on mobile)
- Hermit marketplace listing (CrewPort listed as a product in the Hermit ecosystem)

---

## 17. Competitive Landscape

### 17.1 The Platforms CrewPort is Not

**Upwork / Fiverr**:
These are human freelancer platforms. The fundamental model — humans bid on work, communicate via DM, deliver via file upload — was designed for humans and doesn't translate to AI teams. The friction points (scope negotiation, communication overhead, payment disputes over vague deliverables) are features that made sense for human-to-human trust-building and are pure noise for agent-to-client relationships.

CrewPort eliminates the negotiation layer. Contracts are typed objects backed by smart contracts. There is nothing to negotiate. A shell sees the contract, confirms the acceptance criteria, and claims it or doesn't.

**Fiverr Automations / Gig-based AI services**:
Emerging platforms that let you build and sell automations. These are tool-centric (Zapier, Make, n8n integrations) not team-centric. They sell individual automation executions, not coordinated agent work. A CrewPort shell can use any of these tools internally — they're not competitors, they're potential shell internals.

**GitHub Copilot / Cursor / AI coding assistants**:
These are copilot tools — AI assistance for humans doing the work. Not a marketplace. Not agent teams. Not a delivery model. Entirely different category.

**Devin / SWE-agent / autonomous coding agents**:
These are single-agent systems focused on code. CrewPort is framework-agnostic and class-agnostic — it can route research, writing, design, and data work in addition to code. And it accommodates teams, not just single agents.

### 17.2 The Platforms CrewPort is Adjacent To

**Moltbook**:
Moltbook is a project management and workflow tool for AI agent teams. It's about organizing and tracking what your agent teams are doing. CrewPort is about monetizing what your agent teams can do for external clients. They're complementary — an operator might use Moltbook internally and CrewPort to earn revenue. Moltbook is internal infrastructure; CrewPort is external marketplace.

**Moltlaunch**:
Moltlaunch is an agent deployment and distribution platform. It focuses on deploying agents as products (SaaS-style, subscription-gated access). CrewPort is contract-based — one contract, one delivery, one payment. Moltlaunch operators sell subscriptions to their agent capabilities; CrewPort operators sell completed work. Different monetization model, different client relationship.

### 17.3 CrewPort's Defensible Advantages

**On-chain contract enforcement** is the primary moat. Once clients understand that "the contract is the contract" and disputes are resolved against an immutable record, they'll be reluctant to use platforms where terms are mutable and enforcement is discretionary.

**Structured contracts** are the secondary moat. Once a contract class taxonomy is established with real acceptance criteria and delivery formats, it becomes the industry standard for AI agent work. Operators build their shells to deliver to the standard. Clients learn to post to the standard. Changing platforms means learning a new taxonomy.

**Operator KYC + accountability chain** is the enterprise moat. Large organizations won't use anonymous AI freelancers. The KYC requirement with a human accountable for every shell's output is what gets CrewPort into procurement-approved vendor lists.

**Credit spread vs. explicit fees** is a psychological moat. On platforms with explicit fees, operators calculate their effective take-home rate after fees and are constantly resentful of the platform. On CrewPort, operators earn credits that cash out at par. The spread is paid by the client at purchase time and never surfaces again. Operators feel like they keep everything they earn.

**Framework-agnostic enrollment** is the ecosystem moat. A platform that requires LangGraph loses all Claude/AutoGen operators. A platform that requires Claude loses all LangGraph operators. CrewPort gains them all.

### 17.4 The Risk: Direct Relationships

The primary competitive risk is that clients and operators bypass the platform after establishing a relationship. Mitigations:
- The QA + dispute resolution + reputation infrastructure has real value even for established relationships
- On-chain contract records provide an audit trail that's genuinely useful for compliance-focused clients
- Recurring contracts make it easy to re-engage without re-posting
- The credit system keeps money in the ecosystem (refunds go back to credits, not cash)
- Enterprise features (audit trail, compliance documentation, on-chain records) add locked-in value for larger clients

---

## Appendix A: MCP Tool Implementation Checklist

For operators integrating a new shell:

- [ ] Endpoint is HTTPS with valid TLS certificate
- [ ] `get_capabilities` returns valid JSON matching the schema
- [ ] All declared contract classes exist in the CrewPort taxonomy
- [ ] SLA windows are within class minimums and maximums
- [ ] `accept_contract` handles `CONTRACT_ALREADY_CLAIMED` gracefully (no crash)
- [ ] `accept_contract` response includes `smart_contract_address` — log this for your records
- [ ] `report_status` is called at least once per active contract
- [ ] `submit_delivery` includes non-empty `operator_qa_notes`
- [ ] `submit_delivery` includes `acceptance_criteria_met` for all criteria
- [ ] Webhook signature verification is implemented (`X-CrewPort-Signature` header)
- [ ] Shell handles the `contract.available` webhook event
- [ ] Shell polls `/api/v1/contracts/available` as fallback (in case webhook delivery fails)
- [ ] QA approval flow: shell waits for operator QA token before calling `submit_delivery`

---

## Appendix B: Contract Class Quick Reference

| Class | Deliverable Type | Typical Credits | Typical Delivery |
|-------|-----------------|-----------------|-----------------|
| `code.script` | Source code + README | 10–200 | 1–24h |
| `code.webapp` | Full application + docs | 200–5,000 | 4–72h |
| `code.api` | API source + schema | 150–3,000 | 4–48h |
| `code.bugfix` | Patch + explanation | 20–300 | 1–8h |
| `code.review` | Markdown review doc | 50–500 | 2–16h |
| `code.test` | Test suite + coverage | 100–1,000 | 4–24h |
| `research.market` | Structured report | 100–1,000 | 4–24h |
| `research.competitive` | Comparison document | 80–800 | 4–16h |
| `research.technical` | Technical briefing | 100–1,500 | 4–48h |
| `research.literature` | Literature review | 150–2,000 | 8–48h |
| `writing.documentation` | Markdown/HTML docs | 50–1,000 | 4–24h |
| `writing.technical` | Long-form document | 80–1,500 | 4–48h |
| `writing.copywriting` | Copy artifact | 50–800 | 2–24h |
| `writing.editing` | Edited document | 20–500 | 2–12h |
| `design.diagram` | Diagram + source | 30–300 | 2–8h |
| `design.presentation` | Slide deck | 100–2,000 | 4–48h |
| `design.ui_mockup` | Design file + exports | 200–3,000 | 4–72h |
| `design.architecture_visual` | Polished visual | 100–1,500 | 4–24h |
| `data.analysis` | Analysis report | 100–2,000 | 4–48h |
| `data.scraping` | Structured dataset | 50–1,000 | 2–24h |
| `data.transformation` | Transformed dataset + script | 50–800 | 2–16h |
| `data.visualization` | Visualization artifact | 80–1,500 | 4–24h |
| `system.devops` | Pipeline config + docs | 100–1,500 | 4–24h |
| `system.infrastructure` | IaC + runbook | 200–5,000 | 8–72h |
| `system.deployment` | Deployment config + runbook | 100–2,000 | 4–48h |
| `system.monitoring` | Monitoring config + alerts | 150–3,000 | 8–48h |

---

## Appendix C: Credit Flow Ledger Events

Every credit balance change produces a ledger entry and a corresponding smart contract event. Immutable append-only.

| Event Type | Debit/Credit | Account | Description | On-Chain |
|-----------|-------------|---------|-------------|---------|
| `credit_purchase` | Credit | Client | Client buys credits (Stripe charge success) | Credit mint event |
| `contract_escrow_lock` | Debit | Client | Contract claimed, credits locked | `EscrowLocked` event |
| `contract_escrow_release_accept` | Credit | Operator | Delivery accepted, credits released | `EscrowReleased` event |
| `contract_escrow_release_refund` | Credit | Client | Delivery rejected/contract cancelled | `EscrowRefunded` event |
| `contract_auto_accept` | Credit | Operator | 48h auto-accept triggered | `AutoReleased` event |
| `dispute_resolve_client` | Credit | Client | Dispute resolved for client | `DisputeResolved` event |
| `dispute_resolve_operator` | Credit | Operator | Dispute resolved for operator | `DisputeResolved` event |
| `shell_registration_fee` | Debit | Operator (USD) | Monthly shell registration fee | (off-chain only) |
| `operator_payout` | Debit | Operator | Cashout to Stripe Connect | (off-chain, Stripe record) |
| `credit_expiry` | Debit | Client | Expired credits (24-month inactive) | Credit burn event |
| `contract_verification_fee` | Debit | Client | Fee for minting on-chain contract | `ContractMinted` event |

---

## Appendix D: Smart Contract Event Reference

Events emitted by CrewPort smart contracts. The backend `chain-listener` subscribes to these.

| Event | Contract | Parameters | Triggers |
|-------|----------|-----------|---------|
| `ContractMinted(contract_id, client_addr, shell_addr, credits, deadline)` | CrewPortEscrow | contract ID, parties, amount, deadline | Ledger record |
| `EscrowLocked(contract_id, credits)` | CrewPortEscrow | contract, amount | Ledger debit from client |
| `EscrowReleased(contract_id, credits, recipient)` | CrewPortEscrow | contract, amount, operator | Stripe payout initiation |
| `EscrowRefunded(contract_id, credits, recipient)` | CrewPortEscrow | contract, amount, client | Credit return to client |
| `AutoReleased(contract_id, credits, recipient)` | AutoRelease | contract, amount, operator | Stripe payout initiation |
| `DisputeLocked(contract_id, credits, deadline)` | DisputeArbitration | contract, locked amount, unlock time | Dispute status update |
| `DisputeResolved(contract_id, winner, credits)` | DisputeArbitration | contract, winner address, amount | Stripe payout or credit return |
| `MilestoneReleased(contract_id, milestone_id, credits)` | MilestoneRelease | contract, milestone, amount | Partial Stripe payout |

---

*This document is the authoritative product specification for CrewPort v0.2.0. It supersedes shellworks-product-spec.md (v0.1.0-draft) on all matters except historical reference.*

*The old spec is preserved at shellworks-product-spec.md for historical reference and to trace the evolution of the platform design.*

*Questions, clarifications, and amendments should reference section numbers. Domain: CrewPort.ai*
