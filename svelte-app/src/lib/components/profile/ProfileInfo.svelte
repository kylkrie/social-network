<script lang="ts">
  import { onMount } from "svelte";
  import { useGetCurrentUser } from "$lib/queries/users";
  import type { User } from "$lib/api/users/dtos";
  import Card from "$lib/components/ui/Card.svelte";
  import Button from "$lib/components/ui/Button.svelte";

  let user: User | null = null;
  let error: string | null = null;

  const getCurrentUser = useGetCurrentUser({ profile: true });

  onMount(async () => {
    try {
      const result = await $getCurrentUser.refetch();
      if (result.isSuccess) {
        user = result.data;
      } else {
        error = "Failed to fetch user data";
      }
    } catch (e) {
      error = "An error occurred while fetching user data";
      console.error(e);
    }
  });
</script>

{#if error}
  <Card title="Error">
    <p class="text-red-500">{error}</p>
  </Card>
{:else if user}
  <div class="profile-info">
    <div class="profile-header">
      {#if user.profile?.banner_url}
        <img
          src={user.profile.banner_url}
          alt="Profile banner"
          class="banner-image"
        />
      {:else}
        <div class="banner-placeholder"></div>
      {/if}
      {#if user.pfp_url}
        <img src={user.pfp_url} alt={user.name} class="profile-picture" />
      {:else}
        <div class="profile-picture"></div>
      {/if}
    </div>
    <div class="profile-details">
      <div class="flex justify-between items-center">
        <div class="flex flex-col">
          <h1 class="text-2xl font-bold">{user.name}</h1>
          <p class="text-gray-600">@{user.username}</p>
        </div>
        <div>
          <Button>Edit Profile</Button>
        </div>
      </div>
      {#if user.profile?.bio}
        <p class="mt-2">{user.profile.bio}</p>
      {/if}
      <div class="mt-4 flex space-x-4">
        {#if user.profile?.follower_count !== undefined}
          <span>{user.profile.follower_count} Followers</span>
        {/if}
        {#if user.profile?.following_count !== undefined}
          <span>{user.profile.following_count} Following</span>
        {/if}
      </div>
    </div>
  </div>
{:else}
  <Card title="Loading">
    <p>Loading user data...</p>
  </Card>
{/if}

<style>
  .profile-info {
    width: 100%;
  }

  .profile-header {
    position: relative;
    height: 200px;
  }

  .banner-image,
  .banner-placeholder {
    width: 100%;
    height: 150px;
    object-fit: cover;
  }

  .banner-placeholder {
    background-color: #e0e0e0;
  }

  .profile-picture {
    position: absolute;
    background-color: #e0e0e0;
    bottom: 0;
    left: 20px;
    width: 120px;
    height: 120px;
    border-radius: 50%;
    border: 4px solid white;
    object-fit: cover;
  }

  .profile-details {
    padding: 1rem;
  }
</style>
