<!-- src/lib/components/post/PostModal.svelte -->
<script lang="ts">
  import { useCreatePost } from "$lib/queries/posts";
  import Button from "$lib/components/ui/Button.svelte";
  import { Users, X, Image, Loader } from "lucide-svelte";
  import Modal from "../ui/Modal.svelte";
  import PostCard from "./PostCard.svelte";
  import { postModalStore } from "$lib/stores/postModal";
  import type { CreatePostRequest } from "$lib/api/posts";

  let textAreaElement: HTMLTextAreaElement;
  let content = "";
  const createPostMutation = useCreatePost();
  const MAX_FILES = 4;

  $: ({ variant, post, user } = $postModalStore);
  $: if ($postModalStore.isOpen && textAreaElement) {
    setTimeout(() => textAreaElement.focus(), 0);
  }
  $: isReply = variant === "reply";
  $: isQuote = variant === "quote";

  let selectedFiles: File[] = [];
  let previewUrls: string[] = [];

  $: isPostValid = content.trim().length > 0;

  function handleInput(event: Event) {
    const target = event.target as HTMLTextAreaElement;
    content = target.value;
  }

  function handleFileInput(event: Event) {
    const target = event.target as HTMLInputElement;
    if (target.files) {
      const newFiles = Array.from(target.files).filter((file) =>
        file.type.startsWith("image/"),
      );
      const availableSlots = MAX_FILES - selectedFiles.length;

      // Filter out duplicate files
      const uniqueNewFiles = newFiles.filter(
        (newFile) =>
          !selectedFiles.some(
            (existingFile) =>
              existingFile.name === newFile.name &&
              existingFile.size === newFile.size &&
              existingFile.type === newFile.type,
          ),
      );

      if (uniqueNewFiles.length === 0) {
        alert("All selected images have already been added.");
        target.value = "";
        return;
      }

      if (uniqueNewFiles.length > availableSlots) {
        alert(
          `You can only add ${availableSlots} more image${availableSlots !== 1 ? "s" : ""}. ${MAX_FILES} images maximum.`,
        );
      }

      const filesToAdd = uniqueNewFiles.slice(0, availableSlots);
      selectedFiles = [...selectedFiles, ...filesToAdd];

      // Create new preview URLs only for the newly added files
      const newPreviewUrls = filesToAdd.map((file) =>
        URL.createObjectURL(file),
      );
      previewUrls = [...previewUrls, ...newPreviewUrls];

      if (filesToAdd.length < uniqueNewFiles.length) {
        alert(
          `${uniqueNewFiles.length - filesToAdd.length} image(s) were not added due to the ${MAX_FILES} image limit.`,
        );
      }
    }

    // Reset the file input so the same files can be selected again if needed
    target.value = "";
  }

  function removeFile(index: number) {
    URL.revokeObjectURL(previewUrls[index]);
    selectedFiles = selectedFiles.filter((_, i) => i !== index);
    previewUrls = previewUrls.filter((_, i) => i !== index);
  }

  async function handleSubmit() {
    if (isPostValid) {
      try {
        let postData: CreatePostRequest = { content, media: selectedFiles };
        if (isReply && post) {
          postData.reply_to_post_id = post.id;
        }
        if (isQuote && post) {
          postData.quote_post_id = post.id;
        }
        await $createPostMutation.mutateAsync(postData);
        handleClose();
      } catch (error) {
        console.error("Failed to create post:", error);
        // Handle error (e.g., show error message to user)
      }
    }
  }

  function handleClose() {
    content = "";
    selectedFiles = [];
    previewUrls.forEach(URL.revokeObjectURL);
    previewUrls = [];
    postModalStore.closeModal();
  }
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
        <PostCard
          data={{ post, user }}
          variant="reply_source"
          showButtons={false}
        />
      </div>
    {/if}
    <div class="flex relative">
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
        {#if previewUrls.length > 0}
          <div class="flex flex-wrap mt-2">
            {#each previewUrls as url, index}
              <div class="relative m-1">
                <img
                  src={url}
                  alt="Preview"
                  class="w-24 h-24 object-cover rounded"
                />
                <button
                  class="absolute top-0 right-0 bg-red-500 text-white rounded-full p-1"
                  on:click={() => removeFile(index)}
                >
                  <X size={12} />
                </button>
              </div>
            {/each}
          </div>
        {/if}
      </div>
    </div>
    {#if isQuote && post && user}
      <div class="mt-4 border border-border rounded-lg">
        <PostCard data={{ post, user }} variant="normal" />
      </div>
    {/if}
    <div class="flex justify-between items-center mt-2">
      <div class="flex items-center">
        <input
          type="file"
          id="file-input"
          accept="image/*"
          multiple
          class="hidden"
          on:change={handleFileInput}
        />
        <label
          for="file-input"
          class="cursor-pointer text-primary hover:text-primary-dark mr-2"
          class:opacity-50={selectedFiles.length >= MAX_FILES}
          class:cursor-not-allowed={selectedFiles.length >= MAX_FILES}
        >
          <Image size={24} />
        </label>
        <input
          type="file"
          id="file-input"
          accept="image/*"
          multiple
          class="hidden"
          on:change={handleFileInput}
        />
        <span class="text-sm text-gray-500">
          {selectedFiles.length} / {MAX_FILES} images selected
        </span>
      </div>
      <Button
        variant={isPostValid ? "default" : "outline"}
        size="sm"
        on:click={isPostValid ? handleSubmit : undefined}
        disabled={$createPostMutation.isPending}
      >
        {#if $createPostMutation.isPending}
          <Loader class="animate-spin mr-2" size={16} />
        {/if}
        {isReply ? "Reply" : "Post"}
      </Button>
    </div>
  </div>
</Modal>
