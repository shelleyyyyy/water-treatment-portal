<script>

    import SectorCard from "../../lib/app/SectorCard.svelte";
    import { onMount } from 'svelte';

    import PocketBase from 'pocketbase';

    const pb = new PocketBase('http://192.168.1.179:8080');

    // you can also fetch all records at once via getFullList
    let records = []
    onMount(async () => {
        records = await pb.collection('sectors').getFullList(200 /* batch size */, {
            sort: '-created',
        });
    })

</script>

<h1 class="font-bold text-xl px-10 py-2">Sectors</h1>

<div class="grid grid-cols-3 gap-5 p-5">
    {#each records as d}
        <SectorCard title={d.title} description={d.description} link={d.id} img={d.link}/>
    {/each}
</div>
