<script lang="ts">
  import { onMount, afterUpdate } from "svelte";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import "../app.css";
  import TopBar from "$lib/components/layout/TopBar.svelte";
  import SideBar from "$lib/components/layout/SideBar.svelte";
  import { getAuthState } from "$lib/auth";
  import { writable } from "svelte/store";

  const isAuthenticated = writable(false);

  let isSidebarOpen = false;

  $: publicRoutes = ["/"];
  $: isPublicRoute = publicRoutes.includes($page.url.pathname);

  function toggleSidebar() {
    isSidebarOpen = !isSidebarOpen;
  }

  function closeSidebar() {
    isSidebarOpen = false;
  }

  function checkAuth() {
    const authState = getAuthState();
    isAuthenticated.set(authState.isAuthenticated);

    if ($isAuthenticated && isPublicRoute) {
      goto("/home");
    } else if (!$isAuthenticated && !isPublicRoute) {
      goto("/");
    }
  }

  onMount(checkAuth);

  afterUpdate(checkAuth);

  $: if ($page.url.pathname === "/home") {
    checkAuth();
  }
</script>

{#if !$isAuthenticated || isPublicRoute}
  <slot />
{:else}
  <div class="flex flex-col h-screen">
    <TopBar onToggleSidebar={toggleSidebar} />
    <div class="flex flex-1 overflow-hidden">
      <SideBar bind:isSidebarOpen onCloseSidebar={closeSidebar} />
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
