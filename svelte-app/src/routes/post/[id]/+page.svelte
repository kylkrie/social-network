<script lang="ts">
  import { page } from "$app/stores";
  import PostCard from "$lib/components/post/PostCard.svelte";
  import { useGetPost } from "$lib/queries";
  import { getUserForPost } from "$lib/util";
  import { ArrowLeft } from "lucide-svelte";

  $: postId = $page.params.id;
  $: getPost = useGetPost(postId);
  $: post = $getPost.data?.data;
  $: users = $getPost.data?.includes.users;
  $: user = post ? getUserForPost(users, post) : undefined;

  function goBack() {
    history.back();
  }
</script>

<div class="page border-x border-border">
  <div class="text-text p-4 font-bold text-lg flex items-center">
    <button on:click={goBack} class="mr-4 hover:text-primary-light">
      <ArrowLeft size={24} />
    </button>
    <h1>Post</h1>
  </div>
  {#if post && user}
    <PostCard {post} {user} />
  {/if}
</div>

<style>
  .page {
    max-width: 800px;
    margin: 0 auto;
  }
</style>
