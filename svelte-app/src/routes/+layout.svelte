<script lang="ts">
  import "../app.css";
  import TopBar from "$lib/components/layout/TopBar.svelte";
  import SideBar, {
    type SidebarItem,
  } from "$lib/components/layout/SideBar.svelte";
  import { Bell, Home, MessageSquare, Settings, Users } from "lucide-svelte";
  import { auth } from "$lib/auth";
  import { onMount } from "svelte";
  import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";

  let isSidebarOpen = false;

  function toggleSidebar() {
    isSidebarOpen = !isSidebarOpen;
  }

  function closeSidebar() {
    isSidebarOpen = false;
  }

  onMount(() => auth.init());

  const sidebarItems: SidebarItem[] = [
    { icon: Home, label: "Home", href: "/" },
    { icon: Users, label: "Friends", href: "/friends" },
    { icon: MessageSquare, label: "Messages", href: "/messages" },
    { icon: Bell, label: "Notifications", href: "/notifications" },
    { icon: Settings, label: "Settings", href: "/settings" },
  ];

  $: isAuthenticated = $auth;

  const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
  {#if !isAuthenticated}
    <slot />
  {:else}
    <div class="flex flex-col h-screen">
      <TopBar onToggleSidebar={toggleSidebar} />
      <div class="flex flex-1 overflow-hidden bg-background">
        <SideBar
          bind:isSidebarOpen
          {sidebarItems}
          onCloseSidebar={closeSidebar}
        />
        <!-- Main Content -->
        <main class="flex-1 relative overflow-y-auto focus:outline-none">
          <div class="py-6">
            <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
              <slot />
            </div>
          </div>
        </main>
      </div>
    </div>
  {/if}
</QueryClientProvider>
