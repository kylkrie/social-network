<script lang="ts">
  import { onMount } from "svelte";
  import { useListPosts } from "$lib/queries/posts";
  import type { ListPostsResponse, Post } from "$lib/api/posts/dtos";
  import Card from "$lib/components/ui/Card.svelte";
  import PostCard from "../post/PostCard.svelte";
  import type { User } from "$lib/api";

  let posts: Post[] = [];
  let user: User;
  let error: string | null = null;
  let loading = true;

  const listPosts = useListPosts({ limit: 10 });

  onMount(async () => {
    try {
      const result = await $listPosts.refetch();
      if (result.isSuccess) {
        const res = result.data.pages[0] as ListPostsResponse;
        posts = res.data;
        user = res.includes.users[0];
      } else {
        error = "Failed to fetch posts";
      }
    } catch (e) {
      error = "An error occurred while fetching posts";
      console.error(e);
    } finally {
      loading = false;
    }
  });
</script>

{#if loading}
  <Card>
    <p class="text-text-secondary">Loading posts...</p>
  </Card>
{:else if error}
  <Card>
    <p class="text-error">{error}</p>
  </Card>
{:else if posts.length === 0}
  <Card>
    <p class="text-text-secondary">No posts yet.</p>
  </Card>
{:else}
  {#each posts as post}
    <PostCard {user} {post} />
  {/each}
{/if}
