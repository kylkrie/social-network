<script lang="ts">
  import { useCreatePost } from "$lib/queries/posts";
  import Button from "$lib/components/ui/Button.svelte";
  import { Users, X } from "lucide-svelte";
  import Modal from "../ui/Modal.svelte";
  import PostCard from "./PostCard.svelte";
  import { postModalStore } from "$lib/stores/postModal";
  import type { CreatePostRequest } from "$lib/api/posts";

  let textAreaElement: HTMLTextAreaElement;
  let content = "";
  let characterCount = 0;
  const maxCharacters = 280;

  const createPostMutation = useCreatePost();

  $: ({ variant, post, user } = $postModalStore);
  $: if ($postModalStore.isOpen && textAreaElement) {
    setTimeout(() => textAreaElement.focus(), 0);
  }

  $: isReply = variant === "reply";
  $: isQuote = variant === "quote";
  $: console.log(variant, post, user);
  $: console.log("test");

  function handleInput(event: Event) {
    const target = event.target as HTMLTextAreaElement;
    content = target.value;
    characterCount = content.length;
  }

  async function handleSubmit() {
    if (content.trim() && characterCount <= maxCharacters) {
      try {
        let postData: CreatePostRequest = { content };
        if (isReply && post) {
          postData.reply_to_post_id = post.id;
        }
        if (isQuote && post) {
          postData.quote_post_id = post.id;
        }
        await $createPostMutation.mutateAsync(postData);
        content = "";
        handleClose();
      } catch (error) {
        console.error("Failed to create post:", error);
        // Handle error (e.g., show error message to user)
      }
    }
  }

  function handleClose() {
    content = "";
    postModalStore.closeModal();
  }

  $: isPostValid = content.trim() && characterCount <= maxCharacters;
</script>

<Modal isOpen={$postModalStore.isOpen} onClose={handleClose} width="max-w-xl">
  <div class="flex items-center p-4 pb-0">
    <button on:click={handleClose} class="text-gray-500 hover:text-gray-700">
      <X size={24} />
    </button>
  </div>

  <div class="p-4">
    {#if isReply && post && user}
      <div class="mb-4">
        <PostCard {post} {user} variant="reply_source" showButtons={false} />
      </div>
    {/if}

    <div class="flex relative">
      <!-- Profile picture -->
      <div
        class="rounded-full bg-primary-light h-12 w-12 ml-2 mr-1 flex items-center justify-center"
      >
        <Users class="h-6 w-6 text-background" />
      </div>
      <div class="flex-grow">
        <textarea
          bind:this={textAreaElement}
          class="w-full h-32 p-2 bg-background text-text resize-none focus:outline-none"
          placeholder={isReply ? "Post your reply" : "What's happening?"}
          bind:value={content}
          on:input={handleInput}
        ></textarea>
      </div>
    </div>

    {#if isQuote && post && user}
      <div class="mt-4 border border-border rounded-lg">
        <PostCard {post} {user} variant="normal" />
      </div>
    {/if}

    <div class="flex justify-between items-center mt-2">
      <div class="text-sm text-gray-500">
        {characterCount} / {maxCharacters}
      </div>
      {#if characterCount > maxCharacters}
        <div class="text-sm text-red-500">Character limit exceeded</div>
      {/if}
      <Button
        variant={isPostValid ? "default" : "outline"}
        size="sm"
        on:click={isPostValid ? handleSubmit : undefined}
      >
        {isReply ? "Reply" : "Post"}
      </Button>
    </div>
  </div>
</Modal>
