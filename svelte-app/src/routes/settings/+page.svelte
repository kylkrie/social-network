<!-- src/routes/settings/+page.svelte -->
<script lang="ts">
  import { onMount } from "svelte";
  import { Bell, Eye, Moon, Globe } from "lucide-svelte";
  import Card from "$lib/components/ui/Card.svelte";
  import Button from "$lib/components/ui/Button.svelte";

  let user = {
    notifications: {
      email: true,
      push: false,
    },
    privacy: {
      profileVisibility: "public",
      showOnlineStatus: true,
    },
    preferences: {
      darkMode: false,
      language: "en",
    },
  };

  let showSuccessMessage = false;

  function updateNotifications() {
    // In a real app, you would send this update to your API
    showSuccessMessage = true;
    setTimeout(() => (showSuccessMessage = false), 3000);
  }

  function updatePrivacy() {
    // In a real app, you would send this update to your API
    showSuccessMessage = true;
    setTimeout(() => (showSuccessMessage = false), 3000);
  }

  function updatePreferences() {
    // In a real app, you would send this update to your API
    showSuccessMessage = true;
    setTimeout(() => (showSuccessMessage = false), 3000);
  }

  onMount(() => {
    // In a real app, you might fetch user settings from your API here
  });
</script>

<div class="container mx-auto px-4 py-8">
  <h1 class="text-3xl font-bold mb-6 text-primary">Settings</h1>

  <div class="space-y-6">
    <Card title="Notification Settings">
      <div class="space-y-4">
        <div>
          <label class="inline-flex items-center">
            <input
              type="checkbox"
              bind:checked={user.notifications.email}
              class="form-checkbox"
            />
            <span class="ml-2">Receive email notifications</span>
          </label>
        </div>
        <div>
          <label class="inline-flex items-center">
            <input
              type="checkbox"
              bind:checked={user.notifications.push}
              class="form-checkbox"
            />
            <span class="ml-2">Receive push notifications</span>
          </label>
        </div>
        <Button on:click={updateNotifications}>
          <Bell class="h-4 w-4 mr-2" />
          Update Notifications
        </Button>
      </div>
    </Card>

    <Card title="Privacy Settings">
      <div class="space-y-4">
        <div>
          <label
            for="profile-visibility"
            class="block text-sm font-medium text-gray-700"
            >Profile Visibility</label
          >
          <select
            id="profile-visibility"
            bind:value={user.privacy.profileVisibility}
            class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
          >
            <option value="public">Public</option>
            <option value="friends">Friends Only</option>
            <option value="private">Private</option>
          </select>
        </div>
        <div>
          <label class="inline-flex items-center">
            <input
              type="checkbox"
              bind:checked={user.privacy.showOnlineStatus}
              class="form-checkbox"
            />
            <span class="ml-2">Show online status</span>
          </label>
        </div>
        <Button on:click={updatePrivacy}>
          <Eye class="h-4 w-4 mr-2" />
          Update Privacy Settings
        </Button>
      </div>
    </Card>

    <Card title="Preferences">
      <div class="space-y-4">
        <div>
          <label class="inline-flex items-center">
            <input
              type="checkbox"
              bind:checked={user.preferences.darkMode}
              class="form-checkbox"
            />
            <span class="ml-2">Dark Mode</span>
          </label>
        </div>
        <div>
          <label for="language" class="block text-sm font-medium text-gray-700"
            >Language</label
          >
          <select
            id="language"
            bind:value={user.preferences.language}
            class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
          >
            <option value="en">English</option>
            <option value="es">Español</option>
            <option value="fr">Français</option>
          </select>
        </div>
        <Button on:click={updatePreferences}>
          <Globe class="h-4 w-4 mr-2" />
          Update Preferences
        </Button>
      </div>
    </Card>
  </div>
</div>
