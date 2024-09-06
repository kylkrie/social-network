// User DTOs
export interface PublicUser {
  id: number;
  name: string;
  username: string;
  pfp_url: string | null;
  protected: boolean;
  profile?: PublicUserProfile;
}

export interface PublicUserProfile {
  banner_url: string | null;
  bio: string | null;
  website: string | null;
  location: string | null;
  birthday: string | null; // ISO 8601 format
  pinned_post_id: number | null;
  follower_count: number | null;
  following_count: number | null;
}

// API Response DTOs
export interface GetUserResponse {
  data: PublicUser;
}

export interface GetCurrentUserResponse {
  data: PublicUser;
}

// Request DTOs
export interface UpdateUserRequest {
  name?: string;
  protected?: boolean;
  bio?: string;
  website?: string;
  location?: string;
  birthday?: string; // ISO 8601 format
  pinned_post_id?: number;
}

// Query Parameter DTOs
export interface GetUserParams {
  profile?: boolean;
}

export interface GetCurrentUserParams {
  profile?: boolean;
}
