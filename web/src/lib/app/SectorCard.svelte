<script>
	import { PUBLIC_POCKETBASE, PUBLIC_POCKETBASE_PORT } from '$env/static/public'
	export let title;
	export let description;
	export let link;
	export let img;
	import { onMount } from 'svelte'
	import Icon from '$lib/img/comingSoon.jpeg';

	let imageSrc = `http://${PUBLIC_POCKETBASE}:${PUBLIC_POCKETBASE_PORT}/api/files/sectors/${link}/${img}`
	
	onMount(() => {
		if(img === ""){
			imageSrc = Icon
			console.log("HITING IT")
		}
	})	

</script>

<div class="card lg:card-side bg-base-300 shadow-xl">
	<figure><img class="w-56 p-5" src={imageSrc} alt={img} /></figure>
		
	<div class="card-body">
		<h2 class="card-title">{title}</h2>
		<p class="w-56">{description}</p>
		<div class="card-actions justify-end">
			<a href={'/app/' + link} class="btn btn-primary">View</a>
		</div>
	</div>
</div>
