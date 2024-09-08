<script lang="ts" module>
export type PostCardVariant = "normal" | "reply_source" | "reply_dest";
</script>
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
  import type { Post, PostReference } from "$lib/api/posts/dtos";
  import type { User } from "$lib/api/users/dtos";

  export let post: Post;
  export let user: User;
  export let variant: PostCardVariant = "normal"
  $: replySource = variant === "reply_source"
  $: replyDest = variant === "reply_dest"

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
  class="p-2 hover:bg-hover cursor-pointer relative border-t-0
  {replySource ? 'border-b-0' : ''} border-y border-border"
  on:click={handlePostClick}
>
  {#if replyDest}
    <div class="absolute top-0 left-8 w-0.5 h-2 bg-border"></div>
  {/if}
  {#if replySource}
    <div class="absolute top-2 bottom-0 left-8 w-0.5 h-full bg-border"></div>
  {/if}
  <div class="flex relative">
    <!-- Profile picture -->
    {#if user.pfp_url}
      <img
        src={user.pfp_url || "/default-avatar.png"}
        alt="Profile"
        class="w-12 h-12 rounded-full mr-3 z-10"
      />
    {:else}
      <div
        class="rounded-full bg-primary-light h-12 w-12 mr-3 flex items-center justify-center z-10"
      >
        <Users class="h-6 w-6 text-background" />
      </div>
    {/if}
    <div class="flex-grow">
      <!-- Header: Name, Username, and More options -->
      <div class="flex justify-between items-center mb-0.5">
        <div>
          <span class="font-bold text-text">{user.name}</span>
          <span class="text-text-secondary">@{user.username}</span>
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
