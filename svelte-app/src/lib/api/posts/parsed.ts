import type {
  ListPostsResponse,
  Post,
  User,
  IncludesData,
  UserPostInteraction,
} from "$lib/api";

export interface ParsedIncludesData {
  posts: Record<string, Post>;
  users: Record<string, User>;
  userInteractions: ParsedUserInteractions;
}

export interface ParsedUserInteractions {
  likedPosts: Record<string, boolean>;
  bookmarkedPosts: Record<string, boolean>;
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
  const userInteractions: ParsedUserInteractions = {
    likedPosts: {},
    bookmarkedPosts: {},
  };

  if (response.includes) {
    parseIncludes(response.includes, includedPosts, users, userInteractions);
  }

  return {
    posts: response.data,
    includes: {
      posts: includedPosts,
      users,
      userInteractions,
    },
    nextCursor: response.next_cursor,
  };
}

function parseIncludes(
  includes: IncludesData,
  posts: Record<string, Post>,
  users: Record<string, User>,
  userInteractions: ParsedUserInteractions,
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

  if (includes.user_interactions) {
    includes.user_interactions.forEach((interaction) => {
      userInteractions.likedPosts[interaction.post_id] = interaction.is_liked;
      userInteractions.bookmarkedPosts[interaction.post_id] =
        interaction.is_bookmarked;
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
  const userInteractions: ParsedUserInteractions = {
    likedPosts: {},
    bookmarkedPosts: {},
  };

  if (response.includes) {
    parseIncludes(response.includes, includedPosts, users, userInteractions);
  }

  return {
    post: response.data,
    includes: {
      posts: includedPosts,
      users,
      userInteractions,
    },
  };
}
