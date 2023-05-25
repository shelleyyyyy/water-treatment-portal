<script>
	// @ts-nocheck

	import SectorCard from '$lib/app/SectorCard.svelte';
	/** @type {import('./$types').PageData} */
	export let data;
	import { onMount } from 'svelte';
	import Control from '$lib/app/Control.svelte';
	import SubCard from '$lib/app/SubCard.svelte';
	import { pb } from '$lib/app/pocketbase.js';
	import ChartModal from '$lib/app/ChartModal.svelte';
	import HistChart from '$lib/app/HistChart.svelte';
	import HistChartModal from '$lib/app/HistChartModal.svelte';
	import CallHistChar from '$lib/app/CallHistChar.svelte';
	import HistTable from '$lib/app/HistTable.svelte';
	let records = [];
	let sector = {};
	let subs = [];
	let isConnected = false;

	const getTitle = async (id) => {
		const record = await pb.collection('subsciptions').getOne(id);
		console.log(record.title, "a;slkdjf;alskdj")
	}

	onMount(async () => {
		try {
			// Attempt to connect to Pocketbase
			await pb.collection('subsciptions').getFullList(200, { sort: '-created' });
			isConnected = true;
		} catch (err) {
			console.error('Failed to connect to Pocketbase:', err.message);
			isConnected = false;
		}

		sector = await pb.collection('sectors').getOne(data.title);
		var res = await pb.collection('devices').getFullList(200, {
			expand: 'subs',
			filter: `sector="${data.title}"`
		});

		console.log(res)

		records = res;
	});


</script>

<h1 class="text-xl font-bold p-5 px-10">{sector.title}</h1>
<div class="p-5 flex flex-wrap justify-center gap-5">
	{#each records as r}
		<div class="card bg-primary text-primary-content">
			<div class="card-body">
				<h2 class="card-title text-3xl">{r.title}</h2>

				<p>{r.description}</p>
				<div class="card-actions pt-5 justify-start grid gap-3 grid-cols-1">
					{#if r.controls.length > 0}
						<div class="font-bold text-center text-2xl">Controls</div>
						<div class="flex gap-1 flex-wrap">
							{#each r.controls as c}
								<Control type={c.type} topic={c.topic} options={c.options} title={c.title} />
							{/each}
						</div>
					{/if}

					{#if r.expand.subs.length > 0}
						<div class="font-bold text-center text-2xl">Subscriptions</div>
						<div class="flex gap-1 flex-wrap">
							{#each r.expand.subs as s}
								<div class="bg-base-100 p-5 card grid gap-5">
									<div>
										<SubCard sub={s} />
									</div>
									{#if s.chart}
										<!-- <div class="justify-center flex w-full">
											<ChartModal sub={s}/>
										</div> -->
										<div class="justify-center flex">
											<!-- <HistChart id={s.topic} title={s.id}/> -->
											<!-- <HistChartModal sub={s}/> -->
											<!-- <CallHistChar sub={s}/> -->
											<HistTable sub={s}/>
										</div>
									{/if}
								</div>
							{/each}
						</div>
					{/if}

					<!-- {#if r.subs.length > 0}
						<div class="font-bold text-center text-2xl">Charts</div>
						{#each r.expand.subs as s}
							
						{/each}
					{/if} -->

					
				</div>
			</div>
		</div>
	{/each}
</div>
