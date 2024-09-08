<script lang="ts">
  import { useCreatePost } from "$lib/queries/posts";
  import Button from "$lib/components/ui/Button.svelte";
  import { X } from "lucide-svelte";
  import Modal from "../ui/Modal.svelte";

  export let isOpen = false;
  let textAreaElement: HTMLTextAreaElement;
  $: if (isOpen && textAreaElement) {
    setTimeout(() => textAreaElement.focus(), 0);
  }

  let content = "";
  let characterCount = 0;
  const maxCharacters = 280;

  const createPostMutation = useCreatePost();

  function handleInput(event: Event) {
    const target = event.target as HTMLTextAreaElement;
    content = target.value;
    characterCount = content.length;
  }

  async function handleSubmit() {
    if (content.trim() && characterCount <= maxCharacters) {
      try {
        $createPostMutation.mutate({ content });
        content = "";
        isOpen = false;
      } catch (error) {
        console.error("Failed to create post:", error);
        // Handle error (e.g., show error message to user)
      }
    }
  }

  $: isPostValid = content.trim() && characterCount <= maxCharacters;
</script>

<Modal bind:isOpen width="max-w-xl">
  <div class="flex justify-between items-center mb-4">
    <button
      on:click={() => (isOpen = false)}
      class="text-gray-500 hover:text-gray-700"
    >
      <X size={24} />
    </button>
    <Button
      variant={isPostValid ? "default" : "outline"}
      size="sm"
      on:click={isPostValid ? handleSubmit : undefined}
    >
      Post
    </Button>
  </div>

  <textarea
    bind:this={textAreaElement}
    class="w-full h-32 p-2 bg-background text-text resize-none focus:outline-none"
    placeholder="What's happening?"
    bind:value={content}
    on:input={handleInput}
  ></textarea>

  <div class="flex justify-between items-center mt-2">
    <div class="text-sm text-gray-500">
      {characterCount} / {maxCharacters}
    </div>
    {#if characterCount > maxCharacters}
      <div class="text-sm text-red-500">Character limit exceeded</div>
    {/if}
  </div>
</Modal>
