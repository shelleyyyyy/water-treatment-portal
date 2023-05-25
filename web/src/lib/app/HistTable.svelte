<script>
    import { onMount } from "svelte";
    import { writable } from "svelte/store";
    import { pb } from '$lib/app/pocketbase.js'

    export let sub



    const state = writable({ loading: false, records: [] });


	onMount(async () => {
        state.update(s => ({ ...s, loading: true }));

		await pb.collection('historical').getFullList(200, {
            filter: `topic = "${sub.topic}"`
        })
        .then((data) => {
            console.log(sub.topic, data)
            state.set({ loading: false, records: data });

        })
	});

    function formatDateString(dateString) {
        const dateObj = new Date(dateString);
        const options = { 
            timeZone: 'America/New_York',
            month: 'long',
            day: 'numeric',
            hour: 'numeric',
            minute: 'numeric'
        };
        return dateObj.toLocaleDateString('en-US', options);
    }

</script>



{#if $state.loading}
    loading...
{:else}
<div class="overflow-x-auto overflow-y-scroll max-h-96 w-96">
    <table class="table table-compact w-full">
      <thead>
        <tr>
          <th></th> 
          <th>Time</th> 
          <th>Message</th> 
        </tr>
      </thead> 
      <tbody>
        {#each $state.records as record}
            <tr>
                <th>1</th> 
                <td>{formatDateString(record.created)}</td> 
                <td>{record.value}</td> 
            </tr>
        {/each}
    </table>
</div>
{/if}