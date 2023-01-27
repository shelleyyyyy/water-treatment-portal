<script>
	import SectorCard from '$lib/app/SectorCard.svelte';
    /** @type {import('./$types').PageData} */
    export let data;
    import PocketBase from 'pocketbase';
    import { onMount } from 'svelte';
    import Control from '$lib/app/Control.svelte';

    const pb = new PocketBase('http://192.168.1.179:8080');

    let records = []
    let sector = {}

    onMount(async () => {
        sector = await pb.collection('sectors').getOne(data.title);

        records = await pb.collection('devices').getFullList(200, {
            expand: 'sector'
        });
        // console.log(records[0].controls)
    })
</script>

<h1 class="text-xl font-bold p-5 px-10">{sector.title}</h1>

<div class="p-5">
    {#each records as r}
        <div class="card w-96 bg-primary text-primary-content">
            <div class="card-body">
                <h2 class="card-title">{r.title}</h2>
                <p>{r.description}</p>
                <div class="card-actions pt-5 justify-start">
                    <!-- {console.log(r.controls[0])} -->

                    {#each r.controls as c}
                        <Control type={c.type} topic={c.topic} options={c.options}/>
                    {/each}

                    <!-- {#each r.contorls as s}
                        <Control type={"switch"}/>
                    {/each} -->
                </div>
            </div>
        </div>
    {/each}
</div>
