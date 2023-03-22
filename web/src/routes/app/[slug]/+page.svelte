<script>
	import SectorCard from '$lib/app/SectorCard.svelte';
    /** @type {import('./$types').PageData} */
    export let data;
    import { onMount } from 'svelte';
    import Control from '$lib/app/Control.svelte';
    import SubCard from '$lib/app/SubCard.svelte';
    import { pb } from "$lib/app/pocketbase.js"

    let records = []
    let sector = {}
    let subs = []

    onMount(async () => {
        sector = await pb.collection('sectors').getOne(data.title);

        var res = await pb.collection('devices').getFullList(200, {
            expand: 'sector',
            expane: 'subs'
        });
        
        console.log(res)

        var forRecords = []

        res.forEach(element => {
            if (element.expand.sector.id == data.title) {
                forRecords.push(element)
            }
        });

        records = forRecords

        // var resSubs = await pb.collection('subsciptions').getFullList(200, {
        //     expand: 'device'
        // })
        
        // console.log(resSubs)

        // console.log(records)

    })

</script>

<h1 class="text-xl font-bold p-5 px-10">{sector.title}</h1>

<div class="p-5 flex gap-5 ">
    {#each records as r}
        <div class="card bg-primary text-primary-content">
            <div class="card-body">
                <h2 class="card-title text-3xl">{r.title}</h2>
                <p>{r.description}</p>
                <div class="card-actions pt-5 justify-start grid gap-3 grid-cols-1">
                    <!-- {console.log(r.controls[0])} -->

                    {#if r.controls.length > 0}
                        <div class="font-bold text-center text-2xl">
                            Controls   
                        </div>
                    

                        <div class="flex gap-1 flex-wrap">
                            {#each r.controls as c}
                                <Control type={c.type} topic={c.topic} options={c.options} title={c.title}/>
                            {/each}
                        </div>
                   {/if}
                    <!-- {#each r.expand as }

                    {/each} -->

                    {#if r.subs.length > 0}
                        <div class="font-bold text-center text-2xl">
                            Subscriptions   
                        </div>

                        <div class="flex gap-1 flex-wrap">
                            {#each r.subs as s}
                                <SubCard id={s}/>
                            {/each}
                        </div>
                    {/if}

                    <!-- {#each r.contorls as s}
                        <Control type={"switch"}/>
                    {/each} -->
                </div>
            </div>
        </div>
    {/each}
</div>
