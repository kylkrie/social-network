<script lang="ts">
  import { Users, UserPlus, UserMinus } from "lucide-svelte";
  import Card from "$lib/components/ui/Card.svelte";
  import Button from "$lib/components/ui/Button.svelte";

  // Placeholder data for friends
  const friends = [
    { id: 1, name: "Alice Johnson", mutualFriends: 15 },
    { id: 2, name: "Bob Smith", mutualFriends: 7 },
    { id: 3, name: "Carol Williams", mutualFriends: 3 },
  ];

  // Placeholder data for friend requests
  const friendRequests = [
    { id: 4, name: "David Brown" },
    { id: 5, name: "Eva Davis" },
  ];
</script>

<div class="container mx-auto px-4 py-8">
  <h1 class="text-3xl font-bold mb-6">Friends</h1>

  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
    <!-- Friends List -->
    <Card title="Your Friends">
      <ul class="divide-y divide-gray-200">
        {#each friends as friend (friend.id)}
          <li class="py-4 flex items-center justify-between">
            <div class="flex items-center">
              <Users class="h-6 w-6 text-gray-400 mr-3" />
              <div>
                <p class="font-medium">{friend.name}</p>
                <p class="text-sm text-gray-500">
                  {friend.mutualFriends} mutual friends
                </p>
              </div>
            </div>
            <Button
              variant="outline"
              size="sm"
              on:click={() => alert(`Unfriend ${friend.name}`)}
            >
              <UserMinus class="h-4 w-4 mr-2" />
              Unfriend
            </Button>
          </li>
        {/each}
      </ul>
    </Card>

    <!-- Friend Requests -->
    <Card title="Friend Requests">
      <ul class="divide-y divide-gray-200">
        {#each friendRequests as request (request.id)}
          <li class="py-4 flex items-center justify-between">
            <div class="flex items-center">
              <Users class="h-6 w-6 text-gray-400 mr-3" />
              <p class="font-medium">{request.name}</p>
            </div>
            <div class="space-x-2">
              <Button
                variant="default"
                size="sm"
                on:click={() => alert(`Accept ${request.name}`)}
              >
                <UserPlus class="h-4 w-4 mr-2" />
                Accept
              </Button>
              <Button
                variant="outline"
                size="sm"
                on:click={() => alert(`Decline ${request.name}`)}
              >
                <UserMinus class="h-4 w-4 mr-2" />
                Decline
              </Button>
            </div>
          </li>
        {/each}
      </ul>
    </Card>
  </div>
</div>
