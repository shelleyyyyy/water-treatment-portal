<script>
	import { pb, currentUser } from '$lib/app/pocketbase.js';
	import { onMount } from 'svelte';
	import Chart from 'chart.js/auto';
	let chart;
	let chartData = {
		labels: [
			'9:00 AM',
			'9:10 AM',
			'9:20 AM',
			'9:30 AM',
			'9:40 AM',
			'9:50 AM',
			'10:00 AM',
			'10:10 AM',
			'10:20 AM',
			'10:30 AM'
		],
		datasets: [
			{
				label: 'Water Temperature',
				borderColor: 'blue',
				data: [23, 24, 25, 26, 26, 27, 28, 28, 29, 30]
			}
		]
	};
	onMount(() => {
		// @ts-ignore
		const ctx = document.getElementById('chart').getContext('2d');
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
