<script lang="ts" module>
  import { goto } from "$app/navigation";
  import type {ParsedUserInteractions, ParsedIncludesData} from "$lib/api"

  export type PostCardVariant = "normal" | "reply_source" | "reply_dest";
  export type PostData = {
    user: User
    post: Post
    is_liked?: boolean
    is_bookmarked?: boolean
  }

</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<script lang="ts">
  import { toPostDate } from "$lib/util/date";
  import { MoreHorizontal, Users } from "lucide-svelte";
  import type { Post, User } from "$lib/api";
  import PostActionButtons from "./PostActionButtons.svelte";

  export let data: PostData;
  export let variant: PostCardVariant = "normal";
  export let quote: PostData = undefined;
  export let showButtons: boolean = true;
  export let clickable: boolean = true;

  $: replySource = variant === "reply_source";
  $: replyDest = variant === "reply_dest";
  $: postDate = toPostDate(new Date(data.post.created_at))

  function handlePostClick() {
    if (clickable) {
      goto(`/post/${data.post.id}`)
    }
  }

  function handleUserClick() {
    goto(`/profile/${data.user.username}`)
  }

  function handleQuoteClick() {
    goto(`/post/${quote.post.id}`)
  }

  function handleMoreClick() {
    console.log("more clicked");
  }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  class="p-3 relative border-t-0
  {replySource ? 'border-b-0' : ''} border-y border-border"
  class:cursor-pointer={clickable}
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
</div>
