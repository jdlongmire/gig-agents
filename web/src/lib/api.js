/**
 * CrewPort API wrapper
 *
 * Base URL = window.location.origin (same-origin, Go server proxies all /api/* routes)
 * Auth = httpOnly cookie (set by Go OAuth callback), plus Authorization: Bearer header
 *        if a token is available in the auth store.
 * 401  = clears auth store and redirects to login.
 */

import { auth } from './auth.js';
import { get as storeGet } from 'svelte/store';

const BASE = typeof window !== 'undefined' ? window.location.origin : '';

/**
 * Core fetch wrapper.
 * @param {string} path  — e.g. '/api/auth/me'
 * @param {object} opts  — fetch options
 */
async function request(path, opts = {}) {
  const authState = storeGet(auth);

  const headers = {
    'Content-Type': 'application/json',
    ...(opts.headers || {}),
  };

  // Attach Bearer token if we have one (fallback for non-cookie flows)
  if (authState.token && authState.token !== '__cookie__') {
    headers['Authorization'] = `Bearer ${authState.token}`;
  }

  const response = await fetch(`${BASE}${path}`, {
    ...opts,
    credentials: 'include', // always send cookies
    headers,
  });

  if (response.status === 401) {
    auth.clearAuth();
    if (typeof window !== 'undefined') {
      window.location.hash = '#/';
    }
    throw new ApiError(401, 'Unauthorized');
  }

  if (!response.ok) {
    let message = `HTTP ${response.status}`;
    try {
      const body = await response.json();
      message = body.error || body.message || message;
    } catch {
      // non-JSON error body — keep default message
    }
    throw new ApiError(response.status, message);
  }

  // 204 No Content
  if (response.status === 204) return null;

  return response.json();
}

class ApiError extends Error {
  constructor(status, message) {
    super(message);
    this.status = status;
    this.name = 'ApiError';
  }
}

// ── Exported helpers ──────────────────────────────────────

export function get(path) {
  return request(path, { method: 'GET' });
}

export function post(path, body) {
  return request(path, { method: 'POST', body: JSON.stringify(body) });
}

export function put(path, body) {
  return request(path, { method: 'PUT', body: JSON.stringify(body) });
}

export function del(path) {
  return request(path, { method: 'DELETE' });
}

export { ApiError };
