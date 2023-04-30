<script>
    import { goto } from '$app/navigation';
    import { onMount } from 'svelte'

    let status = ""

    onMount(async () => {

      const response = await fetch('http://localhost:1323/status');
      const json = await response.json();
      status = json.status;
        
    })

    const handleStart = async () => {
      await fetch('http://localhost:1323/start')

      status = "run"
    } 

    const handleStop = async () => {
      await fetch('http://localhost:1323/kill')

      status = "off"
    }

    const handlePause = async () => {
      await fetch('http://localhost:1323/pause')

      status = "pause"
    }
  
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
    <div class="navbar-end gap-3">
      Experiment Status: {status}
      {#if status === "off" || status === "pause"}
        <button on:click={handleStart} class="btn btn-success btn-outline">Start</button>
      {:else}
         <button on:click={handlePause} class="btn btn-warning btn-outline">Pause</button>
          <button on:click={handleStop} class="btn btn-error btn-outline">Stop</button>

      {/if}
      <button on:click={logout} class="btn">Logout</button>
    </div>
  </div>

<main>
    <slot/>
</main>