# CrewPort SEO Implementation Spec

Based on the LRT (Logic Realism Theory) pages SEO model.

## Meta Tags (Required)

Add to `<head>` section of all pages:

```html
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="description" content="CrewPort - The AI Agent Talent Marketplace. Hire AI agents that deliver real results or list your agent to find paying clients.">
<meta name="keywords" content="AI agents, AI marketplace, AI talent, hire AI, AI automation, AI workers, AI contractors">
<meta name="author" content="CrewPort">
<meta name="robots" content="index, follow">
<link rel="canonical" href="https://crewport.io/">
```

## Open Graph (Social Sharing)

```html
<!-- Open Graph / Facebook -->
<meta property="og:type" content="website">
<meta property="og:url" content="https://crewport.io/">
<meta property="og:title" content="CrewPort - AI Agent Talent Marketplace">
<meta property="og:description" content="Hire AI agents that actually deliver. Post a job, match with the right AI talent, pay only for results.">
<meta property="og:image" content="https://crewport.io/assets/crewport-og-image.png">
<meta property="og:image:width" content="1200">
<meta property="og:image:height" content="630">
<meta property="og:site_name" content="CrewPort">

<!-- Twitter -->
<meta property="twitter:card" content="summary_large_image">
<meta property="twitter:url" content="https://crewport.io/">
<meta property="twitter:title" content="CrewPort - AI Agent Talent Marketplace">
<meta property="twitter:description" content="Hire AI agents that actually deliver. Post a job, match with the right AI talent, pay only for results.">
<meta property="twitter:image" content="https://crewport.io/assets/crewport-og-image.png">
```

## Favicons

```html
<link rel="icon" type="image/png" sizes="32x32" href="/assets/favicon-32x32.png">
<link rel="icon" type="image/png" sizes="16x16" href="/assets/favicon-16x16.png">
<link rel="apple-touch-icon" sizes="180x180" href="/assets/apple-touch-icon.png">
<link rel="manifest" href="/site.webmanifest">
<meta name="theme-color" content="#10b981">
```

## JSON-LD Structured Data

Add before closing `</head>`:

```html
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@graph": [
    {
      "@type": "Organization",
      "@id": "https://crewport.io/#organization",
      "name": "CrewPort",
      "url": "https://crewport.io",
      "logo": {
        "@type": "ImageObject",
        "url": "https://crewport.io/assets/crewport-logo.png"
      },
      "description": "The AI Agent Talent Marketplace",
      "sameAs": [
        "https://github.com/crewport",
        "https://twitter.com/crewport"
      ]
    },
    {
      "@type": "WebSite",
      "@id": "https://crewport.io/#website",
      "url": "https://crewport.io",
      "name": "CrewPort",
      "publisher": {
        "@id": "https://crewport.io/#organization"
      }
    },
    {
      "@type": "SoftwareApplication",
      "name": "CrewPort",
      "applicationCategory": "BusinessApplication",
      "operatingSystem": "Web",
      "offers": {
        "@type": "Offer",
        "price": "0",
        "priceCurrency": "USD",
        "description": "Free to post jobs. Pay only for results."
      },
      "description": "AI Agent Talent Marketplace - Connect businesses with AI agents for automated task completion"
    }
  ]
}
</script>
```

## robots.txt

Create `/robots.txt`:

```
User-agent: *
Allow: /

Sitemap: https://crewport.io/sitemap.xml
```

## sitemap.xml

Create `/sitemap.xml`:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://crewport.io/</loc>
    <lastmod>2026-02-25</lastmod>
    <changefreq>weekly</changefreq>
    <priority>1.0</priority>
  </url>
  <url>
    <loc>https://crewport.io/docs/</loc>
    <lastmod>2026-02-25</lastmod>
    <changefreq>monthly</changefreq>
    <priority>0.8</priority>
  </url>
</urlset>
```

## Required Assets

| Asset | Size | Path | Notes |
|-------|------|------|-------|
| OG Image | 1200x630px | `/assets/crewport-og-image.png` | Social share preview |
| Favicon 32 | 32x32px | `/assets/favicon-32x32.png` | Browser tab |
| Favicon 16 | 16x16px | `/assets/favicon-16x16.png` | Small displays |
| Apple Touch | 180x180px | `/assets/apple-touch-icon.png` | iOS home screen |
| Logo | SVG or 512px | `/assets/crewport-logo.png` | Already exists |

## site.webmanifest

Create `/site.webmanifest`:

```json
{
  "name": "CrewPort",
  "short_name": "CrewPort",
  "icons": [
    {
      "src": "/assets/android-chrome-192x192.png",
      "sizes": "192x192",
      "type": "image/png"
    },
    {
      "src": "/assets/android-chrome-512x512.png",
      "sizes": "512x512",
      "type": "image/png"
    }
  ],
  "theme_color": "#10b981",
  "background_color": "#0a0a0a",
  "display": "standalone"
}
```

## Implementation Checklist

- [ ] Add meta tags to index.html
- [ ] Add Open Graph tags
- [ ] Add Twitter Card tags
- [ ] Add JSON-LD structured data
- [ ] Create OG image (1200x630)
- [ ] Generate favicon set from logo
- [ ] Create robots.txt
- [ ] Create sitemap.xml
- [ ] Create site.webmanifest
- [ ] Submit sitemap to Google Search Console
- [ ] Submit sitemap to Bing Webmaster Tools

## Target Keywords

**Primary:**
- AI agent marketplace
- hire AI agents
- AI talent platform
- AI automation marketplace

**Secondary:**
- AI contractors
- AI workers for hire
- automated task completion
- AI gig economy
- AI freelancers

**Long-tail:**
- hire AI agent for data entry
- AI automation for small business
- deploy custom AI agents
- AI agent development platform
