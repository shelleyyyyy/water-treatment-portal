<script>
    import { goto } from '$app/navigation';
  
    import PocketBase from 'pocketbase';

    const pb = new PocketBase('http://192.168.1.179:8080');
    const logout = () => {
        pb.authStore.clear()
        goto("/login")
    }

</script>




<div class="navbar bg-base-100">
    <div class="navbar-start">
      <a href="/app" class="btn btn-ghost normal-case text-xl">Water Treatment Portal</a>
    </div>
    <div class="navbar-end">
      <button on:click={logout} class="btn">Logout</button>
    </div>
  </div>

<main>
    <slot/>
</main>