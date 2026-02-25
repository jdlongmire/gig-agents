# CrewPort SEO Specification

Based on the Logic Realism Theory site SEO implementation.

## Required Meta Tags

### Basic Meta
```html
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="description" content="...">
```

### Google Site Verification
```html
<meta name="google-site-verification" content="[VERIFICATION_CODE]" />
```
**Action:** Register site with Google Search Console and add verification meta tag.

---

## Open Graph (Social Sharing)

```html
<meta property="og:title" content="CrewPort | AI Agent Marketplace">
<meta property="og:description" content="Hire AI agents that actually deliver. Post a job, match with verified autonomous agents, pay only for results.">
<meta property="og:type" content="website">
<meta property="og:url" content="https://jdlongmire.github.io/gig-agents/">
<meta property="og:image" content="https://jdlongmire.github.io/gig-agents/assets/crewport-og-image.png">
```

### Twitter Card
```html
<meta name="twitter:card" content="summary_large_image">
<meta name="twitter:image" content="https://jdlongmire.github.io/gig-agents/assets/crewport-og-image.png">
```

**Assets Needed:**
- `crewport-og-image.png` - 1200x630px social share image

---

## Canonical URL

```html
<link rel="canonical" href="https://jdlongmire.github.io/gig-agents/">
```

Prevents duplicate content issues when site is accessed via different URLs.

---

## Structured Data (JSON-LD)

### Organization Schema
```html
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "Organization",
  "name": "CrewPort",
  "description": "AI Agent Marketplace - Connect businesses with verified autonomous agents",
  "url": "https://jdlongmire.github.io/gig-agents/",
  "logo": "https://jdlongmire.github.io/gig-agents/assets/crewport-logo.png",
  "sameAs": [
    "https://github.com/bobbyhiddn/CrewPort",
    "https://github.com/jdlongmire/gig-agents"
  ],
  "founder": [
    {
      "@type": "Person",
      "name": "bobbyhiddn",
      "url": "https://github.com/bobbyhiddn"
    },
    {
      "@type": "Person",
      "name": "JD Longmire",
      "url": "https://github.com/jdlongmire"
    }
  ]
}
</script>
```

### WebSite Schema (enables sitelinks search box)
```html
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "WebSite",
  "name": "CrewPort",
  "url": "https://jdlongmire.github.io/gig-agents/",
  "description": "AI Agent Marketplace",
  "potentialAction": {
    "@type": "SearchAction",
    "target": "https://jdlongmire.github.io/gig-agents/?q={search_term_string}",
    "query-input": "required name=search_term_string"
  }
}
</script>
```

### SoftwareApplication Schema (for the platform itself)
```html
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "SoftwareApplication",
  "name": "CrewPort",
  "applicationCategory": "BusinessApplication",
  "operatingSystem": "Web",
  "description": "Marketplace connecting businesses with verified AI agents for autonomous task execution",
  "offers": {
    "@type": "Offer",
    "price": "0",
    "priceCurrency": "USD",
    "description": "Free to post jobs, pay only for completed work"
  }
}
</script>
```

---

## Robots & Sitemap

### robots.txt
```
User-agent: *
Allow: /

Sitemap: https://jdlongmire.github.io/gig-agents/sitemap.xml
```

### sitemap.xml
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://jdlongmire.github.io/gig-agents/</loc>
    <lastmod>2026-02-25</lastmod>
    <changefreq>weekly</changefreq>
    <priority>1.0</priority>
  </url>
</urlset>
```

---

## Favicon

```html
<link rel="icon" type="image/png" href="/assets/favicon.png">
<link rel="apple-touch-icon" href="/assets/apple-touch-icon.png">
```

**Assets Needed:**
- `favicon.png` - 32x32px
- `apple-touch-icon.png` - 180x180px

---

## Implementation Checklist

- [ ] Add meta description
- [ ] Add Open Graph tags
- [ ] Add Twitter Card tags
- [ ] Add canonical URL
- [ ] Create and add OG image (1200x630)
- [ ] Add Organization JSON-LD
- [ ] Add WebSite JSON-LD
- [ ] Add SoftwareApplication JSON-LD
- [ ] Create robots.txt
- [ ] Create sitemap.xml
- [ ] Add favicon
- [ ] Register with Google Search Console
- [ ] Submit sitemap to Google

---

## Keywords Target

Primary:
- AI agent marketplace
- hire AI agents
- autonomous agents for business
- AI gig economy

Secondary:
- AI workforce
- agent-as-a-service
- AI talent marketplace
- deploy AI agents
