<script lang="ts">
  import { onMount } from "svelte";
  import { useListPosts } from "$lib/queries/posts";
  import type { Post } from "$lib/api/posts/dtos";
  import Card from "$lib/components/ui/Card.svelte";
  import PostCard from "../post/PostCard.svelte";
  import { useGetCurrentUser } from "$lib/queries";

  let posts: Post[] = [];
  let error: string | null = null;
  let loading = true;

  const listPosts = useListPosts({ limit: 10 });
  const getCurrentUser = useGetCurrentUser({ profile: true });
  const user = $getCurrentUser.data;

  onMount(async () => {
    try {
      const result = await $listPosts.refetch();
      if (result.isSuccess) {
        //@ts-ignore
        posts = result.data.pages[0].data;
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
    <p>Loading posts...</p>
  </Card>
{:else if error}
  <Card>
    <p class="text-red-500">{error}</p>
  </Card>
{:else if posts.length === 0}
  <Card>
    <p>No posts yet.</p>
  </Card>
{:else}
  {#each posts as post}
    <PostCard {user} {post} />
  {/each}
{/if}
