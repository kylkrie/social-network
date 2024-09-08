<script lang="ts">
  import "../app.css";
  import TopBar from "$lib/components/layout/TopBar.svelte";
  import SideBar, {
    type SidebarItem,
  } from "$lib/components/layout/SideBar.svelte";
  import { Bell, Home, MessageSquare, User, Bookmark } from "lucide-svelte";
  import { auth } from "$lib/stores";
  import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import PostModal from "$lib/components/post/PostModal.svelte";

  let isSidebarOpen = false;

  function toggleSidebar() {
    isSidebarOpen = !isSidebarOpen;
  }

  function closeSidebar() {
    isSidebarOpen = false;
  }

  const sidebarItems: SidebarItem[] = [
    { icon: Home, label: "Home", href: "/" },
    { icon: User, label: "Profile", href: "/profile" },
    { icon: MessageSquare, label: "Messages", href: "/messages" },
    { icon: Bookmark, label: "Bookmarks", href: "/bookmarks" },
    { icon: Bell, label: "Notifications", href: "/notifications" },
  ];

  $: isAuthenticated = $auth;

  $: if (!isAuthenticated && $page.url.pathname !== "/") {
    goto("/");
  }

  const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
  {#if !isAuthenticated}
    <slot />
  {:else}
    <div class="h-screen flex flex-col">
      <!-- Top Bar (kept as is) -->
      <div class="z-10">
        <TopBar onToggleSidebar={toggleSidebar} />
      </div>

      <!-- Main content area with shared background -->
      <div class="flex flex-1 bg-background overflow-hidden">
        <!-- Sidebar (no background color of its own) -->
        <div class="fixed left-0 top-16 bottom-0 z-10 w-64">
          <SideBar
            bind:isSidebarOpen
            {sidebarItems}
            onCloseSidebar={closeSidebar}
          />
        </div>

        <!-- Scrollable Main Content (no background color of its own) -->
        <main class="flex-1 overflow-y-auto ml-64">
          <!-- Adjust top padding as needed -->
          <div class="mx-auto max-w-[600px]">
            <slot />
          </div>
        </main>
      </div>
    </div>
  {/if}
  <PostModal />
</QueryClientProvider>
