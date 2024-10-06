<script lang="ts" module>
    import { fade } from 'svelte/transition';

  import { type ComponentType } from 'svelte';
  import { type Icon } from 'lucide-svelte';
    import { auth, authModalStore } from '$lib/stores';
  export interface SidebarItem {
    icon: ComponentType<Icon>;
    label: string;
    href: string;
  };
</script>

<script lang="ts">
  import { page } from "$app/stores";
  import Button from "$lib/components/ui/Button.svelte";
  import { postModalStore } from "$lib/stores";
  import { X } from 'lucide-svelte';

  export let sidebarItems: SidebarItem[];
  export let isSidebarOpen = false;
  export let onCloseSidebar: () => void;

  $: isAuthenticated = !!$auth?.accessToken;

  function handlePostClick() {
    if (isAuthenticated) {
      postModalStore.openModal("normal")
    } else {
      authModalStore.openModal("post")
    }
  }
</script>

<nav
  class="fixed md:static inset-y-0 left-0 z-30 w-64 bg-background shadow-lg md:shadow-none transition-transform duration-300 ease-in-out transform {isSidebarOpen ? 'translate-x-0' : '-translate-x-full'} md:translate-x-0"
>
  <div class="h-full flex flex-col">
    <div class="flex justify-end p-4 md:hidden">
      <button on:click={onCloseSidebar} class="text-text-secondary hover:text-text">
        <X size={24} />
      </button>
    </div>
    <div class="flex-1 flex flex-col pt-5 pb-4 overflow-y-auto">
      <div class="flex-1 px-2 space-y-1">
        {#each sidebarItems as item}
          <a
            href={item.href}
            on:click={onCloseSidebar}
            class="text-text-secondary hover:bg-surface hover:text-text group flex items-center px-2 py-2 text-sm font-medium rounded-md"
            class:bg-surface={$page.url.pathname === item.href}
            class:text-text={$page.url.pathname === item.href}
          >
            <svelte:component
              this={item.icon}
              class="mr-3 flex-shrink-0 h-6 w-6"
            />
            {item.label}
          </a>
        {/each}
        <div class="mx-auto max-w-full p-4">
          <Button on:click={handlePostClick}>
            Post
          </Button>
        </div>
      </div>
    </div>
  </div>
</nav>

{#if isSidebarOpen}
  <div
    transition:fade={{ duration: 200 }}
    class="fixed inset-0 bg-gray-800 bg-opacity-75 z-20 md:hidden transition-colors duration-300"
    on:click={onCloseSidebar}
  ></div>
{/if}
