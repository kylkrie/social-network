<script lang="ts" module>
    import ReplyInput from "./ReplyInput.svelte";

  import { goto } from "$app/navigation";
  import type { Media } from "$lib/api"
  export type PostCardVariant = "normal" | "reply_start" | "reply_mid" | "reply_end"
  export type PostData = {
    user: User
    post: Post
    is_liked?: boolean
    is_bookmarked?: boolean
    media?: Media[] 
  }
</script>

<script lang="ts">
  import { toPostDate } from "$lib/util/date";
  import { MoreHorizontal, Users } from "lucide-svelte";
  import type { Post, User } from "$lib/api";
  import PostActionButtons from "./PostActionButtons.svelte";
  import Lightbox from "../layout/Lightbox.svelte"

  export let data: PostData;
  export let variant: PostCardVariant = "normal";
  export let quote: PostData = undefined;
  export let showButtons: boolean = true;
  export let clickable: boolean = true;
  export let showReplyInput = false;

  $: replyStart = variant === "reply_start";
  $: replyMid = variant === "reply_mid";
  $: replyEnd = variant === "reply_end";
  $: postDate = toPostDate(new Date(data.post.created_at));
  $: mediaItems = data.media || [];

  function handlePostClick() {
    if (clickable) {
      goto(`/post/${data.post.id}`);
    }
  }

  function handleUserClick() {
    goto(`/profile/${data.user.username}`);
  }

  function handleQuoteClick() {
    goto(`/post/${quote.post.id}`);
  }

  function handleMoreClick() {
    console.log("more clicked");
  }

  let lightboxIndex = 0;
  let lightboxIsOpen = false;

  function handleMediaClick(event: Event, index: number) {
    event.stopPropagation();
    lightboxIndex = index;
    lightboxIsOpen = true;
  }


  function handleLightboxClose() {
    lightboxIsOpen = false;
  }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  class="p-3 relative border-t-0
  {replyStart ? 'border-b-0' : ''} border-y border-border"
  class:cursor-pointer={clickable}
  on:click={handlePostClick}
>
  {#if replyStart}
    <div class="absolute top-4 left-8 w-0.5 h-full bg-border"></div>
  {/if}
  {#if replyMid}
    <div class="absolute left-8 w-0.5 h-full bg-border"></div>
  {/if}
  {#if replyEnd}
    <div class="absolute left-8 w-0.5 h-4 bg-border"></div>
  {/if}
  <div class="flex relative">
    <!-- Profile picture -->
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div class="mr-3">
      <div
        class="w-12 h-12 rounded-full overflow-hidden cursor-pointer"
        on:click|stopPropagation={handleUserClick}
      >
        {#if data.user?.pfp_url}
          <img
            src={data.user.pfp_url || "/default-avatar.png"}
            alt=""
            class="w-full h-full object-cover"
          />
        {:else}
          <div
            class="w-full h-full bg-primary-light flex items-center justify-center"
          >
            <Users class="h-6 w-6 text-background" />
          </div>
        {/if}
      </div>
    </div>
    <div class="flex-grow">
      <!-- Header: Name, Username, and More options -->
      <div class="flex justify-between items-center mb-0.5">
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <div on:click|stopPropagation={handleUserClick}>
          <span class="font-bold text-text hover:underline cursor-pointer">{data.user?.name}</span>
          <span class="text-text-tertiary hover:underline cursor-pointer">@{data.user?.username}</span>
          <span class="text-text-tertiary text-sm"> - {postDate}</span>
        </div>
        {#if showButtons}
          <button
            class="text-text-secondary hover:text-text"
            on:click|stopPropagation={() => handleMoreClick()}
          >
            <MoreHorizontal size={20} />
          </button>
        {/if}
      </div>
      <!-- Post content -->
      <p class="mb-3 text-text whitespace-pre-wrap break-words">
        {data.post.content}
      </p>


      <!-- Media Grid -->
      {#if mediaItems.length > 0}
        <div class="mt-2 mb-3 grid gap-2" class:grid-cols-2={mediaItems.length > 1}>
          {#each mediaItems as media, index}
            <div 
              class="relative overflow-hidden rounded-lg cursor-pointer"
              class:col-span-2={mediaItems.length === 3 && index === 0}
              on:click={(e) => handleMediaClick(e, index)}
            >
              <img 
                src={media.url} 
                alt="Post media" 
                class="w-full h-full object-cover"
                style="aspect-ratio: {media.width} / {media.height};"
              />
            </div>
          {/each}
        </div>
      {/if}
      
      <!-- Quote section -->
      {#if quote}
        <div 
          class="mt-2 mb-3 border border-border rounded-lg p-3 cursor-pointer"
          on:click|stopPropagation={handleQuoteClick}
        >
          <div class="flex items-center mb-2">
            {#if quote.user.pfp_url}
              <img
                src={quote.user.pfp_url || "/default-avatar.png"}
                alt="Quoted Profile"
                class="w-5 h-5 rounded-full mr-2"
              />
            {:else}
              <div class="rounded-full bg-primary-light h-5 w-5 mr-2 flex items-center justify-center">
                <Users class="h-3 w-3 text-background" />
              </div>
            {/if}
            <span class="font-bold text-sm text-text">{quote.user.name}</span>
            <span class="text-sm text-text-secondary ml-1">@{quote.user.username}</span>
          </div>
          <p class="text-sm text-text-secondary whitespace-pre-wrap break-words">
            {quote.post.content}
          </p>
        </div>
      {/if}

      {#if showButtons}
        <PostActionButtons data={data} />
      {/if}
    </div>
  </div>
  {#if showReplyInput}
    <div class="mx-2 mt-4 px-2 py-4 border-t border-border">
      <ReplyInput postId={data.post.id} />
    </div>
  {/if}

</div>


<Lightbox 
  images={mediaItems} 
  initialIndex={lightboxIndex} 
  bind:isOpen={lightboxIsOpen} 
  on:close={handleLightboxClose}
/>

<style>
  .grid-cols-2 > div:first-child:nth-last-child(3),

  .grid-cols-2 > div:first-child:nth-last-child(3) ~ div {
    aspect-ratio: 16 / 9;
  }

</style>
