<script lang="ts" module>
  import { type ComponentType } from 'svelte';
  import { type Icon } from 'lucide-svelte';
  export interface SidebarItem {
    icon: ComponentType<Icon>;
    label: string;
    href: string;
  };
</script>

<script lang="ts">
  import { page } from "$app/stores";
  import Button from "$lib/components/ui/Button.svelte";
  import PostModal from "$lib/components/post/PostModal.svelte";

  export let sidebarItems: SidebarItem[];
  export let isSidebarOpen = false;
  export let onCloseSidebar: () => void;

  let isPostModalOpen = false;

  function openPostModal() {
    isPostModalOpen = true;
    if (isSidebarOpen) {
      onCloseSidebar();
    }
  }
</script>

<nav
  class="sidebar-content {isSidebarOpen
    ? 'translate-x-0'
    : '-translate-x-full'} md:translate-x-0 fixed md:static inset-y-0 left-0 transform transition ease-in-out z-30 w-64 flex-shrink-0"
>
  <div class="h-full flex flex-col">
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
              this={item.icon as any}
              class="mr-3 flex-shrink-0 h-6 w-6"
            />
            {item.label}
          </a>
        {/each}
      </div>
      <div class="px-4 mb-4">
        <Button on:click={openPostModal} variant="default" size="lg">
          Post
        </Button>
      </div>
    </div>
  </div>
</nav>

{#if isSidebarOpen}
  <div
    class="fixed inset-0 bg-background bg-opacity-75 z-20 md:hidden"
    on:click={onCloseSidebar}
  ></div>
{/if}

<PostModal bind:isOpen={isPostModalOpen} />
