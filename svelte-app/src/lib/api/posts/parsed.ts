import type { ListPostsResponse, Post, User, IncludesData } from "$lib/api";

export interface ParsedIncludesData {
  posts: Record<string, Post>;
  users: Record<string, User>;
}
export interface ParsedListPostsResponse {
  posts: Post[];
  includes: ParsedIncludesData;
  nextCursor?: string;
}

export function parseListPostsResponse(
  response: ListPostsResponse,
): ParsedListPostsResponse {
  const includedPosts: Record<string, Post> = {};
  const users: Record<string, User> = {};

  if (response.includes) {
    parseIncludes(response.includes, includedPosts, users);
  }

  return {
    posts: response.data,
    includes: {
      posts: includedPosts,
      users,
    },
    nextCursor: response.next_cursor,
  };
}

function parseIncludes(
  includes: IncludesData,
  posts: Record<string, Post>,
  users: Record<string, User>,
) {
  if (includes.posts) {
    includes.posts.forEach((post) => {
      posts[post.id] = post;
    });
  }

  if (includes.users) {
    includes.users.forEach((user) => {
      users[user.id] = user;
    });
  }
}

export interface ParsedGetPostResponse {
  post: Post;
  includes: ParsedIncludesData;
}

export function parseGetPostResponse(response: any): ParsedGetPostResponse {
  const users: Record<string, User> = {};
  const includedPosts: Record<string, Post> = {};

  if (response.includes) {
    parseIncludes(response.includes, includedPosts, users);
  }

  return {
    post: response.data,
    includes: {
      posts: includedPosts,
      users,
    },
  };
}
