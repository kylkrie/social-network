<script lang="ts">
  import { useListPosts } from "$lib/queries/posts";
  import PostCard from "../post/PostCard.svelte";
  import type { Post } from "$lib/api/posts/dtos";
  import type { User } from "$lib/api/users/dtos";
  import { getQuoteForPost, getUserForPost } from "$lib/util";

  export let username: string = undefined;
  const listPostsQuery = useListPosts({ username, replies: true });

  $: allReplies =
    $listPostsQuery.data?.pages.flatMap((page, pageIndex) =>
      page.data.map((post, postIndex) => ({
        ...post,
        uniqueKey: `${pageIndex}-${postIndex}-${post.id}`,
      })),
    ) ?? [];
  $: includePosts =
    $listPostsQuery.data?.pages.flatMap((page) => page.includes.posts) ?? [];
  $: users =
    $listPostsQuery.data?.pages.flatMap((page) => page.includes.users) ?? [];
  $: isLoading = $listPostsQuery.isLoading;
  $: error = $listPostsQuery.error;

  function handleLoadMore() {
    if ($listPostsQuery.hasNextPage) {
      $listPostsQuery.fetchNextPage();
    }
  }

  function getReplySourceForDest(dest: Post): Post | undefined {
    const sourceId = dest.references.find(
      (p) => p.reference_type === "reply_to",
    ).referenced_post_id;
    return includePosts.find((s) => s.id === sourceId);
  }
</script>

{#if isLoading && allReplies.length === 0}
  <p class="text-text-secondary">Loading posts...</p>
{:else if error}
  <p class="text-error">{error}</p>
{:else if allReplies.length === 0}
  <p class="text-text-secondary">No posts yet.</p>
{:else}
  {#each allReplies as dest (dest.uniqueKey)}
    {@const source = getReplySourceForDest(dest)}
    {@const sourceQuote = getQuoteForPost(users, includePosts, source)}
    {@const destQuote = getQuoteForPost(users, includePosts, dest)}
    <PostCard
      user={getUserForPost(users, source)}
      post={source}
      quoteUser={sourceQuote?.user}
      quotePost={sourceQuote?.post}
      variant="reply_source"
    />
    <PostCard
      user={getUserForPost(users, dest)}
      post={dest}
      quoteUser={destQuote?.user}
      quotePost={destQuote?.post}
      variant="reply_dest"
    />
  {/each}
  {#if $listPostsQuery.hasNextPage}
    <button
      on:click={handleLoadMore}
      class="w-full p-2 bg-primary text-white rounded mt-4"
    >
      {$listPostsQuery.isFetchingNextPage ? "Loading more..." : "Load More"}
    </button>
  {/if}
{/if}

{#if $listPostsQuery.isFetchingNextPage}
  <p class="text-text-secondary">Loading more posts...</p>
{/if}
