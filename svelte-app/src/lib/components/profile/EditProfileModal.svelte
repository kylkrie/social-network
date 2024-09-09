<script lang="ts">
  import { useGetCurrentUser, useUpdateCurrentUser } from "$lib/queries";
  import Modal from "$lib/components/ui/Modal.svelte";
  import Input from "$lib/components/ui/Input.svelte";
  import Button from "$lib/components/ui/Button.svelte";

  export let isOpen = false;
  export let onClose: () => void;

  const currentUserQuery = useGetCurrentUser({ profile: true });
  const updateUserMutation = useUpdateCurrentUser();

  let editableUser: {
    name?: string;
    bio?: string;
    website?: string;
    location?: string;
  } = {};

  $: if ($currentUserQuery.data) {
    editableUser = {
      name: $currentUserQuery.data.name,
      bio: $currentUserQuery.data.profile?.bio,
      website: $currentUserQuery.data.profile?.website,
      location: $currentUserQuery.data.profile?.location,
    };
  }

  function handleUpdateUser() {
    $updateUserMutation.mutate(editableUser, {
      onSuccess: () => {
        onClose();
      },
    });
  }
</script>

<Modal {isOpen} {onClose}>
  <div class="p-6">
    <h2 class="text-text text-2xl font-bold mb-4">Edit Profile</h2>
    <form on:submit|preventDefault={handleUpdateUser} class="space-y-4">
      <div>
        <label for="name" class="block text-sm font-medium text-text-secondary"
          >Name</label
        >
        <Input type="text" id="name" bind:value={editableUser.name} />
      </div>
      <div>
        <label for="bio" class="block text-sm font-medium text-text-secondary"
          >Bio</label
        >
        <textarea
          id="bio"
          bind:value={editableUser.bio}
          class="w-full px-3 py-2 text-text bg-surface border border-input rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
        ></textarea>
      </div>
      <div>
        <label
          for="website"
          class="block text-sm font-medium text-text-secondary">Website</label
        >
        <Input type="text" id="website" bind:value={editableUser.website} />
      </div>
      <div>
        <label
          for="location"
          class="block text-sm font-medium text-text-secondary">Location</label
        >
        <Input type="text" id="location" bind:value={editableUser.location} />
      </div>
      <div class="flex justify-end space-x-2">
        <Button variant="outline" on:click={onClose}>Cancel</Button>
        <Button type="submit" disabled={$updateUserMutation.isPending}>
          {$updateUserMutation.isPending ? "Saving..." : "Save Changes"}
        </Button>
      </div>
    </form>
  </div>
</Modal>
