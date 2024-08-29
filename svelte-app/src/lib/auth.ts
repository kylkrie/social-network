import { goto } from "$app/navigation";

const CLIENT_ID = import.meta.env.VITE_CLIENT_ID;
const BASE_URL = import.meta.env.VITE_BASE_URL;
const REDIRECT_URI = `${BASE_URL}/callback`;
const OAUTH_BASE_URL = import.meta.env.VITE_OAUTH_BASE_URL;

interface AuthState {
  isAuthenticated: boolean;
  token: string | null;
}

function getStoredAuthState(): AuthState {
  const storedState = localStorage.getItem("authState");
  if (storedState) {
    return JSON.parse(storedState);
  }
  return { isAuthenticated: false, token: null };
}

function setStoredAuthState(state: AuthState): void {
  localStorage.setItem("authState", JSON.stringify(state));
}

export function getAuthState(): AuthState {
  return getStoredAuthState();
}

function setToken(token: string): void {
  const newState = { isAuthenticated: true, token };
  setStoredAuthState(newState);
}

function clearToken(): void {
  const newState = { isAuthenticated: false, token: null };
  setStoredAuthState(newState);
}

function generateState(): string {
  return (
    Math.random().toString(36).substring(2, 15) +
    Math.random().toString(36).substring(2, 15)
  );
}

export function logout(): void {
  clearToken();
  goto("/");
}

export function startAuthLogin(): void {
  const state = generateState();
  sessionStorage.setItem("oauth_state", state);

  const authUrl = new URL(`${OAUTH_BASE_URL}/auth`);
  authUrl.searchParams.append("client_id", CLIENT_ID);
  authUrl.searchParams.append("redirect_uri", REDIRECT_URI);
  authUrl.searchParams.append("response_type", "code");
  authUrl.searchParams.append("state", state);

  window.location.href = authUrl.toString();
}

export function startAuthRegister(): void {
  const state = generateState();
  sessionStorage.setItem("oauth_state", state);

  const authUrl = new URL(`${OAUTH_BASE_URL}/registrations`);
  authUrl.searchParams.append("client_id", CLIENT_ID);
  authUrl.searchParams.append("redirect_uri", REDIRECT_URI);
  authUrl.searchParams.append("response_type", "code");
  authUrl.searchParams.append("state", state);

  window.location.href = authUrl.toString();
}

export async function handleCallback(
  code: string,
  state: string,
): Promise<boolean> {
  const storedState = sessionStorage.getItem("oauth_state");
  sessionStorage.removeItem("oauth_state");

  if (!state || state !== storedState) {
    return false;
  }

  try {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/token`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ code }),
    });

    if (response.ok) {
      const data = await response.json();
      setToken(data.access_token);
      return true;
    }
  } catch (error) {
    console.error("Error during token exchange:", error);
  }

  return false;
}
