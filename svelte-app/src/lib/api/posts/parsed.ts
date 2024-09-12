import type {
  ListPostsResponse,
  Post,
  User,
  IncludesData,
  Media,
} from "$lib/api";

export interface ParsedIncludesData {
  posts: Record<string, Post>;
  users: Record<string, User>;
  likedPosts: Record<string, boolean>;
  bookmarkedPosts: Record<string, boolean>;
  media: Record<string, Media>;
}

export interface ParsedUserInteractions {}

export interface ParsedListPostsResponse {
  posts: Post[];
  includes: ParsedIncludesData;
  nextCursor?: string;
}

export function parseListPostsResponse(
  response: ListPostsResponse,
): ParsedListPostsResponse {
  const parsedIncludes: ParsedIncludesData = {
    posts: {},
    users: {},
    likedPosts: {},
    bookmarkedPosts: {},
    media: {},
  };

  if (response.includes) {
    parseIncludes(response.includes, parsedIncludes);
  }

  return {
    posts: response.data,
    includes: parsedIncludes,
    nextCursor: response.next_cursor,
  };
}

function parseIncludes(
  includes: IncludesData,
  parsedIncludes: ParsedIncludesData,
) {
  if (includes.posts) {
    includes.posts.forEach((post) => {
      parsedIncludes.posts[post.id] = post;
    });
  }

  if (includes.users) {
    includes.users.forEach((user) => {
      parsedIncludes.users[user.id] = user;
    });
  }

  if (includes.user_interactions) {
    includes.user_interactions.forEach((interaction) => {
      parsedIncludes.likedPosts[interaction.post_id] = interaction.is_liked;
      parsedIncludes.bookmarkedPosts[interaction.post_id] =
        interaction.is_bookmarked;
    });
  }

  if (includes.media) {
    includes.media.forEach((mediaItem) => {
      parsedIncludes.media[mediaItem.media_key] = mediaItem;
    });
  }
}

export interface ParsedGetPostResponse {
  post: Post;
  includes: ParsedIncludesData;
}

export function parseGetPostResponse(response: any): ParsedGetPostResponse {
  const parsedIncludes: ParsedIncludesData = {
    posts: {},
    users: {},
    likedPosts: {},
    bookmarkedPosts: {},
    media: {},
  };

  if (response.includes) {
    parseIncludes(response.includes, parsedIncludes);
  }

  return {
    post: response.data,
    includes: parsedIncludes,
  };
}
