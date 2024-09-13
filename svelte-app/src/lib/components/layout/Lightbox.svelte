<!-- src/lib/components/Lightbox.svelte -->
<script lang="ts">
  import { fade, scale } from "svelte/transition";
  import { X, ChevronLeft, ChevronRight } from "lucide-svelte";
  import type { Media } from "$lib/api";

  export let images: Media[] = [];
  export let initialIndex: number = 0;

  export let isOpen = false;

  let currentIndex = initialIndex;

  $: currentImage = images[currentIndex];

  function open() {
    isOpen = true;
  }

  function close() {
    isOpen = false;
  }

  function next() {
    currentIndex = (currentIndex + 1) % images.length;
  }

  function previous() {
    currentIndex = (currentIndex - 1 + images.length) % images.length;
  }

  function handleKeydown(event: KeyboardEvent) {
    if (!isOpen) return;

    switch (event.key) {
      case "ArrowRight":
        next();
        break;
      case "ArrowLeft":
        previous();
        break;
      case "Escape":
        close();
        break;
    }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

{#if isOpen}
  <!-- svelte-ignore a11y_click_events_have_key_events -->
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  <div
    class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-75"
    transition:fade={{ duration: 200 }}
    on:click={close}
  >
    <div
      class="relative max-w-4xl w-full h-full flex items-center justify-center"
      on:click|stopPropagation
    >
      <button
        on:click={close}
        class="absolute top-4 right-4 text-white hover:text-gray-300"
      >
        <X size={24} />
      </button>

      <button
        on:click={previous}
        class="absolute left-4 text-white hover:text-gray-300"
      >
        <ChevronLeft size={36} />
      </button>

      <img
        src={currentImage.url}
        alt="Lightbox image"
        class="max-h-full max-w-full object-contain"
        transition:scale={{ duration: 200 }}
      />

      <button
        on:click={next}
        class="absolute right-4 text-white hover:text-gray-300"
      >
        <ChevronRight size={36} />
      </button>

      <div
        class="absolute bottom-4 left-1/2 transform -translate-x-1/2 text-white"
      >
        {currentIndex + 1} / {images.length}
      </div>
    </div>
  </div>
{/if}

<slot {open} />
