<script lang="ts">
  import { onMount } from "svelte";
  import { startOAuth, getAuthState, clearToken } from "$lib/auth";

  let isAuthenticated = false;
  let userInfo: any = null;

  async function fetchUserInfo() {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/userinfo`, {
      headers: {
        Authorization: `Bearer ${getAuthState().token}`,
      },
    });
    if (response.ok) {
      userInfo = await response.json();
    } else {
      console.error("Failed to fetch user info");
    }
  }

  function handleLogin() {
    startOAuth();
  }

  function handleLogout() {
    clearToken();
    isAuthenticated = false;
    userInfo = null;
  }

  onMount(() => {
    isAuthenticated = getAuthState().isAuthenticated;
    if (isAuthenticated) {
      fetchUserInfo();
    }
  });

  $: if (isAuthenticated) {
    fetchUserInfo();
  }
</script>

<main>
  <h1>Welcome to Our App</h1>

  {#if isAuthenticated}
    <div>
      <h2>User Information</h2>
      {#if userInfo}
        <p>Username: {userInfo.preferred_username}</p>
        <p>Email: {userInfo.email}</p>
        <p>Full Name: {userInfo.name}</p>
      {:else}
        <p>Loading user information...</p>
      {/if}
      <button on:click={handleLogout}>Logout</button>
    </div>
  {:else}
    <button on:click={handleLogin}>Login</button>
  {/if}
</main>

<style>
  main {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
  }

  h1 {
    color: #333;
    text-align: center;
  }

  button {
    background-color: #4caf50;
    border: none;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    margin: 4px 2px;
    cursor: pointer;
  }
</style>
