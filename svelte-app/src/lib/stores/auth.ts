import { writable, get } from 'svelte/store';

const CLIENT_ID = import.meta.env.VITE_CLIENT_ID;
const BASE_URL = import.meta.env.VITE_BASE_URL;
const REDIRECT_URI = `${BASE_URL}/callback`;
const OAUTH_BASE_URL = import.meta.env.VITE_OAUTH_BASE_URL;
const API_BASE_URL = import.meta.env.VITE_API_URL;

export type AuthToken = {
  accessToken: string;
  refreshToken: string;
  expiresAt: number;
} | null;

function createAuthStore() {
  const { subscribe, set, update } = writable<AuthToken>(null);

  return {
    subscribe,
    login: (token: AuthToken) => {
      localStorage.setItem('authToken', JSON.stringify(token));
      set(token);
    },
    logout: () => {
      localStorage.removeItem('authToken');
      set(null);
    },
    init: () => {
      const storedToken = localStorage.getItem('authToken');
      if (storedToken) {
        set(JSON.parse(storedToken));
      }
    },
    refresh: async () => {
      const currentToken = get({ subscribe });
      if (!currentToken?.refreshToken) return false;

      try {
        const response = await fetch(`${API_BASE_URL}/auth/v1/token`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ refresh_token: currentToken.refreshToken }),
        });

        if (response.ok) {
          const data = await response.json();
          const newToken: AuthToken = {
            accessToken: data.access_token,
            refreshToken: data.refresh_token,
            expiresAt: new Date(data.expiry).getTime(),
          };
          localStorage.setItem('authToken', JSON.stringify(newToken));
          set(newToken);
          return true;
        }
      } catch (error) {
        console.error('Error refreshing token:', error);
      }
      return false;
    },
  };
}

export const auth = createAuthStore();

function generateRandomString(length: number): string {
  const charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~';
  return Array.from(crypto.getRandomValues(new Uint8Array(length)))
    .map(x => charset[x % charset.length])
    .join('');
}

async function generateCodeChallenge(verifier: string): Promise<string> {
  const encoder = new TextEncoder();
  const data = encoder.encode(verifier);
  const hash = await crypto.subtle.digest('SHA-256', data);
  return btoa(String.fromCharCode(...new Uint8Array(hash)))
    .replace(/=/g, '')
    .replace(/\+/g, '-')
    .replace(/\//g, '_');
}

async function startAuth(isRegister: boolean): Promise<void> {
  const state = generateRandomString(32);
  const codeVerifier = generateRandomString(64);
  const codeChallenge = await generateCodeChallenge(codeVerifier);

  sessionStorage.setItem('oauth_state', state);
  sessionStorage.setItem('code_verifier', codeVerifier);

  const authUrl = new URL(`${OAUTH_BASE_URL}/${isRegister ? 'registrations' : 'auth'}`);
  authUrl.searchParams.append('client_id', CLIENT_ID);
  authUrl.searchParams.append('redirect_uri', REDIRECT_URI);
  authUrl.searchParams.append('response_type', 'code');
  authUrl.searchParams.append('state', state);
  authUrl.searchParams.append('code_challenge', codeChallenge);
  authUrl.searchParams.append('code_challenge_method', 'S256');

  window.location.href = authUrl.toString();
}

export const startAuthLogin = () => startAuth(false);
export const startAuthRegister = () => startAuth(true);

export async function handleCallback(code: string, state: string): Promise<boolean> {
  const storedState = sessionStorage.getItem('oauth_state');
  const codeVerifier = sessionStorage.getItem('code_verifier');

  sessionStorage.removeItem('oauth_state');
  sessionStorage.removeItem('code_verifier');

  if (!state || state !== storedState || !codeVerifier) {
    console.error('Invalid state or missing code verifier');
    return false;
  }

  try {
    const response = await fetch(`${API_BASE_URL}/auth/v1/token`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ code, code_verifier: codeVerifier }),
    });

    if (response.ok) {
      const data = await response.json();
      const token: AuthToken = {
        accessToken: data.access_token,
        refreshToken: data.refresh_token,
        expiresAt: new Date(data.expiry).getTime(),
      };
      auth.login(token);
      return true;
    } else {
      console.error('Token exchange failed:', await response.json());
      return false;
    }
  } catch (error) {
    console.error('Error during token exchange:', error);
    return false;
  }
}

// Initialize auth store
auth.init();
