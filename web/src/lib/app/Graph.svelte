<script>
	import { pb } from '$lib/app/pocketbase.js';
	import { onMount } from 'svelte';
	import Chart from 'chart.js/auto';

	let chart;
	let chartData = {
		labels: [],
		datasets: [
			{
				label: 'Data',
				borderColor: 'blue',
				data: []
			}
		]
	};

	onMount(async () => {
		// const myCollection = await pb.collection('devices');
		// const documents = await myCollection.get();

		const ctx = document.getElementById('chart').getContext('2d');

		// const subscription = await pb.subscribe('my_data_source');

		// subscription.on('message', (message) => {
		// 	const { timestamp, value } = message.data;
		// 	chartData.labels.push(timestamp);
		// 	chartData.datasets[0].data.push(value);

		// 	if (chart) {
		// 		chart.update();
		// 	}
		// });

		chart = new Chart(ctx, {
			type: 'line',
			data: chartData,
			options: {
				scales: {
					y: {
						suggestedMin: 0,
						suggestedMax: 100
					}
				}
			}
		});
	});
</script>

<div>
	<canvas id="chart" />
</div>
