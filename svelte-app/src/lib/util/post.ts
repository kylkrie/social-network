import type {
  ParsedIncludesData,
  ParsedUserInteractions,
  Post,
} from "$lib/api";
import type { PostData } from "$lib/components/post/PostCard.svelte";

export function getReplyForPost(
  post: Post,
  includes: ParsedIncludesData,
): PostData {
  const ref = post.references?.find((r) => r.reference_type === "reply_to");
  if (!ref) {
    return undefined;
  }

  const replyPost = includes.posts[ref.referenced_post_id];
  return buildPostData(replyPost, includes);
}

export function getQuoteForPost(
  post: Post,
  includes: ParsedIncludesData,
): PostData {
  const ref = post.references?.find((r) => r.reference_type === "quote");
  if (!ref) {
    return undefined;
  }

  const quotePost = includes.posts[ref.referenced_post_id];
  return buildPostData(quotePost, includes);
}

export function buildPostData(
  post: Post,
  includes: ParsedIncludesData,
): PostData {
  if (!post) {
    return undefined;
  }

  return {
    post: post,
    user: includes.users[post.author_id],
    is_liked: includes.userInteractions.likedPosts[post.id],
    is_bookmarked: includes.userInteractions.bookmarkedPosts[post.id],
  };
}
