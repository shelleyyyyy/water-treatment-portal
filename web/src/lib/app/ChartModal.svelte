<script>
	import Chart from './Chart.svelte';
    import { onMount } from 'svelte'
    import { pb } from '$lib/app/pocketbase.js'
    export let sub;

    // let title = ""
    // let chart = false

    // onMount(async () => {
    //     await pb.collection('subsciptions').getOne(id)
    //     .then((e) => {
    //         console.log(e)
    //         chart = e.chart
    //         title = e.title
    //     })
        
    // })

</script>


{#if sub.chart}
    <label for={sub.id + "s"} class="btn w-full"><h1 class="font-bold uppercase px-2 text-green-500"> Live Chart</h1>{sub.title} </label>

    <input type="checkbox" id={sub.id + "s"} class="modal-toggle" />
    <div class="modal">
    <div class="card p-5 bg-base-100">
        <Chart id={sub.id} title={sub.title} />
        <div class="modal-action">
            <label for={sub.id + "s"} class="btn w-full">Close</label>
        </div>
    </div>
    </div>
{/if}
