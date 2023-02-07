<script>
    import "../app.css";

    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';
	import { pb } from "$lib/app/pocketbase.js"

	onMount(() => {
		if(!pb.authStore.isValid){
			goto('/login')
		} else {
            // goto('/app')
        }
	})
</script>

<main>
    <slot/>
</main>

