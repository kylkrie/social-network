<script lang="ts">
  import { useGetCurrentUser } from "$lib/queries/users";
  import type { User } from "$lib/api/users/dtos";
  import { auth } from "$lib/stores";

  $: isAuthenticated = !!$auth?.accessToken;
  $: getCurrentUser = isAuthenticated
    ? useGetCurrentUser({ profile: true })
    : null;

  $: user = $getCurrentUser?.data as User;
  $: name = user?.name ?? "Guest";
  $: username = user?.username ? `@${user.username}` : "";
</script>

<div class="text-text border border-x-0 border-t-0 border-border p-2">
  <div class="flex justify-between items-center">
    <div class="flex flex-col">
      <h1 class="text-2xl font-bold">{name}</h1>
      <p class="text-text-tertiary">{username}</p>
    </div>
  </div>
</div>
