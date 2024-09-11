import { get } from "svelte/store";
import { auth, type AuthToken } from "../stores";

const API_BASE_URL = import.meta.env.VITE_API_URL;
const REFRESH_THRESHOLD = 10000; // 10 seconds

async function getValidToken(): Promise<string | null> {
  const token = get(auth) as AuthToken;
  if (!token) return null;

  const now = Date.now();
  if (token.expiresAt - now <= REFRESH_THRESHOLD) {
    const refreshSuccess = await auth.refresh();
    if (!refreshSuccess) {
      auth.logout();
      throw new Error("Session expired. Please log in again.");
    }
    return get(auth)?.accessToken ?? null;
  }
  return token.accessToken;
}

type RequestOptions = {
  headers?: Record<string, string>;
  body?: any;
};

async function request(
  endpoint: string,
  method: string,
  options: RequestOptions = {},
): Promise<any> {
  const url = `${API_BASE_URL}/api/v1${endpoint}`;
  const headers = new Headers(options.headers);

  const accessToken = await getValidToken();
  if (accessToken) {
    headers.set("Authorization", `Bearer ${accessToken}`);
  }

  let body: string | FormData | undefined;
  if (options.body) {
    if (options.body instanceof FormData) {
      body = options.body;
    } else if (typeof options.body === "object") {
      headers.set("Content-Type", "application/json");
      body = JSON.stringify(options.body);
    }
  }

  const config: RequestInit = {
    method,
    headers,
    body,
  };

  try {
    let response = await fetch(url, config);

    if (response.status === 401 && accessToken) {
      const refreshSuccess = await auth.refresh();
      if (refreshSuccess) {
        const newToken = get(auth);
        headers.set("Authorization", `Bearer ${newToken?.accessToken}`);
        config.headers = headers;
        response = await fetch(url, config);
      } else {
        auth.logout();
        throw new Error("Session expired. Please log in again.");
      }
    }

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const contentType = response.headers.get("content-type");
    if (contentType && contentType.includes("application/json")) {
      return await response.json();
    } else {
      return {
        status: response.status,
        statusText: response.statusText,
      };
    }
  } catch (error) {
    console.error("API request failed:", error);
    throw error;
  }
}

export const api = {
  get: (endpoint: string) => request(endpoint, "GET"),
  post: (endpoint: string, body?: any, options: RequestOptions = {}) =>
    request(endpoint, "POST", { ...options, body }),
  put: (endpoint: string, body?: any) => request(endpoint, "PUT", { body }),
  delete: (endpoint: string) => request(endpoint, "DELETE"),
};
