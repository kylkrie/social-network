<script lang="ts">
  import {
    MessageCircle,
    Repeat2,
    Heart,
    Bookmark,
    Share,
    Link,
  } from "lucide-svelte";
  import { auth, authModalStore, postModalStore } from "$lib/stores";
  import type { PostData } from "./PostCard.svelte";
  import { postsApi } from "$lib/api/posts";
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";

  export let data: PostData;

  $: isAuthenticated = !!$auth?.accessToken;

  let isShareDropdownOpen = false;
  let shareButtonRef: HTMLButtonElement;
  let dropdownRef: HTMLDivElement;

  function handleReply() {
    if (isAuthenticated) {
      postModalStore.openModal("reply", data.post, data.user);
    } else {
      authModalStore.openModal("post");
    }
  }

  function handleQuote() {
    if (isAuthenticated) {
      postModalStore.openModal("quote", data.post, data.user);
    } else {
      authModalStore.openModal("post");
    }
  }

  async function handleLike() {
    if (!isAuthenticated) {
      authModalStore.openModal("like posts");
      return;
    }
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
    }
  }

  async function handleBookmark() {
    if (!isAuthenticated) {
      authModalStore.openModal("bookmark posts");
      return;
    }
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
    }
  }

  function toggleShareDropdown(event: Event) {
    event.stopPropagation();
    isShareDropdownOpen = !isShareDropdownOpen;
  }

  function handleCopyLink() {
    const postUrl = `${window.location.origin}/post/${data.post.id}`;
    navigator.clipboard.writeText(postUrl).catch((error) => {
      console.error("Failed to copy link:", error);
    });
    isShareDropdownOpen = false;
  }

  function handleClickOutside(event: MouseEvent) {
    if (
      isShareDropdownOpen &&
      !shareButtonRef.contains(event.target as Node) &&
      !dropdownRef.contains(event.target as Node)
    ) {
      isShareDropdownOpen = false;
    }
  }

  onMount(() => {
    document.addEventListener("click", handleClickOutside);
    return () => {
      document.removeEventListener("click", handleClickOutside);
    };
  });
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
  <div class="flex flex-row space-x-2">
    <button
      class="hover:text-blue-500"
      class:text-blue-500={data.is_bookmarked}
      on:click|stopPropagation={handleBookmark}
    >
      <Bookmark size={18} fill={data.is_bookmarked ? "currentColor" : "none"} />
    </button>
    <div class="relative">
      <button
        bind:this={shareButtonRef}
        class="hover:text-primary"
        on:click|stopPropagation={toggleShareDropdown}
      >
        <Share size={18} />
      </button>
      {#if isShareDropdownOpen}
        <div
          bind:this={dropdownRef}
          class="absolute right-0 mt-2 w-48 rounded-md shadow-lg bg-background border border-border z-10"
          transition:fade={{ duration: 100 }}
        >
          <div class="py-1">
            <button
              class="flex items-center w-full px-4 py-2 text-sm text-text hover:bg-surface"
              on:click|stopPropagation={handleCopyLink}
            >
              <Link size={16} class="mr-2" />
              Copy link
            </button>
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>
