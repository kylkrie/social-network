<script lang="ts">
  import { onMount } from "svelte";
  import { MessageCircle, Send } from "lucide-svelte";
  import Card from "$lib/components/ui/Card.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import Input from "$lib/components/ui/Input.svelte";

  let conversations = [
    {
      id: 1,
      name: "Alice Johnson",
      lastMessage: "Hey, how are you?",
      timestamp: "10:30 AM",
    },
    {
      id: 2,
      name: "Bob Smith",
      lastMessage: "Did you see the game last night?",
      timestamp: "Yesterday",
    },
    {
      id: 3,
      name: "Carol Williams",
      lastMessage: "Let's meet for coffee",
      timestamp: "2 days ago",
    },
  ];

  let selectedConversation = null;
  let messages = [];
  let newMessage = "";

  function selectConversation(conversation) {
    selectedConversation = conversation;
    // In a real app, you would fetch messages for this conversation from your API
    messages = [
      {
        id: 1,
        sender: conversation.name,
        content: "Hey there!",
        timestamp: "10:30 AM",
      },
      {
        id: 2,
        sender: "You",
        content: "Hi! How are you?",
        timestamp: "10:31 AM",
      },

      {
        id: 3,
        sender: conversation.name,
        content: "I'm good, thanks! How about you?",
        timestamp: "10:32 AM",
      },
    ];
  }

  function sendMessage() {
    if (newMessage.trim() && selectedConversation) {
      messages = [
        ...messages,
        {
          id: messages.length + 1,
          sender: "You",
          content: newMessage,
          timestamp: "Just now",
        },
      ];

      newMessage = "";
      // In a real app, you would send this message to your API
    }
  }

  onMount(() => {
    // In a real app, you might fetch the list of conversations from your API here
  });
</script>

<div class="container mx-auto px-4 py-8">
  <h1 class="text-3xl font-bold mb-6">Messages</h1>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    <!-- Conversations List -->
    <Card title="Conversations">
      <ul class="divide-y divide-gray-200">
        {#each conversations as conversation (conversation.id)}
          <li
            class="py-4 cursor-pointer hover:bg-gray-50"
            on:click={() => selectConversation(conversation)}
          >
            <div class="flex items-center space-x-3">
              <MessageCircle class="h-6 w-6 text-gray-400" />
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-gray-900 truncate">
                  {conversation.name}
                </p>
                <p class="text-sm text-gray-500 truncate">
                  {conversation.lastMessage}
                </p>
              </div>
              <div class="text-xs text-gray-500">
                {conversation.timestamp}
              </div>
            </div>
          </li>
        {/each}
      </ul>
    </Card>

    <!-- Selected Conversation -->
    <div class="md:col-span-2">
      {#if selectedConversation}
        <Card title={selectedConversation.name}>
          <div class="h-96 overflow-y-auto mb-4">
            {#each messages as message (message.id)}
              <div class="mb-4 {message.sender === 'You' ? 'text-right' : ''}">
                <p class="text-sm text-gray-600">{message.sender}</p>

                <div
                  class="inline-block bg-gray-100 rounded-lg px-4 py-2 max-w-xs"
                >
                  {message.content}
                </div>
                <p class="text-xs text-gray-500 mt-1">{message.timestamp}</p>
              </div>
            {/each}
          </div>
          <div class="flex space-x-2">
            <Input
              type="text"
              placeholder="Type a message..."
              bind:value={newMessage}
              on:keypress={(e) => e.key === "Enter" && sendMessage()}
            />
            <Button on:click={sendMessage}>
              <Send class="h-4 w-4 mr-2" />
              Send
            </Button>
          </div>
        </Card>
      {:else}
        <Card>
          <p class="text-center text-gray-500">
            Select a conversation to start messaging
          </p>
        </Card>
      {/if}
    </div>
  </div>
</div>
