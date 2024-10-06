<script lang="ts">
  import { useCreatePost } from "$lib/queries/posts";
  import Button from "$lib/components/ui/Button.svelte";
  import { Users, Loader } from "lucide-svelte";
  import type { CreatePostRequest } from "$lib/api/posts";
  import { onMount } from "svelte";
  import { auth, authModalStore } from "$lib/stores";

  export let postId: string;

  let content = "";
  let textareaElement: HTMLTextAreaElement;
  const createPostMutation = useCreatePost();
  $: isAuthenticated = !!$auth?.accessToken;

  $: isSubmitting = $createPostMutation.isPending;
  $: isDisabled = !content.trim() || isSubmitting;

  function autoResize() {
    if (textareaElement) {
      textareaElement.style.height = "auto";
      textareaElement.style.height = textareaElement.scrollHeight + "px";
    }
  }

  onMount(() => {
    if (textareaElement) {
      textareaElement.setAttribute(
        "style",
        "height:" + textareaElement.scrollHeight + "px;overflow-y:hidden;",
      );
    }
  });

  async function handleSubmit() {
    if (!isAuthenticated) {
      authModalStore.openModal("post");
      return;
    }
    if (content.trim()) {
      const postData: CreatePostRequest = {
        content,
        reply_to_post_id: postId,
      };

      try {
        await $createPostMutation.mutateAsync(postData);
        content = "";
        if (textareaElement) {
          textareaElement.style.height = "auto";
        }
        // The query cache will be automatically invalidated and updated
      } catch (error) {
        console.error("Failed to post reply:", error);
        // Handle error (e.g., show error message to user)
      }
    }
  }
</script>

<div class="flex">
  <div class="flex-shrink-0 mr-3">
    <div
      class="w-10 h-10 rounded-full bg-primary-light flex items-center justify-center"
    >
      <Users class="h-6 w-6 text-background" />
    </div>
  </div>
  <div class="flex-grow">
    <textarea
      bind:value={content}
      bind:this={textareaElement}
      on:input={autoResize}
      placeholder="Post your reply"
      class="w-full p-2 bg-background text-text resize-none focus:outline-none min-h-[40px] max-h-[300px] overflow-y-auto"
      rows="1"
    ></textarea>
    <div class="flex justify-end mt-2">
      <Button on:click={handleSubmit} disabled={isDisabled}>
        {#if isSubmitting}
          <Loader class="animate-spin mr-2" size={16} />
        {/if}
        Reply
      </Button>
    </div>
  </div>
</div>
