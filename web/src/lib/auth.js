import { writable } from 'svelte/store';

// Auth store — { user, token, isAuthenticated, loading }
function createAuthStore() {
  const { subscribe, set, update } = writable({
    user: null,
    token: null,
    isAuthenticated: false,
    loading: true,
  });

  return {
    subscribe,

    /**
     * Initialize auth state on app mount.
     * JWT lives in an httpOnly cookie so JS can't read it directly.
     * We validate by hitting /api/auth/me — 200 = authed, 401 = not.
     */
    async init() {
      try {
        const response = await fetch('/api/auth/me', {
          credentials: 'include',
          headers: { 'Content-Type': 'application/json' },
        });
        if (response.ok) {
          const me = await response.json();
          set({
            user: me,
            token: '__cookie__',
            isAuthenticated: true,
            loading: false,
          });
        } else {
          set({ user: null, token: null, isAuthenticated: false, loading: false });
        }
      } catch {
        set({ user: null, token: null, isAuthenticated: false, loading: false });
      }
    },

    setUser(user) {
      update(s => ({ ...s, user, isAuthenticated: true, loading: false }));
    },

    async logout() {
      try {
        await fetch('/api/auth/logout', { method: 'POST', credentials: 'include' });
      } catch {
        // clear client state regardless
      }
      set({ user: null, token: null, isAuthenticated: false, loading: false });
    },

    clearAuth() {
      set({ user: null, token: null, isAuthenticated: false, loading: false });
    },
  };
}

export const auth = createAuthStore();
