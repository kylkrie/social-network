import type { IncludesData } from "..";

// Post DTOs
export interface Post {
  id: string;
  content: string;
  author_id: string;
  conversation_id?: string;
  created_at: string; // ISO 8601 format
  updated_at: string; // ISO 8601 format
  deleted_at?: string; // ISO 8601 format
  public_metrics?: PostPublicMetrics;
  edits?: PostEdit[];
  tags?: PostTag[];
  references?: PostReference[];
}

export interface PostPublicMetrics {
  reposts: number;
  replies: number;
  likes: number;
  views: number;
}

export interface PostEdit {
  content: string;
  edited_at: string; // ISO 8601 format
}

export interface PostTag {
  entity_type: string;
  start_index?: number;
  end_index?: number;
  tag?: string;
}

export interface PostReference {
  referenced_post_id: string;
  reference_type: string;
}

// API Response DTOs
export interface GetPostResponse {
  data: Post;
  includes?: IncludesData;
}

export interface ListPostsResponse {
  data: Post[];
  includes?: IncludesData;
  next_cursor?: string;
}

// Request DTOs
export interface CreatePostRequest {
  content: string;
  reply_to_post_id?: string;
  quote_post_id?: string;
}

export interface UpdatePostRequest {
  content: string;
}

// Query Parameter DTOs
export interface ListPostsParams {
  limit?: number;
  cursor?: string;
  replies?: boolean;
  conversation_id?: string;
}

export interface GetPostParams {
  include_metrics?: boolean;
}
