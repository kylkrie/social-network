<script lang="ts">
  import { useGetCurrentUser, useGetUser } from "$lib/queries/users";
  import type { User } from "$lib/api/users/dtos";
  import Button from "$lib/components/ui/Button.svelte";
  import EditProfileModal from "./EditProfileModal.svelte";
  import { Calendar, Link, MapPin } from "lucide-svelte";

  export let profile: string = undefined;

  $: getUser = profile
    ? useGetUser(profile, { profile: true })
    : useGetCurrentUser({ profile: true });

  const getCurrentUser = useGetCurrentUser();

  $: user = $getUser.data as User;
  $: isCurrentUserProfile =
    $getCurrentUser.data?.id && $getCurrentUser.data.id === user?.id;

  $: name = user?.name ?? "";
  $: username = user?.username ? `@${user.username}` : "";
  $: bio = user?.profile?.bio ?? "";
  $: website = user?.profile?.website ?? "";
  $: location = user?.profile?.location ?? "";
  $: postCount = user?.profile?.post_count ?? 0;
  $: followerCount = user?.profile?.follower_count ?? 0;
  $: followingCount = user?.profile?.following_count ?? 0;
  $: joinedDate = user?.created_at
    ? new Date(user.created_at).toLocaleDateString("en-US", {
        month: "long",
        year: "numeric",
      })
    : "";

  let isEditModalOpen = false;

  function openEditModal() {
    isEditModalOpen = true;
  }

  function closeEditModal() {
    isEditModalOpen = false;
  }
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
    <div class="profile-picture-container">
      {#if user?.pfp_url}
        <img src={user.pfp_url} alt={user.name} class="profile-picture" />
      {:else}
        <div class="profile-picture-placeholder"></div>
      {/if}
    </div>
  </div>
  <div class="profile-details p-4">
    <div class="flex justify-between items-start mb-4">
      <div class="flex flex-col">
        <h1 class="text-2xl font-bold">{name}</h1>
        <p class="text-text-tertiary">{username}</p>
      </div>
      {#if isCurrentUserProfile}
        <div>
          <Button on:click={openEditModal}>Edit Profile</Button>
        </div>
      {/if}
    </div>
    {#if bio}
      <p class="mb-4">{bio}</p>
    {/if}
    <div class="flex flex-wrap gap-y-2 mb-4">
      {#if location}
        <div class="flex items-center mr-4 text-text-secondary">
          <MapPin size={16} class="mr-1" />
          <span>{location}</span>
        </div>
      {/if}
      {#if website}
        <div class="flex items-center mr-4 text-text-secondary">
          <Link size={16} class="mr-1" />
          <a
            href={website.startsWith("http") ? website : `https://${website}`}
            target="_blank"
            rel="noopener noreferrer"
            class="text-primary hover:underline"
          >
            {website.replace(/^https?:\/\//, "")}
          </a>
        </div>
      {/if}
      {#if joinedDate}
        <div class="flex items-center text-text-secondary">
          <Calendar size={16} class="mr-1" />
          <span>Joined {joinedDate}</span>
        </div>
      {/if}
    </div>
    <div class="flex space-x-4 text-text-secondary">
      <span><strong class="text-text">{followingCount}</strong> Following</span>
      <span><strong class="text-text">{followerCount}</strong> Followers</span>
      <span><strong class="text-text">{postCount}</strong> Posts</span>
    </div>
  </div>
</div>

<EditProfileModal isOpen={isEditModalOpen} onClose={closeEditModal} />

<style>
  .profile-header {
    position: relative;
    height: 200px;
  }

  .banner-image,
  .banner-placeholder {
    width: 100%;
    height: 200px;
    object-fit: cover;
  }

  .banner-placeholder {
    background-color: var(--primary-light);
  }

  .profile-picture-container {
    position: absolute;
    bottom: -64px;
    left: 16px;
    border: 4px solid var(--background);
    border-radius: 50%;
    overflow: hidden;
  }

  .profile-picture,
  .profile-picture-placeholder {
    width: 128px;
    height: 128px;
    object-fit: cover;
  }

  .profile-picture-placeholder {
    background-color: var(--primary-light);
  }

  .profile-details {
    margin-top: 64px;
  }
</style>
