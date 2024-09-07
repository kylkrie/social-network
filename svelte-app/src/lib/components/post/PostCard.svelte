<!-- src/lib/components/Post.svelte -->
<script lang="ts">
  import {
    MoreHorizontal,
    MessageCircle,
    Repeat2,
    Heart,
    Eye,
    Bookmark,
    Share,
    Users,
  } from "lucide-svelte";
  import type { Post } from "$lib/api/posts/dtos";
  import type { User } from "$lib/api/users/dtos";

  export let post: Post;
  export let user: User;

  function handlePostClick(event: MouseEvent) {
    // Prevent the post click if the click was on a button
    if (!(event.target as HTMLElement).closest("button")) {
      console.log("Post clicked, navigate to post detail view");
    }
  }

  function handleIconClick(action: string) {
    console.log(`${action} clicked`);
  }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  class="border-y border-border p-2 hover:bg-hover cursor-pointer"
  on:click={handlePostClick}
>
  <div class="flex">
    <!-- Profile picture -->
    {#if user.pfp_url}
      <img
        src={user.pfp_url || "/default-avatar.png"}
        alt="Profile"
        class="w-12 h-12 rounded-full mr-3"
      />
    {:else}
      <div
        class="rounded-full bg-primary-light h-12 w-12 mr-3 flex items-center justify-center"
      >
        <Users class="h-6 w-6 text-background" />
      </div>
    {/if}

    <div class="flex-grow">
      <!-- Header: Name, Username, and More options -->
      <div class="flex justify-between items-center mb-0.5">
        <div>
          <span class="font-bold text-text">{user.name}</span>
          <span class="text-text-secondary ml-2">@{user.username}</span>
        </div>
        <button
          class="text-text-secondary hover:text-text"
          on:click={() => handleIconClick("more")}
        >
          <MoreHorizontal size={20} />
        </button>
      </div>

      <!-- Post content -->
      <p class="mb-3 text-text whitespace-pre-wrap break-words">
        {post.content}
      </p>

      <!-- Action icons -->
      <div class="flex justify-between text-text-secondary">
        <button
          class="flex items-center hover:text-primary-light"
          on:click={() => handleIconClick("reply")}
        >
          <MessageCircle size={18} />
          <span class="ml-2">{post.public_metrics?.replies || 0}</span>
        </button>
        <button
          class="flex items-center hover:text-green-500"
          on:click={() => handleIconClick("repost")}
        >
          <Repeat2 size={18} />
          <span class="ml-2">{post.public_metrics?.reposts || 0}</span>
        </button>
        <button
          class="flex items-center hover:text-red-500"
          on:click={() => handleIconClick("like")}
        >
          <Heart size={18} />
          <span class="ml-2">{post.public_metrics?.likes || 0}</span>
        </button>
        <button
          class="flex items-center hover:text-primary"
          on:click={() => handleIconClick("view")}
        >
          <Eye size={18} />
          <span class="ml-2">{post.public_metrics?.views || 0}</span>
        </button>
        <div class="flex flex-row space-x-2">
          <button
            class="hover:text-blue-500"
            on:click={() => handleIconClick("bookmark")}
          >
            <Bookmark size={18} />
          </button>
          <button
            class="hover:text-primary"
            on:click={() => handleIconClick("share")}
          >
            <Share size={18} />
          </button>
        </div>
      </div>
    </div>
  </div>
</div>
