<script lang="ts">
  import { goto } from "$app/navigation";
  import { useGetCurrentUser } from "$lib/queries";

  $: user = useGetCurrentUser();
  $: username = $user?.data?.username;

  $: if (username) {
    goto(`/profile/${username}`);
  }
</script>
