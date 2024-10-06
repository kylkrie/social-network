<script lang="ts">
  import "../app.css";
  import TopBar from "$lib/components/layout/TopBar.svelte";
  import SideBar from "$lib/components/layout/SideBar.svelte";
  import { Bell, Home, MessageSquare, User, Bookmark } from "lucide-svelte";
  import { auth } from "$lib/stores";
  import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import PostModal from "$lib/components/post/PostModal.svelte";
  import RightBar from "$lib/components/layout/RightBar.svelte";
  import AuthModal from "$lib/components/layout/AuthModal.svelte";

  let isSidebarOpen = false;

  function toggleSidebar() {
    isSidebarOpen = !isSidebarOpen;
  }

  function closeSidebar() {
    isSidebarOpen = false;
  }

  $: isAuthenticated = !!$auth?.accessToken;
  $: sidebarItems = isAuthenticated
    ? [
        { icon: Home, label: "Home", href: "/" },
        { icon: User, label: "Profile", href: "/profile" },
        { icon: MessageSquare, label: "Messages", href: "/messages" },
        { icon: Bookmark, label: "Bookmarks", href: "/bookmarks" },
        { icon: Bell, label: "Notifications", href: "/notifications" },
      ]
    : [{ icon: Home, label: "Home", href: "/" }];

  $: if (!isAuthenticated) {
    const path = $page.url.pathname;
    if (path !== "/" && !path.startsWith("/post")) {
      goto("/");
    }
  }

  const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
  <div class="bg-background text-text">
    <!-- {#if !isAuthenticated} -->
    <!--   <slot /> -->
    <!-- {:else} -->
    <div class="h-screen flex flex-col">
      <!-- Top Bar -->
      <div class="z-10">
        <TopBar onToggleSidebar={toggleSidebar} />
      </div>

      <!-- Main content area with shared background -->
      <div class="flex flex-1 bg-background overflow-hidden">
        <!-- Sidebar (visible on all screen sizes, positioned absolutely on small screens) -->
        <SideBar
          bind:isSidebarOpen
          {sidebarItems}
          onCloseSidebar={closeSidebar}
        />

        <!-- Scrollable Main Content -->
        <main class="flex-1 overflow-y-auto w-full">
          <div class="flex justify-center max-w-screen-xl mx-auto">
            <!-- Center content on small to large screens, align left on xl -->
            <div class="w-full max-w-[600px] xl:px-0">
              <slot />
            </div>

            <!-- Spacer for xl screens -->
            <div class="hidden xl:block w-8"></div>

            <!-- Right sidebar for xl screens -->
            <div class="hidden xl:block w-80">
              <div class="fixed top-16 bottom-0 w-80 overflow-y-auto">
                <RightBar />
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
    <!-- {/if} -->
    <PostModal />
    <AuthModal />
  </div>
</QueryClientProvider>
