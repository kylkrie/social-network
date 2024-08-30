<script lang="ts">
  import { Home, Users, MessageSquare, Bell, Settings } from "lucide-svelte";
  import { page } from "$app/stores";

  export let isSidebarOpen = false;
  export let onCloseSidebar: () => void;

  let sidebarItems = [
    { icon: Home, label: "Home", href: "/home" },
    { icon: Users, label: "Friends", href: "/friends" },
    { icon: MessageSquare, label: "Messages", href: "/messages" },
    { icon: Bell, label: "Notifications", href: "/notifications" },
    { icon: Settings, label: "Settings", href: "/settings" },
  ];
</script>

<nav
  class="sidebar-content {isSidebarOpen
    ? 'translate-x-0'
    : '-translate-x-full'} md:translate-x-0 fixed md:static inset-y-0 left-0 transform transition duration-200 ease-in-out z-30 bg-gray-800 w-64 flex-shrink-0"
>
  <div class="h-full flex flex-col">
    <div class="flex-1 flex flex-col pt-5 pb-4 overflow-y-auto">
      <div class="flex-1 px-2 space-y-1">
        {#each sidebarItems as item}
          <a
            href={item.href}
            on:click={onCloseSidebar}
            class="text-gray-300 hover:bg-gray-700 hover:text-white group flex items-center px-2 py-2 text-sm font-medium rounded-md"
            class:bg-gray-900={$page.url.pathname === item.href}
          >
            <svelte:component
              this={item.icon}
              class="mr-3 flex-shrink-0 h-6 w-6"
            />
            {item.label}
          </a>
        {/each}
      </div>
    </div>
  </div>
</nav>

{#if isSidebarOpen}
  <div
    class="fixed inset-0 bg-gray-600 bg-opacity-75 z-20 md:hidden"
    on:click={onCloseSidebar}
  ></div>
{/if}
