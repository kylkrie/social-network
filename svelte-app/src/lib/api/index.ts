import type { Post } from "./posts";
import type { User } from "./users";

export * from "./users";
export * from "./posts";
export * from "./feed";
export * from "./api";

export interface IncludesData {
  users?: User[];
  posts?: Post[];
  user_interactions?: UserPostInteraction[];
}

export interface UserPostInteraction {
  post_id: string;
  is_liked: boolean;
  is_bookmarked: boolean;
}

export function cleanUrlParams(params: Record<string, any>): string {
  const cleanParams = Object.fromEntries(
    Object.entries(params).filter(([_, value]) => value !== undefined),
  );
  return new URLSearchParams(cleanParams as Record<string, string>).toString();
}
