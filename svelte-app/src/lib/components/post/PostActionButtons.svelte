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
  import type { Post } from "$lib/api/posts/dtos";

  export let post: Post;

  function handleReply() {
    postModalStore.openModal("reply", post);
  }

  function handleQuote() {
    postModalStore.openModal("quote", post);
  }

  function handleIconClick(action: string) {
    console.log(`${action} clicked`);
  }
</script>

<div class="flex justify-between text-text-secondary">
  <button
    class="flex items-center hover:text-primary-light"
    on:click|stopPropagation={handleReply}
  >
    <MessageCircle size={18} />
    <span class="ml-2">{post.public_metrics?.replies || 0}</span>
  </button>
  <button
    class="flex items-center hover:text-green-500"
    on:click|stopPropagation={handleQuote}
  >
    <Repeat2 size={18} />
    <span class="ml-2">{post.public_metrics?.reposts || 0}</span>
  </button>
  <button
    class="flex items-center hover:text-red-500"
    on:click|stopPropagation={() => handleIconClick("like")}
  >
    <Heart size={18} />
    <span class="ml-2">{post.public_metrics?.likes || 0}</span>
  </button>
  <button
    class="flex items-center hover:text-primary"
    on:click|stopPropagation={() => handleIconClick("view")}
  >
    <Eye size={18} />
    <span class="ml-2">{post.public_metrics?.views || 0}</span>
  </button>
  <div class="flex flex-row space-x-2">
    <button
      class="hover:text-blue-500"
      on:click|stopPropagation={() => handleIconClick("bookmark")}
    >
      <Bookmark size={18} />
    </button>
    <button
      class="hover:text-primary"
      on:click|stopPropagation={() => handleIconClick("share")}
    >
      <Share size={18} />
    </button>
  </div>
</div>
