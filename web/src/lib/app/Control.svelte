<script>
	import Switch from "./Switch.svelte";

    export let type;
    export let topic;
    export let options;

    // const publish = async (msg) => {
    //     const response = await fetch("http://localhost:1323/save");
    //     const data = await response.json();
    //     console.log(data);

    //     console.log(msg, topic)
    // }

    let result = null

    async function publish (msg) {
		const res = await fetch('http://192.168.1.179:1323/save', {
			method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
            },
			body: JSON.stringify({
				"topic": topic,
                "message": msg
			})
		})
		
		const json = await res.json()
		result = JSON.stringify(json)
        console.log(result)
	}
    
</script>

<div class="p-5 bg-secondary rounded-lg flex ">
    {#if type == "switch"}
        <Switch publish={publish} options={options}/>
    {/if}
</div>
