<script lang="ts">
  import { onMount } from "svelte";
  import { ArrowRight, Users, MessageCircle, Share2 } from "lucide-svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import Card from "$lib/components/ui/Card.svelte";
  import { getAuthState, startAuthLogin, startAuthRegister } from "$lib/auth";
  import { goto } from "$app/navigation";

  let isAuthenticated = false;

  function handleLogin() {
    startAuthLogin();
  }

  function handleSignUp() {
    startAuthRegister();
  }

  onMount(() => {
    isAuthenticated = getAuthState().isAuthenticated;
    if (isAuthenticated) {
      goto("/dashboard");
    }
  });
</script>

<div class="min-h-screen bg-gradient-to-b from-indigo-100 to-white">
  <header class="container mx-auto px-4 py-8">
    <h1 class="text-4xl font-bold text-center text-indigo-800">Yabro Social</h1>
    <p class="text-xl text-center mt-4 text-gray-600">
      Connect, Share, and Engage with Your Community
    </p>
  </header>

  <main class="container mx-auto px-4 py-12">
    <div class="flex justify-center space-x-4 mb-12">
      {#if isAuthenticated}
        <Button
          variant="default"
          size="lg"
          on:click={() => (window.location.href = "/dashboard")}
        >
          Go to Dashboard <ArrowRight class="ml-2 h-4 w-4" />
        </Button>
      {:else}
        <Button variant="default" size="lg" on:click={handleLogin}>
          Log In <ArrowRight class="ml-2 h-4 w-4" />
        </Button>
        <Button variant="outline" size="lg" on:click={handleSignUp}>
          Sign Up
        </Button>
      {/if}
    </div>

    <section class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-12">
      {#each [{ icon: Users, title: "Build Connections", description: "Connect with friends, family, and like-minded individuals in your community." }, { icon: MessageCircle, title: "Engaging Discussions", description: "Participate in meaningful conversations on topics that matter to you." }, { icon: Share2, title: "Share Your World", description: "Share your experiences, thoughts, and moments with your network." }] as feature}
        <Card title={feature.title}>
          <div class="flex flex-col items-center">
            <svelte:component
              this={feature.icon}
              class="w-12 h-12 text-indigo-500 mb-4"
            />
            <p class="text-gray-600 text-center">{feature.description}</p>
          </div>
        </Card>
      {/each}
    </section>

    <section class="text-center">
      <h2 class="text-3xl font-semibold mb-4">Join Yabro Social Today</h2>
      <p class="text-xl text-gray-600 mb-8">
        Experience a new way of social networking that puts community first.
      </p>
      <Button variant="default" size="lg" on:click={handleSignUp}>
        Get Started
      </Button>
    </section>
  </main>

  <footer class="bg-indigo-800 text-white py-8 mt-12">
    <div class="container mx-auto px-4 text-center">
      <p>&copy; 2024 Yabro Social. All rights reserved.</p>
    </div>
  </footer>
</div>
