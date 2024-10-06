const API_BASE_URL = import.meta.env.VITE_API_URL;

type RequestOptions = {
  headers?: Record<string, string>;
  body?: any;
};

async function request(
  endpoint: string,
  method: string,
  options: RequestOptions = {},
): Promise<any> {
  const url = `${API_BASE_URL}/public/v1${endpoint}`;
  const headers = new Headers(options.headers);

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

export const publicApi = {
  get: (endpoint: string) => request(endpoint, "GET"),
  post: (endpoint: string, body?: any, options: RequestOptions = {}) =>
    request(endpoint, "POST", { ...options, body }),
  put: (endpoint: string, body?: any) => request(endpoint, "PUT", { body }),
  delete: (endpoint: string) => request(endpoint, "DELETE"),
};
