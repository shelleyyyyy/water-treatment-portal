<script>
    import { onMount } from "svelte";
    import { pb } from '$lib/app/pocketbase.js'
    import HistChart from "./HistChart.svelte";

    export let sub

    let data = [1, 1]
    let loading = false

    onMount(async () => {
        loading = true
        const records = await pb.collection('historical').getFullList(20000 /* batch size */, {
            sort: 'created',
            filter: `topic = "${sub.topic}"`
        });
        let clone = records.map(person => person.value);
        data = clone

        loading = false
    })


    const testing = () => {
        console.log("HIST ChART DATA", data)
    }

</script>

<button on:click={testing} class="btn">Testing</button>

{#if loading}
   <div></div>
{:else}
    <HistChart data={data} id={sub.id + "1"}/> 
{/if}