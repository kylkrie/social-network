<script lang="ts">
  import { useListPosts } from "$lib/queries/posts";
  import PostCard from "../post/PostCard.svelte";
  import { getQuoteForPost, getUserForPost } from "$lib/util";

  export let username: string = undefined;
  const listPostsQuery = useListPosts({ username });

  $: allPosts =
    $listPostsQuery.data?.pages.flatMap((page, pageIndex) =>
      page.data.map((post, postIndex) => ({
        ...post,
        uniqueKey: `${pageIndex}-${postIndex}-${post.id}`,
      })),
    ) ?? [];
  $: users =
    $listPostsQuery.data?.pages.flatMap((page) => page.includes.users) ?? [];
  $: includePosts =
    $listPostsQuery.data?.pages.flatMap((page) => page.includes.posts) ?? [];
  $: isLoading = $listPostsQuery.isLoading;
  $: error = $listPostsQuery.error;

  function handleLoadMore() {
    if ($listPostsQuery.hasNextPage) {
      $listPostsQuery.fetchNextPage();
    }
  }
</script>

{#if isLoading && allPosts.length === 0}
  <p class="text-text-secondary">Loading posts...</p>
{:else if error}
  <p class="text-error">{error}</p>
{:else if allPosts.length === 0}
  <p class="text-text-secondary">No posts yet.</p>
{:else}
  {#each allPosts as post (post.uniqueKey)}
    {@const quote = getQuoteForPost(users, includePosts, post)}
    <PostCard
      user={getUserForPost(users, post)}
      {post}
      quotePost={quote?.post}
      quoteUser={quote?.user}
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
