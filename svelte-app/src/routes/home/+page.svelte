<script lang="ts">
  import { onMount } from "svelte";
  import { api } from "$lib/api";
  import Card from "$lib/components/ui/Card.svelte";
  import { Users, MessageSquare, Bell } from "lucide-svelte";
  import { auth } from "$lib/auth";

  let userInfo: any = null;
  let posts: any[] = [];
  let notifications: any[] = [];

  async function fetchUserInfo() {
    try {
      userInfo = await api.get("/userinfo");
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
</script>

<h1 class="text-2xl font-semibold text-gray-900 mb-6">Home</h1>

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
        <h2 class="text-xl font-semibold">{userInfo?.name || "Loading..."}</h2>
        <p class="text-gray-600">{userInfo?.email || ""}</p>
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
