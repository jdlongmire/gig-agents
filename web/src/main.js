import './lib/theme.css';
import App from './App.svelte';

// Prevent flash of wrong theme — set data-theme BEFORE Svelte mounts.
// Reads localStorage preference, defaults to light.
const saved = localStorage.getItem('crewport-theme');
document.documentElement.setAttribute('data-theme', saved || 'light');

const app = new App({
  target: document.getElementById('app'),
});

export default app;
