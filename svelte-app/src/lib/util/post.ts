import type { ParsedIncludesData, Post, User } from "$lib/api";

export function getReplyForPost(
  post: Post,
  includes: ParsedIncludesData,
): { user: User; post: Post } {
  const ref = post.references?.find((r) => r.reference_type === "reply_to");
  if (!ref) {
    return undefined;
  }
  const replyPost = includes.posts[ref.referenced_post_id];
  const replyUser = includes.users[replyPost.author_id];

  return { user: replyUser, post: replyPost };
}

export function getQuoteForPost(
  post: Post,
  includes: ParsedIncludesData,
): { user: User; post: Post } {
  const ref = post.references?.find((r) => r.reference_type === "quote");
  if (!ref) {
    return undefined;
  }
  const quotePost = includes.posts[ref.referenced_post_id];
  const quoteUser = includes.users[quotePost.author_id];

  return { user: quoteUser, post: quotePost };
}
