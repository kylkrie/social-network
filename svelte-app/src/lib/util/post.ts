import type { IncludesData, Post, User } from "$lib/api";

export function getUserForPost(users: User[], post: Post): User | undefined {
  return users.find((user) => user.id === post.author_id);
}
export function getQuoteForPost(
  users: User[],
  posts: Post[],
  post: Post,
): { user: User; post: Post } {
  const ref = post.references?.find((r) => r.reference_type === "quote");
  if (!ref) {
    return undefined;
  }
  const quotePost = posts.find((p) => p.id === ref.referenced_post_id);
  const quoteUser = getUserForPost(users, quotePost);

  return { user: quoteUser, post: quotePost };
}
