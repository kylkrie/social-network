// User DTOs
export interface User {
  id: string;
  name: string;
  username: string;
  pfp_url?: string;
  protected: boolean;
  profile?: UserProfile;
}

export interface UserProfile {
  banner_url?: string;
  bio?: string;
  website?: string;
  location?: string;
  birthday?: string; // ISO 8601 format
  pinned_post_id?: string;
  follower_count?: number;
  following_count?: number;
}

// API Response DTOs
export interface GetUserResponse {
  data: User;
}

export interface GetCurrentUserResponse {
  data: User;
}

// Request DTOs
export interface UpdateUserRequest {
  name?: string;
  protected?: boolean;
  bio?: string;
  website?: string;
  location?: string;
  birthday?: string; // ISO 8601 format
  pinned_post_id?: string;
}

// Query Parameter DTOs
export interface GetUserParams {
  profile?: boolean;
}

export interface GetCurrentUserParams {
  profile?: boolean;
}
