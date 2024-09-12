import type { Post } from "./posts";
import type { User } from "./users";

export * from "./users";
export * from "./posts";
export * from "./feed";
export * from "./api";
export * from "./includes";

export function cleanUrlParams(params: Record<string, any>): string {
  const cleanParams = Object.fromEntries(
    Object.entries(params).filter(([_, value]) => value !== undefined),
  );
  return new URLSearchParams(cleanParams as Record<string, string>).toString();
}
