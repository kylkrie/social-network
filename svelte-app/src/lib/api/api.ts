import { get } from "svelte/store";
import { auth } from "../auth";

const API_BASE_URL = import.meta.env.VITE_API_URL;
const REFRESH_THRESHOLD = 10000; // 10 seconds

interface RequestOptions {
  method: "GET" | "POST" | "PUT" | "DELETE";
  headers?: Record<string, string>;
  body?: any;
}

let refreshPromise: Promise<boolean> | null = null;

async function refreshToken(): Promise<boolean> {
  if (!refreshPromise) {
    refreshPromise = auth.refresh().finally(() => {
      refreshPromise = null;
    });
  }
  return refreshPromise;
}

async function getValidToken(): Promise<string | null> {
  if (refreshPromise) {
    await refreshPromise;
  }
  let token = get(auth);
  if (token) {
    const now = Date.now();
    const timeUntilExpiry = token.expiresAt - now;
    if (timeUntilExpiry <= REFRESH_THRESHOLD) {
      const refreshSuccess = await refreshToken();
      if (refreshSuccess) {
        token = get(auth);
        if (!token) {
          throw new Error("Token not found after refresh");
        }
      } else {
        auth.logout();
        throw new Error("Session expired. Please log in again.");
      }
    }
    return token.accessToken;
  }
  return null;
}

async function request(
  endpoint: string,
  options: RequestOptions,
): Promise<any> {
  const url = `${API_BASE_URL}/api${endpoint}`;
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...options.headers,
  };

  const accessToken = await getValidToken();
  if (accessToken) {
    headers["Authorization"] = `Bearer ${accessToken}`;
  }

  const config: RequestInit = {
    method: options.method,
    headers,
  };

  if (options.body) {
    config.body = JSON.stringify(options.body);
  }

  try {
    let response = await fetch(url, config);
    if (response.status === 401 && accessToken) {
      // Token might be expired, try to refresh
      const refreshSuccess = await refreshToken();
      if (refreshSuccess) {
        // Retry the request with the new token
        const newToken = get(auth);
        headers["Authorization"] = `Bearer ${newToken?.accessToken}`;
        config.headers = headers;
        response = await fetch(url, config);
      } else {
        // Refresh failed, logout the user
        auth.logout();
        throw new Error("Session expired. Please log in again.");
      }
    }
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    // Check if the response has content
    const contentType = response.headers.get("content-type");
    if (contentType && contentType.includes("application/json")) {
      return await response.json();
    } else {
      // For non-JSON responses, return an object with status and statusText
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
  get: (endpoint: string) => request(endpoint, { method: "GET" }),
  post: (endpoint: string, body: any) =>
    request(endpoint, { method: "POST", body }),
  put: (endpoint: string, body: any) =>
    request(endpoint, { method: "PUT", body }),
  delete: (endpoint: string) => request(endpoint, { method: "DELETE" }),
};
