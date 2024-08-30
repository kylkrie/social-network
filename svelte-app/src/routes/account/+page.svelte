<!-- src/routes/account/+page.svelte -->

<script lang="ts">
  import { onMount } from "svelte";
  import { User, Lock, Shield } from "lucide-svelte";
  import Card from "$lib/components/ui/Card.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import Input from "$lib/components/ui/Input.svelte";
  import Alert from "$lib/components/ui/Alert.svelte";

  let user = {
    name: "John Doe",
    email: "john.doe@example.com",
  };

  let newPassword = "";
  let confirmPassword = "";
  let currentPassword = "";
  let showSuccessMessage = false;

  function updateProfile() {
    // In a real app, you would send this update to your API
    showSuccessMessage = true;
    setTimeout(() => (showSuccessMessage = false), 3000);
  }

  function updatePassword() {
    if (
      newPassword === confirmPassword &&
      newPassword.length > 0 &&
      currentPassword.length > 0
    ) {
      // In a real app, you would send this update to your API
      newPassword = "";
      confirmPassword = "";
      currentPassword = "";
      showSuccessMessage = true;
      setTimeout(() => (showSuccessMessage = false), 3000);
    } else {
      alert("Please fill all password fields correctly");
    }
  }

  onMount(() => {
    // In a real app, you might fetch user data from your API here
  });
</script>

<div class="container mx-auto px-4 py-8">
  <h1 class="text-3xl font-bold mb-6">My Account</h1>

  {#if showSuccessMessage}
    <Alert title="Success" variant="default" class="mb-4">
      Your account information has been updated successfully.
    </Alert>
  {/if}

  <div class="space-y-6">
    <Card title="Profile Information">
      <div class="space-y-4">
        <div>
          <label for="name" class="block text-sm font-medium text-gray-700"
            >Name</label
          >
          <Input type="text" id="name" bind:value={user.name} />
        </div>
        <div>
          <label for="email" class="block text-sm font-medium text-gray-700"
            >Email</label
          >
          <Input type="email" id="email" bind:value={user.email} />
        </div>
        <Button on:click={updateProfile}>
          <User class="h-4 w-4 mr-2" />
          Update Profile
        </Button>
      </div>
    </Card>

    <Card title="Change Password">
      <div class="space-y-4">
        <div>
          <label
            for="current-password"
            class="block text-sm font-medium text-gray-700"
            >Current Password</label
          >
          <Input
            type="password"
            id="current-password"
            bind:value={currentPassword}
          />
        </div>
        <div>
          <label
            for="new-password"
            class="block text-sm font-medium text-gray-700">New Password</label
          >
          <Input type="password" id="new-password" bind:value={newPassword} />
        </div>

        <div>
          <label
            for="confirm-password"
            class="block text-sm font-medium text-gray-700"
            >Confirm New Password</label
          >
          <Input
            type="password"
            id="confirm-password"
            bind:value={confirmPassword}
          />
        </div>
        <Button on:click={updatePassword}>
          <Lock class="h-4 w-4 mr-2" />
          Change Password
        </Button>
      </div>
    </Card>

    <Card title="Account Security">
      <div class="space-y-4">
        <p class="text-sm text-gray-600">
          Enhance your account security with these additional features:
        </p>
        <Button variant="outline">
          <Shield class="h-4 w-4 mr-2" />
          Enable Two-Factor Authentication
        </Button>
        <Button variant="outline">View Login History</Button>
      </div>
    </Card>
  </div>
</div>
