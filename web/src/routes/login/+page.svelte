<script>
    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';
    import { pb } from "$lib/app/pocketbase.js"
  
    onMount(() => {
      if(pb.authStore.isValid){
        goto('/app')
      }
    })
  
    let username = ""
    let password = ""
  
    $: failedLogin = false
  
    const login = async () => {
      await pb.collection('users').authWithPassword(
        username,
        password,
      ).catch(() => {
        console.log("handle error")
      }).then(() => {
        if(pb.authStore.isValid){
          goto("/app")
        } else {
          failedLogin = true
        }
      })
    }
  
  </script>
  
  <div class="hero min-h-screen bg-base-200">
  <div class="hero-content flex-col lg:flex-row-reverse">
    <div class="text-center lg:text-left w-60">
      <h1 class="text-5xl font-bold">Sign In</h1>
      <!-- <p class="py-6">Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.</p> -->
    </div>
    <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
      <div class="card-body">
        {#if failedLogin}
          <h1 class="text-red-500">Invalid Credentials Try Again</h1>
        {/if}
        <div class="form-control">
          <!-- svelte-ignore a11y-label-has-associated-control -->
          <label class="label">
            <span class="label-text">username</span>
          </label>
          <input bind:value={username} type="text" placeholder="username" class="input input-bordered" />
        </div>
        <div class="form-control">
          <!-- svelte-ignore a11y-label-has-associated-control -->
          <label class="label">
            <span class="label-text">password</span>
          </label>
          <input bind:value={password} type="password" placeholder="password" class="input input-bordered" />
        </div>
        <div class="form-control mt-6">
          <button on:click={login} class="btn btn-primary">Login</button>
        </div>
      </div>
    </div>
  </div>
  </div>
  