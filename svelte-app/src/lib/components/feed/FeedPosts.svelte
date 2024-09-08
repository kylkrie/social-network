<script lang="ts">
  import Card from "$lib/components/ui/Card.svelte";
  import PostCard from "../post/PostCard.svelte";
  import type { Post } from "$lib/api/posts/dtos";
  import type { User } from "$lib/api/users/dtos";
  import { useListFeed } from "$lib/queries/posts/feed";

  const listPostsQuery = useListFeed();

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
  $: includePosts;
  $: isLoading = $listPostsQuery.isLoading;
  $: error = $listPostsQuery.error;

  function handleLoadMore() {
    if ($listPostsQuery.hasNextPage) {
      $listPostsQuery.fetchNextPage();
    }
  }

  function getUserForPost(post: Post): User | undefined {
    return users.find((user) => user.id === post.author_id);
  }
  function getQuoteForPost(post: Post): { user: User; post: Post } {
    const ref = post.references?.find((r) => r.reference_type === "quote");
    if (!ref) {
      return undefined;
    }
    const quotePost = includePosts.find((p) => p.id === ref.referenced_post_id);
    const quoteUser = getUserForPost(quotePost);

    return { user: quoteUser, post: quotePost };
  }
</script>

{#if isLoading && allPosts.length === 0}
  <Card>
    <p class="text-text-secondary">Loading posts...</p>
  </Card>
{:else if error}
  <Card>
    <p class="text-error">{error}</p>
  </Card>
{:else if allPosts.length === 0}
  <Card>
    <p class="text-text-secondary">No posts yet.</p>
  </Card>
{:else}
  {#each allPosts as post (post.uniqueKey)}
    {@const quote = getQuoteForPost(post)}
    <PostCard
      user={getUserForPost(post)}
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
  <Card>
    <p class="text-text-secondary">Loading more posts...</p>
  </Card>
{/if}
