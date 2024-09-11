<script lang="ts">
  import { useGetCurrentUser } from "$lib/queries";
  import Modal from "$lib/components/ui/Modal.svelte";
  import Input from "$lib/components/ui/Input.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import { X, Loader, Upload } from "lucide-svelte";
  import {
    updateProfile,
    createObjectURL,
    revokeObjectURL,
    type ProfileUpdateData,
  } from "$lib/profile";
  import { onDestroy } from "svelte";

  export let isOpen = false;
  export let onClose: () => void;

  const currentUserQuery = useGetCurrentUser({ profile: true });

  let profileUpdateData: ProfileUpdateData = {};
  let profilePicturePreview: string | null = null;
  let bannerPreview: string | null = null;
  let isUpdating = false;
  let isNewProfilePicture = false;
  let isNewBanner = false;

  $: if ($currentUserQuery.data) {
    profileUpdateData = {
      name: $currentUserQuery.data.name,
      bio: $currentUserQuery.data.profile?.bio,
      website: $currentUserQuery.data.profile?.website,
      location: $currentUserQuery.data.profile?.location,
    };
    // Set initial previews from current user data
    profilePicturePreview = $currentUserQuery.data.pfp_url || null;
    bannerPreview = $currentUserQuery.data.profile?.banner_url || null;
  }

  function handleFileChange(
    event: Event,
    fileType: "profilePicture" | "bannerImage",
  ) {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files[0]) {
      profileUpdateData[fileType] = target.files[0];
      if (fileType === "profilePicture") {
        revokeObjectURL(profilePicturePreview);
        profilePicturePreview = createObjectURL(
          profileUpdateData.profilePicture,
        );
        isNewProfilePicture = true;
      } else if (fileType === "bannerImage") {
        revokeObjectURL(bannerPreview);
        bannerPreview = createObjectURL(profileUpdateData.bannerImage);
        isNewBanner = true;
      }
    }
  }

  function revertFile(fileType: "profilePicture" | "bannerImage") {
    if (fileType === "profilePicture") {
      revokeObjectURL(profilePicturePreview);
      profilePicturePreview = $currentUserQuery.data?.pfp_url || null;
      profileUpdateData.profilePicture = undefined;
      isNewProfilePicture = false;
    } else if (fileType === "bannerImage") {
      revokeObjectURL(bannerPreview);
      bannerPreview = $currentUserQuery.data?.profile?.banner_url || null;
      profileUpdateData.bannerImage = undefined;
      isNewBanner = false;
    }
    const fileInput = document.getElementById(
      `${fileType}-input`,
    ) as HTMLInputElement;
    if (fileInput) fileInput.value = "";
  }

  async function handleUpdateProfile() {
    isUpdating = true;
    try {
      await updateProfile(profileUpdateData);
      $currentUserQuery.refetch();
      isUpdating = false;
      handleClose();
    } catch (error) {
      console.error("Failed to update profile:", error);
      alert("Failed to update profile. Please try again.");
    } finally {
      isUpdating = false;
    }
  }

  function handleClose() {
    if (!isUpdating) {
      revertFile("profilePicture");
      revertFile("bannerImage");
      onClose();
    }
  }

  onDestroy(() => {
    revokeObjectURL(profilePicturePreview);
    revokeObjectURL(bannerPreview);
  });
</script>

