import './lib/theme.css';
import App from './App.svelte';

// Prevent flash of wrong theme — set data-theme BEFORE Svelte mounts.
// Reads localStorage preference, falls back to system preference, defaults to dark.
const saved = localStorage.getItem('crewport-theme');
const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
document.documentElement.setAttribute('data-theme', saved || (prefersDark ? 'dark' : 'light'));

const app = new App({
  target: document.getElementById('app'),
});

export default app;
