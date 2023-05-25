<script>
	import { pb } from '$lib/app/pocketbase.js';
	import { onMount } from 'svelte';
	import ChartModal from './Chart.svelte';

	export let sub;
	console.log("What I am looking for", sub)

	$: record = 'tmp';

	var title = '';

	onMount(async () => {
		// var res = await pb.collection('subsciptions').getOne(id);
		// console.log(res);
		// title = res.title;
		// record = res.value;

		pb.collection('subsciptions').subscribe(sub.id, function (e) {
			record = e.record.value;
		});

		console.log(title, ": asdfas", value)
	});
</script>

<div class="p-5 bg-secondary rounded-lg">
	<div class="font-bold text-center">
		{sub.title}
	</div>
	<div class="text-center p-3">
		{sub.value}
	</div>
</div>

<!-- <ChartModal sub={sub}/> -->