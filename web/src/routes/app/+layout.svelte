<script>
    import { goto } from '$app/navigation';
  
    import { pb, currentUser } from "$lib/app/pocketbase.js"
    
    const logout = () => {
        pb.authStore.clear()
        goto("/login")
    }

    console.log(currentUser.username, "********")

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