<script lang="ts">
  import Card from "$lib/components/ui/Card.svelte";
  import PostCard from "../post/PostCard.svelte";
  import type { ListPostsQueryResult } from "$lib/queries";
  import type { Readable } from "svelte/store";
  import { buildPostData, getQuoteForPost, getReplyForPost } from "$lib/util";

  export let postData: Readable<ListPostsQueryResult>;
  export let showReplySource = true;
  $: includes = $postData.data.includes;
  $: allPosts = $postData.data.posts.map((p) => ({
    post: buildPostData(p, includes),
    reply: getReplyForPost(p, includes),
    quote: getQuoteForPost(p, includes),
  }));
  $: query = $postData.query;

  function handleLoadMore() {
    if ($postData.query.hasNextPage) {
      $postData.query.fetchNextPage();
    }
  }
</script>

{#if query.isLoading && allPosts.length === 0}
  <Card>
    <p class="text-text-secondary">Loading posts...</p>
  </Card>
{:else if query.error}
  <Card>
    <p class="text-error">{query.error}</p>
  </Card>
{:else if allPosts.length === 0}
  <Card>
    <p class="text-text-secondary">No posts yet.</p>
  </Card>
{:else}
  {#each allPosts as post}
    {#if showReplySource && post.reply}
      <PostCard data={post.reply} variant="reply_start" />
      <PostCard data={post.post} variant="reply_end" />
    {:else}
      <PostCard data={post.post} quote={post.quote} />
    {/if}
  {/each}
  {#if query.hasNextPage}
    <button
      on:click={handleLoadMore}
      class="w-full p-2 bg-primary text-white rounded mt-4"
    >
      {query.isFetchingNextPage ? "Loading more..." : "Load More"}
    </button>
  {/if}
{/if}

{#if query.isFetchingNextPage}
  <Card>
    <p class="text-text-secondary">Loading more posts...</p>
  </Card>
{/if}
