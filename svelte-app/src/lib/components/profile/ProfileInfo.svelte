<script lang="ts">
  import { useGetCurrentUser, useGetUser } from "$lib/queries/users";
  import type { User } from "$lib/api/users/dtos";
  import Button from "$lib/components/ui/Button.svelte";

  export let profile: string = undefined;

  const getUser = profile
    ? useGetUser(profile, { profile: true })
    : useGetCurrentUser({ profile: true });

  const getCurrentUser = useGetCurrentUser();

  $: user = $getUser.data as User;
  $: isCurrentUserProfile =
    $getCurrentUser.data?.id && $getCurrentUser.data.id === user?.id;

  $: name = user?.name ?? "";
  $: username = user?.username ? `@${user.username}` : "";
  $: bio = user?.profile?.bio ?? "";
  $: postCount = user?.profile?.post_count ?? 0;
  $: followerCount = user?.profile?.follower_count ?? 0;
  $: followingCount = user?.profile?.following_count ?? 0;
</script>

<div class="profile-info">
  <div class="profile-header">
    {#if user?.profile?.banner_url}
      <img
        src={user.profile.banner_url}
        alt="Profile banner"
        class="banner-image"
      />
    {:else}
      <div class="banner-placeholder"></div>
    {/if}
    {#if user?.pfp_url}
      <img src={user.pfp_url} alt={user.name} class="profile-picture" />
    {:else}
      <div class="profile-picture"></div>
    {/if}
  </div>
  <div class="profile-details">
    <div class="flex justify-between items-center">
      <div class="flex flex-col">
        <h1 class="text-2xl font-bold">{name}</h1>
        <p class="text-gray-600">{username}</p>
      </div>
      {#if isCurrentUserProfile}
        <div>
          <Button>Edit Profile</Button>
        </div>
      {/if}
    </div>
    <p class="mt-2">{bio}</p>
    <div class="mt-4 flex space-x-4">
      <span>{postCount} Posts</span>
      <span>{followerCount} Followers</span>
      <span>{followingCount} Following</span>
    </div>
  </div>
</div>

<style>
  .profile-info {
    width: 100%;
    color: var(--text);
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
    background-color: var(--primary-light);
  }

  .profile-picture {
    position: absolute;
    background-color: var(--primary-light);
    bottom: 0;
    left: 20px;
    width: 120px;
    height: 120px;
    border-radius: 50%;
    border: 4px solid var(--background);
    object-fit: cover;
  }

  .profile-details {
    padding: 1rem;
  }
</style>
