<!-- src/lib/components/layout/UserProfileMenu.svelte -->
<script lang="ts">
  import { goto } from "$app/navigation";
  import { logout } from "$lib/auth";
  import { User, LogOut, Users } from "lucide-svelte";
  import { onMount } from "svelte";

  let isOpen = false;
  let menuRef: HTMLDivElement;

  function handleMyAccount() {
    isOpen = false;
    goto("/account");
  }

  function handleLogout() {
    isOpen = false;
    logout();
    goto("/");
  }

  function toggleMenu() {
    isOpen = !isOpen;
  }

  function clickOutside(node: HTMLElement) {
    const handleClick = (event: MouseEvent) => {
      if (
        node &&
        !node.contains(event.target as Node) &&
        !event.defaultPrevented
      ) {
        node.dispatchEvent(new CustomEvent("outclick"));
      }
    };

    document.addEventListener("click", handleClick, true);

    return {
      destroy() {
        document.removeEventListener("click", handleClick, true);
      },
    };
  }

  onMount(() => {
    if (menuRef) {
      menuRef.addEventListener("outclick", () => {
        isOpen = false;
      });
    }
  });
</script>

<div class="relative user-profile-menu" bind:this={menuRef} use:clickOutside>
  <button
    on:click={toggleMenu}
    class="flex items-center text-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-indigo-600 focus:ring-white"
    aria-haspopup="true"
    aria-expanded={isOpen}
  >
    <span class="sr-only">Open user menu</span>
    <div
      class="rounded-full bg-gray-300 h-10 w-10 flex items-center justify-center"
    >
      <Users class="h-6 w-6 text-gray-600" />
    </div>
  </button>
  {#if isOpen}
    <div
      class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none z-50"
      role="menu"
      aria-orientation="vertical"
      aria-labelledby="user-menu"
    >
      <button
        on:click={handleMyAccount}
        class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
        role="menuitem"
      >
        <User class="inline-block mr-2 h-4 w-4" />
        My Account
      </button>
      <button
        on:click={handleLogout}
        class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
        role="menuitem"
      >
        <LogOut class="inline-block mr-2 h-4 w-4" />
        Logout
      </button>
    </div>
  {/if}
</div>
