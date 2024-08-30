import { get } from "svelte/store";
import { auth } from "./auth";

const API_BASE_URL = import.meta.env.VITE_API_URL;

interface RequestOptions {
  method: "GET" | "POST" | "PUT" | "DELETE";
  headers?: Record<string, string>;
  body?: any;
}

async function request(
  endpoint: string,
  options: RequestOptions,
): Promise<any> {
  const token = get(auth);
  const url = `${API_BASE_URL}${endpoint}`;

  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...options.headers,
  };

  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const config: RequestInit = {
    method: options.method,
    headers,
  };

  if (options.body) {
    config.body = JSON.stringify(options.body);
  }

  try {
    const response = await fetch(url, config);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return await response.json();
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

export async function getUserInfo() {
  return api.get("/userinfo");
}
