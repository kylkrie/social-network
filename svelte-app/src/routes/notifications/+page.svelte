<!-- src/routes/notifications/+page.svelte -->
<script lang="ts">
  import { onMount } from "svelte";

  import { Bell, ThumbsUp, UserPlus, MessageCircle } from "lucide-svelte";
  import Card from "$lib/components/ui/Card.svelte";

  import Button from "$lib/components/ui/Button.svelte";

  let notifications = [
    {
      id: 1,
      type: "like",
      content: "Alice Johnson liked your post",
      timestamp: "2 minutes ago",
      read: false,
    },

    {
      id: 2,
      type: "friend_request",
      content: "Bob Smith sent you a friend request",
      timestamp: "1 hour ago",
      read: false,
    },
    {
      id: 3,
      type: "comment",
      content: "Carol Williams commented on your photo",

      timestamp: "Yesterday",
      read: true,
    },

    {
      id: 4,
      type: "like",
      content: "David Brown liked your comment",
      timestamp: "2 days ago",
      read: true,
    },
  ];

  function markAsRead(id: number) {
    notifications = notifications.map((notification) =>
      notification.id === id ? { ...notification, read: true } : notification,
    );

    // In a real app, you would send this update to your API
  }

  function markAllAsRead() {
    notifications = notifications.map((notification) => ({
      ...notification,
      read: true,
    }));
    // In a real app, you would send this update to your API
  }

  function getIcon(type: string) {
    switch (type) {
      case "like":
        return ThumbsUp;
      case "friend_request":
        return UserPlus;
      case "comment":
        return MessageCircle;
      default:
        return Bell;
    }
  }

  onMount(() => {
    // In a real app, you might fetch notifications from your API here
  });
</script>

<div class="container mx-auto px-4 py-8">
  <div class="flex justify-between items-center mb-6">
    <h1 class="text-3xl font-bold">Notifications</h1>
    <Button variant="outline" on:click={markAllAsRead}>Mark All as Read</Button>
  </div>

  <Card>
    <ul class="divide-y divide-gray-200">
      {#each notifications as notification (notification.id)}
        <li
          class="py-4 flex items-start space-x-3 {notification.read
            ? 'opacity-50'
            : ''}"
        >
          <div class="flex-shrink-0">
            <svelte:component
              this={getIcon(notification.type)}
              class="h-6 w-6 text-blue-500"
            />
          </div>
          <div class="flex-grow">
            <p class="text-sm font-medium text-gray-900">
              {notification.content}
            </p>
            <p class="text-sm text-gray-500">{notification.timestamp}</p>
          </div>
          {#if !notification.read}
            <Button
              variant="outline"
              size="sm"
              on:click={() => markAsRead(notification.id)}
            >
              Mark as Read
            </Button>
          {/if}
        </li>
      {/each}
    </ul>
  </Card>
</div>
