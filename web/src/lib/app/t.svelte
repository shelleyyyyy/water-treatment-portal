<script>
	import { onMount } from 'svelte';
	// library that creates chart objects in page
	import Chart from 'chart.js';
	export let title;
	export let id;
	import { connectSocket, sendMessage, listener } from '../../networking/socketSend';

	$: list = [];

	// init chart
	onMount(async () => {
		var config = {
			type: 'line',
			data: {
				labels: list,
				datasets: [
					{
						label: 'Temperature',
						backgroundColor: 'yellow',
						borderColor: 'yellow',
						data: list,
						fill: false,
						barThickness: 8
					},
					{
						label: 'Temperature',
						backgroundColor: 'hot-pink',
						borderColor: 'hot-pink',
						data: list,
						fill: false,
						barThickness: 8
					},
					{
						label: 'Temperature',
						backgroundColor: '#3182ce',
						borderColor: '#3182ce',
						data: list,
						fill: false,
						barThickness: 8
					}
				]
			},
			options: {
				maintainAspectRatio: false,
				responsive: true,
				title: {
					display: false,
					text: 'Sales Charts',
					fontColor: 'Black'
				},
				legend: {
					labels: {
						fontColor: 'white'
					},
					align: 'end',
					position: 'bottom'
				},
				tooltips: {
					mode: 'index',
					intersect: false
				},
				hover: {
					mode: 'nearest',
					intersect: true
				},
				scales: {
					xAxes: [
						{
							ticks: {
								fontColor: 'rgba(255,255,255,.7)'
							},
							display: false,
							scaleLabel: {
								display: false,
								labelString: 'Month',
								fontColor: 'white'
							},
							gridLines: {
								display: false,
								borderDash: [2],
								borderDashOffset: [2],
								color: 'rgba(33, 37, 41, 0.3)',
								zeroLineColor: 'rgba(0, 0, 0, 0)',
								zeroLineBorderDash: [2],
								zeroLineBorderDashOffset: [2]
							}
						}
					],
					yAxes: [
						{
							ticks: {
								fontColor: 'rgba(255,255,255,.7)'
							},
							display: true,
							scaleLabel: {
								display: false,
								labelString: 'Value',
								fontColor: 'white'
							},
							gridLines: {
								borderDash: [3],
								borderDashOffset: [3],
								drawBorder: false,
								color: 'rgba(255, 255, 255, 0.15)',
								zeroLineColor: 'rgba(33, 37, 41, 0)',
								zeroLineBorderDash: [2],
								zeroLineBorderDashOffset: [2]
							}
						}
					]
				}
			}
		};
		var ctx = document.getElementById(id).getContext('2d');
		var myLine = new Chart(ctx, config);

		const updateData = (uData) => {
			list.push(uData);
			if (list.length > 20) {
				list.shift();
			}
			myLine.update();
		};

		const socket = connectSocket('');
		sendMessage(socket, 'get_temp');
		listener(socket, 'temp_response', updateData);
	});
</script>

<div
	class="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded bg-blueGray-700"
>
	<!-- <span class="bg-white " on:click={toggleType}>Click Me</span> -->
	<div class="rounded-t mb-0 px-4 py-3 bg-transparent">
		<div class="flex flex-wrap items-center">
			<div class="relative w-full max-w-full flex-grow flex-1">
				<h6 class="uppercase text-blueGray-100 mb-1 text-xs font-semibold">Temperature Data</h6>
				<h2 class="text-white text-xl font-semibold">
					{title}
				</h2>
			</div>
		</div>
	</div>
	<div class="p-4 flex-auto">
		<!-- Chart -->
		<div class="relative h-96">
			<canvas {id} />
		</div>
	</div>
</div>
