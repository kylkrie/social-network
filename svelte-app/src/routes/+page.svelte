<script lang="ts">
  import {
    ArrowRight,
    Users,
    MessageCircle,
    Share2,
    MessageSquare,
    Bell,
  } from "lucide-svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import Card from "$lib/components/ui/Card.svelte";
  import { auth, startAuthLogin, startAuthRegister } from "$lib/auth";
  import { api } from "$lib/api";
  import { onMount } from "svelte";

  let userInfo: any = null;
  let posts: any[] = [];
  let notifications: any[] = [];

  async function fetchUserInfo() {
    try {
      userInfo = (await api.get("/api/users/v1")).data;
    } catch (error) {
      console.error("Failed to fetch user info", error);
    }
  }

  async function fetchPosts() {
    // Simulated API call
    posts = [
      {
        id: 1,
        author: "Jane Doe",
        content: "Just joined Yabro Social! Excited to connect with everyone!",
      },
      {
        id: 2,
        author: "John Smith",
        content:
          "Beautiful day for a hike. Who's up for an adventure this weekend?",
      },
      {
        id: 3,
        author: "Alice Johnson",
        content:
          "Working on a new project. Can't wait to share it with the community!",
      },
    ];
  }

  async function fetchNotifications() {
    // Simulated API call
    notifications = [
      { id: 1, content: "Jane Doe liked your post" },
      { id: 2, content: "New friend request from John Smith" },
      { id: 3, content: "Your post has 10 new comments" },
    ];
  }

  onMount(() => {
    if ($auth) {
      fetchUserInfo();
      fetchPosts();
      fetchNotifications();
    }
  });

  function handleLogin() {
    startAuthLogin();
  }

  function handleSignUp() {
    startAuthRegister();
  }
</script>

{#if !$auth}
  <div class="min-h-screen bg-gradient-to-b from-indigo-100 to-white">
    <header class="container mx-auto px-4 py-8">
      <h1 class="text-4xl font-bold text-center text-indigo-800">
        Yabro Social
      </h1>
      <p class="text-xl text-center mt-4 text-gray-600">
        Connect, Share, and Engage with Your Community
      </p>
    </header>

    <main class="container mx-auto px-4 py-12">
      <div class="flex justify-center space-x-4 mb-12">
        <Button variant="default" size="lg" on:click={handleLogin}>
          Log In <ArrowRight class="ml-2 h-4 w-4" />
        </Button>
        <Button variant="outline" size="lg" on:click={handleSignUp}>
          Sign Up
        </Button>
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
{:else}
  <h1 class="text-2xl font-semibold text-gray-900 mb-6 text-primary">Home</h1>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    <!-- User Profile -->
    <Card>
      <div class="flex items-center space-x-4">
        <div
          class="rounded-full bg-gray-300 h-12 w-12 flex items-center justify-center"
        >
          <Users class="h-6 w-6 text-gray-600" />
        </div>
        <div>
          <h2 class="text-xl font-semibold">
            {userInfo?.name || "Loading..."}
          </h2>
          <p class="text-gray-600">@{userInfo?.username || ""}</p>
        </div>
      </div>
    </Card>

    <!-- Recent Posts -->
    <Card>
      <h2 class="text-xl font-semibold mb-4 flex items-center">
        <MessageSquare class="mr-2 h-5 w-5" /> Recent Posts
      </h2>
      <ul class="space-y-4">
        {#each posts as post}
          <li>
            <p class="font-medium">{post.author}</p>
            <p class="text-gray-600">{post.content}</p>
          </li>
        {/each}
      </ul>
    </Card>

    <!-- Notifications -->
    <Card>
      <h2 class="text-xl font-semibold mb-4 flex items-center">
        <Bell class="mr-2 h-5 w-5" /> Notifications
      </h2>
      <ul class="space-y-2">
        {#each notifications as notification}
          <li class="text-gray-600">{notification.content}</li>
        {/each}
      </ul>
    </Card>
  </div>
{/if}
