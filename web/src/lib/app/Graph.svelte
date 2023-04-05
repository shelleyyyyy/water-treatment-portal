<script>
	import { pb } from '$lib/app/pocketbase.js';
	import { onMount } from 'svelte';
	import ApexCharts from 'apexcharts';

	export let id;
	let title = '';
	let chart = null;
	let data = [];
	const XAXISRANGE = 10 * 1000;

	onMount(async () => {
		var res = await pb.collection('subsciptions').getOne(id);
		title = res.title;

		const getNewSeries = async (baseval, yrange) => {
			const subscription = await pb.collection('subsciptions').getOne(id);
			const value = subscription.value;
			const newDate = baseval + XAXISRANGE;
			data.push({
				x: newDate,
				y: value
			});
		};

		await getNewSeries(new Date().getTime(), {
			min: 10,
			max: 90
		});

		const options = {
			series: [
				{
					data: data.slice()
				}
			],
			chart: {
				id: `realtime-${id}`,
				height: 350,
				type: 'line',
				animations: {
					enabled: true,
					easing: 'linear',
					dynamicAnimation: {
						speed: 1000
					}
				},
				toolbar: {
					show: false
				},
				zoom: {
					enabled: false
				}
			},
			dataLabels: {
				enabled: false
			},
			stroke: {
				curve: 'smooth'
			},
			title: {
				text: `Real-time graph for ${title}`,
				style: {
					color: '#000'
				}
			},
			markers: {
				size: 0
			},
			xaxis: {
				type: 'datetime',
				range: XAXISRANGE,
				labels: {
					datetimeUTC: false,
					datetimeFormatter: {
						hour: 'HH:mm:ss',
						timeZone: 'America/New_York'
					}
				},
				tooltip: {
					enabled: false
				},
				title: {
					text: 'Time',
					style: {
						color: '#000',
						fontSize: '14px'
					}
				}
			},
			yaxis: {
				max: 400,
				min: 0, // adjust the max value to accommodate the new data range
				tickAmount: 15,
				title: {
					text: 'Subscription Data',
					style: {
						color: '#000',
						fontSize: '14px'
					}
				}
			},
			legend: {
				show: false
			}
		};

		const chartElement = document.getElementById(`chart-${id}`);
		chart = new ApexCharts(chartElement, options);
		chart.render();

		pb.collection('subsciptions').subscribe(id, function (e) {
			getNewSeries(new Date().getTime(), {
				min: 10,
				max: 100 // adjust the max value to accommodate the new data range
			}).then(() => {
				chart.updateSeries([
					{
						data: data
					}
				]);
			});
			console.log(e.record.value);
		});
	});
</script>

<div class="p-5 bg-secondary rounded-lg">
	<div class="p-3 bg-white">
		<div id="chart-{id}" />
	</div>
</div>
