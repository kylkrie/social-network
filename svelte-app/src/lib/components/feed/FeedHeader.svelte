<script lang="ts">
  import { useGetCurrentUser } from "$lib/queries/users";
  import type { User } from "$lib/api/users/dtos";

  const getCurrentUser = useGetCurrentUser({ profile: true });

  $: user = $getCurrentUser.data as User;
  $: name = user?.name ?? "";
  $: username = user?.username ? `@${user.username}` : "";
</script>

<div class="profile-info border border-x-0 border-t-0 border-border">
  <div class="profile-details">
    <div class="flex justify-between items-center">
      <div class="flex flex-col">
        <h1 class="text-2xl font-bold">{name}</h1>
        <p class="text-gray-600">{username}</p>
      </div>
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
