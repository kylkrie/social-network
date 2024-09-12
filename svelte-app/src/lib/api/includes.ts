import type { Post } from "./posts";
import type { User } from "./users";

export interface IncludesData {
  users?: User[];
  posts?: Post[];
  user_interactions?: UserPostInteraction[];
  media?: Media[];
}

export interface UserPostInteraction {
  post_id: string;
  is_liked: boolean;
  is_bookmarked: boolean;
}

export interface Media {
  media_key: string;
  type: "photo" | "video" | "animated_gif";
  url: string;
  width: number;
  height: number;
}
