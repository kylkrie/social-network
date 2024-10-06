<script lang="ts">
  import Modal from "$lib/components/ui/Modal.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import {
    authModalStore,
    startAuthLogin,
    startAuthRegister,
  } from "$lib/stores";
  import { X } from "lucide-svelte";

  function onClose() {
    authModalStore.closeModal();
  }

  function onLogin() {
    startAuthLogin();
  }

  function onSignUp() {
    startAuthRegister();
  }
</script>

<Modal isOpen={$authModalStore.isOpen} {onClose} width="max-w-md">
  <div class="p-6">
    <div class="flex justify-between items-center mb-4">
      <h2 class="text-xl font-semibold">Account Required</h2>
      <button
        on:click={onClose}
        class="text-gray-500 hover:text-gray-700 focus:outline-none"
        aria-label="Close modal"
      >
        <X size={24} />
      </button>
    </div>
    <p class="mb-6">
      You need an account to {$authModalStore.action}. Would you like to sign up
      or log in?
    </p>
    <div class="flex justify-end space-x-4">
      <Button variant="outline" on:click={onSignUp}>Sign Up</Button>

      <Button on:click={onLogin}>Log In</Button>
    </div>
  </div>
</Modal>
