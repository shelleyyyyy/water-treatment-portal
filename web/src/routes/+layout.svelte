<script>
    import "../app.css";

    import { goto } from '$app/navigation';
	import PocketBase from 'pocketbase';
    import { onMount } from 'svelte';

	const pb = new PocketBase('http://192.168.1.179:8080');

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

