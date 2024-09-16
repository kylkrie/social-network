<script lang="ts">
  import {
    MessageCircle,
    Repeat2,
    Heart,
    Eye,
    Bookmark,
    Share,
  } from "lucide-svelte";
  import { postModalStore } from "$lib/stores";
  import type { PostData } from "./PostCard.svelte";
  import { postsApi } from "$lib/api/posts";

  export let data: PostData;

  function handleReply() {
    postModalStore.openModal("reply", data.post, data.user);
  }

  function handleQuote() {
    postModalStore.openModal("quote", data.post, data.user);
  }

  async function handleLike() {
    try {
      if (data.is_liked) {
        await postsApi.unlikePost(data.post.id);
        data.is_liked = false;
        data.post.public_metrics.likes--;
      } else {
        await postsApi.likePost(data.post.id);
        data.is_liked = true;
        data.post.public_metrics.likes++;
      }
    } catch (error) {
      console.error("Error toggling like:", error);
      // Optionally, revert the optimistic update here
    }
  }

  async function handleBookmark() {
    try {
      if (data.is_bookmarked) {
        await postsApi.unbookmarkPost(data.post.id);
        data.is_bookmarked = false;
      } else {
        await postsApi.bookmarkPost(data.post.id);
        data.is_bookmarked = true;
      }
    } catch (error) {
      console.error("Error toggling bookmark:", error);
      // Optionally, revert the optimistic update here
    }
  }

  function handleShare() {
    // Implement share functionality
    console.log("Share clicked");
  }
</script>

<div class="flex justify-between text-text-secondary">
  <button
    class="flex items-center hover:text-primary-light"
    on:click|stopPropagation={handleReply}
  >
    <MessageCircle size={18} />
    <span class="ml-2">{data.post.public_metrics?.replies || 0}</span>
  </button>
  <button
    class="flex items-center hover:text-green-500"
    on:click|stopPropagation={handleQuote}
  >
    <Repeat2 size={18} />
    <span class="ml-2">{data.post.public_metrics?.reposts || 0}</span>
  </button>
  <button
    class="flex items-center hover:text-red-500"
    class:text-red-500={data.is_liked}
    on:click|stopPropagation={handleLike}
  >
    <Heart size={18} fill={data.is_liked ? "currentColor" : "none"} />
    <span class="ml-2">{data.post.public_metrics?.likes || 0}</span>
  </button>
  <div></div>
  <!-- <button -->
  <!--   class="flex items-center hover:text-primary" -->
  <!--   on:click|stopPropagation={() => console.log("View clicked")} -->
  <!-- > -->
  <!--   <Eye size={18} /> -->
  <!--   <span class="ml-2">{data.post.public_metrics?.views || 0}</span> -->
  <!-- </button> -->
  <div class="flex flex-row space-x-2">
    <button
      class="hover:text-blue-500"
      class:text-blue-500={data.is_bookmarked}
      on:click|stopPropagation={handleBookmark}
    >
      <Bookmark size={18} fill={data.is_bookmarked ? "currentColor" : "none"} />
    </button>
    <button class="hover:text-primary" on:click|stopPropagation={handleShare}>
      <Share size={18} />
    </button>
  </div>
</div>
