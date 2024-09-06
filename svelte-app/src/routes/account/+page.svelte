<script lang="ts">
  import { useGetCurrentUser, useUpdateCurrentUser } from "$lib/queries";

  const currentUserQuery = useGetCurrentUser({ profile: true });
  const updateUserMutation = useUpdateCurrentUser();

  let editMode = false;
  let editableUser: {
    name?: string;
    bio?: string;
    website?: string;
    location?: string;
  } = {};

  $: if ($currentUserQuery.data) {
    editableUser = {
      name: $currentUserQuery.data.data.name,
      bio: $currentUserQuery.data.data.profile?.bio,
      website: $currentUserQuery.data.data.profile?.website,
      location: $currentUserQuery.data.data.profile?.location,
    };
  }

  function handleEditToggle() {
    editMode = !editMode;
  }

  function handleUpdateUser() {
    $updateUserMutation.mutate(editableUser, {
      onSuccess: () => {
        editMode = false;
      },
    });
  }
</script>

<div>
  <h1>User Profile</h1>

  {#if $currentUserQuery.isLoading}
    <p>Loading user data...</p>
  {:else if $currentUserQuery.isError}
    <p>Error: {$currentUserQuery.error.message}</p>
  {:else}
    <div>
      {#if editMode}
        <input bind:value={editableUser.name} placeholder="Name" />
        <textarea bind:value={editableUser.bio} placeholder="Bio"></textarea>
        <input bind:value={editableUser.website} placeholder="Website" />
        <input bind:value={editableUser.location} placeholder="Location" />
        <button
          on:click={handleUpdateUser}
          disabled={$updateUserMutation.isPending}
        >
          Save Changes
        </button>
      {:else}
        <p>Name: {$currentUserQuery.data.data.name}</p>
        <p>Username: {$currentUserQuery.data.data.username}</p>
        <p>Bio: {$currentUserQuery.data.data.profile?.bio}</p>
        <p>Website: {$currentUserQuery.data.data.profile?.website}</p>
        <p>Location: {$currentUserQuery.data.data.profile?.location}</p>
        <p>Followers: {$currentUserQuery.data.data.profile?.follower_count}</p>
        <p>Following: {$currentUserQuery.data.data.profile?.following_count}</p>
      {/if}
      <button on:click={handleEditToggle}>
        {editMode ? "Cancel" : "Edit Profile"}
      </button>
    </div>
  {/if}
</div>
