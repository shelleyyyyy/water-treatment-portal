<script>
    import { onMount } from 'svelte'
    import { pb } from '$lib/app/pocketbase.js'
	import HistChart from './HistChart.svelte';
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
    <label for={sub.id + "h"} class="btn w-full"><h1 class="font-bold uppercase pr-2 text-red-500"> Historical Chart</h1>{sub.title} </label>

    <input type="checkbox" id={sub.id + "h"} class="modal-toggle" />
    <div class="modal">
    <div class="card p-5 bg-base-100">
        <HistChart id={sub.id} title={sub.title} />
        <div class="modal-action">
            <label for={sub.id + "h"} class="btn w-full">Close</label>
        </div>
    </div>
    </div>
{/if}