<Modal {isOpen} onClose={handleClose}>
  <div class="p-6 relative">
    {#if isUpdating}
      <div
        class="absolute inset-0 bg-background bg-opacity-50 flex items-center justify-center z-10"
      >
        <Loader class="animate-spin text-primary" size={48} />
      </div>
    {/if}
    <h2 class="text-text text-2xl font-bold mb-4">Edit Profile</h2>
    <form on:submit|preventDefault={handleUpdateProfile} class="space-y-4">
      <!-- Banner Image Section -->
      <div>
        <label
          for="banner-image-input"
          class="block text-sm font-medium text-text-secondary mb-2"
        >
          Banner Image
        </label>
        <div
          class="relative w-full h-32 bg-primary-light rounded-md overflow-hidden"
        >
          {#if bannerPreview}
            <img
              src={bannerPreview}
              alt="Banner preview"
              class="w-full h-full object-cover"
            />
            {#if isNewBanner}
              <Button
                variant="outline"
                on:click={() => revertFile("bannerImage")}
                class="absolute top-2 right-2 p-1"
                disabled={isUpdating}
              >
                <X size={16} />
              </Button>
            {/if}
          {:else}
            <div
              class="w-full h-full flex items-center justify-center text-background"
            >
              No banner image
            </div>
          {/if}
          <input
            type="file"
            id="banner-image-input"
            accept="image/jpeg,image/png"
            on:change={(e) => handleFileChange(e, "bannerImage")}
            class="hidden"
            disabled={isUpdating}
          />
          <label
            for="banner-image-input"
            class="absolute bottom-2 right-2 cursor-pointer inline-flex items-center px-2 py-1 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-primary hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary disabled:opacity-50 disabled:cursor-not-allowed"
            class:pointer-events-none={isUpdating}
          >
            <Upload size={16} class="mr-1" />
            Upload Banner
          </label>
        </div>
      </div>

      <!-- Profile Picture Section -->
      <div>
        <label
          for="profile-picture-input"
          class="block text-sm font-medium text-text-secondary mb-2"
        >
          Profile Picture
        </label>
        <div class="flex items-center space-x-4">
          <div
            class="w-16 h-16 rounded-full overflow-hidden bg-primary-light flex items-center justify-center text-background"
          >
            {#if profilePicturePreview}
              <img
                src={profilePicturePreview}
                alt="Profile preview"
                class="w-full h-full object-cover"
              />
            {:else}
              {$currentUserQuery.data?.name?.[0].toUpperCase() ?? "U"}
            {/if}
          </div>
          <div class="flex items-center space-x-2">
            {#if isNewProfilePicture}
              <Button
                variant="outline"
                on:click={() => revertFile("profilePicture")}
                class="p-2"
                disabled={isUpdating}
              >
                <X size={16} />
              </Button>
            {/if}
            <div>
              <input
                type="file"
                id="profile-picture-input"
                accept="image/jpeg,image/png"
                on:change={(e) => handleFileChange(e, "profilePicture")}
                class="hidden"
                disabled={isUpdating}
              />
              <label
                for="profile-picture-input"
                class="cursor-pointer inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-primary hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary disabled:opacity-50 disabled:cursor-not-allowed"
                class:pointer-events-none={isUpdating}
              >
                <Upload size={16} class="mr-1" />
                Upload Picture
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- Other form fields -->
      <div>
        <label for="name" class="block text-sm font-medium text-text-secondary"
          >Name</label
        >
        <Input
          type="text"
          id="name"
          bind:value={profileUpdateData.name}
          disabled={isUpdating}
        />
      </div>
      <div>
        <label for="bio" class="block text-sm font-medium text-text-secondary"
          >Bio</label
        >
        <textarea
          id="bio"
          bind:value={profileUpdateData.bio}
          class="w-full px-3 py-2 text-text bg-surface border border-input rounded-md focus:outline-none focus:ring-2 focus:ring-primary disabled:opacity-50"
          disabled={isUpdating}
        ></textarea>
      </div>
      <div>
        <label
          for="website"
          class="block text-sm font-medium text-text-secondary">Website</label
        >
        <Input
          type="text"
          id="website"
          bind:value={profileUpdateData.website}
          disabled={isUpdating}
        />
      </div>
      <div>
        <label
          for="location"
          class="block text-sm font-medium text-text-secondary">Location</label
        >
        <Input
          type="text"
          id="location"
          bind:value={profileUpdateData.location}
          disabled={isUpdating}
        />
      </div>
      <div class="flex justify-end space-x-2">
        <Button variant="outline" on:click={handleClose} disabled={isUpdating}
          >Cancel</Button
        >
        <Button type="submit" disabled={isUpdating}>
          {#if isUpdating}
            <Loader class="animate-spin mr-2" size={16} />
          {/if}
          Save Changes
        </Button>
      </div>
    </form>
  </div>
</Modal>
