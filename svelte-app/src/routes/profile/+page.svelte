<script lang="ts">
  import { onMount } from "svelte";
  import { api } from "$lib/api";
  import { ArrowLeft, Calendar, Link, MapPin } from "lucide-svelte";
  import Button from "$lib/components/ui/Button.svelte";

  let userData = {
    id: 1831507521534574592,
    name: "Karl Smith",
    username: "kylkrie",
    pfp_url: null,
    protected: true,
    profile: {
      bio: "hell yeah brother",
      website: "https://yabrio.io",
      location: null,
      birthday: null,
      pinned_post_id: null,
    },
    metrics: {
      FollowersCount: 0,
      FollowingCount: 0,
      PostCount: 0,
      ListedCount: 0,
    },
  };

  // Placeholder for future API call
  // onMount(async () => {
  //   try {
  //     const response = await api.get('/api/users/v1/profile');
  //     userData = response.data;
  //   } catch (error) {
  //     console.error('Failed to fetch user profile', error);
  //   }
  // });
</script>

<div class="max-w-2xl mx-auto text-text">
  <!-- Header -->
  <div class="border-b border-gray-200 p-4 flex items-center">
    <a href="/" class="mr-4">
      <ArrowLeft class="h-5 w-5" />
    </a>
    <div>
      <h1 class="text-xl font-bold">{userData.name}</h1>
      <p class="text-sm text-text-secondary">
        {userData.metrics.PostCount} posts
      </p>
    </div>
  </div>

  <!-- Profile Header -->
  <div class="relative">
    <div class="h-32 bg-gray-300"></div>
    <!-- Placeholder for banner image -->
    <div class="absolute bottom-0 left-4 transform translate-y-1/2">
      <div
        class="w-24 h-24 rounded-full bg-gray-400 border-4 border-surface"
      ></div>
      <!-- Placeholder for profile picture -->
    </div>
  </div>

  <!-- Profile Info -->
  <div class="mt-16 px-4">
    <div class="flex justify-end mb-4">
      <Button variant="outline">Edit profile</Button>
    </div>
    <h2 class="text-xl font-bold">{userData.name}</h2>
    <p class="text-text-secondary">@{userData.username}</p>
    <p class="mt-2">{userData.profile.bio}</p>
    <div class="flex flex-wrap mt-2 text-sm text-text-secondary">
      {#if userData.profile.website}
        <span class="flex items-center mr-4">
          <Link class="h-4 w-4 mr-1" />
          <a
            href={userData.profile.website}
            target="_blank"
            rel="noopener noreferrer"
            class="text-primary">{userData.profile.website}</a
          >
        </span>
      {/if}
      {#if userData.profile.location}
        <span class="flex items-center mr-4">
          <MapPin class="h-4 w-4 mr-1" />
          {userData.profile.location}
        </span>
      {/if}
      <span class="flex items-center">
        <Calendar class="h-4 w-4 mr-1" />
        Joined January 2023
      </span>
    </div>
    <div class="flex mt-4 text-sm">
      <span class="mr-4"
        ><strong>{userData.metrics.FollowingCount}</strong> Following</span
      >
      <span><strong>{userData.metrics.FollowersCount}</strong> Followers</span>
    </div>
  </div>

  <!-- Tabs -->
  <div class="border-b border-gray-200 mt-4">
    <nav class="flex">
      {#each ["Posts", "Replies", "Highlights", "Media", "Likes"] as tab}
        <button
          class="px-4 py-2 font-medium text-sm {tab === 'Posts'
            ? 'text-primary border-b-2 border-primary'
            : 'text-text-secondary'}"
        >
          {tab}
        </button>
      {/each}
    </nav>
  </div>

  <!-- Posts will be added here later -->
  <div class="p-4">
    <p class="text-center text-text-secondary">Posts will be displayed here.</p>
  </div>
</div>
