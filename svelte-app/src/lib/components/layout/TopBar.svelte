<script lang="ts">
  import { Menu } from "lucide-svelte";
  import UserProfileMenu from "./UserProfileMenu.svelte";
  import ThemeToggle from "./ThemeToggle.svelte";

  export let onToggleSidebar: () => void;
</script>

<header class="bg-indigo-600 text-white">
  <div class="mx-auto px-2 sm:px-6 lg:px-8">
    <div class="flex items-center justify-between h-16">
      <div class="flex items-center">
        <!-- Hamburger menu button -->
        <button
          on:click={onToggleSidebar}
          class="inline-flex items-center justify-center p-2 rounded-md text-white hover:text-gray-200 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white md:hidden"
          aria-label="Toggle sidebar"
        >
          <Menu class="h-7 w-7" />
        </button>
        <span class="font-bold text-2xl ml-2 md:ml-0">Yabro Social</span>
      </div>
      <div class="flex items-center space-x-4">
        <ThemeToggle />
        <UserProfileMenu />
      </div>
    </div>
  </div>
</header>
