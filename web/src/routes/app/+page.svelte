<script>

    import SectorCard from "../../lib/app/SectorCard.svelte";
    import { onMount } from 'svelte';
    import { pb, currentUser } from "$lib/app/pocketbase.js"

    let records = []
    onMount(async () => {
        records = await pb.collection('sectors').getFullList(200)
    })


</script>

<h1 class="font-bold text-xl px-10 py-2">Sectors</h1>

<div class="grid grid-cols-3 gap-5 p-5">
    {#each records as d}
        <SectorCard title={d.title} description={d.description} link={d.id} img={d.link}/>
    {/each}
</div>
